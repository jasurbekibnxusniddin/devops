# let’s do a full, hands-on step-by-step to build a tiny Go “Hello” API, containerize it, and deploy it to Kubernetes (Minikube).
# Follow the commands in order — they work on Linux and macOS.

## Plan (what we’ll do)
1. Create a simple Go HTTP app (main.go, go.mod).
2. Build a Docker image (two options: local→minikube or push→Docker Hub).
3. Create Kubernetes manifests (Deployment + Service).
4. kubectl apply and verify.
5. Access the app (browser or curl) and show basic debugging (logs, exec, scaling, cleanup).

### 0 — Prereqs
You should have:
1. kubectl installed and configured (minikube context).
2. minikube installed and running (minikube start).
3. docker installed (if you want to build/push to Docker Hub).
4. go (for local test) — optional but helpful.

Start with minikube:
```bash
minikube start
```

### 1 — Project files
Create a directory:
```bash
mkdir hello-k8s
cd hello-k8s
go mod init hello-k8s
touch main.go
```
write to main.go file:
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello k8s!")
	})

	log.Printf("Listening on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
```
(You can go run main.go to test locally: curl http://localhost:8080)

### 2 — Dockerfile (multi-stage)
Create Dockerfile:
```Dockerfile
# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /src
COPY go.mod .
COPY main.go .
RUN go build -o /hello

# Run stage
FROM alpine:3.18
RUN adduser -D -u 1000 appuser
COPY --from=builder /hello /hello
USER appuser
EXPOSE 8080
ENTRYPOINT ["/hello"]
```
### 3 — Build image — two options

**Option A — Build image inside Minikube (no push)**

This is easiest for local dev: minikube will use the image directly.
```bash
# Point your Docker CLI to minikube's Docker daemon
eval $(minikube -p minikube docker-env)

# Build image
docker build -t hello-k8s:1.0 .

# (Optional) verify image exists in minikube's docker:
docker images | grep hello-k8s

# Switch Docker env back to normal (optional)
eval $(minikube -p minikube docker-env -u)
```
**Important**: If you build in minikube's docker env, use image: hello-k8s:1.0 in your Deployment and keep imagePullPolicy: IfNotPresent (or omit).

**Option B — Build locally and push to Docker Hub**

Use this if you want a registry image (or if not using minikube).
1. Tag and push:
    ```bash
    docker build -t <your-dockerhub-username>/hello-k8s:1.0 .
    docker push <your-dockerhub-username>/hello-k8s:1.0
    ```
2. In the Deployment YAML, set image: <your-dockerhub-username>/hello-k8s:1.0 and imagePullPolicy: Always.

(You must be docker logined to push.)

### 4 — Kubernetes manifests
Create deployment.yaml:
```yml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-deployment
spec:
  replicas: 2
  selector:
    matchLabels:
      app: hello
  template:
    metadata:
      labels:
        app: hello
    spec:
      containers:
      - name: hello
        image: hello-k8s:1.0         # <- if using Minikube build
        # image: <your-dockerhub-username>/hello-k8s:1.0  # <- if pushing to Docker Hub
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8080
        env:
        - name: PORT
          value: "8080"
```
Create service.yaml:
```yml
apiVersion: v1
kind: Service
metadata:
  name: hello-service
spec:
  type: NodePort
  selector:
    app: hello
  ports:
  - port: 8080
    targetPort: 8080
    nodePort: 30080    # fixed NodePort (optional) — you can omit nodePort and let k8s pick
```
If you used Docker Hub images, set imagePullPolicy: Always (or remove the explicit IfNotPresent and rely on default).

### 5 — Deploy to Kubernetes
Apply manifests:
```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```
Check resources:
```bash
kubectl get deployments
kubectl get pods -l app=hello
kubectl get svc hello-service
```
Example expected output:
```bash
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
hello-deployment   2/2     2            2           30s

NAME                                TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
hello-service                       NodePort    10.96.129.213   <none>        8080:30080/TCP   10s
```

### 6 — Access the app
**Option 1 — minikube service (opens browser)**
```bash
minikube service hello-service
```
This opens a browser to the right URL (minikube handles routing).

**Option 2 — Use minikube ip + NodePort**
```bash
minikube ip
# suppose returns 192.168.49.2
# then open http://192.168.49.2:30080
curl http://$(minikube ip):30080
```

**Option 3 — Port-forward (works even without NodePort)**
```bash
kubectl port-forward svc/hello-service 8080:8080
# then curl localhost:8080
curl http://localhost:8080
```
You should see:
```txt
hello k8s!
```
### 7 — Useful ops & debugging
Get pod names:
```bash
kubectl get pods -l app=hello -o wide
```

View logs:
```bash
kubectl logs <pod-name>
# or for a deployment's first pod:
kubectl logs -l app=hello
```
Exec into a pod:
```bash
kubectl exec -it <pod-name> -- sh
# inside you can run: wget -qO- http://localhost:8080
```

Scale up/down:
```bash
kubectl scale deployment hello-deployment --replicas=4
kubectl get pods
```

Update image (rolling update):
```bash
# If you pushed a new image tag to Docker Hub:
kubectl set image deployment/hello-deployment hello=<your-dockerhub-username>/hello-k8s:1.1
kubectl rollout status deployment/hello-deployment

# Rollback:
kubectl rollout undo deployment/hello-deployment
```

Describe resource to see events:
```bash
kubectl describe pod <pod-name>
kubectl get events --sort-by=.metadata.creationTimestamp
```

### 8 — Cleanup
When done:
```bash
kubectl delete -f deployment.yaml
kubectl delete -f service.yaml
# or
kubectl delete deployment hello-deployment
kubectl delete svc hello-service
```
If you built inside minikube and want to stop:
```bash
minikube stop
```
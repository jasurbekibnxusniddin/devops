# Kubernetes

## list of the main Kubernetes terms

### Before diving deep into Kubernetes, it’s smart to build a solid vocabulary — understanding the core terms will make the rest of the learning journey much easier. Here’s a structured list of the main Kubernetes terms grouped by category:

#### 🧱 Core Concepts

| Term              | Description                                                                                                            |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------- |
| **Cluster**       | A set of machines (nodes) that run containerized applications managed by Kubernetes.                                   |
| **Node**          | A single machine (virtual or physical) that runs workloads. It can be a **Master (Control Plane)** or **Worker Node**. |
| **Pod**           | The smallest deployable unit in Kubernetes — usually runs one or more tightly coupled containers.                      |
| **Container**     | A lightweight, standalone executable package that includes everything needed to run an application.                    |
| **Namespace**     | A virtual cluster inside a Kubernetes cluster used to isolate resources (e.g., dev, staging, prod).                    |
| **Control Plane** | The brain of the cluster — manages scheduling, scaling, and the cluster state.                                         |
| **Kubelet**       | Agent running on each node ensuring that containers are running as expected.                                           |
| **kubectl**       | Command-line tool to interact with your cluster (`kubectl get pods`, `kubectl apply -f`, etc.).                        |

#### ⚙️ Workloads
| Term            | Description                                                                                          |
| --------------- | ---------------------------------------------------------------------------------------------------- |
| **Deployment**  | Defines how to deploy and manage a set of identical Pods (handles scaling, rolling updates, etc.).   |
| **ReplicaSet**  | Ensures a specified number of pod replicas are running at any time (usually managed by Deployments). |
| **DaemonSet**   | Ensures a copy of a pod runs on **every** (or specific) node. Common for system agents.              |
| **StatefulSet** | Manages **stateful** applications (e.g., databases) with stable identities and storage.              |
| **Job**         | Runs a pod to completion for batch tasks.                                                            |
| **CronJob**     | Like a Job, but runs on a schedule (e.g., backups every night).                                      |

#### 🌐 Networking
| Term              | Description                                                                            |
| ----------------- | -------------------------------------------------------------------------------------- |
| **Service**       | Provides stable networking and load balancing for pods.                                |
| **ClusterIP**     | Internal-only service accessible within the cluster.                                   |
| **NodePort**      | Exposes a service on each node’s IP at a static port.                                  |
| **LoadBalancer**  | Exposes a service externally via a cloud provider’s load balancer.                     |
| **Ingress**       | Manages external HTTP/HTTPS access to services, usually with routing rules.            |
| **NetworkPolicy** | Defines how pods are allowed to communicate with each other and with external systems. |

#### 💾 Storage
| Term                            | Description                                                               |
| ------------------------------- | ------------------------------------------------------------------------- |
| **Volume**                      | Storage attached to a pod (lifecycle tied to the pod).                    |
| **PersistentVolume (PV)**       | A piece of storage in the cluster provisioned by an admin or dynamically. |
| **PersistentVolumeClaim (PVC)** | A request for storage by a user (binds to a PV).                          |
| **StorageClass**                | Defines different types of storage (e.g., SSD vs HDD).                    |

#### 🔒 Security & Configuration
| Term                                 | Description                                                                             |
| ------------------------------------ | --------------------------------------------------------------------------------------- |
| **ConfigMap**                        | Stores non-confidential configuration data (e.g., environment variables, config files). |
| **Secret**                           | Stores sensitive data like passwords, tokens, or keys.                                  |
| **ServiceAccount**                   | Provides an identity for processes running in pods.                                     |
| **Role & ClusterRole**               | Define sets of permissions.                                                             |
| **RoleBinding & ClusterRoleBinding** | Grant roles to users or service accounts.                                               |

#### 🧩 Other Important Terms
| Term                   | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| **Scheduler**          | Decides which node a new pod should run on.                                   |
| **Controller Manager** | Runs controllers that manage the state of the cluster (e.g., node lifecycle). |
| **API Server**         | Central management endpoint — all operations go through it.                   |
| **etcd**               | Key-value store used by Kubernetes to persist the cluster state.              |
| **Helm**               | A package manager for Kubernetes (think “apt” or “yum” for K8s).              |
| **Operator**           | A custom controller extending Kubernetes to manage complex apps.              |

## What Is Kubernetes?
### 🧭 1. What Exactly Is Kubernetes?

Kubernetes (a.k.a. “K8s”) is an open-source system for **automating the deployment**, **scaling**, and **management** of **containerized applications**.

####  Let’s break that into plain words 👇
| Concept                        | Meaning                                                                           |
| ------------------------------ | --------------------------------------------------------------------------------- |
| **Automating deployment**      | It runs your containers (apps) for you — no manual starting/stopping.             |
| **Scaling**                    | It can add or remove containers automatically when load changes.                  |
| **Management**                 | It keeps your apps healthy (auto-restarts if they crash, balances traffic, etc.). |
| **Containerized applications** | Apps that run inside containers (like Docker containers).                         |

So, Kubernetes is like a “container orchestrator” — a manager that controls a group of containers across multiple machines.

### ⚙️ 2. Why Do We Need It?
Docker alone is great for running a single container on one machine, but problems appear when you grow:

| Problem with Docker alone                       | How Kubernetes solves it                            |
| ----------------------------------------------- | --------------------------------------------------- |
| How do I run the same app on multiple servers?  | K8s schedules and runs containers across a cluster. |
| What if one container crashes?                  | K8s automatically restarts it.                      |
| How do I update an app without downtime?        | K8s does rolling updates.                           |
| How do I handle 10,000 requests instead of 100? | K8s scales pods automatically.                      |
| How do containers communicate securely?         | K8s has built-in networking and service discovery.  |

> Think of it like this:
>
> Docker = runs containers 
> Kubernetes = manages lots of containers, everywhere.

### 🏗️ 3. How Kubernetes Works (Big Picture)
Imagine a Kubernetes Cluster as a small city:

* The Control Plane is the city hall 🏛️ — it makes all the decisions.
* The Nodes (workers) are the buildings 🏢 — where the real work happens.
* The Pods are the rooms 🛠️ inside those buildings — each running your containers.

Architecture Summary
| Component              | Description                                                                     |
| ---------------------- | ------------------------------------------------------------------------------- |
| **API Server**         | The “front door” — everything goes through it.                                  |
| **etcd**               | Database that stores all cluster info (state).                                  |
| **Scheduler**          | Decides which Node runs which Pod.                                              |
| **Controller Manager** | Keeps the system running in the desired state (e.g., ensures 3 replicas exist). |
| **Kubelet**            | Runs on every Node, talks to API Server, runs containers.                       |
| **kube-proxy**         | Handles network routing between Pods/Services.                                  |

So when you say:
```sh
kubectl apply -f deployment.yaml
```
Kubernetes:

1. Reads your desired state from the YAML file (e.g., 3 pods of nginx).
0. Stores it in etcd.
0. Scheduler picks which nodes to use.
0. Kubelets on those nodes create the containers using Docker (or another runtime).
0. Controller keeps checking that 3 pods are running — if one dies, it restarts it.

### 🧩 4. How Kubernetes Relates to Docker
Historically:

* Kubernetes used Docker as its default container runtime.
* But Kubernetes doesn’t require Docker anymore — it uses a generic interface called CRI (Container Runtime Interface).

👉 So now, K8s can use:

* containerd (used inside Docker)
* CRI-O
* Docker Engine (through a compatibility layer)

So the relationship today looks like this:
```
Developer → builds container with Docker → pushes to registry → 
Kubernetes → pulls image → runs it via containerd or Docker runtime
```
In short:
- Docker is how you build and package apps.
- Kubernetes is how you run and manage them at scale.

### 🛠️ 5. Key Kubernetes Concepts (In One Picture)
```
+-----------------------------------------------------------+
|                     Kubernetes Cluster                    |
|                                                           |
|  +--------------------+        +--------------------+      |
|  |      Node 1        |        |      Node 2        |      |
|  |  (worker machine)  |        |  (worker machine)  |      |
|  |                    |        |                    |      |
|  |  +--------------+  |        |  +--------------+  |      |
|  |  |   Pod (nginx) | |        |  |   Pod (api)   | |      |
|  |  +--------------+  |        |  +--------------+  |      |
|  |  |   Pod (redis) | |        |  |   Pod (db)    | |      |
|  |  +--------------+  |        |  +--------------+  |      |
|  +--------------------+        +--------------------+      |
|                                                           |
+-----------------------------------------------------------+
           ↑
           |
        Control Plane
     (API Server, Scheduler,
      etcd, Controller Manager)
```

### 💬 6. How People Usually Call It
- Kubernetes → full name
- K8s → nickname (because there are 8 letters between “K” and “s”)
- Cluster → your group of servers managed by K8s
- Pods / Deployments / Services → your building blocks inside the cluster

Example sentence in tech talk:
> “We deployed our microservices on a GKE Kubernetes cluster. Each service runs as a Deployment with 3 pods behind a LoadBalancer Service.”
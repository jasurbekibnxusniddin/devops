#

# list of the main Kubernetes terms

## Before diving deep into Kubernetes, it’s smart to build a solid vocabulary — understanding the core terms will make the rest of the learning journey much easier. Here’s a structured list of the main Kubernetes terms grouped by category:

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

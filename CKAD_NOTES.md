# CKAD Exam Study Notes

Comprehensive, structured study guide for the **Certified Kubernetes Application Developer (CKAD)** exam, organized according to the official CKAD curriculum domains.

> **Note:** The original unorganized notes remain preserved in [README.md](README.md). Examples referenced throughout this guide correspond to the manifests in [`manifests/`](manifests/) and [`probes/`](probes/).

---

## Table of Contents

- [CKAD Exam Domains Overview](#ckad-exam-domains-overview)
- [1. Application Design and Build (20%)](#1-application-design-and-build-20)
  - [Building & Managing Container Images](#building--managing-container-images)
  - [Pods & Pod Lifecycle](#pods--pod-lifecycle)
  - [Multi-Container Pod Patterns](#multi-container-pod-patterns)
  - [Workload Controllers](#workload-controllers)
  - [Storage & Volumes](#storage--volumes)
- [2. Application Deployment (20%)](#2-application-deployment-20)
  - [Deployment Strategies](#deployment-strategies)
  - [Rollouts & Rollbacks](#rollouts--rollbacks)
  - [Pod Scheduling & Placement](#pod-scheduling--placement)
  - [Horizontal Pod Autoscaling (HPA)](#horizontal-pod-autoscaling-hpa)
  - [Helm Package Manager](#helm-package-manager)
- [3. Application Observability and Maintenance (15%)](#3-application-observability-and-maintenance-15)
  - [Container Probes (Startup, Readiness, Liveness)](#container-probes-startup-readiness-liveness)
  - [Monitoring & Debugging](#monitoring--debugging)
  - [API Deprecations](#api-deprecations)
  - [Custom Resource Definitions (CRDs) & Operators](#custom-resource-definitions-crds--operators)
- [4. Application Environment, Configuration and Security (25%)](#4-application-environment-configuration-and-security-25)
  - [ConfigMaps](#configmaps)
  - [Secrets](#secrets)
  - [Resource Management (Requests, Limits, Quotas, LimitRanges, QoS, PriorityClass)](#resource-management)
  - [Security Context & Linux Security Fundamentals](#security-context--linux-security-fundamentals)
  - [ServiceAccounts & RBAC](#serviceaccounts--rbac)
- [5. Services and Networking (20%)](#5-services-and-networking-20)
  - [Kubernetes Networking Fundamentals](#kubernetes-networking-fundamentals)
  - [Services](#services)
  - [Ingress](#ingress)
  - [NetworkPolicies](#networkpolicies)
- [6. Essential `kubectl` Cheat Sheet](#6-essential-kubectl-cheat-sheet)
- [7. Repository Manifests & Practice Resources](#7-repository-manifests--practice-resources)

---

## CKAD Exam Domains Overview

| Domain | Weight | Key Focus Areas |
| :--- | :---: | :--- |
| **Application Design and Build** | 20% | Container images, Jobs/CronJobs, Multi-container Pods, Volumes/PVCs |
| **Application Deployment** | 20% | Deployments, Rolling/Blue-Green/Canary strategies, Helm |
| **Application Observability and Maintenance** | 15% | Probes, Logs, Debugging, API deprecations, CRDs |
| **Application Environment, Configuration and Security** | 25% | ConfigMaps, Secrets, Resources/Quotas/LimitRanges, SecurityContext, RBAC/ServiceAccounts |
| **Services and Networking** | 20% | Services, Ingress, NetworkPolicies, Troubleshooting connectivity |

---

## 1. Application Design and Build (20%)

### Building & Managing Container Images

Understanding how container images are built and configured is required on the CKAD exam (using `docker` or `podman`).

- **Multi-Stage Builds:** Reduce final image size and attack surface by compiling in a builder stage and copying only the binary into a minimal runtime image (see [`probes/Dockerfile`](probes/Dockerfile)):
  ```dockerfile
  FROM golang:1.25 AS builder
  WORKDIR /app
  COPY go.mod go.sum ./
  RUN go mod download
  COPY . ./
  RUN CGO_ENABLED=0 go build -o ./bin/probes .

  FROM ubuntu:22.04
  COPY --from=builder /app/bin/probes /opt/app/probes
  EXPOSE 3000
  ENTRYPOINT ["/opt/app/probes"]
  ```
- **Common Image Commands (`docker` / `podman`):**
  ```bash
  docker build -t myapp:1.0 .                   # Build an image from Dockerfile
  docker tag myapp:1.0 registry.io/myapp:1.0    # Tag for a registry
  docker push registry.io/myapp:1.0             # Push to a registry
  docker save -o myapp.tar myapp:1.0            # Export image as a tar archive
  docker load -i myapp.tar                      # Load image from a tar archive
  ```
- **`ENTRYPOINT` vs `CMD` vs Kubernetes `command`/`args`:**
  - Dockerfile `ENTRYPOINT` maps to Kubernetes `spec.containers[].command`.
  - Dockerfile `CMD` maps to Kubernetes `spec.containers[].args`.
  - Setting `command` in the Pod spec overrides both `ENTRYPOINT` and `CMD` from the image. Setting only `args` overrides `CMD` while keeping `ENTRYPOINT`.
- **`imagePullPolicy`:**
  - `IfNotPresent` (default when tag is not `:latest`): Pull only if not already cached on the node.
  - `Always` (default when tag is `:latest` or omitted): Always contact the registry to check the image digest.
  - `Never`: Never pull; only use locally cached images.

---

### Pods & Pod Lifecycle

- **Pod:** The smallest deployable unit in Kubernetes. It runs one or more containers that share the same **network namespace** (one Pod IP, shared `localhost`, shared port space) and can share **storage volumes**, scheduled together onto a single Node.
- **`restartPolicy`** (applies to all containers in a Pod):
  - `Always` (default, required for Deployments)
  - `OnFailure` (restarts only if container exits with non-zero code; common for Jobs)
  - `Never` (never restarts; common for Jobs / one-off debug Pods)
- **Pod Termination & Lifecycle Hooks:**
  - `terminationGracePeriodSeconds` (default `30`s): Time Kubernetes waits after sending `SIGTERM` before forcibly killing the container with `SIGKILL`.
  - `lifecycle.postStart`: Executed immediately after a container is created (runs asynchronously with entrypoint).
  - `lifecycle.preStop`: Executed synchronously before `SIGTERM` is sent—useful for graceful connection draining.

---

### Multi-Container Pod Patterns

Multiple containers in a single Pod share lifecycle, network (`localhost`), and volumes (`emptyDir`).

1. **Sidecar:** Runs alongside the main application container to provide supporting functionality such as logging, monitoring, or proxying.
   - *Problem:* A web server writes logs to `/var/log/app`, and you want to ship logs to a central system without modifying the web server code.
   - *Solution:* Mount a shared `emptyDir` volume at `/var/log/app` in both containers. The sidecar tails the log file and forwards it externally.
2. **Ambassador:** Acts as an outbound proxy for the main container, handling service discovery, routing, TLS termination, or circuit breaking.
   - *Problem:* Your app must connect to a sharded database, and sharding/routing logic shouldn't live inside the app.
   - *Solution:* An ambassador proxy container listens on `localhost:5432`. The main container connects to `localhost:5432`, and the ambassador routes requests to the appropriate external database shard.
3. **Adapter:** Standardizes or transforms the output/metrics of the main container into a format expected by external consumers.
   - *Problem:* A legacy app exposes metrics in a non-standard XML format over HTTP, while your monitoring stack only scrapes Prometheus text format.
   - *Solution:* An adapter container queries the main container via `localhost`, converts XML to Prometheus metrics, and exposes them on another port for scraping.
4. **Init Container (`initContainers`):** Runs to completion sequentially *before* any main application containers start.
   - *Problem:* The main app needs a dynamically fetched `config.json` or needs to wait for a dependent database service before booting.
   - *Solution:* An init container runs a setup script, fetches the configuration into a shared `emptyDir` volume (or polls the dependency), and exits `0`. Only after all init containers succeed do the main containers start.
5. **Native Sidecar Containers (Kubernetes 1.28+ / GA 1.33):**
   - Configured under `initContainers` with `restartPolicy: Always`.
   - Starts *before* main containers (like an init container), keeps running alongside them, and is terminated *after* main containers exit—solving the classic problem where normal sidecars could block Jobs from completing.

> **Note:** Ambassador and Adapter are specialized variations of the Sidecar pattern.

---

### Workload Controllers

#### ReplicaSet
Ensures a specified number of identical Pod replicas are running at any given time via a label selector.
- Maintains pod count stability (replaces failed/deleted pods).
- Does **not** provide automated rolling updates or rollback history (use a `Deployment` instead).

#### Deployment
Manages ReplicaSets declaratively to keep a desired number of identical Pods running.
- Creates a new **Revision** (and new ReplicaSet) whenever `spec.template` (`PodTemplateSpec`) changes (e.g., image, env vars, args, resource requests/limits).
- Provides automated rolling updates and rollbacks to previous revisions.

#### DaemonSet
Ensures that a single copy of a Pod runs on **every Node** (or a subset via `nodeSelector`/affinity/tolerations) in the cluster. Automatically adds Pods to newly joined nodes and cleans up when nodes are removed.
- *Common Use Cases:* Node log collectors (Fluentd, Filebeat), node monitoring agents (Prometheus Node Exporter), CNI network plugins, storage daemons.

#### StatefulSet
Manages stateful applications that require **stable, persistent identity and storage**.
- **Predictable Pod Names & Ordering:** Pods are named `<statefulset-name>-0`, `<statefulset-name>-1`, etc., created sequentially (`0 -> 1 -> 2`) and terminated in reverse order (`2 -> 1 -> 0`).
- **Stable Network Identity:** Paired with a **Headless Service** (`clusterIP: None`), giving each Pod a predictable DNS name: `<pod-name>.<service-name>.<namespace>.svc.cluster.local`.
- **Stable Storage:** Uses `volumeClaimTemplates` to automatically provision a dedicated PVC per Pod ordinal. If `<pod>-0` restarts or moves to another node, it reattaches to its original PVC.
- *Common Use Cases:* Databases (PostgreSQL, MySQL clusters), distributed systems (Kafka, ZooKeeper, etcd, Redis Sentinel).

#### Job
Creates one or more Pods and ensures a specified number of them **successfully terminate** (exit `0`). Jobs treat each Pod execution as an immutable attempt for clean retry counting and tracking.
- **Key Configuration Parameters:**
  - `completions`: Number of Pods that must succeed before the Job is marked complete (default `1`).
  - `parallelism`: Maximum number of Pods allowed to run concurrently (default `1`).
  - `backoffLimit`: Maximum retries on failure before marking the Job as Failed (default `6`).
  - `activeDeadlineSeconds`: Hard timeout for the entire Job; terminates running Pods and marks Job failed with `DeadlineExceeded`.
  - `ttlSecondsAfterFinished`: Automatically deletes the completed (Succeeded or Failed) Job and its Pods after this many seconds.

#### CronJob
Creates Jobs on a repeating cron schedule (backups, report generation, DB maintenance).
- **Schedule Format:**
  ```text
  # ┌───────────── minute (0 - 59)
  # │ ┌───────────── hour (0 - 23)
  # │ │ ┌───────────── day of the month (1 - 31)
  # │ │ │ ┌───────────── month (1 - 12)
  # │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday)
  # * * * * *
  ```
- **Key Configuration Parameters:**
  - `concurrencyPolicy`:
    - `Allow` (default): Multiple Jobs can run concurrently.
    - `Forbid`: Skips new run if the previous Job is still running.
    - `Replace`: Cancels the currently running Job and starts a new one.
  - `startingDeadlineSeconds`: Max seconds a Job can be delayed from its scheduled time before the missed schedule is skipped.
  - `successfulJobsHistoryLimit` (default `3`) / `failedJobsHistoryLimit` (default `1`): Number of finished Job records to retain.
  - `suspend`: Boolean to pause future scheduled executions without deleting the CronJob.
  - `jobTemplate`: Embeds the standard `Job` specification.

---

### Storage & Volumes

#### Ephemeral & Node-Local Volumes
- **`emptyDir`:** Ephemeral storage created when a Pod is assigned to a Node and deleted when the Pod is removed.
  - Shared among containers in the same Pod.
  - Backed by node disk (or memory if `emptyDir.medium: Memory`).
- **`hostPath`:** Mounts a file or directory from the host Node's filesystem into the Pod.
  - Exposes host security risks, lacks cross-node consistency, and is not managed for disk quotas.
  - **Avoid in production** unless building node-level system DaemonSets.
- **`downwardAPI`:** Exposes Pod or Container metadata (Pod name, namespace, IP, labels, annotations, resource limits) as files inside the container.
- **`projected`:** Combines multiple volume sources (`secret`, `configMap`, `downwardAPI`, `serviceAccountToken`) into a single mount directory.

#### Persistent Storage (PV, PVC, StorageClass)
- **PersistentVolume (PV):** A cluster-wide storage resource provisioned statically by an admin or dynamically via a StorageClass.
- **PersistentVolumeClaim (PVC):** A namespaced request for storage by a Pod (similar to how a Pod requests CPU/memory from a Node). Pods use PVCs to decouple workloads from underlying storage implementation details.
- **StorageClass:** A template defining a storage tier (e.g., SSD, standard HDD) and its provisioner, enabling dynamic PV creation.
  - `reclaimPolicy`: `Delete` (default for dynamic provisioning; deletes PV and underlying storage when PVC is deleted) or `Retain` (keeps data after PVC deletion for manual recovery).
  - `volumeBindingMode`: `Immediate` (binds PVC right away) vs `WaitForFirstConsumer` (delays binding until a Pod using the PVC is scheduled, ensuring topology/zone alignment).
- **PVC Access Modes:**
  - `ReadWriteOnce` (`RWO`): Mounted read-write by a **single node** (multiple Pods on that same node can share it).
  - `ReadOnlyMany` (`ROX`): Mounted read-only by **many nodes**.
  - `ReadWriteMany` (`RWX`): Mounted read-write by **many nodes** simultaneously (e.g., NFS, EFS).
  - `ReadWriteOncePod` (`RWOP`): Mounted read-write by a **single Pod** across the entire cluster.

> **Important Scaling Caveat:** The first Pod that mounts a `ReadWriteOnce` PVC causes its PV to attach to that Pod's Node. A Deployment with `replicas > 1` sharing a single `RWO` PVC cannot scale across multiple Nodes—Pods scheduled on other Nodes will fail to attach the volume.

---

## 2. Application Deployment (20%)

### Deployment Strategies

| Strategy | Behavior | Downtime | Rollback Speed | Use Case |
| :--- | :--- | :---: | :---: | :--- |
| **RollingUpdate** *(default)* | Gradually replaces old Pods with new ones | Zero | Moderate | Standard stateless apps |
| **Recreate** | Terminates all existing Pods before creating new ones | Yes | Moderate | Apps that cannot run two versions concurrently (e.g., schema locks, single RWO PVC) |
| **Blue-Green** | Runs two full environments; switches Service selector from Blue (`v1`) to Green (`v2`) | Zero | Instant | Zero-downtime cutovers with instant rollback |
| **Canary** | Routes a small percentage of traffic to `v2` alongside `v1` before full rollout | Zero | Instant | Validating risky releases on real traffic |

- **RollingUpdate Tuning (`spec.strategy.rollingUpdate`):**
  - `maxSurge`: Max Pods allowed *above* desired `replicas` during update (number or `%`, default `25%`).
  - `maxUnavailable`: Max Pods allowed to be *unavailable* during update (number or `%`, default `25%`).
  - *Rule of thumb:* Deployments prioritize availability over exact replica count during updates—temporary over- or under-provisioning is expected.
- **Blue-Green in Kubernetes:**
  1. Deploy `v1` (`labels: { app: myapp, version: blue }`) behind a `Service` selecting `app: myapp, version: blue`.
  2. Deploy `v2` (`labels: { app: myapp, version: green }`) and test it internally.
  3. Patch the Service selector to `version: green` for instant cutover.
- **Canary in Kubernetes:**
  - *Replica-ratio method:* Both `v1` (e.g., 9 replicas) and `v2` (e.g., 1 replica) share the Service's label selector (`app: myapp`), sending ~10% of traffic to `v2`.
  - *Ingress-based method:* Use Ingress controller annotations (e.g., NGINX `canary-weight`) for exact percentage traffic splitting.

---

### Rollouts & Rollbacks

```bash
kubectl rollout status deployment <name>                     # Watch rollout progress
kubectl rollout history deployment <name>                    # List all revisions
kubectl rollout history deployment <name> --revision=2       # Inspect a specific revision
kubectl rollout undo deployment <name>                       # Rollback to previous revision
kubectl rollout undo deployment <name> --to-revision=1       # Rollback to a specific revision
kubectl rollout pause deployment <name>                      # Pause rollout (batch multiple edits)
kubectl rollout resume deployment <name>                     # Resume rollout (triggers single new revision)
kubectl rollout restart deployment <name>                    # Rolling restart of all Pods
```

> **Tip:** `rollout pause` does not affect Pods that have already entered their terminating state. Use `kubectl annotate deployment <name> kubernetes.io/change-cause="Updated to v2"` to record descriptive notes in rollout history.

---

### Pod Scheduling & Placement

- **Labels & Selectors:** Key-value pairs attached to objects (`metadata.labels`) used by Services, Deployments, and NetworkPolicies to select targets.
  - Equality-based: `-l env=prod,tier!=frontend`
  - Set-based: `-l 'env in (prod, staging)'`
- **Annotations (`metadata.annotations`):** Non-identifying metadata (build info, Ingress controller settings, change causes); cannot be used for label selection.
- **`nodeSelector`:** Simplest node constraint; requires exact key-value match in `node.metadata.labels`.
- **Node Affinity (`spec.affinity.nodeAffinity`):**
  - `requiredDuringSchedulingIgnoredDuringExecution`: Hard rule (Pod stays `Pending` if no node matches).
  - `preferredDuringSchedulingIgnoredDuringExecution`: Soft rule with weight `1-100` (scheduler tries to match, schedules elsewhere if needed).
- **Pod Affinity / Anti-Affinity:** Co-locate Pods together (e.g., cache next to app) or spread Pods across failure domains (`topologyKey: kubernetes.io/hostname` or `topology.kubernetes.io/zone`).
- **Taints & Tolerations:**
  - **Taint** (on Node): Repels Pods unless they have a matching **Toleration** (`kubectl taint nodes node1 key=value:NoSchedule`).
  - Effects: `NoSchedule` (won't schedule new Pods), `PreferNoSchedule` (soft), `NoExecute` (evicts existing Pods without matching toleration, unless `tolerationSeconds` delays eviction).

---

### Horizontal Pod Autoscaling (HPA)

HPA (`autoscaling/v2`) automatically scales the number of Pods in a Deployment, ReplicaSet, or StatefulSet based on observed metrics (CPU, memory, or custom/external metrics).

- **Prerequisites:**
  1. **Metrics Server** must be installed in the cluster to collect CPU/memory metrics from each node's kubelet.
  2. Target Pods **must** define resource `requests` (e.g., `requests.cpu`) so HPA can calculate utilization percentage (`currentUsage / requestedResource`).
- **Control Loop:**
  1. HPA controller queries the Metrics API every **15 seconds** (default `--horizontal-pod-autoscaler-sync-period`).
  2. Computes desired replicas: `desiredReplicas = ceil[currentReplicas * (currentMetricValue / desiredMetricValue)]`.
  3. Updates the target's `replicas` within `[minReplicas, maxReplicas]`.
- **Stabilization Windows (Preventing Thrashing):**
  - **Scale-down stabilization window:** Defaults to **300 seconds (5 minutes)**—HPA picks the highest recommended replica count over the last 5 minutes before scaling down, preventing premature scale-ins during brief traffic dips.
  - **Scale-up stabilization window:** Defaults to **0 seconds** (scales up immediately, bounded by default scale-up rate policies).
  - Customizable via `spec.behavior.scaleUp` and `spec.behavior.scaleDown` in `autoscaling/v2`.
- **Imperative Command & Manifest:**
  ```bash
  kubectl autoscale deployment webapp --min=10 --max=20 --cpu-percent=85
  ```
  See [`manifests/autoscaler/horizontal.yaml`](manifests/autoscaler/horizontal.yaml) for an `autoscaling/v2` manifest example.

---

### Helm Package Manager

Helm is the standard package manager for Kubernetes. It packages Kubernetes manifests into **Charts** parameterized with **Values**, deploying them into clusters as versioned **Releases**.

- **Key Concepts:**
  - **Chart:** A directory or `.tgz` archive containing `Chart.yaml` (metadata), `values.yaml` (defaults), and `templates/` (Go-templated Kubernetes manifests).
  - **Repository:** An HTTP server storing packaged charts.
  - **Release:** A specific deployed instance of a Chart in a namespace (tracks revision history for upgrades/rollbacks).
- **Essential Helm Commands for CKAD:**
  ```bash
  # Repositories
  helm repo add bitnami https://charts.bitnami.com/bitnami
  helm repo update
  helm search repo nginx                              # Search charts in added repos
  helm show values bitnami/nginx > values.yaml        # Inspect configurable values

  # Installing & Upgrading
  helm install my-rel bitnami/nginx -n web --create-namespace
  helm install my-rel bitnami/nginx --set replicaCount=3 -f custom-values.yaml
  helm upgrade --install my-rel bitnami/nginx --set image.tag=1.25

  # Inspecting & Rendering (great for dry-runs)
  helm template my-rel bitnami/nginx > rendered.yaml  # Render manifests locally without cluster
  helm list -A                                        # List releases across all namespaces
  helm list --pending -A                              # Find stuck/pending releases
  helm get values my-rel                              # View user-supplied values for a release

  # Rollback & Cleanup
  helm history my-rel
  helm rollback my-rel 1                              # Rollback to revision 1
  helm uninstall my-rel
  ```

---

## 3. Application Observability and Maintenance (15%)

### Container Probes (Startup, Readiness, Liveness)

Probes are diagnostic checks performed by the **kubelet** to monitor container health.

> **Mental Model:** *Startup protects boot, readiness protects traffic, liveness protects uptime.*

- **Probe Types:**
  1. **`startupProbe`:** Checks if the application inside the container has finished starting. Disables `livenessProbe` and `readinessProbe` until it succeeds. Ideal for slow-starting legacy or JVM apps so they aren't prematurely killed by liveness checks.
  2. **`readinessProbe`:** Checks if the container is ready to accept network traffic. On failure, Kubernetes removes the Pod IP from matching Service `Endpoints` (does **not** restart the container).
  3. **`livenessProbe`:** Checks if the container is still alive and healthy (not deadlocked). On failure, kubelet kills and restarts the container according to `restartPolicy`.
- **Probe Handlers:**
  - **`httpGet`:** Sends an HTTP GET request to a specified `path` and `port`. Status codes `200–399` indicate success; `4xx/5xx` or connection errors indicate failure.
  - **`tcpSocket`:** Attempts to open a TCP connection on the specified `port`. Succeeds if the connection is established (useful for databases, message queues).
  - **`exec`:** Runs a command inside the container. Exit code `0` indicates success; non-zero or timeout indicates failure.
  - **`grpc`:** Checks a gRPC health endpoint implementing the gRPC Health Checking Protocol.
- **Timing & Threshold Fields:**
  - `initialDelaySeconds` (default `0`): Seconds to wait after container start before running the first probe.
  - `periodSeconds` (default `10`): How often (in seconds) to perform the probe.
  - `timeoutSeconds` (default `1`): Seconds after which the probe times out.
  - `failureThreshold` (default `3`): Consecutive failures required to mark the probe failed.
  - `successThreshold` (default `1`): Consecutive successes required after a failure to mark the probe healthy (must be `1` for liveness and startup probes).

See [`probes/deployment.yaml`](probes/deployment.yaml) for a full working example.

---

### Monitoring & Debugging

- **Resource Metrics (requires Metrics Server):**
  ```bash
  kubectl top nodes
  kubectl top pods -A --sort-by=cpu
  kubectl top pods -n prod --containers               # Per-container breakdown
  ```
- **Logs:**
  ```bash
  kubectl logs <pod>                                  # Standard output of single-container Pod
  kubectl logs <pod> -c <container>                   # Specific container in multi-container Pod
  kubectl logs <pod> --previous                       # Logs from the previous crashed container instance
  kubectl logs -l app=web --tail=50 -f                # Tail & follow logs across all Pods matching label
  ```
- **Events & Diagnostics:**
  ```bash
  kubectl get events --sort-by='.lastTimestamp' -A
  kubectl describe pod <pod>                          # Check Events, Conditions, State, Exit Codes
  ```
- **Common Pod Failure States:**
  - `Pending`: Scheduler cannot place the Pod (insufficient CPU/RAM, node taint, unbound PVC, or `nodeSelector` mismatch). Check `kubectl describe pod`.
  - `ImagePullBackOff` / `ErrImagePull`: Typo in image name/tag, private registry missing `imagePullSecrets`, or network issue.
  - `CrashLoopBackOff`: Container starts and exits repeatedly. Check `kubectl logs <pod> --previous` and exit code (`0` means app finished instead of staying running; `137` means `SIGKILL` / `OOMKilled`).
  - `CreateContainerConfigError`: Referenced ConfigMap or Secret is missing (unless `optional: true`).
  - `OOMKilled` (Exit Code `137`): Container exceeded its memory `limit`.
- **Interactive Debugging & Port Forwarding:**
  ```bash
  kubectl exec -it <pod> -c <container> -- sh         # Open shell inside running container
  kubectl port-forward pod/<pod> 8080:80              # Forward local port 8080 to Pod port 80
  kubectl port-forward svc/<svc> 8080:80              # Forward local port 8080 to Service port 80
  kubectl debug -it <pod> --image=busybox --target=<container>  # Attach ephemeral debug container (even for distroless images!)
  kubectl run tmp-dbg --image=busybox --rm -it --restart=Never -- sh  # Temporary standalone network/DNS debug Pod
  ```

---

### API Deprecations

Kubernetes regularly evolves API versions (`v1alpha1` -> `v1beta1` -> `v1`) and removes deprecated APIs in newer cluster versions.

- **Discovering Available APIs on the Cluster:**
  ```bash
  kubectl api-resources                               # Lists all resource types, shortnames, apiGroups, namespaced status
  kubectl api-versions                                # Lists all available group/version pairs (e.g., apps/v1)
  kubectl explain deployment.spec.strategy            # Inspect valid fields and schema for the cluster's API version
  ```
- **Common Historical Migrations to Watch For:**
  - `extensions/v1beta1` or `networking.k8s.io/v1beta1` `Ingress` -> `networking.k8s.io/v1` (changes `backend.serviceName`/`servicePort` to `backend.service.name`/`port.number`, requires `pathType`).
  - `autoscaling/v2beta1` / `v2beta2` `HorizontalPodAutoscaler` -> `autoscaling/v2`.
  - `batch/v1beta1` `CronJob` -> `batch/v1`.
  - `policy/v1beta1` `PodDisruptionBudget` -> `policy/v1`.
- **Checking & Updating Old Manifests:**
  - Always check `kubectl explain <kind>` on the exam cluster if unsure about an `apiVersion` or field structure.
  - Tools like `kubectl get <resource> -o yaml` return resources in the preferred server version, and `pluto` / `kubent` are commonly used in real clusters to audit deprecated manifests before upgrades.

---

### Custom Resource Definitions (CRDs) & Operators

- **Custom Resource Definition (CRD):** Extends the Kubernetes API by registering a new resource `Kind` (e.g., `Certificate`, `Prometheus`, `Database`). Once a CRD is created, users can manage **Custom Resources (CRs)** using standard `kubectl get/apply/describe/delete`.
  ```bash
  kubectl get crds                                    # List installed CRDs
  kubectl explain <custom-resource-kind>              # Inspect CRD schema fields
  ```
- **Controller & Operator Pattern:**
  - A CRD alone only stores data in `etcd`; it needs a **Controller** running a continuous reconciliation loop (*observe current state -> compare with desired CR spec -> take action*) to have real behavior.
  - An **Operator** = **CRD + Custom Controller** encoding domain-specific operational knowledge (backups, failover, upgrades, scaling) for stateful or complex applications (e.g., Postgres Operator, Cert-Manager).

---

## 4. Application Environment, Configuration and Security (25%)

### ConfigMaps

A ConfigMap stores **non-confidential** configuration data as key-value pairs in `etcd`.

- **Creating ConfigMaps:**
  ```bash
  kubectl create configmap app-config --from-literal=ENV=prod --from-literal=PORT=8080
  kubectl create configmap file-config --from-file=config.txt=manifests/configmap/config.txt
  ```
- **Consuming ConfigMaps in Pods** (see [`manifests/configmap/deployment.yaml`](manifests/configmap/deployment.yaml)):
  1. **As Environment Variables (`env` / `envFrom`):** Best for simple, static startup values.
  2. **As Mounted Volume Files (`volumes` + `volumeMounts`):** Best for large/structured configs (YAML, JSON, `.conf`), or configs that need dynamic updates without restarting the Pod.
  3. **As Command-Line Arguments:** Reference env vars via `$(VAR_NAME)`.
- **Dynamic Updates & `subPath` Caveat:**
  - Volume-mounted ConfigMaps **automatically update** inside the running container (after kubelet sync period) when the ConfigMap changes.
  - Environment variables loaded from ConfigMaps **do NOT update** until the Pod is recreated.
  - **Trap:** Mounting a single key using `volumeMounts[].subPath` **prevents automatic updates** of that file when the ConfigMap is modified!

---

### Secrets

Secrets store **sensitive** data (passwords, API tokens, TLS certs, registry credentials).

#### Why "Encrypted ConfigMap != Secret"
1. **Intent & Treatment:** Kubernetes treats ConfigMaps as plain config—they may appear in debug dumps or logs. Secrets signal sensitive intent.
2. **RBAC Isolation:** RBAC permissions can restrict `secrets` separately from `configmaps`. Storing credentials in ConfigMaps lets anyone with `get configmaps` read them.
3. **Special Handling:**
   - Secrets are mounted in **in-memory `tmpfs`** volumes rather than written to node disk.
   - Integrates with **Encryption at Rest** in `etcd` and external secret stores (Vault, AWS Secrets Manager).
   - Values are base64-encoded in YAML (note: base64 is encoding, **not** encryption—anyone with `get secret` access can decode it).

#### Common Secret Types & Usage
- `Opaque` (default): Arbitrary key-value pairs.
- `kubernetes.io/tls`: Stores `tls.crt` and `tls.key` for HTTPS/Ingress TLS.
- `kubernetes.io/dockerconfigjson`: Registry credentials referenced by `spec.imagePullSecrets`.

```bash
kubectl create secret generic db-secret --from-literal=USER=admin --from-literal=PASS=s3cr3t
kubectl create secret tls my-tls --cert=tls.crt --key=tls.key
kubectl create secret docker-registry regcred \
  --docker-server=registry.io --docker-username=user --docker-password=pass
```
See [`manifests/secret/deployment.yaml`](manifests/secret/deployment.yaml) for mounting a Secret via both `envFrom` and `volumeMounts` (`readOnly: true`).

---

### Resource Management

> **Golden Rule:** *Requests decide placement. Limits decide enforcement. QoS & PriorityClass decide eviction.*

#### Requests & Limits
- **Units:**
  - **CPU:** `1` = 1 vCPU/core; `1000m` = 1 core (`250m` = 0.25 core).
  - **Memory:** Binary units `Ki` (Kibibyte, 1024 B), `Mi` (Mebibyte, $1024^2$ B), `Gi` (Gibibyte, $1024^3$ B) vs decimal `K`, `M`, `G` ($1000^3$ B).
- **Requests:** Minimum resources guaranteed to a container. Used by `kube-scheduler` to find a Node with enough unallocated capacity. If omitted, treated as `0` (unless defaulted by a `LimitRange`).
- **Limits:** Maximum resources a container can consume, enforced on the Node by **kubelet** using Linux **cgroups**:
  - **CPU (compressible):** Throttled by the Linux kernel—container runs slower, **not** killed.
  - **Memory (incompressible):** Triggers Linux **OOM killer**—container is terminated and reported as `OOMKilled` (exit code `137`).

#### ResourceQuota
Limits the **aggregate** resource consumption or object counts across an entire **Namespace** (prevents one team from exhausting cluster capacity). See [`manifests/resource-quota/resource-quota.yaml`](manifests/resource-quota/resource-quota.yaml).
1. **Compute Resource Quota:** `requests.cpu`, `requests.memory`, `limits.cpu`, `limits.memory`
2. **Storage Resource Quota:** `requests.storage`, `persistentvolumeclaims`, `<storage-class>.storageclass.storage.k8s.io/requests.storage`
3. **Object Count Quota:** `pods`, `services`, `configmaps`, `secrets`, `deployments.apps`, etc.

> **Important:** When a Compute `ResourceQuota` is active in a Namespace, every Pod created in that Namespace **must** specify those corresponding `requests`/`limits` (either explicitly in the Pod spec or automatically injected by a `LimitRange`).

#### LimitRange
Enforces resource constraints and injects defaults for **individual** Pods, Containers, or PVCs within a Namespace. See [`manifests/limitrange/limitrange.yaml`](manifests/limitrange/limitrange.yaml).
- `min` / `max`: Allowed range for requests and limits.
- `default`: Automatically injected `limits` when omitted.
- `defaultRequest`: Automatically injected `requests` when omitted.
- `maxLimitRequestRatio`: Caps burst ratio (e.g., `limit / request <= 2`).
- Scope `type`: `Container`, `Pod`, or `PersistentVolumeClaim`.

#### Quality of Service (QoS) Classes
Assigned automatically by Kubernetes based on container `requests` and `limits` to determine eviction order under Node memory/disk pressure (listed highest to lowest protection):

| QoS Class | Criteria | Eviction Priority | Typical Workload |
| :--- | :--- | :--- | :--- |
| **Guaranteed** | *Every* container defines both CPU & memory `requests` and `limits`, and `request == limit` for each | Killed **last** | Databases, stateful services |
| **Burstable** | Does not qualify for Guaranteed, and at least one container defines a CPU or memory `request` or `limit` | Killed second | Web servers, APIs, microservices |
| **BestEffort** | *No* container defines any CPU or memory `request` or `limit` | Killed **first** | Dev/test pods, non-critical batch jobs |

*(Examples in [`manifests/resource-management/`](manifests/resource-management/))*

#### PriorityClass
Cluster-scoped resource (`priority.scheduling.k8s.io/v1`) assigning an integer priority value to Pods (`spec.priorityClassName`). See [`manifests/priority-class/`](manifests/priority-class/).
1. **Scheduling Order:** Higher-priority pending Pods are scheduled before lower-priority ones.
2. **Preemption:** If a high-priority Pod cannot fit, the scheduler can evict lower-priority Pods from a Node to make room.
3. **Node Pressure Eviction:** Under node resource pressure, kubelet groups Pods exceeding their requests by priority, evicting the lowest-priority group first (and breaking ties within the same priority using QoS class).

---

### Security Context & Linux Security Fundamentals

- **Linux UID/GID Conventions:**
  - `0`: `root` user/group.
  - `1–999`: Reserved for system services and daemons.
  - `1000+`: Standard range for non-privileged users.
- **SetUID Bit & Privilege Escalation:**
  - Normally, a child process runs with the Effective User ID (EUID) of the user who executed it.
  - When an executable has the `setuid` bit set (e.g., `/usr/bin/passwd`, owned by `root`), any user running it temporarily gains the **file owner's EUID (`root`)** so `passwd` can update `/etc/shadow`.
  - Setting `allowPrivilegeEscalation: false` prevents processes inside a container from gaining additional privileges via `setuid` binaries or file capabilities.
- **SecurityContext Fields (Pod-level vs Container-level):**
  - Container-level `securityContext` overrides Pod-level `securityContext` for overlapping fields. See [`manifests/security-context/security-context.yaml`](manifests/security-context/security-context.yaml).
  - `runAsUser` / `runAsGroup`: UID/GID for container processes.
  - `fsGroup` *(Pod-level only)*: Supplementary GID applied as group owner of mounted volumes so non-root processes can read/write them.
  - `runAsNonRoot: true`: Blocks container startup if the image defaults to UID `0` without a non-zero `runAsUser`.
  - `allowPrivilegeEscalation: false`: Blocks `setuid`/`setgid` elevation (also required if `readOnlyRootFilesystem` / restricted PSS is enforced).
  - `privileged: true`: Grants host-level device access and disables isolation—**avoid**.
  - `capabilities`: Fine-grained kernel privileges. Best practice: `drop: ["ALL"]`, then `add` only what is needed (e.g., `NET_BIND_SERVICE` to bind ports `< 1024` as non-root).
  - `readOnlyRootFilesystem: true`: Prevents anyone (even UID `0`) from writing to the container's root filesystem. Pair with `emptyDir` mounts for `/tmp` or writable cache dirs.
  - `seccompProfile`: Restricts system calls (`type: RuntimeDefault` or `Localhost`).
- **Pod Security Standards (PSS):** Built-in admission policies applied at the namespace level via labels (`pod-security.kubernetes.io/enforce=restricted`):
  - `Privileged`: Unrestricted.
  - `Baseline`: Blocks known privilege escalations (hostNetwork, hostPID, privileged containers, hostPath).
  - `Restricted`: Enforces hardening best practices (`runAsNonRoot: true`, `allowPrivilegeEscalation: false`, `drop: ["ALL"]`, `seccompProfile: RuntimeDefault`).

---

### ServiceAccounts & RBAC

#### ServiceAccounts
Provides an identity for **processes running inside a Pod** when communicating with the Kubernetes API. Each namespace has a `default` ServiceAccount.
- **Token Mounting:** Kubelet mounts a short-lived, projected API token into the Pod at `/var/run/secrets/kubernetes.io/serviceaccount/token` (along with `ca.crt` and `namespace`). *(Note: Since Kubernetes 1.24, long-lived token Secrets are no longer auto-created for ServiceAccounts; tokens are issued dynamically via the TokenRequest API.)*
- **Hardening:** Set `automountServiceAccountToken: false` on the Pod or ServiceAccount if the application does not need to talk to the Kubernetes API.

#### Role-Based Access Control (RBAC)
ServiceAccounts have **zero permissions** by default. Grant permissions using RBAC (`rbac.authorization.k8s.io/v1`):
- **`Role` (Namespaced) / `ClusterRole` (Cluster-wide):** Defines *what* actions (`verbs`: `get`, `list`, `watch`, `create`, `update`, `patch`, `delete`) are allowed on *which* `resources` (`pods`, `deployments`, `configmaps`) in which `apiGroups` (`""` for core API group, `"apps"`, etc.).
- **`RoleBinding` (Namespaced) / `ClusterRoleBinding` (Cluster-wide):** Binds a `Role`/`ClusterRole` to `subjects` (`User`, `Group`, or `ServiceAccount`).
  - *Tip:* A `RoleBinding` can reference a `ClusterRole` to grant those permissions **only within that single Namespace**.

```bash
kubectl create serviceaccount my-sa
kubectl create role pod-reader --verb=get,list,watch --resource=pods
kubectl create rolebinding read-pods --role=pod-reader --serviceaccount=default:my-sa
kubectl auth can-i get pods --as=system:serviceaccount:default:my-sa
```
See full YAML examples in [`manifests/service-account/`](manifests/service-account/).

---

## 5. Services and Networking (20%)

### Kubernetes Networking Fundamentals

- **Flat Pod Network:** Every Pod gets its own cluster-wide routable IP address. Any Pod can communicate with any other Pod IP directly without NAT (unless blocked by a NetworkPolicy).
- **Ephemeral Pod IPs:** Pod IPs change whenever a Pod is recreated. Never hardcode Pod IPs.
- **CoreDNS & Service Discovery:** Kubernetes DNS creates records for every Service:
  - Same namespace: `<service-name>`
  - Cross-namespace: `<service-name>.<namespace>.svc.cluster.local`
- **`kube-proxy`:** Runs on every Node and programs iptables/IPVS rules so traffic sent to a Service's virtual `ClusterIP` is load-balanced directly to ready backend `Endpoint` (Pod) IPs.

```text
Pod → web (DNS via CoreDNS)
  ↓
ClusterIP (Virtual IP)
  ↓
kube-proxy (Node-level forwarding rules)
  ↓
Ready Endpoint (Pod IP)
```

---

### Services

Services provide a **stable IP**, **stable DNS name**, and **L4 load balancing** across dynamic Pods matching `spec.selector`.

| Service Type | Scope | Purpose |
| :--- | :--- | :--- |
| **`ClusterIP`** *(default)* | Internal only | Exposes the Service on an internal virtual IP reachable inside the cluster. |
| **`NodePort`** | Internal + Node IPs | Builds on `ClusterIP` and opens a static port (`30000–32767`) on **every Node**. `kube-proxy` forwards incoming traffic to any ready backend Pod—even on a different Node. See [`manifests/service/nodeport.yaml`](manifests/service/nodeport.yaml). |
| **`LoadBalancer`** | External (Cloud) | Builds on `NodePort` + `ClusterIP` and provisions a cloud provider load balancer with a public IP. |
| **`ExternalName`** | Internal alias | Maps the Service to an external DNS name via a `CNAME` record (no selector, no proxying). |
| **Headless (`clusterIP: None`)** | Internal DNS | Returns the individual Pod IPs directly in DNS queries instead of a single virtual IP (used with StatefulSets). |

- **Port Mapping Reminder:**
  - `port`: Port the Service listens on.
  - `targetPort`: Port on the container/Pod where traffic is forwarded (can be a port number or named container port).
  - `nodePort`: External port opened on Nodes (`30000–32767`).
- **Troubleshooting Services:** Run `kubectl get endpoints <service>` (or `kubectl get endpointslice -l kubernetes.io/service-name=<service>`). If `<none>`, verify Service `selector` labels match Pod `metadata.labels` and that Pod `readinessProbe` is passing!

---

### Ingress

An **Ingress** (`networking.k8s.io/v1`) manages external HTTP/HTTPS access to internal Services through a **single entry point** (implemented by an **Ingress Controller** like NGINX, Traefik, or HAProxy), avoiding the cost of provisioning a separate `LoadBalancer` per microservice.

- **Routing Capabilities** (see [`manifests/ingress/ingress.yaml`](manifests/ingress/ingress.yaml)):
  - **Path-based routing:**
    - `http://my-cool-shop.com/` -> `frontend-service`
    - `http://my-cool-shop.com/api` -> `api-service`
    - `http://my-cool-shop.com/auth` -> `auth-service`
  - **Host-based routing:**
    - `http://my-cool-shop.com` -> `frontend-service`
    - `http://api.my-cool-shop.com` -> `api-service`
- **Important `networking.k8s.io/v1` Fields:**
  - `spec.ingressClassName`: Specifies which Ingress Controller handles this Ingress (e.g., `nginx`).
  - `pathType` *(Required)*:
    - `Exact`: Matches the URL path strictly (case-sensitive).
    - `Prefix`: Matches based on URL path prefix split by `/`.
    - `ImplementationSpecific`: Delegated to the Ingress Controller.
  - `spec.tls`: Configures HTTPS termination using a `kubernetes.io/tls` Secret (`hosts` + `secretName`).

---

### NetworkPolicies

NetworkPolicies (`networking.k8s.io/v1`) act as L3/L4 firewalls controlling Pod-to-Pod and Pod-to-CIDR traffic. Requires a CNI plugin that enforces NetworkPolicies (e.g., Calico, Cilium). See [`manifests/generic/network-policy.yaml`](manifests/generic/network-policy.yaml).

- **Core Isolation Rules:**
  - By default, **all ingress and egress traffic is allowed** if no NetworkPolicy selects a Pod.
  - A Pod becomes **isolated for Ingress** if at least one NetworkPolicy selects it with `"Ingress"` in `policyTypes`. Unlisted incoming traffic is denied (an empty `ingress: []` = deny all incoming).
  - A Pod becomes **isolated for Egress** if at least one NetworkPolicy selects it with `"Egress"` in `policyTypes`.
  - NetworkPolicies are **namespaced** and **additive** (union of all allowed rules across policies selecting the Pod).
- **Selectors:**
  - `podSelector`: Matches Pods in the **same namespace** as the NetworkPolicy.
  - `namespaceSelector`: Matches Pods in namespaces with matching labels.
  - `ipBlock`: Matches IP CIDR ranges (`cidr` + optional `except` list) for external IPs.
- **Critical CKAD Trap — `AND` vs `OR` Logic in `from` / `to`:**
  - **Two separate array items (`-`) = `OR`:**
    ```yaml
    from:
    - namespaceSelector:
        matchLabels:
          env: prod
    - podSelector:                # Separate list item -> OR!
        matchLabels:
          role: frontend
    # Allows traffic from ANY pod in namespace env=prod OR any pod with role=frontend in THIS namespace.
    ```
  - **Two fields in the same array item = `AND`:**
    ```yaml
    from:
    - namespaceSelector:
        matchLabels:
          env: prod
      podSelector:                # Same list item (no leading dash) -> AND!
        matchLabels:
          role: frontend
    # Allows traffic ONLY from pods with role=frontend INSIDE namespaces labeled env=prod.
    ```
- **DNS Egress Gotcha:** When applying a restrictive `Egress` policy, Pods won't be able to resolve Service names unless you explicitly allow UDP/TCP port `53` egress to `kube-system` (`coredns`)!

---

## 6. Essential `kubectl` Cheat Sheet

### Exam Speed & Productivity Setup
```bash
alias k=kubectl
export do="--dry-run=client -o yaml"                 # Generate YAML templates fast: k run web --image=nginx $do > pod.yaml
export now="--force --grace-period=0"                # Fast Pod deletion during practice: k delete pod web $now
kubectl explain pod.spec.containers.securityContext  # Built-in API reference (no internet needed)
```

### Pods, Imperative Creation & Debugging
```bash
# Quick Pod & multi-purpose YAML generation
kubectl run nginx --image=nginx --port=80 --labels="app=web"
kubectl run busy --image=busybox --command -- sh -c "sleep 3600"
kubectl run tmp --image=busybox --rm -it --restart=Never -- wget -qO- http://web-svc

# Inspect & Debug
kubectl get pods -A -o wide --show-labels
kubectl describe pod <pod>
kubectl logs <pod> -c <container> --previous
kubectl exec -it <pod> -- sh
kubectl top pods --sort-by=memory
```

### Deployments, Jobs, CronJobs & Scaling
```bash
kubectl create deployment web --image=nginx:1.25 --replicas=3
kubectl set image deployment/web nginx=nginx:1.26
kubectl scale deployment web --replicas=5
kubectl rollout status|history|undo|pause|resume deployment web
kubectl autoscale deployment web --min=2 --max=10 --cpu-percent=80

kubectl create job pi --image=perl -- perl -Mbignum=bpi -wle 'print bpi(2000)'
kubectl create cronjob hello --image=busybox --schedule="*/5 * * * *" -- date
kubectl create job manual-run --from=cronjob/hello   # Trigger a CronJob immediately
```

### ConfigMaps, Secrets, Services & Labels
```bash
kubectl create configmap cfg --from-literal=key=val --from-file=app.conf
kubectl create secret generic sec --from-literal=pass=secret
kubectl expose deployment web --port=80 --target-port=80 --type=NodePort
kubectl label pod <pod> env=prod --overwrite
kubectl label pod <pod> env-                         # Remove label 'env'
kubectl get pods -l 'env in (prod,qa)'
```

### Formatting & JSONPath / Custom Columns
```bash
kubectl get pods -o custom-columns=NAME:.metadata.name,IMAGE:.spec.containers[*].image,NODE:.spec.nodeName
kubectl get nodes -o jsonpath='{.items[*].status.addresses[?(@.type=="InternalIP")].address}{"\n"}'
kubectl get events --sort-by='.lastTimestamp'
```

---

## 7. Repository Manifests & Practice Resources

### Local Examples in This Repository
- **Probes & Multi-Stage Dockerfile:** [`probes/`](probes/) (`Dockerfile`, `main.go`, `deployment.yaml`)
- **Workloads & Jobs:** [`manifests/generic/deployment.yaml`](manifests/generic/deployment.yaml), [`manifests/jobs/job.yaml`](manifests/jobs/job.yaml), [`manifests/jobs/cronjob.yaml`](manifests/jobs/cronjob.yaml)
- **Autoscaling:** [`manifests/autoscaler/horizontal.yaml`](manifests/autoscaler/horizontal.yaml)
- **ConfigMaps & Secrets:** [`manifests/configmap/`](manifests/configmap/), [`manifests/secret/deployment.yaml`](manifests/secret/deployment.yaml)
- **Storage (PVC):** [`manifests/pvc/`](manifests/pvc/)
- **Resource Management, Quotas & Priority:** [`manifests/resource-management/`](manifests/resource-management/), [`manifests/resource-quota/resource-quota.yaml`](manifests/resource-quota/resource-quota.yaml), [`manifests/limitrange/limitrange.yaml`](manifests/limitrange/limitrange.yaml), [`manifests/priority-class/`](manifests/priority-class/)
- **Security & RBAC:** [`manifests/security-context/security-context.yaml`](manifests/security-context/security-context.yaml), [`manifests/service-account/`](manifests/service-account/)
- **Networking, Services & Ingress:** [`manifests/service/nodeport.yaml`](manifests/service/nodeport.yaml), [`manifests/ingress/`](manifests/ingress/), [`manifests/generic/network-policy.yaml`](manifests/generic/network-policy.yaml)

### External Practice & Docs
- Practice questions: https://medium.com/bb-tutorials-and-thoughts/practice-enough-with-these-questions-for-the-ckad-exam-2f42d1228552
- Official Kubernetes Documentation: https://kubernetes.io/docs/
- Official CKAD Curriculum (CNCF): https://github.com/cncf/curriculum

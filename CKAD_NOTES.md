# CKAD Complete Reference & Study Handbook

A comprehensive, visual, and exam-focused reference manual for the **Certified Kubernetes Application Developer (CKAD)** exam. This handbook reorganizes, corrects, and expands your original study notes from [`README.md`](README.md) across all 5 official CKAD domains, complete with architecture diagrams, line-by-line annotated YAML templates, exam traps, and an audit of the manifests in [`manifests/`](manifests/) and [`probes/`](probes/).

---

## Table of Contents

- [CKAD Exam Domains & High-Level Mental Models](#ckad-exam-domains--high-level-mental-models)
- [1. Application Design and Build (20%)](#1-application-design-and-build-20)
  - [1.1 Building & Managing Container Images](#11-building--managing-container-images)
  - [1.2 Pods, Network Namespace & Pod Lifecycle](#12-pods-network-namespace--pod-lifecycle)
  - [1.3 Multi-Container Pod Patterns (Sidecar, Ambassador, Adapter, Init & Native Sidecars)](#13-multi-container-pod-patterns-sidecar-ambassador-adapter-init--native-sidecars)
  - [1.4 Workload Controllers (ReplicaSet, Deployment, DaemonSet, StatefulSet, Job, CronJob)](#14-workload-controllers-replicaset-deployment-daemonset-statefulset-job-cronjob)
  - [1.5 Storage & Volumes (`emptyDir`, `hostPath`, `projected`, `downwardAPI`, `PV`, `PVC`, `StorageClass`)](#15-storage--volumes-emptydir-hostpath-projected-downwardapi-pv-pvc-storageclass)
- [2. Application Deployment (20%)](#2-application-deployment-20)
  - [2.1 Deployment Strategies (`RollingUpdate`, `Recreate`, `Blue-Green`, `Canary`)](#21-deployment-strategies-rollingupdate-recreate-blue-green-canary)
  - [2.2 Rollouts, Revisions & Rollbacks](#22-rollouts-revisions--rollbacks)
  - [2.3 Pod Scheduling & Placement (Labels, Selectors, Affinity, Taints & Tolerations)](#23-pod-scheduling--placement-labels-selectors-affinity-taints--tolerations)
  - [2.4 Horizontal Pod Autoscaling (`autoscaling/v2`)](#24-horizontal-pod-autoscaling-autoscalingv2)
  - [2.5 Helm Package Manager (v3)](#25-helm-package-manager-v3)
- [3. Application Observability and Maintenance (15%)](#3-application-observability-and-maintenance-15)
  - [3.1 Container Probes (`startupProbe`, `readinessProbe`, `livenessProbe`)](#31-container-probes-startupprobe-readinessprobe-livenessprobe)
  - [3.2 Monitoring, Logging & Interactive Debugging](#32-monitoring-logging--interactive-debugging)
  - [3.3 API Deprecations & Version Migrations](#33-api-deprecations--version-migrations)
  - [3.4 Custom Resource Definitions (CRDs) & The Operator Pattern](#34-custom-resource-definitions-crds--the-operator-pattern)
- [4. Application Environment, Configuration and Security (25%)](#4-application-environment-configuration-and-security-25)
  - [4.1 ConfigMaps (Env Vars vs Volume Mounts vs `subPath` Trap)](#41-configmaps-env-vars-vs-volume-mounts-vs-subpath-trap)
  - [4.2 Secrets (Why "Encrypted ConfigMap != Secret", Types & Consumption)](#42-secrets-why-encrypted-configmap--secret-types--consumption)
  - [4.3 Resource Management (Requests, Limits, `ResourceQuota`, `LimitRange`, QoS & `PriorityClass`)](#43-resource-management-requests-limits-resourcequota-limitrange-qos--priorityclass)
  - [4.4 Security Context, Linux UID/GID, SetUID, Capabilities & Pod Security Standards](#44-security-context-linux-uidgid-setuid-capabilities--pod-security-standards)
  - [4.5 ServiceAccounts & Role-Based Access Control (RBAC)](#45-serviceaccounts--role-based-access-control-rbac)
- [5. Services and Networking (20%)](#5-services-and-networking-20)
  - [5.1 Kubernetes Networking Fundamentals & CoreDNS](#51-kubernetes-networking-fundamentals--coredns)
  - [5.2 Services (`ClusterIP`, `NodePort`, `LoadBalancer`, `ExternalName`, Headless)](#52-services-clusterip-nodeport-loadbalancer-externalname-headless)
  - [5.3 Ingress & TLS Termination (`networking.k8s.io/v1`)](#53-ingress--tls-termination-networkingk8siov1)
  - [5.4 NetworkPolicies (Isolation Rules, `AND` vs `OR` Selector Trap & DNS Egress)](#54-networkpolicies-isolation-rules-and-vs-or-selector-trap--dns-egress)
- [6. Complete Exam-Day `kubectl` & `helm` Cheat Sheet](#6-complete-exam-day-kubectl--helm-cheat-sheet)
- [7. Repository Manifests Index & Bug Audit of `manifests/`](#7-repository-manifests-index--bug-audit-of-manifests)

---

## CKAD Exam Domains & High-Level Mental Models

| Domain | Weight | Core Objects & Skills Tested |
| :--- | :---: | :--- |
| **1. Application Design and Build** | **20%** | Container images (`Dockerfile`), `Pod`, Multi-container patterns (`initContainers`, sidecars), `Job`, `CronJob`, `PV`, `PVC`, `StorageClass` |
| **2. Application Deployment** | **20%** | `Deployment`, `RollingUpdate`/`Recreate`/Blue-Green/Canary, Affinity, Taints/Tolerations, `HorizontalPodAutoscaler`, `Helm` |
| **3. Application Observability and Maintenance** | **15%** | `startupProbe`/`readinessProbe`/`livenessProbe`, `logs`, `exec`, `debug`, `top`, `events`, API deprecations, `CustomResourceDefinition` |
| **4. Application Environment, Configuration and Security** | **25%** | `ConfigMap`, `Secret`, `resources` (requests/limits), `ResourceQuota`, `LimitRange`, QoS, `PriorityClass`, `securityContext`, `ServiceAccount`, `Role`/`ClusterRole`/`RoleBinding` |
| **5. Services and Networking** | **20%** | `Service` (`ClusterIP`, `NodePort`, `LoadBalancer`, `ExternalName`, Headless), `Endpoints`, `Ingress`, `NetworkPolicy` |

### Six Core Rules to Keep in Your Head
1. **One Pod = One Network Namespace = One IP:** All containers inside a Pod share the same IP address, the same `localhost` loopback interface, and the same port space (two containers in the same Pod cannot bind to the same port!).
2. **Requests decide placement. Limits decide enforcement. PriorityClass & QoS decide eviction:** The scheduler ignores `limits` and only checks `requests`; the node kubelet enforces `limits` (throttling CPU via Linux cgroups, killing memory hogs via Linux OOM Killer -> `OOMKilled`).
3. **Startup protects boot, Readiness protects traffic, Liveness protects uptime:** A failing `readinessProbe` cuts Service traffic (`Endpoints`) without restarting the container; a failing `livenessProbe` restarts the container; a `startupProbe` disables both until slow initialization finishes.
4. **Services are virtual IPs programmed by `kube-proxy`:** Pod IPs are ephemeral and flat (every Pod can reach every other Pod IP without NAT unless blocked by a NetworkPolicy). Services provide a stable virtual IP (`ClusterIP`) and DNS name that `kube-proxy` translates into ready backend Pod IPs (`Endpoints`).
5. **Env vars are static; Volume-mounted ConfigMaps/Secrets are dynamic (except `subPath`!):** Mounted ConfigMaps and Secrets update automatically inside a running container (~60s sync), whereas environment variables and `subPath` mounts never update until the Pod is recreated.
6. **NetworkPolicies default to allow-all, become default-deny on selection, and are additive:** If no policy selects a Pod, all traffic is allowed. As soon as *any* policy selects a Pod for `Ingress` (or `Egress`), all unlisted traffic in that direction is blocked.

---

## 1. Application Design and Build (20%)

### 1.1 Building & Managing Container Images

#### Multi-Stage Dockerfiles
Multi-stage builds compile your application in a build-time image (with compilers/SDKs) and copy **only** the compiled artifact into a tiny runtime base image (`alpine`, `ubuntu`, `distroless`). This minimizes image size, pull time, and security vulnerabilities.

See [`probes/Dockerfile`](probes/Dockerfile) in this repository:
```dockerfile
# Stage 1: Builder image with Go toolchain
FROM golang:1.25 AS builder
WORKDIR /app
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o ./bin/probes .

# Stage 2: Lightweight runtime image
FROM ubuntu:22.04
RUN apt-get update && \
    apt-get install -y ca-certificates curl && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/bin/probes /opt/app/probes
EXPOSE 3000
ENTRYPOINT ["/opt/app/probes"]
```

#### Essential `docker` / `podman` CLI Commands
*(Note: On the CKAD exam, `docker` and/or `podman` may be installed—their CLI flags for building, tagging, pushing, and saving images are identical.)*
```bash
docker build -t my-app:1.0 -f Dockerfile .               # Build image from current directory
docker build --build-arg APP_VERSION=1.0 -t my-app:1.0 . # Pass build-time argument
docker images                                            # List local images
docker tag my-app:1.0 registry.example.com/my-app:1.0    # Add a new tag to an existing image
docker push registry.example.com/my-app:1.0              # Push image to remote registry
docker save -o /tmp/my-app.tar my-app:1.0                # Export image to a tarball archive
docker load -i /tmp/my-app.tar                           # Load image from a tarball archive
docker rmi my-app:1.0                                    # Remove local image
```

#### How Dockerfile `ENTRYPOINT` & `CMD` Map to Kubernetes `command` & `args`

```text
+-----------------------------------+------------------------------------+
| Dockerfile Field                  | Kubernetes Container Field         |
+-----------------------------------+------------------------------------+
| ENTRYPOINT ["/bin/app"]           | command: ["/bin/app"]              |
| CMD ["--port", "8080"]            | args: ["--port", "8080"]           |
+-----------------------------------+------------------------------------+
```

| What You Define in Pod YAML | Resulting Command Executed in Container |
| :--- | :--- |
| Neither `command` nor `args` | `/bin/app --port 8080` (uses image defaults) |
| Only `args: ["--port", "9090"]` | `/bin/app --port 9090` (keeps `ENTRYPOINT`, overrides `CMD`) |
| Only `command: ["/bin/custom"]` | `/bin/custom` (overrides **both** `ENTRYPOINT` and `CMD`!) |
| Both `command: ["/bin/sh", "-c"]` and `args: ["date"]` | `/bin/sh -c date` (overrides both) |

> **Imperative CLI Trap (`kubectl run`):**
> - `kubectl run p --image=busybox -- sh -c "sleep 3600"` (without `--command`) sets **`args:`**!
> - `kubectl run p --image=busybox --command -- sh -c "sleep 3600"` (with `--command`) sets **`command:`**!

#### `imagePullPolicy` Rules
- `IfNotPresent`: Pulls the image only if it is not already present on the Node (default when a specific tag like `nginx:1.25` is specified).
- `Always`: Queries the container registry on every Pod start (default when `:latest` is used or tag is omitted).
- `Never`: Only uses local images on the Node; never attempts a network pull.

---

### 1.2 Pods, Network Namespace & Pod Lifecycle

- **Definition:** A **Pod** is the smallest deployable unit in Kubernetes. It encapsulates one or more tightly coupled containers that share:
  1. **Network Namespace (`IPC` / `NET`):** One IP address per Pod, one loopback interface (`localhost`), and one shared port space.
  2. **Storage Volumes:** Any volume defined in `spec.volumes` can be mounted into multiple containers inside the Pod at different `mountPath`s.

#### Pod `restartPolicy`
Configured at `pod.spec.restartPolicy` (applies to all containers in the Pod):
- `Always` *(Default)*: Restart container whenever it exits (regardless of exit code `0` or non-zero). **Required for `Deployment`, `ReplicaSet`, `DaemonSet`, and `StatefulSet`.**
- `OnFailure`: Restart container **only** if it exits with a non-zero exit code. Valid for standalone Pods and `Job`/`CronJob`.
- `Never`: Never restart terminated containers. Valid for standalone Pods and `Job`/`CronJob`.

#### Pod Lifecycle Hooks & Graceful Termination Sequence
When a Pod is deleted, Kubernetes runs the **graceful shutdown sequence** in parallel with endpoint removal:

```text
[kubectl delete pod]
   │
   ├──► Control Plane marks Pod "Terminating" -> Removed from Service Endpoints
   │
   └──► Kubelet starts terminationGracePeriodSeconds countdown (default 30s):
          1. Executes `lifecycle.preStop` hook (if defined) synchronously.
          2. Sends `SIGTERM` to PID 1 inside the container.
          3. Waits for process to exit cleanly.
          4. If timer expires before exit -> Sends `SIGKILL` (forced termination).
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: lifecycle-pod
spec:
  terminationGracePeriodSeconds: 20
  containers:
  - name: web
    image: nginx:1.25-alpine
    lifecycle:
      postStart:                       # Runs asynchronously right after container creation
        exec:
          command: ["sh", "-c", "echo 'Booted' > /tmp/boot.log"]
      preStop:                         # Runs synchronously BEFORE SIGTERM is sent
        exec:
          command: ["sh", "-c", "nginx -s quit; sleep 3"]
```

---

### 1.3 Multi-Container Pod Patterns (Sidecar, Ambassador, Adapter, Init & Native Sidecars)

```text
+-----------------------------------------------------------------------------+
|                              POD EXECUTION FLOW                             |
|                                                                             |
|  [initContainers: 1..N] (Run sequentially; each must exit 0 before next)    |
|            │                                                                |
|            ├──► Normal InitContainer: Runs setup task -> Exits 0            |
|            └──► Native Sidecar (`restartPolicy: Always`): Starts & stays up |
|                                    │                                        |
|                                    ▼                                        |
|  [containers: 1..M] (Start in parallel once all initContainers succeed)     |
|            ├──► Main App Container                                          |
|            ├──► Sidecar (e.g., tails logs from shared emptyDir)             |
|            ├──► Ambassador (proxies localhost:5432 -> external DB shards)   |
|            └──► Adapter (converts XML on localhost:8080 -> Prometheus:9100) |
+-----------------------------------------------------------------------------+
```

1. **Sidecar Pattern:** Runs alongside the main application container to provide supporting functionality (logging, file syncing, security agents) without bloating the main application's code.
   - *Problem:* Your web server writes logs to `/var/log/nginx/access.log`. You want to stream or ship those logs to a central system without modifying the web server image.
   - *Solution:* Mount a shared `emptyDir` volume at `/var/log/nginx` in both the web container and a sidecar container running `tail -n+1 -F /var/log/nginx/access.log`.
2. **Ambassador Pattern:** Acts as a local network proxy (`localhost`) handling outbound traffic from the main container (service discovery, database sharding, TLS encryption, circuit breaking).
   - *Problem:* Your app connects to a sharded database, and the routing logic for choosing the right shard is complex or changes per environment.
   - *Solution:* Add an ambassador proxy container (e.g., Envoy or custom proxy) listening on `localhost:5432`. The main container connects to `localhost:5432`, and the ambassador inspects requests and forwards them to the proper remote shard.
3. **Adapter Pattern:** Standardizes or normalizes the output/monitoring format of the main container to match an external interface standard.
   - *Problem:* A legacy application exposes monitoring stats in an outdated XML format on `localhost:8080`, while your cluster's Prometheus server requires Prometheus text format.
   - *Solution:* Add an adapter container that scrapes `http://localhost:8080/xml`, transforms the data into Prometheus metrics, and exposes them on port `9100`.
4. **InitContainer Pattern (`spec.initContainers`):** Runs to completion *before* any main containers start. If an InitContainer fails, kubelet restarts it according to the Pod's `restartPolicy` until it succeeds (`exit 0`).
   - *Problem:* Your main app crashes immediately on boot if the database service isn't reachable yet or if `/config/config.json` hasn't been fetched from a remote config server.
   - *Solution:* Add an `initContainer` that polls the database (`until nslookup db-svc; do sleep 2; done`) or fetches `/config/config.json` onto a shared `emptyDir` volume and exits `0`.
5. **Native Sidecar Containers (Kubernetes 1.28+ / GA 1.33):**
   - Defined inside `spec.initContainers` with **`restartPolicy: Always`**.
   - **Why it matters on CKAD:** If you attach a traditional sidecar under `spec.containers` inside a **`Job`**, the Job's Pod never completes (`1/2 Ready`) because the sidecar keeps running after the main worker exits! By placing the sidecar under `spec.initContainers` with `restartPolicy: Always`, Kubernetes starts the sidecar *before* the main container and **automatically terminates the sidecar** as soon as the main container exits `0`.

#### Complete Multi-Container Pod Template (Init + Main + Sidecar + Shared `emptyDir`)
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: web-with-sidecar
spec:
  initContainers:
  - name: config-fetcher
    image: busybox:1.36
    command: ["sh", "-c", "echo '<h1>Loaded by InitContainer</h1>' > /work-dir/index.html"]
    volumeMounts:
    - name: html-vol
      mountPath: /work-dir
  containers:
  - name: main-web
    image: nginx:1.25-alpine
    ports:
    - containerPort: 80
    volumeMounts:
    - name: html-vol
      mountPath: /usr/share/nginx/html
    - name: log-vol
      mountPath: /var/log/nginx
  - name: log-sidecar
    image: busybox:1.36
    command: ["sh", "-c", "touch /logs/access.log && tail -n+1 -F /logs/access.log"]
    volumeMounts:
    - name: log-vol
      mountPath: /logs
  volumes:
  - name: html-vol
    emptyDir: {}
  - name: log-vol
    emptyDir: {}
```

---

### 1.4 Workload Controllers (ReplicaSet, Deployment, DaemonSet, StatefulSet, Job, CronJob)

#### Workload Controller Comparison Matrix

| Controller | Purpose | Valid `restartPolicy` | Key Spec Fields |
| :--- | :--- | :---: | :--- |
| **`ReplicaSet`** | Maintains a stable count of identical Pod replicas | `Always` | `replicas`, `selector`, `template` (No rolling updates/rollbacks) |
| **`Deployment`** | Manages ReplicaSets for declarative rolling updates & rollbacks | `Always` | `replicas`, `strategy` (`RollingUpdate`/`Recreate`), `revisionHistoryLimit` |
| **`DaemonSet`** | Runs 1 copy of a Pod on every Node (or matching Node subset) | `Always` | `selector`, `template` (No `replicas` field!) |
| **`StatefulSet`** | Stateful apps needing sticky ordinal names, DNS, and per-Pod PVCs | `Always` | `serviceName` (Headless Service), `replicas`, `volumeClaimTemplates` |
| **`Job`** | Batch tasks that run to completion ($N$ Pods exit `0`) | `OnFailure` or `Never` | `completions`, `parallelism`, `backoffLimit`, `activeDeadlineSeconds`, `ttlSecondsAfterFinished` |
| **`CronJob`** | Spawns `Job` objects on a recurring cron schedule | `OnFailure` or `Never` | `schedule`, `concurrencyPolicy`, `startingDeadlineSeconds`, `jobTemplate` |

---

#### StatefulSet + Headless Service In-Depth
Unlike a `Deployment` where replicas are interchangeable (`web-6d8f9-x2p9k`), a **`StatefulSet`** gives each Pod a permanent ordinal identity (`0` to $N-1$):
1. **Ordered Naming & Lifecycle:** Pods are named `<statefulset>-0`, `<statefulset>-1`, `<statefulset>-2`. By default (`podManagementPolicy: OrderedReady`), Pod `1` is not created until Pod `0` is `Running` and `Ready`. Scaling down terminates in reverse order (`2 -> 1 -> 0`).
2. **Predictable Per-Pod DNS via Headless Service (`clusterIP: None`):**
   Each Pod gets its own A/AAAA record in CoreDNS:
   ```text
   <pod-name>.<headless-service-name>.<namespace>.svc.cluster.local
   Example: db-0.db-headless.prod.svc.cluster.local
   ```
3. **Sticky Per-Pod Storage (`volumeClaimTemplates`):**
   Automatically creates a dedicated PVC for each ordinal (`<pvc-template-name>-<statefulset-name>-<ordinal>`, e.g., `data-db-0`, `data-db-1`). Deleting a StatefulSet or scaling it down **does not delete its PVCs** (protecting database state).

```yaml
apiVersion: v1
kind: Service
metadata:
  name: db-headless
spec:
  clusterIP: None                  # Defines a Headless Service (returns Pod IPs directly)
  selector:
    app: postgres
  ports:
  - port: 5432
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: db
spec:
  serviceName: db-headless         # REQUIRED: Must match Headless Service name
  replicas: 2
  selector:
    matchLabels:
      app: postgres
  template:
    metadata:
      labels:
        app: postgres
    spec:
      containers:
      - name: postgres
        image: postgres:16-alpine
        ports:
        - containerPort: 5432
        env:
        - name: POSTGRES_PASSWORD
          value: "secret"
        volumeMounts:
        - name: data
          mountPath: /var/lib/postgresql/data
  volumeClaimTemplates:
  - metadata:
      name: data                   # Creates PVCs named 'data-db-0' and 'data-db-1'
    spec:
      accessModes: ["ReadWriteOnce"]
      resources:
        requests:
          storage: 1Gi
```

---

#### DaemonSet Tip (Fastest Exam Generation Shortcut)
`kubectl create` has **no** `kubectl create daemonset` command. To generate a DaemonSet in 10 seconds on the exam:
1. Run `kubectl create deploy my-ds --image=busybox:1.36 $do > ds.yaml`
2. Change `kind: Deployment` to `kind: DaemonSet`, and delete `replicas: 1` and `strategy: {}`!

---

#### Job & CronJob Configuration Reference
See [`manifests/jobs/job.yaml`](manifests/jobs/job.yaml) and [`manifests/jobs/cronjob.yaml`](manifests/jobs/cronjob.yaml).

- **Job Parameters (`Job.spec`):**
  1. `completions` *(default `1`)*: Total number of Pods that must terminate with exit code `0` for the Job to be marked complete.
  2. `parallelism` *(default `1`)*: Maximum number of Pods allowed to run at the same time.
  3. `backoffLimit` *(default `6`)*: Maximum number of retries if a Pod fails (exits non-zero). Once exceeded, the Job is marked `Failed` (`BackoffLimitExceeded`).
  4. `activeDeadlineSeconds`: Hard time limit for the entire Job. If the Job is still running after this many seconds, Kubernetes terminates all running Pods and marks the Job `Failed` (`DeadlineExceeded`).
  5. `ttlSecondsAfterFinished`: Automatically deletes the Job (and its Pods) $N$ seconds after the Job finishes (either `Complete` or `Failed`).

- **CronJob Parameters (`CronJob.spec`):**
  1. `schedule` *(Required)*:
     ```text
     # ┌───────────── minute (0 - 59)        e.g., */5 = every 5 minutes
     # │ ┌───────────── hour (0 - 23)        e.g., 0   = midnight
     # │ │ ┌───────────── day of month (1 - 31)
     # │ │ │ ┌───────────── month (1 - 12)
     # │ │ │ │ ┌───────────── day of week (0 - 6) (0 = Sunday, 6 = Saturday)
     # * * * * *
     ```
  2. `concurrencyPolicy`:
     - `Allow` *(default)*: Allows multiple overlapping Jobs to run at the same time.
     - `Forbid`: Skips starting the new Job if the previous Job is still running.
     - `Replace`: Cancels the currently running Job and starts a new one.
  3. `startingDeadlineSeconds`: Maximum seconds a Job can be delayed past its scheduled time (e.g., if the controller was down) before that missed execution is skipped.
  4. `successfulJobsHistoryLimit` *(default `3`)* / `failedJobsHistoryLimit` *(default `1`)*: Number of finished Job objects to retain for log inspection.
  5. `suspend` *(default `false`)*: If `true`, pauses all subsequent scheduled runs without deleting the CronJob.
  6. `jobTemplate`: Contains `spec` (which is a full `Job.spec`, including `template.spec` for the Pod!).

```bash
# Trigger a manual one-off Job from an existing CronJob to test it immediately:
kubectl create job test-cron-run --from=cronjob/my-cronjob
```

---

### 1.5 Storage & Volumes (`emptyDir`, `hostPath`, `projected`, `downwardAPI`, `PV`, `PVC`, `StorageClass`)

#### Ephemeral & Metadata Volume Types
- **`emptyDir`:** Created when a Pod is assigned to a Node; deleted permanently when the Pod is removed from that Node.
  - Shared across containers in the same Pod (each container can mount it at a different path).
  - Set `emptyDir: { medium: Memory, sizeLimit: 64Mi }` to use a RAM-backed `tmpfs` filesystem.
- **`hostPath`:** Mounts a directory or file from the host Node's filesystem into the Pod (`type: DirectoryOrCreate`, `File`, `Socket`, etc.).
  - *Discouraged in production:* Exposes host filesystem security risks, lacks multi-node portability, and bypasses cluster storage quotas.
- **`downwardAPI`:** Injects Pod metadata (`metadata.name`, `metadata.namespace`, `metadata.labels`, `metadata.annotations`, `spec.nodeName`, `status.podIP`) or container resource limits/requests (`limits.cpu`, `requests.memory`) as files or environment variables.
- **`projected`:** Merges multiple volume sources (`secret`, `configMap`, `downwardAPI`, `serviceAccountToken`) into a single unified directory.

#### Persistent Storage Architecture (`PV` -> `PVC` -> `Pod`)
See [`manifests/pvc/pvc.yaml`](manifests/pvc/pvc.yaml) and [`manifests/pvc/pod-with-pvc.yaml`](manifests/pvc/pod-with-pvc.yaml).

```text
+-----------------------------------+          +-----------------------------------+
| PersistentVolume (PV)             |          | PersistentVolumeClaim (PVC)       |
| - Cluster-scoped storage asset    | ◄──────► | - Namespaced storage request      |
| - Provisioned by Admin or         |  Binds   | - Matches PV on:                  |
|   dynamically by StorageClass     |   1:1    |   1. storageClassName             |
| - Lifecycle phases:               |          |   2. accessModes                  |
|   Available -> Bound -> Released  |          |   3. PV capacity >= PVC request   |
+-----------------------------------+          +-----------------+-----------------+
                                                                 │
                                                                 │ claimName
                                                                 ▼
                                               +-----------------------------------+
                                               | Pod (`volumes[].persistentVolumeClaim`)
                                               +-----------------------------------+
```

- **Access Modes:**
  - `ReadWriteOnce` (`RWO`): Mounted read-write by a **single Node** (multiple Pods on that **same Node** can mount it simultaneously; Pods on other Nodes cannot!).
  - `ReadOnlyMany` (`ROX`): Mounted read-only by **multiple Nodes**.
  - `ReadWriteMany` (`RWX`): Mounted read-write by **multiple Nodes** simultaneously (e.g., NFS, CephFS, AWS EFS).
  - `ReadWriteOncePod` (`RWOP`): Mounted read-write by a **single Pod** across the entire cluster.

- **Reclaim Policies (`persistentVolumeReclaimPolicy`):**
  - `Delete` *(default for dynamically provisioned PVs)*: Deleting the PVC automatically deletes both the PV object and the backing storage volume.
  - `Retain` *(default for manually created PVs)*: Deleting the PVC leaves the PV and its data intact, moving the PV status to **`Released`**.
    - **Critical Recovery Trick:** A `Released` PV **cannot** bind to a new PVC because it still contains `spec.claimRef` pointing to the old deleted PVC. To make a `Released` PV `Available` again:
      ```bash
      kubectl patch pv <pv-name> --type=json -p='[{"op": "remove", "path": "/spec/claimRef"}]'
      ```

- **`StorageClass` & `volumeBindingMode`:**
  - `Immediate`: Binds/provisions the PV as soon as the PVC is created (even before any Pod uses it).
  - `WaitForFirstConsumer`: Delays binding/provisioning the PV until a Pod using the PVC is scheduled, ensuring the storage volume is created in the exact same cloud availability zone or node as the Pod.

#### Complete PV + PVC + Pod Manifest Template
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: app-pv
spec:
  capacity:
    storage: 2Gi
  accessModes:
  - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  storageClassName: fast-local
  hostPath:
    path: /mnt/data/app
    type: DirectoryOrCreate
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: app-pvc
spec:
  accessModes:
  - ReadWriteOnce
  storageClassName: fast-local       # Must match PV's storageClassName! (Set "" for no class)
  resources:
    requests:
      storage: 1Gi                   # Must be <= PV capacity (2Gi)
---
apiVersion: v1
kind: Pod
metadata:
  name: app-with-storage
spec:
  containers:
  - name: app
    image: nginx:1.25-alpine
    volumeMounts:
    - name: data-vol
      mountPath: /usr/share/nginx/html
      subPath: html                  # Optional: mount a subdirectory inside the volume
  volumes:
  - name: data-vol
    persistentVolumeClaim:
      claimName: app-pvc
```


---

## 2. Application Deployment (20%)

### 2.1 Deployment Strategies (`RollingUpdate`, `Recreate`, `Blue-Green`, `Canary`)

| Strategy | How It Works | Zero Downtime? | Instant Rollback? | Best For |
| :--- | :--- | :---: | :---: | :--- |
| **`RollingUpdate`** *(default)* | Creates a new ReplicaSet and gradually scales it up while scaling the old ReplicaSet down | **Yes** | Fast (`rollout undo`) | Standard stateless web/API workloads |
| **`Recreate`** | Terminates **all** old Pods completely before creating any new Pods | **No** (brief outage) | Moderate | Legacy apps that cannot run two versions simultaneously or share a single `ReadWriteOnce` lock |
| **`Blue-Green`** | Runs two full environments (`blue` = `v1`, `green` = `v2`); switches the `Service` selector from `blue` to `green` | **Yes** | **Instant** (1-line selector patch) | Major releases requiring full pre-cutover testing and instant rollback |
| **`Canary`** | Exposes a small fraction of live traffic to `v2` alongside `v1` before scaling `v2` to 100% | **Yes** | **Instant** (scale canary to 0) | Testing risky changes on real production traffic |

#### `RollingUpdate` Capacity Math (`maxSurge` & `maxUnavailable`)
- **Rule:** Deployments prioritize **availability** over exact replica count during updates—temporary over-provisioning or under-provisioning is normal.
- Given `replicas: 4`, `maxSurge: 25%` (`1` Pod), and `maxUnavailable: 25%` (`1` Pod):
  - **Max total Pods during rollout:** $4 + 1 = \mathbf{5\text{ Pods}}$
  - **Min available Ready Pods during rollout:** $4 - 1 = \mathbf{3\text{ Pods}}$
- Setting `maxUnavailable: 0` guarantees capacity **never drops below `replicas`** during an update.

#### Visual Architecture: Blue-Green vs Canary

```text
BLUE-GREEN (Switch 100% Traffic by Patching Service Selector)
=============================================================
  [Service: shop-svc]
  selector: { app: shop, version: green }   <-- Flipped from 'blue' to 'green'!
                 │
    +------------+------------+
    │ 0%                      │ 100%
    ▼                         ▼
 [Deployment: shop-blue]   [Deployment: shop-green]
 labels: app=shop          labels: app=shop
         version=blue              version=green
 (3 Ready Pods - Standby)  (3 Ready Pods - Live)


CANARY (Split Traffic Proportionally via Shared Service Label)
==============================================================
  [Service: shop-svc]
  selector: { app: shop }                   <-- Selects ONLY `app: shop`!
                 │
    +------------+------------+
    │ 80% (4 of 5 Pods)       │ 20% (1 of 5 Pods)
    ▼                         ▼
 [Deployment: shop-v1]     [Deployment: shop-v2-canary]
 selector: app=shop,v=v1   selector: app=shop,v=v2
 (4 Ready Pods)            (1 Ready Pod)
```

> **Critical Canary Selector Rule:**
> In a Canary setup, the **Service** selects only `app: shop`, while each **Deployment's `spec.selector.matchLabels`** MUST include both `app: shop` AND its version label (`v: v1` vs `v: v2`) so the two Deployments do not adopt each other's Pods!

---

### 2.2 Rollouts, Revisions & Rollbacks

- **What Creates a New Revision?**
  A **Revision** is a snapshot of a Deployment's `spec.template` (`PodTemplateSpec`), stored inside an underlying `ReplicaSet`.
  - **Triggers a new Revision:** Any change inside `spec.template` (container image, env vars, command/args, resource requests/limits, volumeMounts, pod labels).
  - **Does NOT trigger a new Revision:** Scaling `spec.replicas`, changing `spec.strategy`, or editing `revisionHistoryLimit`!

```bash
# Inspect status & history
kubectl rollout status deployment/webapp
kubectl rollout history deployment/webapp
kubectl rollout history deployment/webapp --revision=2

# Record a descriptive change-cause in revision history
kubectl annotate deployment/webapp kubernetes.io/change-cause="Upgraded nginx to 1.26-alpine"

# Pause a deployment to batch multiple changes into a SINGLE revision
kubectl rollout pause deployment/webapp
kubectl set image deployment/webapp nginx=nginx:1.26-alpine
kubectl set resources deployment/webapp -c=nginx --limits=cpu=200m,memory=256Mi
kubectl rollout resume deployment/webapp

# Rollback to previous revision or a specific revision number
kubectl rollout undo deployment/webapp
kubectl rollout undo deployment/webapp --to-revision=1

# Trigger a rolling restart of all Pods (e.g., after updating an env-var ConfigMap)
kubectl rollout restart deployment/webapp
```

---

### 2.3 Pod Scheduling & Placement (Labels, Selectors, Affinity, Taints & Tolerations)

#### 1. `nodeSelector` (Simplest Exact Match)
Schedules the Pod only onto Nodes possessing all specified labels:
```yaml
spec:
  nodeSelector:
    disktype: ssd
    gpu: "true"
```

#### 2. `nodeAffinity` & `podAffinity` / `podAntiAffinity`
- **`requiredDuringSchedulingIgnoredDuringExecution`:** Hard requirement. Pod remains `Pending` if no Node satisfies the expression.
- **`preferredDuringSchedulingIgnoredDuringExecution`:** Soft preference (`weight: 1–100`). Scheduler scores matching Nodes higher, but still schedules the Pod elsewhere if no match is available.
- *Why `IgnoredDuringExecution`?* If a Node's label is removed *after* the Pod is already running, the running Pod is **not** evicted.
- **Supported Operators:** `In`, `NotIn`, `Exists`, `DoesNotExist`, `Gt`, `Lt`.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: affinity-demo
  labels:
    app: cache
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: topology.kubernetes.io/zone
            operator: In
            values: ["eu-west-1a", "eu-west-1b"]
      preferredDuringSchedulingIgnoredDuringExecution:
      - weight: 80
        preference:
          matchExpressions:
          - key: disktype
            operator: In
            values: ["nvme"]
    podAntiAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
      - labelSelector:
          matchExpressions:
          - key: app
            operator: In
            values: ["cache"]
        topologyKey: kubernetes.io/hostname   # Prevents 2 'app=cache' Pods on the same Node!
  containers:
  - name: redis
    image: redis:7.2-alpine
```

#### 3. Taints (on Nodes) & Tolerations (on Pods)
- **Mental Model:** *Affinity attracts Pods to Nodes; Taints repel Pods from Nodes unless the Pod has a matching Toleration.*
- **Important:** A Toleration **permits** a Pod to schedule onto a tainted Node, but does **not** force it onto that Node (combine Toleration + `nodeSelector`/`nodeAffinity` if you want dedicated placement!).

| Taint Effect | Effect on New Pending Pods | Effect on Already Running Pods on that Node |
| :--- | :--- | :--- |
| **`NoSchedule`** | Rejected unless Pod has matching toleration | **Kept running** (not evicted) |
| **`PreferNoSchedule`** | Scheduler tries to avoid Node, but places Pod there if no other Node fits | **Kept running** (not evicted) |
| **`NoExecute`** | Rejected unless Pod has matching toleration | **Evicted immediately** (or evicted after `tolerationSeconds` if Pod tolerates `NoExecute` with a timer!) |

```bash
# Add a taint to a node
kubectl taint nodes node01 dedicated=special-user:NoSchedule

# Remove a taint from a node (append a trailing '-' minus sign!)
kubectl taint nodes node01 dedicated=special-user:NoSchedule-
```

```yaml
spec:
  tolerations:
  # 1. Exact key + value + effect match
  - key: "dedicated"
    operator: "Equal"                  # Equal (default) | Exists
    value: "special-user"
    effect: "NoSchedule"
  # 2. Match ANY value for key 'node-role.kubernetes.io/control-plane'
  - key: "node-role.kubernetes.io/control-plane"
    operator: "Exists"                 # Omit 'value' when operator is 'Exists'!
    effect: "NoSchedule"
  # 3. Stay running for 60s if node becomes unreachable/tainted with NoExecute
  - key: "node.kubernetes.io/unreachable"
    operator: "Exists"
    effect: "NoExecute"
    tolerationSeconds: 60
```

---

### 2.4 Horizontal Pod Autoscaling (`autoscaling/v2`)

Horizontal Pod Autoscaling (HPA) automatically increases or decreases `replicas` on a `Deployment`, `ReplicaSet`, or `StatefulSet` to match observed CPU, memory, or custom metrics to a target value. See [`manifests/autoscaler/horizontal.yaml`](manifests/autoscaler/horizontal.yaml).

#### Primary Goals of HPA
1. **Maintain Performance & Availability:** Automatically adds replicas when traffic spikes.
2. **Improve Cost Efficiency:** Scales down unneeded replicas during low-traffic periods (down to `minReplicas`).
3. **Automate Operations:** Eliminates manual scaling intervention.

#### How the HPA Control Loop Works
```text
[Node Kubelets (cAdvisor)] ──► [Metrics Server Add-on] ──► [metrics.k8s.io API]
                                                                    │
                                            Queries every 15s (default)
                                                                    ▼
[Deployment / StatefulSet] ◄── Updates `.spec.replicas` ── [HPA Controller]
```

1. **Metrics Server Dependency:** HPA requires the cluster **Metrics Server** add-on. Additionally, target Pods **must define `resources.requests`** (e.g., `requests.cpu`) when scaling on `Utilization` percentage—otherwise HPA cannot compute $\frac{\text{Current Usage}}{\text{Requested Resource}}$ and displays `<unknown>/80%`!
2. **Scaling Formula:**
   $$\text{desiredReplicas} = \left\lceil \text{currentReplicas} \times \left( \frac{\text{currentMetricValue}}{\text{desiredMetricValue}} \right) \right\rceil$$
   *(Example: If `currentReplicas = 4`, current average CPU utilization is `100%`, and target is `50%`, $\text{desiredReplicas} = \lceil 4 \times (100 / 50) \rceil = 8$.)*
3. **Cooldown / Stabilization Windows (Preventing "Thrashing"):**
   - **Scale-Up Stabilization Window:** Defaults to **`0` seconds** (scales up immediately when load spikes).
   - **Scale-Down Stabilization Window:** Defaults to **`300` seconds (5 minutes)**—HPA tracks recommendations over a rolling 5-minute window and chooses the highest replica count in that window before scaling down, ensuring scale-in decisions reflect sustained trends rather than brief dips.

#### Imperative CLI & Full `autoscaling/v2` Manifest
```bash
# Quickly create or export an autoscaling/v2 HPA manifest:
kubectl autoscale deployment webapp --min=2 --max=10 --cpu-percent=80
```
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: webapp
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: webapp
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization              # Utilization (%) | AverageValue (e.g. 200m)
        averageUtilization: 80
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 75
  behavior:
    scaleDown:
      stabilizationWindowSeconds: 300  # Wait 300s before scaling down
      policies:
      - type: Percent
        value: 50                      # Scale down by at most 50% of current pods per 60s
        periodSeconds: 60
```

---

### 2.5 Helm Package Manager (v3)

Helm packages Kubernetes YAML manifests into reusable **Charts**, parameterizes them via **`values.yaml`**, and manages their lifecycle inside the cluster as versioned **Releases** (stored as Kubernetes Secrets in the release's namespace in Helm v3).

- **Chart Directory Structure:**
  ```text
  mychart/
  ├── Chart.yaml          # Chart metadata (name, version, appVersion, dependencies)
  ├── values.yaml         # Default configuration values
  └── templates/          # Go-templated Kubernetes YAML manifests
      ├── deployment.yaml
      ├── service.yaml
      └── _helpers.tpl
  ```
- **Value Override Precedence (Lowest to Highest):**
  $$\text{Chart } \texttt{values.yaml} \quad < \quad \texttt{-f custom-values.yaml} \quad < \quad \texttt{--set key=value}$$

#### Complete Helm CLI Reference for CKAD
```bash
# 1. Repositories & Searching
helm repo list                                              # View existing repos
helm repo add bitnami https://charts.bitnami.com/bitnami    # Add a repo
helm repo update                                            # Refresh repo indexes
helm search repo nginx                                      # Search latest version
helm search repo bitnami/nginx --versions                   # List all available chart versions

# 2. Inspecting & Downloading Charts
helm show chart bitnami/nginx                               # View Chart.yaml metadata
helm show values bitnami/nginx > custom-values.yaml         # Export default values to edit
helm pull bitnami/nginx --untar --untardir /tmp/charts      # Download & extract chart locally

# 3. Rendering Locally (Dry-Run Without Installing)
helm template my-rel bitnami/nginx --set replicaCount=3     # Render raw K8s YAML to stdout

# 4. Installing & Upgrading Releases (Remember `-n <namespace>`!)
helm install my-rel bitnami/nginx -n web --create-namespace
helm install my-rel ./local-chart -n web -f custom-values.yaml --set service.port=8080
helm upgrade my-rel bitnami/nginx -n web --reuse-values --set image.tag=1.26
helm upgrade --install my-rel bitnami/nginx -n web          # Installs if missing, upgrades if present

# 5. Auditing, Rollbacks & Uninstalling
helm list -A                                                # List deployed releases in ALL namespaces
helm list -a -A                                             # Include failed/pending/uninstalling releases!
helm list --pending -A                                      # Filter only stuck pending releases
helm status my-rel -n web
helm get values my-rel -n web                               # Show user-supplied overrides
helm get manifest my-rel -n web                             # Dump exact K8s YAML deployed by release
helm history my-rel -n web                                  # View revision history
helm rollback my-rel 1 -n web                               # Roll back to revision 1
helm uninstall my-rel -n web                                # Delete release and all its K8s resources
```

---

## 3. Application Observability and Maintenance (15%)

### 3.1 Container Probes (`startupProbe`, `readinessProbe`, `livenessProbe`)

Probes are periodic diagnostic checks executed by the **Node's kubelet** against individual containers. See [`probes/deployment.yaml`](probes/deployment.yaml) and [`probes/main.go`](probes/main.go).

```text
Container Starts
       │
       ▼
+-------------------------------------------------------------------------+
| 1. startupProbe (Gatekeeper for Slow-Starting Apps)                     |
|    - Disables liveness & readiness checks until it succeeds ONCE.       |
|    - Max startup time allowed = failureThreshold * periodSeconds.       |
|    - If it fails `failureThreshold` times -> Container is killed/restarted
+-------------------------------------------------------------------------+
       │ (Succeeds once)
       ▼
+------------------------------------+------------------------------------+
| 2. readinessProbe (Traffic Valve)  | 3. livenessProbe (Deadlock Watch)  |
|    - Runs throughout lifecycle.    |    - Runs throughout lifecycle.    |
|    - On failure: Removes Pod IP    |    - On failure: Kubelet kills &   |
|      from Service `Endpoints`.     |      RESTARTS the container!       |
|    - Does NOT restart container!   |                                    |
+------------------------------------+------------------------------------+
```

#### Probe Handlers (Action Mechanisms)
1. **`httpGet`:** Sends an HTTP GET request to `path` and `port` (optional `scheme: HTTP|HTTPS`, `httpHeaders`). Status code **`200–399`** = success; `4xx/5xx` or timeout = failure.
2. **`tcpSocket`:** Attempts to establish a TCP socket connection on `port`. Connection established = success (ideal for Redis, MySQL, Postgres, non-HTTP TCP services).
3. **`exec`:** Runs a command inside the container (`command: ["sh", "-c", "..."]`). Exit status **`0`** = success; non-zero = failure.
4. **`grpc`:** Calls a gRPC health check service on `port`.

#### Probe Timing & Threshold Fields
- `initialDelaySeconds` *(default `0`)*: Seconds to wait after container start before running the very first probe.
- `periodSeconds` *(default `10`)*: How often (in seconds) to execute the probe. *(Note: in the original `README.md`, this was accidentally labeled `failureThreshold`!)*
- `timeoutSeconds` *(default `1`)*: Seconds after which the probe attempt times out and counts as a failure.
- `failureThreshold` *(default `3`)*: Number of consecutive probe failures required before taking action (removing endpoint or restarting container).
- `successThreshold` *(default `1`)*: Consecutive successes required after a failure to mark the probe healthy again (**must be `1` for `startupProbe` and `livenessProbe`**).

---

### 3.2 Monitoring, Logging & Interactive Debugging

#### Pod Failure States & Rapid Diagnostic Decision Tree

| Status Shown in `kubectl get pods` | What It Means | Exact Command to Diagnose | Common Root Causes |
| :--- | :--- | :--- | :--- |
| **`Pending`** | Pod accepted by API server, but **not scheduled** onto a Node yet | `kubectl describe pod <name>` (look at `Events` at the bottom) | 1. Insufficient Node CPU/Memory (`requests` too high)<br>2. Unmatched `nodeSelector` / `nodeAffinity`<br>3. Node has a **Taint** Pod doesn't tolerate<br>4. Referenced **PVC** is unbound (`Pending`) |
| **`Init:0/2`** or **`Init:Error`** | InitContainer is still running or failing | `kubectl logs <pod> -c <init-container-name>` | InitContainer waiting on a Service/DB (`nslookup`) or failing its setup command |
| **`ImagePullBackOff`** / **`ErrImagePull`** | Kubelet cannot pull the container image | `kubectl describe pod <name>` | Typo in image repository/tag, nonexistent tag, or private registry missing `imagePullSecrets` |
| **`CreateContainerConfigError`** | Kubelet cannot build the container configuration | `kubectl describe pod <name>` | 1. Missing `ConfigMap` or `Secret` (or missing key in `secretKeyRef`/`configMapKeyRef`)<br>2. `runAsNonRoot: true` set on an image whose default `USER` is `root` (UID `0`) without `runAsUser` |
| **`CrashLoopBackOff`** | Container starts, crashes/exits, and is backed off before next restart | `kubectl logs <pod> --previous` and `kubectl describe pod <pod>` | 1. Syntax error in `command`/`args`<br>2. Container command finished immediately (`Exit Code 0`) under `restartPolicy: Always`<br>3. Failing `livenessProbe` killing the container<br>4. App crash due to `readOnlyRootFilesystem: true` |
| **`OOMKilled`** (`Exit Code 137`) | Container exceeded `resources.limits.memory` | `kubectl describe pod <name>` (check `Last State: Terminated, Reason: OOMKilled`) | Memory limit set lower than the process's working memory allocation |
| **`Running` (`0/1 Ready`)** | Container is running, but **not receiving Service traffic** | `kubectl describe pod <name>` | `readinessProbe` is failing (wrong `port` or wrong HTTP `path`!) |

#### Essential Observability & Debugging Commands
```bash
# Resource Usage (Metrics Server)
kubectl top nodes --sort-by=cpu
kubectl top pods -A --sort-by=memory
kubectl top pod <pod-name> --containers                     # Per-container CPU & RAM inside a Pod

# Logs
kubectl logs <pod>                                          # Single-container Pod logs
kubectl logs <pod> -c <container-name>                      # Specific container in multi-container Pod
kubectl logs <pod> --previous                               # Logs from previous crashed instance (-p)
kubectl logs -l app=backend --all-containers=true --tail=50 # Logs across all Pods matching label
kubectl logs deploy/webapp -f                               # Stream logs from a Deployment

# Cluster Events
kubectl get events -n <ns> --sort-by='.lastTimestamp'
kubectl get events -A --field-selector type=Warning

# Interactive Shell, Ephemeral Debug Containers & Port Forwarding
kubectl exec -it <pod> -c <container> -- sh
kubectl debug -it <pod> --image=busybox:1.36 --target=<container>  # Debug distroless/minimal containers!
kubectl debug <pod> -it --copy-to=dbg-pod --container=main -- sh   # Clone Pod with overridden command
kubectl port-forward pod/<pod> 8080:80                             # Forward local 8080 -> Pod 80
kubectl port-forward svc/<service> 8080:80                         # Forward local 8080 -> Service 80
```

---

### 3.3 API Deprecations & Version Migrations

Kubernetes APIs mature through **`v1alpha1` -> `v1beta1` -> `v1` (GA)**. When upgrading clusters or fixing older YAML files on the exam, use the API server's built-in discovery commands:

```bash
# 1. List all resource types, shortnames, API groups/versions, and whether they are namespaced:
kubectl api-resources
kubectl api-resources --api-group=networking.k8s.io

# 2. List all enabled API group/version combinations on the cluster:
kubectl api-versions

# 3. Check the exact schema and preferred apiVersion of any resource:
kubectl explain ingress
kubectl explain hpa.spec.metrics
```

#### High-Frequency Historical API Migrations to Recognize

| Resource `Kind` | Deprecated / Removed `apiVersion` | Current Stable `apiVersion` | Key Schema Changes in Stable Version |
| :--- | :--- | :--- | :--- |
| **`Ingress`** | `extensions/v1beta1`, `networking.k8s.io/v1beta1` | **`networking.k8s.io/v1`** | - `pathType` (`Prefix`\|`Exact`) is now **required** on every path.<br>- `spec.ingressClassName` replaces `kubernetes.io/ingress.class` annotation.<br>- `backend.serviceName` + `backend.servicePort` changed to:<br>`backend: { service: { name: <svc>, port: { number: 80 } } }` |
| **`HorizontalPodAutoscaler`** | `autoscaling/v1`, `autoscaling/v2beta1`, `autoscaling/v2beta2` | **`autoscaling/v2`** | - `targetCPUUtilizationPercentage` replaced by `spec.metrics[]` array (`type: Resource`, `target: { type: Utilization, averageUtilization: 80 }`). |
| **`CronJob`** | `batch/v1beta1` | **`batch/v1`** | Schema structure unchanged; update `apiVersion: batch/v1`. |
| **`PodDisruptionBudget`** | `policy/v1beta1` | **`policy/v1`** | Schema structure unchanged; update `apiVersion: policy/v1`. |

*(Note: `kubectl convert` was removed from core `kubectl` binaries years ago; on the exam, use `kubectl explain <kind>` to inspect and update deprecated fields directly in Vim.)*

---

### 3.4 Custom Resource Definitions (CRDs) & The Operator Pattern

- **Custom Resource Definition (`CustomResourceDefinition`, `apiextensions.k8s.io/v1`):**
  Extends the Kubernetes API server by defining a brand-new resource `Kind` (e.g., `Backup`, `Certificate`, `KafkaTopic`) with an OpenAPI v3 validation schema.
- **Custom Resource (CR):** An actual instance of that new `Kind` created by a user via `kubectl apply -f`.
- **Custom Controller / Operator Pattern:**
  - A **CRD on its own is just structured data stored in `etcd`**.
  - An **Operator** pairs a **CRD** with a **Custom Controller** (a Pod running inside the cluster watching the Kubernetes API for changes to that CR) that continuously runs a reconciliation loop (*Observe CR spec -> Compare with real world -> Mutate cluster resources to match*).

```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: crontabs.stable.example.com          # MUST be <spec.names.plural>.<spec.group>
spec:
  group: stable.example.com
  scope: Namespaced                          # Namespaced | Cluster
  names:
    plural: crontabs
    singular: crontab
    kind: CronTab
    shortNames:
    - ct
  versions:
  - name: v1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              cronSpec:
                type: string
              image:
                type: string
              replicas:
                type: integer
```
Once applied, you can run `kubectl get ct`, `kubectl explain crontab.spec`, and `kubectl describe ct <name>` just like any built-in Kubernetes object!


---

## 4. Application Environment, Configuration and Security (25%)

### 4.1 ConfigMaps (Env Vars vs Volume Mounts vs `subPath` Trap)

A **ConfigMap** is an API object that stores **non-confidential** configuration data as key-value pairs in `etcd` (up to `1 MiB` max size).

#### When to Use Env Vars vs Volume-Mounted ConfigMaps

| Criteria | Environment Variables (`env` / `envFrom`) | Volume Mount (`volumes[].configMap`) |
| :--- | :--- | :--- |
| **Best for** | Static startup flags (`LOG_LEVEL`, `PORT`, `APP_ENV`) | Large or structured config files (`nginx.conf`, `config.yaml`, `settings.json`) |
| **Live Updates while Pod runs?** | **NO** (env vars are injected only at process creation; requires Pod restart) | **YES** (kubelet updates mounted files via atomic symlink swap within ~60s) |
| **Exception (`subPath` Trap!)** | N/A | **NO** — If mounted using `volumeMounts[].subPath`, the file is **frozen at startup** and will **NOT** auto-update! |
| **Overwrites existing directory?**| N/A | Mounting a volume at `/etc/app` hides any existing files in `/etc/app` unless `subPath` is used (or `items:` is mounted to a dedicated directory). |

#### All 4 Ways to Consume a ConfigMap in YAML
See [`manifests/configmap/deployment.yaml`](manifests/configmap/deployment.yaml) and [`manifests/configmap/config.txt`](manifests/configmap/config.txt).

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
immutable: false                     # Set true to lock against accidental edits & reduce apiserver watch load
data:
  APP_ENV: "production"
  PORT: "8080"
  app.properties: |
    cache.enabled=true
    max.connections=100
---
apiVersion: v1
kind: Pod
metadata:
  name: cm-consumer
spec:
  containers:
  - name: app
    image: busybox:1.36
    # 1. Use ConfigMap env var inside command-line arguments via $(VAR_NAME) syntax:
    command: ["sh", "-c", "echo Running in $(APP_ENV) on port $(LISTEN_PORT) && sleep 3600"]
    # 2. Import ALL key-value pairs as environment variables:
    envFrom:
    - configMapRef:
        name: app-config
        optional: false              # If true, Pod starts even if ConfigMap doesn't exist
    # 3. Import a SINGLE key and assign it a custom environment variable name:
    env:
    - name: LISTEN_PORT
      valueFrom:
        configMapKeyRef:
          name: app-config
          key: PORT
    # 4. Mount as files in a volume:
    volumeMounts:
    - name: full-config-vol
      mountPath: /etc/config         # Creates /etc/config/APP_ENV, /etc/config/PORT, /etc/config/app.properties (Auto-updates!)
    - name: single-file-vol
      mountPath: /etc/only-props     # Creates ONLY /etc/only-props/custom.properties via `items:` (Auto-updates!)
    - name: full-config-vol
      mountPath: /etc/app.properties # Uses `subPath` to mount a single file without hiding /etc (Does NOT auto-update!)
      subPath: app.properties
  volumes:
  - name: full-config-vol
    configMap:
      name: app-config
      defaultMode: 0644              # File permission bits
  - name: single-file-vol
    configMap:
      name: app-config
      items:
      - key: app.properties
        path: custom.properties
```

---

### 4.2 Secrets (Why "Encrypted ConfigMap != Secret", Types & Consumption)

A **Secret** stores **sensitive** data (passwords, OAuth tokens, SSH keys, TLS certificates, private registry credentials).

#### Why "Encrypted ConfigMap != Secret"
1. **Kubernetes Treats ConfigMaps as Plain Config:** ConfigMaps may be logged, exposed in debug outputs, or mounted with broader default permissions.
2. **Separate RBAC Isolation:** Kubernetes lets you restrict RBAC access to `secrets` independently from `configmaps`. If you put sensitive data inside a ConfigMap, anyone with `get configmaps` permission can read it.
3. **Special Runtime & Storage Handling:**
   - Secret volumes are mounted into containers using **in-memory `tmpfs`** (RAM) on the Node rather than written to physical disk, and wiped immediately when the Pod terminates.
   - Secrets can be encrypted at rest in `etcd` (`EncryptionConfiguration`) and integrated with external secret stores via CSI Secret Store drivers (HashiCorp Vault, AWS Secrets Manager).
4. **Communicates Intent:** Declares explicitly to operators and tooling: *"This data is confidential."*
   - *Reminder:* In a Secret's YAML manifest, values under `data:` are **base64-encoded** (not encrypted!). Use `stringData:` when writing YAML manually so Kubernetes encodes plain strings into base64 automatically.

#### Built-In Secret Types & Imperative CLI Creation
See [`manifests/secret/deployment.yaml`](manifests/secret/deployment.yaml).

| Secret Type | Purpose | Imperative `kubectl create secret` Command |
| :--- | :--- | :--- |
| **`Opaque`** *(default)* | Arbitrary user-defined key-value pairs | `k create secret generic db-sec --from-literal=USER=admin --from-file=key.pem` |
| **`kubernetes.io/tls`** | TLS certificate (`tls.crt`) and private key (`tls.key`) for Ingress/HTTPS | `k create secret tls web-tls --cert=path/to/tls.crt --key=path/to/tls.key` |
| **`kubernetes.io/dockerconfigjson`** | Private container registry credentials used by `imagePullSecrets` | `k create secret docker-registry regcred --docker-server=registry.io --docker-username=u --docker-password=p` |
| **`kubernetes.io/basic-auth`** | Basic auth (`username` and `password` keys) | `k create secret generic auth-sec --type=kubernetes.io/basic-auth --from-literal=username=u --from-literal=password=p` |
| **`kubernetes.io/ssh-auth`** | SSH private key (`ssh-privatekey` key) | `k create secret generic ssh-sec --type=kubernetes.io/ssh-auth --from-file=ssh-privatekey=~/.ssh/id_rsa` |

```bash
# Decode a base64-encoded secret value from the cluster in one line:
kubectl get secret db-sec -o jsonpath='{.data.USER}' | base64 --decode ; echo
```

---

### 4.3 Resource Management (Requests, Limits, `ResourceQuota`, `LimitRange`, QoS & `PriorityClass`)

> **Core Mental Model:** *Requests decide placement. Limits decide enforcement. PriorityClass & QoS decide eviction.*

#### 1. CPU & Memory Units, Requests & Limits
- **CPU Units:** `1` = `1000m` (1000 millicores) = 1 vCPU/core. `250m` = `0.25` core; `500m` = `0.5` core.
- **Memory Units:**
  - **Binary (Base-2 — Standard in K8s):** `1Ki` = $1024\text{ B}$, `1Mi` (Mebibyte) = $1024^2\text{ B}$, `1Gi` (Gibibyte) = $1024^3\text{ B}$.
  - **Decimal (Base-10):** `1K` = $1000\text{ B}$, `1M` (Megabyte) = $1000^2\text{ B}$, `1G` (Gigabyte) = $1000^3\text{ B}$.
- **`requests` (Minimum Guaranteed):** Used by `kube-scheduler` to choose a Node with sufficient allocatable resources. If no Node has enough unallocated capacity, the Pod stays **`Pending`**.
- **`limits` (Hard Ceiling):** Ignored by the scheduler; enforced on the Node by **kubelet** via Linux kernel **cgroups**:
  - **CPU Limit Exceeded (Compressible Resource):** Kernel CFS throttles the container's CPU cycles. The app runs slower, **not** killed.
  - **Memory Limit Exceeded (Incompressible Resource):** Kernel triggers the **Out-Of-Memory (OOM) Killer**, terminating the container process with **`Exit Code 137`** and reporting status **`OOMKilled`**.

---

#### 2. `ResourceQuota` (Namespace Aggregate Budget)
Limits the **sum total** of resources or object counts across all objects in a Namespace. See [`manifests/resource-quota/resource-quota.yaml`](manifests/resource-quota/resource-quota.yaml).

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: team-quota
  namespace: dev
spec:
  hard:
    # 1. Compute Resource Quotas
    requests.cpu: "2"
    requests.memory: 2Gi
    limits.cpu: "4"
    limits.memory: 4Gi
    # 2. Storage Resource Quotas
    requests.storage: 50Gi
    persistentvolumeclaims: "5"
    gold.storageclass.storage.k8s.io/requests.storage: 20Gi
    # 3. Object Count Quotas
    pods: "10"
    services: "5"
    services.nodeports: "2"
    configmaps: "10"
    secrets: "10"
    count/deployments.apps: "5"
```

> **Exam Rule:** If a Namespace has an active `ResourceQuota` for `requests.cpu`, `requests.memory`, `limits.cpu`, or `limits.memory`, **every new Pod created in that Namespace MUST specify those fields** (or have them auto-injected by a `LimitRange`), or the API server will reject the Pod with `403 Forbidden: failed quota`!

---

#### 3. `LimitRange` (Per-Pod / Per-Container Defaults & Min/Max Bounds)
Enforces minimum/maximum resource bounds on **individual** Containers, Pods, or PVCs, and automatically injects default `requests` and `limits` into new containers that omit them. See [`manifests/limitrange/limitrange.yaml`](manifests/limitrange/limitrange.yaml).

```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: cpu-mem-limit-range
  namespace: dev
spec:
  limits:
  - type: Container                  # Container | Pod | PersistentVolumeClaim
    default:                         # Injected as container `limits` if omitted
      cpu: 500m
      memory: 512Mi
    defaultRequest:                  # Injected as container `requests` if omitted
      cpu: 200m
      memory: 256Mi
    max:                             # Maximum limit any single container can specify
      cpu: "1"
      memory: 1Gi
    min:                             # Minimum request any single container must specify
      cpu: 100m
      memory: 128Mi
    maxLimitRequestRatio:            # Optional: limit / request ratio cap
      cpu: "4"
  - type: PersistentVolumeClaim
    min:
      storage: 1Gi
    max:
      storage: 10Gi
```

---

#### 4. Quality of Service (QoS) Classes & `PriorityClass`
See [`manifests/resource-management/`](manifests/resource-management/) and [`manifests/priority-class/`](manifests/priority-class/).

```text
+---------------------------------------------------------------------------------+
|                         POD QOS CLASS ASSIGNMENT RULES                          |
+----------------+---------------------------------------------+------------------+
| QoS Class      | Exact Criteria                              | Eviction Order   |
+----------------+---------------------------------------------+------------------+
| 1. Guaranteed  | EVERY container in Pod defines BOTH CPU &   | Evicted LAST     |
|                | memory requests and limits, AND             | (Most protected) |
|                | `request == limit` for every resource.      |                  |
|                | *(Best for: DBs, message queues, stateful)* |                  |
+----------------+---------------------------------------------+------------------+
| 2. Burstable   | Does not qualify for Guaranteed, AND at     | Evicted SECOND   |
|                | least ONE container in the Pod defines a    |                  |
|                | CPU or memory request or limit.             |                  |
|                | *(Best for: Web servers, APIs, stateless)*  |                  |
+----------------+---------------------------------------------+------------------+
| 3. BestEffort  | ZERO CPU or memory requests or limits on    | Evicted FIRST!   |
|                | ANY container in the Pod.                   | (Least protected)|
|                | *(Best for: Non-critical dev/batch tasks)*  |                  |
+----------------+---------------------------------------------+------------------+
```
*(Note: look at [`manifests/resource-management/burstable-guaranteed.yaml`](manifests/resource-management/burstable-guaranteed.yaml) and [`manifests/resource-management/burstable-bestEffort.yaml`](manifests/resource-management/burstable-bestEffort.yaml) in this repo: if Container 1 is Guaranteed and Container 2 is Burstable, the overall **Pod** QoS is **`Burstable`**!)*

- **`PriorityClass` (`scheduling.k8s.io/v1`):**
  1. **Scheduling Order:** Higher `value` pending Pods in the scheduling queue are scheduled first.
  2. **Preemption:** If a high-priority Pod cannot fit on any Node, the scheduler evicts (preempts) lower-priority Pods from a Node to free capacity (unless `preemptionPolicy: Never` is set on the PriorityClass).
  3. **Node Pressure Eviction:** When a Node runs low on memory or disk, kubelet identifies Pods using more resources than they requested, ranks them by **PriorityClass** (lowest priority evicted first), and breaks ties among equal-priority Pods using **QoS Class** (`BestEffort` -> `Burstable` -> `Guaranteed`).

---

### 4.4 Security Context, Linux UID/GID, SetUID, Capabilities & Pod Security Standards

#### 1. Linux User/Group IDs & The `setuid` Privilege Escalation Mechanism
- **Linux UID/GID Ranges:**
  - **`0`:** `root` superuser / group.
  - **`1–999`:** Reserved for system services and background daemons.
  - **`1000+`:** Standard non-privileged user accounts.
- **How `setuid` Privilege Escalation Works (Real-World Linux Example: `/usr/bin/passwd`):**
  - Normally, when a user executes a binary, the resulting process runs with the Effective User ID (**EUID**) of the user who launched it.
  - However, when an executable file has the **`setuid` permission bit** enabled, any user who executes it spawns a process running with the **EUID of the file's owner** instead!
  - *Example:* In Linux, encrypted user passwords live in `/etc/shadow`, which is readable/writable **only by `root`**. How can a normal user (`developer`, UID `1000`) change their own password using `/usr/bin/passwd`? Because `/usr/bin/passwd` is owned by `root` and has the **`setuid` bit set** (`-rwsr-xr-x`). When `developer` runs `passwd`, the kernel spawns that child process as **`root` (UID `0`)** so it can update `/etc/shadow`.
  - **In Kubernetes:** Setting **`allowPrivilegeEscalation: false`** on a container's `securityContext` sets the Linux kernel `no_new_privs` flag, blocking child processes from gaining elevated privileges via `setuid`/`setgid` binaries or file capabilities!

---

#### 2. Pod-Level vs Container-Level `securityContext` Matrix
See [`manifests/security-context/security-context.yaml`](manifests/security-context/security-context.yaml).

| Field | Pod (`spec.securityContext`) | Container (`containers[].securityContext`) | Behavior & Exam Notes |
| :--- | :---: | :---: | :--- |
| **`runAsUser`** | YES | YES (overrides Pod) | Runs container process as specified UID. |
| **`runAsGroup`** | YES | YES (overrides Pod) | Runs container process with specified primary GID. |
| **`runAsNonRoot`** | YES | YES (overrides Pod) | Kubelet validates that effective UID != `0` before starting. *(If image defaults to `root` and `runAsUser` is not set -> fails with `CreateContainerConfigError`!)* |
| **`seccompProfile`** | YES | YES (overrides Pod) | Filters Linux syscalls (`type: RuntimeDefault`, `Localhost`, `Unconfined`). |
| **`fsGroup`** | **YES (Pod ONLY!)** | **NO (Invalid!)** | Kubelet `chown`s mounted volumes to this GID and adds this GID to the container user's supplementary groups. |
| **`supplementalGroups`** | **YES (Pod ONLY!)** | **NO (Invalid!)** | Array of additional GIDs attached to the container user. |
| **`allowPrivilegeEscalation`** | **NO (Invalid!)** | **YES (Container ONLY!)** | Prevents `setuid` elevation (`false` required when `privileged: false` in Restricted PSS). |
| **`readOnlyRootFilesystem`** | **NO (Invalid!)** | **YES (Container ONLY!)** | Mounts container `/` as read-only (blocks writes even by UID `0`!). Mount `emptyDir`s for writable temp paths (`/tmp`, `/var/cache/nginx`, `/var/run`). |
| **`privileged`** | **NO (Invalid!)** | **YES (Container ONLY!)** | Grants full host root & device access (`false` by default). |
| **`capabilities`** | **NO (Invalid!)** | **YES (Container ONLY!)** | Granular Linux kernel capabilities (`drop: ["ALL"]`, `add: ["NET_BIND_SERVICE", "SYS_TIME"]`). |

---

### 4.5 ServiceAccounts & Role-Based Access Control (RBAC)

See [`manifests/service-account/`](manifests/service-account/).

#### ServiceAccounts
- Every namespace automatically has a `default` ServiceAccount. If a Pod does not set `spec.serviceAccountName`, it uses `default`.
- Kubelet mounts a time-bound, projected JWT token into each Pod at `/var/run/secrets/kubernetes.io/serviceaccount/token` (along with `ca.crt` and `namespace`).
- To prevent a Pod from receiving an API token mount, set **`automountServiceAccountToken: false`** on the ServiceAccount or on `pod.spec` (`pod.spec` takes precedence).

#### RBAC (`rbac.authorization.k8s.io/v1`): The 4 Objects

```text
+-----------------------------------+          +-----------------------------------+
| Subject (WHO)                     |          | Role / ClusterRole (WHAT)         |
| - ServiceAccount (name + ns)      |          | - apiGroups: [""], ["apps"], ...  |
| - User (string)                   |          | - resources: ["pods", "secrets"]  |
| - Group (string)                  |          | - verbs: ["get", "list", "watch"] |
+-----------------+-----------------+          +-----------------+-----------------+
                  │                                              │
                  └──────────────► Bound Together By ◄───────────┘
                                 +-------------------+
                                 | RoleBinding       | (Scoped to 1 Namespace)
                                 |        OR         |
                                 | ClusterRoleBinding| (Cluster-Wide across ALL Namespaces)
                                 +-------------------+
```

- **The 3 Valid Binding Combinations:**
  1. **`Role` + `RoleBinding`:** Grants permissions in **one namespace**.
  2. **`ClusterRole` + `ClusterRoleBinding`:** Grants permissions **cluster-wide** (all namespaces + non-namespaced resources like `nodes`, `persistentvolumes`, `namespaces`).
  3. **`ClusterRole` + `RoleBinding` (High-Frequency Exam Pattern!):** Reuses a cluster-wide `ClusterRole` definition (like `view` or `edit`), but scopes the granted permissions **only to the `RoleBinding`'s single namespace**!
  *(Note: A `ClusterRoleBinding` can NEVER reference a namespaced `Role`.)*

- **Subresources Worth Knowing:**
  - To allow reading Pod logs (`kubectl logs`), grant `get` on **`pods/log`** (`--resource=pods,pods/log`).
  - To allow `kubectl exec`, grant `create` on **`pods/exec`**.
  - To restrict a Role to specific named objects only, add **`resourceNames: ["my-configmap"]`** (`--resource-name=my-configmap`).

```bash
# Fast Imperative RBAC Creation & Verification:
kubectl create serviceaccount app-sa -n prod
kubectl create role pod-log-reader -n prod --verb=get,list,watch --resource=pods,pods/log,deployments.apps
kubectl create rolebinding app-sa-bind -n prod --role=pod-log-reader --serviceaccount=prod:app-sa

# Always test RBAC with `kubectl auth can-i`:
kubectl auth can-i get pods/log -n prod --as=system:serviceaccount:prod:app-sa        # yes
kubectl auth can-i delete deployments.apps -n prod --as=system:serviceaccount:prod:app-sa # no
```

---

## 5. Services and Networking (20%)

### 5.1 Kubernetes Networking Fundamentals & CoreDNS

- **Flat, NAT-Free Pod Network:**
  - Every Pod gets its own unique, cluster-wide routable IP address.
  - Every Pod can reach every other Pod IP directly across any Node without Network Address Translation (NAT), unless restricted by a `NetworkPolicy`.
  - **Pod IPs are ephemeral:** When a Pod restarts or is replaced by a Deployment, it receives a brand-new IP address.
- **What Problem Does a `Service` Solve?**
  1. **Stable Virtual IP (`ClusterIP`)** that never changes for the life of the Service.
  2. **Stable DNS Name** in CoreDNS (`<svc>` within the same namespace; `<svc>.<ns>.svc.cluster.local` across namespaces).
  3. **L4 Load Balancing** across all healthy (`Ready`) Pods matching the Service's label `selector`.

```text
Client Pod
   │ 1. Resolves `web-svc` via CoreDNS -> Gets Service ClusterIP (e.g. 10.96.45.12)
   ▼
Service ClusterIP (10.96.45.12:80) [Virtual IP — does not exist on any physical NIC!]
   │ 2. Node kernel (iptables / IPVS rules programmed by `kube-proxy`) intercepts packet
   ▼
kube-proxy DNAT Translation
   │ 3. Rewrites destination IP:port to one Ready backend Endpoint (`targetPort`)
   ▼
Destination Pod IP (10.244.2.18:8080)
```

---

### 5.2 Services (`ClusterIP`, `NodePort`, `LoadBalancer`, `ExternalName`, Headless)

See [`manifests/service/nodeport.yaml`](manifests/service/nodeport.yaml) and [`manifests/ingress/service.yaml`](manifests/ingress/service.yaml).

| Service Type | How It Works | Accessible From |
| :--- | :--- | :--- |
| **`ClusterIP`** *(default)* | Allocates an internal-only virtual IP address. | Inside the cluster only |
| **`NodePort`** | Creates a `ClusterIP` **AND** opens a static port (`30000–32767`) on **every Node's IP**. Traffic arriving at *any* Node on `<NodeIP>:<nodePort>` is routed by `kube-proxy` to any ready backend Pod—even if that Pod runs on a different Node! | Inside cluster + Outside via `<AnyNodeIP>:<nodePort>` |
| **`LoadBalancer`** | Creates a `NodePort` + `ClusterIP` **AND** requests an external cloud LoadBalancer forwarding to the NodePorts. | Inside + Outside via Cloud LB Public IP |
| **`ExternalName`** | Creates a CoreDNS `CNAME` alias pointing to an external DNS name (`spec.externalName: api.stripe.com`). No `ClusterIP`, no `selector`, no proxying. | Inside cluster via DNS alias |
| **Headless (`clusterIP: None`)** | Allocates no `ClusterIP`. CoreDNS returns the individual backend Pod IPs directly (and creates `<pod-name>.<svc>` DNS records for StatefulSets). | Inside cluster via direct Pod DNS |

- **The 3 Service Port Fields:**
  - **`port`:** The port the **Service** listens on inside the cluster (`ClusterIP:port`).
  - **`targetPort`:** The port (or named `containerPort`) on the **Pod** where traffic is sent (defaults to `port` if omitted).
  - **`nodePort`:** The static external port (`30000–32767`) opened on every Node when `type: NodePort` or `LoadBalancer`.

- **Debugging Services in 10 Seconds:**
  Always check `kubectl get endpoints <service-name>`!
  - If `ENDPOINTS` is `<none>`, either:
    1. The Service's `spec.selector` does not match the Pods' `metadata.labels` (check with `k get pods --show-labels`), OR
    2. The Pods are failing their `readinessProbe` (`0/1 Ready`).

---

### 5.3 Ingress & TLS Termination (`networking.k8s.io/v1`)

An **Ingress** exposes HTTP and HTTPS routes from outside the cluster to internal Services through a single public entry point (managed by an **Ingress Controller** such as NGINX, Traefik, or HAProxy).

#### Real-World Ingress Example: An E-Commerce Microservices Shop
Suppose your e-commerce website runs three services in Kubernetes (`frontend-service`, `api-service`, `auth-service`) under `my-cool-shop.com`:
- **Without Ingress:** Exposing each microservice via its own `LoadBalancer` requires three separate public IPs (`203.0.113.10`, `.11`, `.12`) and triples cloud load balancer costs.
- **With Ingress:** A single Ingress Controller public IP routes traffic intelligently at Layer 7:
  - **Path-based routing:**
    - `https://my-cool-shop.com/` -> `frontend-service:80`
    - `https://my-cool-shop.com/api` -> `api-service:8080`
    - `https://my-cool-shop.com/auth` -> `auth-service:9000`
  - **Host-based routing:**
    - `https://my-cool-shop.com` -> `frontend-service:80`
    - `https://api.my-cool-shop.com` -> `api-service:8080`

#### Complete Annotated `networking.k8s.io/v1` Ingress Template
See [`manifests/ingress/ingress.yaml`](manifests/ingress/ingress.yaml).

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-shop-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  ingressClassName: nginx                  # Selects which IngressController implements this Ingress
  defaultBackend:                          # Optional fallback when no host/path rule matches
    service:
      name: default-404-service
      port:
        number: 80
  tls:                                     # Optional HTTPS termination
  - hosts:
    - my-shop.com
    - api.my-shop.com
    secretName: shop-tls-secret            # Must be a `kubernetes.io/tls` Secret in the same namespace
  rules:
  - host: api.my-shop.com
    http:
      paths:
      - path: /order
        pathType: Prefix                   # Prefix | Exact | ImplementationSpecific
        backend:
          service:
            name: api-service
            port:
              number: 80
      - path: /auth
        pathType: Exact
        backend:
          service:
            name: auth-service
            port:
              number: 80
  - host: my-shop.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: frontend-service
            port:
              number: 80
```

```bash
# Fast Imperative Ingress YAML Generator:
# Note: '/path*' with an asterisk generates pathType: Prefix; '/path' without an asterisk generates pathType: Exact!
kubectl create ingress my-shop-ingress \
  --class=nginx \
  --rule="api.my-shop.com/order*=api-service:80,tls=shop-tls-secret" \
  --rule="api.my-shop.com/auth=auth-service:80,tls=shop-tls-secret" \
  --rule="my-shop.com/*=frontend-service:80,tls=shop-tls-secret" \
  $do > ingress.yaml
```

---

### 5.4 NetworkPolicies (Isolation Rules, `AND` vs `OR` Selector Trap & DNS Egress)

`NetworkPolicy` (`networking.k8s.io/v1`) controls L3/L4 IP/Port traffic to (`Ingress`) and from (`Egress`) selected Pods. See [`manifests/generic/network-policy.yaml`](manifests/generic/network-policy.yaml).

#### Five Fundamental Rules of NetworkPolicies
1. **Default Allow when Unselected:** If no `NetworkPolicy` selects a Pod, all incoming and outgoing traffic to/from that Pod is allowed.
2. **Directional Isolation on Selection:**
   - A Pod becomes isolated for **Ingress** as soon as at least one `NetworkPolicy` selects it and includes `"Ingress"` in `spec.policyTypes`.
   - A Pod becomes isolated for **Egress** as soon as at least one `NetworkPolicy` selects it and includes `"Egress"` in `spec.policyTypes`.
3. **Allow-List Only (No Deny Rules):** NetworkPolicies have no "deny" rules—once isolated, **everything not explicitly listed in `ingress` / `egress` is denied**. An empty `ingress: []` (or omitting `ingress:` when `policyTypes: ["Ingress"]` is set) denies all incoming traffic.
4. **Additive Union:** If multiple NetworkPolicies select the same Pod, their allowed rules are combined (union/OR).
5. **Namespaced Scope:** `spec.podSelector` only selects Pods in the NetworkPolicy's **own namespace**. When `ingress[].from[].podSelector` is used alone (without `namespaceSelector`), it only allows matching Pods from the **same namespace**.

#### The #1 CKAD Exam Trap: `AND` vs `OR` in `from:` / `to:` Selectors

```yaml
# =========================================================================
# PATTERN A: ONE DASH (-) = LOGICAL AND
# =========================================================================
  ingress:
  - from:
    - namespaceSelector:           # <--- Dash #1 starts a single list item
        matchLabels:
          env: prod
      podSelector:                 # <--- NO dash! Part of the SAME list item = AND!
        matchLabels:
          role: frontend
# Meaning: Allows traffic ONLY from Pods labeled `role: frontend` that are
#          running INSIDE namespaces labeled `env: prod`.

# =========================================================================
# PATTERN B: TWO DASHES (- ... -) = LOGICAL OR
# =========================================================================
  ingress:
  - from:
    - namespaceSelector:           # <--- Dash #1: List item 1
        matchLabels:
          env: prod
    - podSelector:                 # <--- Dash #2: Separate list item = OR!
        matchLabels:
          role: frontend
# Meaning: Allows traffic from ANY Pod inside a namespace labeled `env: prod`
#          OR from any Pod labeled `role: frontend` in THIS policy's namespace!
```

#### Complete Ingress + Egress NetworkPolicy Template (with CoreDNS Port 53 Rule!)
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: backend-full-netpol
  namespace: app-ns
spec:
  podSelector:
    matchLabels:
      app: backend                 # Applies to 'app=backend' Pods in 'app-ns'
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    # Allow from 'app=frontend' in 'app-ns'
    - podSelector:
        matchLabels:
          app: frontend
    # OR from 'role=scraper' inside namespace 'monitoring'
    - namespaceSelector:
        matchLabels:
          kubernetes.io/metadata.name: monitoring
      podSelector:
        matchLabels:
          role: scraper
    ports:
    - protocol: TCP
      port: 8080
  egress:
  # 1. CRITICAL: Allow UDP & TCP port 53 so the Pod can resolve Service DNS names!
  - ports:
    - protocol: UDP
      port: 53
    - protocol: TCP
      port: 53
  # 2. Allow outbound traffic to 'app=postgres' in 'app-ns' on TCP 5432
  - to:
    - podSelector:
        matchLabels:
          app: postgres
    ports:
    - protocol: TCP
      port: 5432
  # 3. Allow outbound traffic to external CIDR block except one restricted subnet
  - to:
    - ipBlock:
        cidr: 172.16.0.0/16
        except:
        - 172.16.8.0/24
    ports:
    - protocol: TCP
      port: 443
```

---

## 6. Complete Exam-Day `kubectl` & `helm` Cheat Sheet

### 6.1 Exam Terminal Setup (Run First Thing!)
```bash
alias k=kubectl
export do="--dry-run=client -o yaml"
export now="--force --grace-period=0"
# In ~/.vimrc: set tabstop=2 shiftwidth=2 expandtab number
```

### 6.2 Contexts, Namespaces & API Discovery
```bash
k config get-contexts                                       # List available cluster contexts
k config use-context <context-name>                         # Switch cluster context
k config set-context --current --namespace=<ns>             # Change default namespace
k create namespace <ns>
k api-resources --namespaced=true                           # List all namespaced resource types & shortnames
k api-versions                                              # List all active API group/versions
k explain pod.spec.containers.securityContext --recursive   # Instant built-in field schema lookup
```

### 6.3 Pods, Deployments, Scaling & Rollouts
```bash
# Pods
k run web --image=nginx:1.25-alpine --port=80 --labels="app=web,env=prod" --env="MODE=live"
k run job-pod --image=busybox:1.36 --restart=Never --command -- sh -c "date && sleep 5"
k get pod web -o yaml > web.yaml && k replace --force -f web.yaml   # Replace immutable Pod fields

# Deployments & Rollouts
k create deployment web --image=nginx:1.25-alpine --replicas=3
k scale deployment web --replicas=5
k set image deployment/web nginx=nginx:1.26-alpine
k set env deployment/web APP_ENV=prod
k set resources deployment/web -c=nginx --requests=cpu=100m,memory=64Mi --limits=cpu=200m,memory=128Mi
k set serviceaccount deployment/web my-sa
k rollout status deployment/web
k rollout history deployment/web --revision=2
k rollout pause deployment/web && k rollout resume deployment/web
k rollout undo deployment/web --to-revision=1
k rollout restart deployment/web
k autoscale deployment web --min=2 --max=10 --cpu-percent=75
```

### 6.4 Jobs, CronJobs, ConfigMaps, Secrets & Quotas
```bash
# Jobs & CronJobs
k create job my-job --image=busybox:1.36 -- sh -c "echo hello"
k create cronjob my-cj --image=busybox:1.36 --schedule="*/5 * * * *" -- date
k create job manual-run --from=cronjob/my-cj

# ConfigMaps & Secrets
k create configmap app-cm --from-literal=K1=V1 --from-file=app.conf=/path/to/app.conf
k create secret generic app-sec --from-literal=PASS=s3cr3t
k create secret tls app-tls --cert=tls.crt --key=tls.key
k create secret docker-registry regcred --docker-server=reg.io --docker-username=u --docker-password=p

# ResourceQuota
k create quota my-quota --hard=pods=5,requests.cpu=1,requests.memory=1Gi,limits.cpu=2,limits.memory=2Gi
```

### 6.5 Services, Ingress, Network Testing & RBAC
```bash
# Services & Ingress
k expose pod web --name=web-svc --port=80 --target-port=80
k expose deployment web --name=web-np --type=NodePort --port=8080 --target-port=80
k create service clusterip headless-svc --clusterip="None" --tcp=80:80
k create service externalname ext-db --external-name=db.example.com
k create ingress app-ing --class=nginx --rule="shop.local/api*=api-svc:8080,tls=shop-tls"

# Quick Ephemeral Network / DNS Test Pod
k run net-test --image=busybox:1.36 --rm -it --restart=Never -- wget -qO- --timeout=2 http://web-svc:80

# RBAC
k create serviceaccount bot-sa
k create role bot-role --verb=get,list,watch,create --resource=pods,pods/log,deployments.apps,jobs.batch
k create rolebinding bot-bind --role=bot-role --serviceaccount=default:bot-sa
k auth can-i create jobs.batch --as=system:serviceaccount:default:bot-sa
```

### 6.6 Labels, Annotations, Taints & JSONPath / Custom Columns
```bash
# Labels, Annotations & Taints
k label pod web tier=frontend --overwrite
k label pod web tier-                                       # Remove label 'tier'
k get pods -l 'env in (prod,staging),tier!=db' --show-labels
k annotate pod web description="frontend pod"
k taint nodes node01 key=val:NoSchedule
k taint nodes node01 key=val:NoSchedule-                    # Remove taint

# JSONPath, Custom Columns & Sorting
k get pods -o custom-columns=NAME:.metadata.name,IMAGE:.spec.containers[*].image,QOS:.status.qosClass
k get nodes -o jsonpath='{.items[*].status.addresses[?(@.type=="InternalIP")].address}{"\n"}'
k get pods -A --field-selector=status.phase=Running --sort-by=.metadata.creationTimestamp
k get events --sort-by='.lastTimestamp'
```

---

## 7. Repository Manifests Index & Bug Audit of `manifests/`

### 7.1 Index of Local Examples in This Repository
- **Probes & Multi-Stage Go Image:** [`probes/Dockerfile`](probes/Dockerfile), [`probes/main.go`](probes/main.go), [`probes/deployment.yaml`](probes/deployment.yaml)
- **Generic Workloads & NetworkPolicy:** [`manifests/generic/deployment.yaml`](manifests/generic/deployment.yaml), [`manifests/generic/config-map.yaml`](manifests/generic/config-map.yaml), [`manifests/generic/network-policy.yaml`](manifests/generic/network-policy.yaml)
- **Batch Jobs & CronJobs:** [`manifests/jobs/job.yaml`](manifests/jobs/job.yaml), [`manifests/jobs/cronjob.yaml`](manifests/jobs/cronjob.yaml)
- **Horizontal Pod Autoscaler:** [`manifests/autoscaler/horizontal.yaml`](manifests/autoscaler/horizontal.yaml)
- **ConfigMaps & Secrets:** [`manifests/configmap/config.txt`](manifests/configmap/config.txt), [`manifests/configmap/deployment.yaml`](manifests/configmap/deployment.yaml), [`manifests/secret/deployment.yaml`](manifests/secret/deployment.yaml)
- **Storage (PVC):** [`manifests/pvc/pvc.yaml`](manifests/pvc/pvc.yaml), [`manifests/pvc/pod-with-pvc.yaml`](manifests/pvc/pod-with-pvc.yaml)
- **QoS Classes, Quotas, LimitRanges & PriorityClass:** [`manifests/resource-management/guaranteed.yaml`](manifests/resource-management/guaranteed.yaml), [`manifests/resource-management/burstable-guaranteed.yaml`](manifests/resource-management/burstable-guaranteed.yaml), [`manifests/resource-management/burstable-bestEffort.yaml`](manifests/resource-management/burstable-bestEffort.yaml), [`manifests/resource-management/bestEffort.yaml`](manifests/resource-management/bestEffort.yaml), [`manifests/resource-quota/resource-quota.yaml`](manifests/resource-quota/resource-quota.yaml), [`manifests/limitrange/limitrange.yaml`](manifests/limitrange/limitrange.yaml), [`manifests/priority-class/high.yaml`](manifests/priority-class/high.yaml), [`manifests/priority-class/deployment.yaml`](manifests/priority-class/deployment.yaml)
- **SecurityContext & RBAC:** [`manifests/security-context/security-context.yaml`](manifests/security-context/security-context.yaml), [`manifests/service-account/service-account.yaml`](manifests/service-account/service-account.yaml), [`manifests/service-account/role.yaml`](manifests/service-account/role.yaml), [`manifests/service-account/role-binding.yaml`](manifests/service-account/role-binding.yaml), [`manifests/service-account/deployment.yaml`](manifests/service-account/deployment.yaml)
- **Services & Ingress:** [`manifests/service/nodeport.yaml`](manifests/service/nodeport.yaml), [`manifests/ingress/deployment.yaml`](manifests/ingress/deployment.yaml), [`manifests/ingress/service.yaml`](manifests/ingress/service.yaml), [`manifests/ingress/ingress.yaml`](manifests/ingress/ingress.yaml)
- **Practice Exams & Study Plan:** [`CKAD_STUDY_PLAN.md`](CKAD_STUDY_PLAN.md), [`quiz/CKAD_QUIZ_MCQ.html`](quiz/CKAD_QUIZ_MCQ.html), [`warmup-exam/CKAD_WARMUP_EXAM.html`](warmup-exam/CKAD_WARMUP_EXAM.html), [`warmup-exam/setup.yaml`](warmup-exam/setup.yaml)

---

### 7.2 Bug Audit of Existing Files in `manifests/` (Great Exam Triage Practice!)

While reviewing the YAML files you wrote a year ago in [`manifests/`](manifests/), four subtle bugs stood out that would fail `kubectl apply` validation or fail to bind on a live cluster. Watch out for these exact patterns:

1. **[`manifests/jobs/cronjob.yaml`](manifests/jobs/cronjob.yaml) (Line 9 — Field Name Typo):**
   - **Current:** `concurrenyPolicy: Allow` (missing the second `c`)
   - **Fix:** `concurrencyPolicy: Allow` (Kubernetes strict schema validation rejects misspelled field names with `unknown field "concurrenyPolicy"`).
2. **[`manifests/service-account/deployment.yaml`](manifests/service-account/deployment.yaml) (Line 6) & [`manifests/service-account/role-binding.yaml`](manifests/service-account/role-binding.yaml) (Line 7):**
   - **Bug A in `deployment.yaml`:** `serviceAccountName: my-service-account` is placed under `Deployment.spec.serviceAccountName` (invalid!) instead of inside the Pod template at **`Deployment.spec.template.spec.serviceAccountName`**.
   - **Bug B in `role-binding.yaml`:** The `RoleBinding` references `name: my-app-sa`, whereas [`manifests/service-account/service-account.yaml`](manifests/service-account/service-account.yaml) creates `name: my-service-account`, and `subjects[0]` should include `namespace: default`.
3. **[`manifests/configmap/deployment.yaml`](manifests/configmap/deployment.yaml) (Lines 29–39 — Unmatched `volumeMounts` Name):**
   - **Current:** `volumeMounts` mounts `name: pvc-volume` at `/data`, which is not defined under `spec.template.spec.volumes` (and defines volume `appconfig` which is unused).
   - **Fix:** Every entry in `containers[].volumeMounts[].name` must match a defined `spec.volumes[].name`, or the Pod fails validation.
4. **[`manifests/pvc/pod-with-pvc.yaml`](manifests/pvc/pod-with-pvc.yaml) (Line 30 — Mismatched `claimName`):**
   - **Current:** `persistentVolumeClaim.claimName: pvc`
   - **Fix:** [`manifests/pvc/pvc.yaml`](manifests/pvc/pvc.yaml) defines `metadata.name: my-pvc`, so `claimName` must be `my-pvc` (otherwise `pod-with-pvc` stays `Pending` forever waiting for a nonexistent PVC named `pvc`).

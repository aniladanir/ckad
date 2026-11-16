# The Complete CKAD 5-Week Masterclass & Hands-On Workbook
### November 16, 2026 – December 20, 2026 (35-Day Guided Bootcamp)

> **Who this book-length guide is built for:**
> You studied for CKAD a year ago (so the high-level concepts are already in your head), you haven't touched `kubectl` or written a YAML manifest in 12 months, and you have **1–2 hours on weekdays** and **3–4 hours on weekends** to get exam-ready by **December 20, 2026**.
>
> Instead of a dry checklist, this is a **complete, self-contained daily textbook + terminal lab manual**. Every single day includes:
> 1. **Deep-Dive Theory & Visual Mental Models** (ASCII architecture diagrams, annotated YAML anatomy, decision tables, and exam traps).
> 2. **Guided Step-by-Step Lab Walkthroughs** (exact commands, full manifests, **expected terminal outputs**, and "Break It & Fix It" experiments).
> 3. **Graded Exam-Style Scenarios** (realistic CKAD prompts with collapsible **Step-by-Step Solutions & Verification Scripts**).

---

## Master Table of Contents

- [Part 0: Exam Blueprint, Terminal Cockpit & Speed Engineering](#part-0-exam-blueprint-terminal-cockpit--speed-engineering)
  - [0.1 The CKAD Exam by the Numbers](#01-the-ckad-exam-by-the-numbers)
  - [0.2 Building Your Local Lab Cluster](#02-building-your-local-lab-cluster)
  - [0.3 Terminal Cockpit: Aliases, Autocomplete & Vim for YAML](#03-terminal-cockpit-aliases-autocomplete--vim-for-yaml)
  - [0.4 The 3 Golden Speed Workflows (`$do`, `explain`, `patch`/`edit`)](#04-the-3-golden-speed-workflows-do-explain-patchedit)
  - [0.5 The 2-Pass Exam Strategy & 15-Second Verification Ritual](#05-the-2-pass-exam-strategy--15-second-verification-ritual)
- [Week 1 (Nov 16–22): CLI Muscle Memory Reboot & Core Workloads](#week-1-nov-1622-cli-muscle-memory-reboot--core-workloads)
  - [Day 1 (Mon Nov 16): Pods, Namespaces & The Imperative YAML Factory](#day-1-mon-nov-16-pods-namespaces--the-imperative-yaml-factory)
  - [Day 2 (Tue Nov 17): Mastering `kubectl explain`, Labels, Selectors & Annotations](#day-2-tue-nov-17-mastering-kubectl-explain-labels-selectors--annotations)
  - [Day 3 (Wed Nov 18): Container Images, Multi-Stage Dockerfiles, `command`/`args` & Lifecycle Hooks](#day-3-wed-nov-18-container-images-multi-stage-dockerfiles-commandargs--lifecycle-hooks)
  - [Day 4 (Thu Nov 19): Multi-Container Pod Patterns (Init, Sidecar, Adapter, Ambassador & Native Sidecars)](#day-4-thu-nov-19-multi-container-pod-patterns-init-sidecar-adapter-ambassador--native-sidecars)
  - [Day 5 (Fri Nov 20): Batch Workloads — Jobs & CronJobs In-Depth](#day-5-fri-nov-20-batch-workloads--jobs--cronjobs-in-depth)
  - [Day 6 (Sat Nov 21 — Weekend Deep Lab): Volumes, PV, PVC, StorageClasses & Projected Volumes](#day-6-sat-nov-21--weekend-deep-lab-volumes-pv-pvc-storageclasses--projected-volumes)
  - [Day 7 (Sun Nov 22 — Weekend Boss Lab): Week 1 Synthesis Arena (5 Multi-Part Scenarios)](#day-7-sun-nov-22--weekend-boss-lab-week-1-synthesis-arena-5-multi-part-scenarios)
- [Week 2 (Nov 23–29): Application Deployment, Scheduling & Helm](#week-2-nov-2329-application-deployment-scheduling--helm)
  - [Day 8 (Mon Nov 23): Deployments, ReplicaSets, RollingUpdates & Rollbacks](#day-8-mon-nov-23-deployments-replicasets-rollingupdates--rollbacks)
  - [Day 9 (Tue Nov 24): Advanced Release Strategies — Recreate, Blue-Green & Canary](#day-9-tue-nov-24-advanced-release-strategies--recreate-blue-green--canary)
  - [Day 10 (Wed Nov 25): Pod Scheduling — `nodeSelector`, Node/Pod Affinity & Taints/Tolerations](#day-10-wed-nov-25-pod-scheduling--nodeselector-nodepod-affinity--taintstolerations)
  - [Day 11 (Thu Nov 26): Horizontal Pod Autoscaling (`autoscaling/v2`) & Metrics](#day-11-thu-nov-26-horizontal-pod-autoscaling-autoscalingv2--metrics)
  - [Day 12 (Fri Nov 27): Helm Package Manager — Charts, Values, Releases & Rollbacks](#day-12-fri-nov-27-helm-package-manager--charts-values-releases--rollbacks)
  - [Day 13 (Sat Nov 28 — Weekend Deep Lab): StatefulSets, Headless Services & DaemonSets](#day-13-sat-nov-28--weekend-deep-lab-statefulsets-headless-services--daemonsets)
  - [Day 14 (Sun Nov 29 — Weekend Boss Lab): Week 2 End-to-End Release Engineering Lab](#day-14-sun-nov-29--weekend-boss-lab-week-2-end-to-end-release-engineering-lab)
- [Week 3 (Nov 30–Dec 6): Configuration, Security & Observability (40% of Exam)](#week-3-nov-30dec-6-configuration-security--observability-40-of-exam)
  - [Day 15 (Mon Nov 30): ConfigMaps & Secrets — Env Vars, Volume Mounts, `subPath` & TLS/Registry Secrets](#day-15-mon-nov-30-configmaps--secrets--env-vars-volume-mounts-subpath--tlsregistry-secrets)
  - [Day 16 (Tue Dec 1): Resource Management — Requests/Limits, QoS, ResourceQuota, LimitRange & PriorityClass](#day-16-tue-dec-1-resource-management--requestslimits-qos-resourcequota-limitrange--priorityclass)
  - [Day 17 (Wed Dec 2): Container & Pod SecurityContexts, Linux Capabilities, `fsGroup` & PSS](#day-17-wed-dec-2-container--pod-securitycontexts-linux-capabilities-fsgroup--pss)
  - [Day 18 (Thu Dec 3): ServiceAccounts & Role-Based Access Control (RBAC)](#day-18-thu-dec-3-serviceaccounts--role-based-access-control-rbac)
  - [Day 19 (Fri Dec 4): Container Probes (`startup`, `readiness`, `liveness`) & Live Debugging (`logs`, `top`, `debug`)](#day-19-fri-dec-4-container-probes-startup-readiness-liveness--live-debugging-logs-top-debug)
  - [Day 20 (Sat Dec 5 — Weekend Clinic): The 7-Scenario Broken Cluster Triage Clinic + CRDs & API Deprecations](#day-20-sat-dec-5--weekend-clinic-the-7-scenario-broken-cluster-triage-clinic--crds--api-deprecations)
  - [Day 21 (Sun Dec 6 — Weekend Boss Lab): Week 3 Hardened Microservice Architecture Exam](#day-21-sun-dec-6--weekend-boss-lab-week-3-hardened-microservice-architecture-exam)
- [Week 4 (Dec 7–13): Services, Networking & Full Mock Exam #1](#week-4-dec-713-services-networking--full-mock-exam-1)
  - [Day 22 (Mon Dec 7): Kubernetes Networking Model, CoreDNS & Service Types (`ClusterIP`, `NodePort`, `LoadBalancer`, `ExternalName`)](#day-22-mon-dec-7-kubernetes-networking-model-coredns--service-types-clusterip-nodeport-loadbalancer-externalname)
  - [Day 23 (Tue Dec 8): Ingress Controllers, `IngressClass`, Path/Host Routing & HTTPS/TLS Termination](#day-23-tue-dec-8-ingress-controllers-ingressclass-pathhost-routing--httpstls-termination)
  - [Day 24 (Wed Dec 9): NetworkPolicies Part 1 — Default Isolation, Ingress Rules & The `AND` vs `OR` Selector Trap](#day-24-wed-dec-9-networkpolicies-part-1--default-isolation-ingress-rules--the-and-vs-or-selector-trap)
  - [Day 25 (Thu Dec 10): NetworkPolicies Part 2 — Egress Rules, `ipBlock` CIDRs & The CoreDNS Port 53 Trap](#day-25-thu-dec-10-networkpolicies-part-2--egress-rules-ipblock-cidrs--the-coredns-port-53-trap)
  - [Day 26 (Fri Dec 11): Advanced `kubectl` Extraction — JSONPath, Custom Columns, Field Selectors & Sorting](#day-26-fri-dec-11-advanced-kubectl-extraction--jsonpath-custom-columns-field-selectors--sorting)
  - [Day 27 (Sat Dec 12 — Weekend Mock #1): Full 15-Task Simulated CKAD Exam #1 (Built-In) + Killer.sh Session 1](#day-27-sat-dec-12--weekend-mock-1-full-15-task-simulated-ckad-exam-1-built-in--killersh-session-1)
  - [Day 28 (Sun Dec 13 — Weekend Post-Mortem): Mock #1 Deep Re-Solve & Gap Elimination](#day-28-sun-dec-13--weekend-post-mortem-mock-1-deep-re-solve--gap-elimination)
- [Week 5 (Dec 14–20): High-Speed Clinics, Mock Exam #2 & Exam Day](#week-5-dec-1420-high-speed-clinics-mock-exam-2--exam-day)
  - [Day 29 (Mon Dec 14): High-Speed Clinic #1 — NetworkPolicy + Ingress + Service Gauntlet](#day-29-mon-dec-14-high-speed-clinic-1--networkpolicy--ingress--service-gauntlet)
  - [Day 30 (Tue Dec 15): High-Speed Clinic #2 — RBAC + SecurityContext + Volumes/PVC Gauntlet](#day-30-tue-dec-15-high-speed-clinic-2--rbac--securitycontext--volumespvc-gauntlet)
  - [Day 31 (Wed Dec 16): Full Mock Exam #2 (Killer.sh Session 2 or Split 1h+1h)](#day-31-wed-dec-16-full-mock-exam-2-killersh-session-2-or-split-1h1h)
  - [Day 32 (Thu Dec 17): Mock #2 Post-Mortem & Surgical Fixes](#day-32-thu-dec-17-mock-2-post-mortem--surgical-fixes)
  - [Day 33 (Fri Dec 18): The 15-Command Muscle-Memory Speed Run](#day-33-fri-dec-18-the-15-command-muscle-memory-speed-run)
  - [Day 34 (Sat Dec 19): Final Dress Rehearsal, Logistics & Mental Rest](#day-34-sat-dec-19-final-dress-rehearsal-logistics--mental-rest)
  - [Day 35 (Sun Dec 20): CKAD EXAM DAY Game Plan](#day-35-sun-dec-20-ckad-exam-day-game-plan)
- [Appendix: Personal Mistake & Speed Tracker](#appendix-personal-mistake--speed-tracker)

---

# Part 0: Exam Blueprint, Terminal Cockpit & Speed Engineering

## 0.1 The CKAD Exam by the Numbers

- **Format:** 100% hands-on command-line tasks inside a remote Linux desktop (XFCE) in the PSI Secure Browser.
- **Duration:** **120 minutes** for **15–20 tasks** (~6–7 minutes per task on average).
- **Passing Score:** **66%**.
- **Allowed Resources During the Exam:** Official documentation at `https://kubernetes.io/docs/`, `https://kubernetes.io/blog/`, and `https://helm.sh/docs/`.

```text
+-------------------------------------------------------------------------+
|                        CKAD CURRICULUM WEIGHTS                          |
+---------------------------------------------------+---------------------+
| 1. Application Environment, Config & Security     | 25%  (1/4 of exam!) |
| 2. Application Design and Build                   | 20%                 |
| 3. Application Deployment                         | 20%                 |
| 4. Services and Networking                        | 20%                 |
| 5. Application Observability and Maintenance      | 15%                 |
+---------------------------------------------------+---------------------+
```

---

## 0.2 Building Your Local Lab Cluster

To practice everything in this guide (including **NetworkPolicies**, **HPA / `kubectl top`**, and **Ingress**), set up `minikube` with the Calico CNI plugin:

```bash
# Start a 2-node cluster (having 2 nodes lets you test nodeSelector, Affinity, Taints, DaemonSets, and RWO PVCs!)
minikube start --nodes=2 --cpus=2 --memory=3072 --cni=calico

# Enable Metrics Server (for HPA & kubectl top) and NGINX Ingress Controller
minikube addons enable metrics-server
minikube addons enable ingress

# Confirm nodes and system pods are healthy
kubectl get nodes
kubectl get pods -n kube-system
```

---

## 0.3 Terminal Cockpit: Aliases, Autocomplete & Vim for YAML

Since you haven't written YAML or run `kubectl` in a year, **bad editor defaults will cost you 20+ minutes on exam day**. Add these to your `~/.bashrc` (or `~/.zshrc`) right now and use them in every session:

```bash
alias k=kubectl
source <(kubectl completion bash)
complete -o default -F __start_kubectl k

export do="--dry-run=client -o yaml"
export now="--force --grace-period=0"
```

And configure `~/.vimrc` so pressing `Tab` inserts 2 spaces and never mixes tabs/spaces:

```vim
set tabstop=2
set shiftwidth=2
set expandtab
set autoindent
set number
set ruler
```

**Essential Vim keystrokes to practice daily:**
- `:set paste` (before pasting YAML from docs so autoindent doesn't staircase-indent it!) -> `:set nopaste`
- `V` (visual line mode) + `j`/`k` (select lines) + `>` or `<` (indent/unindent whole block by 2 spaces!) -> `.` (repeat indent)
- `dd` (delete line), `dG` (delete from cursor to end of file), `u` (undo), `Ctrl+r` (redo), `/pattern` (search).

---

## 0.4 The 3 Golden Speed Workflows (`$do`, `explain`, `patch`/`edit`)

1. **Workflow 1 — Generate, Edit, Apply (`$do`):**
   Never type `apiVersion`, `kind`, `metadata` from scratch. Let `kubectl` generate 80% of the YAML:
   ```bash
   k create deploy web --image=nginx:1.25-alpine --replicas=3 $do > web.yaml
   vim web.yaml
   k apply -f web.yaml
   ```
2. **Workflow 2 — Instant Schema Lookup (`kubectl explain`):**
   Forgot whether `securityContext` uses `runAsNonRoot` or `nonRoot`, or where `fsGroup` lives?
   ```bash
   k explain pod.spec.securityContext
   k explain pod.spec.containers.securityContext.capabilities
   k explain netpol.spec.ingress.from --recursive
   ```
3. **Workflow 3 — Modifying Existing Live Objects:**
   - If the field is **mutable** (e.g., Deployment `image`, `replicas`, Service `selector`, `port`): use `k edit deploy <name>` or `k set image` / `k scale`.
   - If the field is **immutable** on a running Pod (e.g., Pod `volumeMounts`, `env`, `securityContext` on a standalone Pod):
     ```bash
     k get pod my-pod -o yaml > my-pod.yaml
     vim my-pod.yaml
     k replace --force -f my-pod.yaml    # Deletes and recreates the Pod in one command!
     ```

---

## 0.5 The 2-Pass Exam Strategy & 15-Second Verification Ritual

```text
 [Question Start]
        │
        ▼
 1. Run `kubectl config use-context <ctx>` (copied directly from prompt!)
 2. Check target `-n <namespace>`
 3. Can I solve this in < 5 mins?
        ├── YES ──► Generate YAML ($do) -> Edit -> Apply -> 15s Verify -> NEXT
        └── NO / STUCK AT 6 MINS ──► Flag Question -> Move on (Return in Pass 2!)
```

**The 15-Second Verification Ritual (Run after EVERY task):**
1. `k get pod,svc,deploy -n <target-namespace>` -> Is the resource actually in the right namespace and `Running` (`1/1` Ready)?
2. `k get endpoints <svc-name> -n <target-namespace>` -> If you touched a Service or NetworkPolicy, are Pod IPs listed under `ENDPOINTS`?
3. `k logs <pod> -n <target-namespace>` -> Did the container command run without syntax errors?


---

# Week 1 (Nov 16–22): CLI Muscle Memory Reboot & Core Workloads

**Weekly Objective:** Rebuild your `kubectl` and YAML editing reflexes from zero to exam speed, and master **Application Design and Build (20% of CKAD)**: Pods, Labels/Selectors, Container Images, Multi-Container Patterns, Jobs/CronJobs, and Persistent Storage.

---

## Day 1 (Mon Nov 16): Pods, Namespaces & The Imperative YAML Factory

### 1. Deep-Dive Theory & Mental Model

#### What Is a Pod Really?
A **Pod** is an atomic wrapper around one or more Linux containers that share the exact same **Network Namespace** (one IP address, one `localhost` loopback, shared TCP/UDP port space) and optional **Storage Volumes**, scheduled together onto a single worker Node.

```text
+-------------------------------------------------------------------+
| Node (Worker VM)                                                  |
|  +-------------------------------------------------------------+  |
|  | Pod (IP: 10.244.1.15)                                       |  |
|  |  +------------------------+   +--------------------------+  |  |
|  |  | Container A (nginx:80) |   | Container B (sidecar)    |  |  |
|  |  | Reaches B on localhost |   | Reaches A on localhost:80|  |  |
|  |  +-----------+------------+   +------------+-------------+  |  |
|  |              |                             |                |  |
|  |              +------> [Shared Volume] <----+                |  |
|  +-------------------------------------------------------------+  |
+-------------------------------------------------------------------+
```

#### Annotated Pod YAML Anatomy
Every Pod manifest has **4 top-level keys**: `apiVersion`, `kind`, `metadata`, and `spec`.

```yaml
apiVersion: v1                     # Core API group v1
kind: Pod                          # Resource type
metadata:
  name: payment-api                # Unique name within the namespace
  namespace: finance               # Defaults to 'default' if omitted!
  labels:                          # Key-value tags used by Services/Selectors
    app: payment
    tier: backend
spec:
  restartPolicy: Always            # Always (default) | OnFailure | Never
  containers:
  - name: api                      # Container name inside the Pod
    image: nginx:1.25-alpine       # Image repository:tag
    imagePullPolicy: IfNotPresent  # Always | IfNotPresent | Never
    ports:
    - containerPort: 80            # Informational port exposed by container
    env:                           # Static environment variables
    - name: LOG_FORMAT
      value: "json"
```

#### Imperative vs Declarative Cheat Table for Pods & Namespaces

| Goal | Fastest Imperative Command |
| :--- | :--- |
| Create a Namespace | `k create ns finance` |
| Switch default namespace in current context | `k config set-context --current --namespace=finance` |
| Run a Pod immediately | `k run web --image=nginx:1.25-alpine -n finance` |
| Run a Pod with labels, port & env | `k run web --image=nginx:1.25-alpine --port=80 --labels="app=web,env=prod" --env="MODE=live"` |
| Export Pod YAML without creating it | `k run web --image=nginx:1.25-alpine $do > web.yaml` |
| Delete a Pod immediately (0s wait) | `k delete pod web $now` |
| Replace an immutable field on a running Pod | `k get pod web -o yaml > web.yaml && vim web.yaml && k replace --force -f web.yaml` |

> **Exam Trap #1 — The Missing `-n` Flag:**
> If an exam prompt says *"In namespace `mercury`, create a Pod..."* and you create it in `default`, the grading script scores **0 points**. Always either pass `-n mercury` on your `k run` command (which bakes `namespace: mercury` right into the generated YAML!) or verify with `k get pods -n mercury`.

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 1A: Generating, Customizing & Applying Pod Manifests
Run these commands in your terminal and compare with the expected outputs:

```bash
# Step 1: Create a dedicated namespace for today's walkthrough
k create ns day1-lab

# Step 2: Generate a Pod YAML skeleton with port, labels, and env vars already populated!
k run billing-pod -n day1-lab \
  --image=nginx:1.25-alpine \
  --port=8080 \
  --labels="app=billing,tier=api" \
  --env="CURRENCY=USD" \
  --env="REGION=eu-central" \
  $do > /tmp/opencode/billing-pod.yaml

# Step 3: Inspect the generated file
cat /tmp/opencode/billing-pod.yaml
```
**Expected Output:**
```yaml
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    app: billing
    tier: api
  name: billing-pod
  namespace: day1-lab
spec:
  containers:
  - env:
    - name: CURRENCY
      value: USD
    - name: REGION
      value: eu-central
    image: nginx:1.25-alpine
    name: billing-pod
    ports:
    - containerPort: 8080
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Always
status: {}
```

Now apply it and inspect its live state:
```bash
k apply -f /tmp/opencode/billing-pod.yaml
k get pods -n day1-lab -o wide --show-labels
```
**Expected Output:**
```text
NAME          READY   STATUS    RESTARTS   AGE   IP            NODE       NOMINATED NODE   READINESS GATES   LABELS
billing-pod   1/1     Running   0          8s    10.244.1.12   minikube   <none>           <none>            app=billing,tier=api
```

#### Lab 1B: Break It & Fix It — Modifying an Immutable Pod Field
On CKAD, you will often be asked to change an environment variable, volume mount, or command on an **already running standalone Pod**. Let's see what happens if you try `kubectl edit`:

```bash
# Try adding a new env var via apply after editing /tmp/opencode/billing-pod.yaml:
sed -i 's/value: USD/value: EUR/' /tmp/opencode/billing-pod.yaml
k apply -f /tmp/opencode/billing-pod.yaml
```
**Expected Error:**
```text
The Pod "billing-pod" is invalid: spec: Forbidden: pod updates may not change fields other than `spec.containers[*].image`, `spec.initContainers[*].image`, `spec.activeDeadlineSeconds`, `spec.tolerations` (only additions to existing tolerations), `spec.terminationGracePeriodSeconds`
```
**The Fix (`k replace --force -f`):**
```bash
k replace --force -f /tmp/opencode/billing-pod.yaml
k exec -n day1-lab billing-pod -- printenv CURRENCY
# Output: EUR
k delete ns day1-lab
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 1.1: Warm-Up Lab Reboot (Tasks 1–6)
Apply [`warmup-exam/setup.yaml`](warmup-exam/setup.yaml) and complete **Tasks 1–6** in [`warmup-exam/CKAD_WARMUP_EXAM.html`](warmup-exam/CKAD_WARMUP_EXAM.html).

#### Scenario 1.2: Standalone Pod Surgery (Exam Weight: 5%)
1. Create namespace `jupiter`.
2. In `jupiter`, create a Pod named `cache-node` running `redis:7.2-alpine` with label `role=cache` and environment variable `ALLOW_EMPTY_PASSWORD=yes`.
3. Once `cache-node` is `Running`, export its YAML to `/tmp/opencode/cache-node.yaml`, change its container name from `cache-node` to `redis-engine`, add a second environment variable `REDIS_PORT=6379`, and replace the running Pod.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 1.2)</strong></summary>

```bash
# 1. Create namespace and initial Pod
k create ns jupiter
k run cache-node -n jupiter --image=redis:7.2-alpine --labels="role=cache" --env="ALLOW_EMPTY_PASSWORD=yes" $do > /tmp/opencode/cache-node.yaml
k apply -f /tmp/opencode/cache-node.yaml
k wait --for=condition=Ready pod/cache-node -n jupiter --timeout=60s

# 2. Update container name and add env var in /tmp/opencode/cache-node.yaml:
cat << 'EOF' > /tmp/opencode/cache-node.yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    role: cache
  name: cache-node
  namespace: jupiter
spec:
  containers:
  - name: redis-engine
    image: redis:7.2-alpine
    env:
    - name: ALLOW_EMPTY_PASSWORD
      value: "yes"
    - name: REDIS_PORT
      value: "6379"
EOF

# 3. Force replace the Pod
k replace --force -f /tmp/opencode/cache-node.yaml

# 4. Verify container name and env vars
k get pod cache-node -n jupiter -o jsonpath='{.spec.containers[0].name}{"\n"}'
k exec -n jupiter cache-node -c redis-engine -- printenv REDIS_PORT
k delete ns jupiter
```
</details>

---

## Day 2 (Tue Nov 17): Mastering `kubectl explain`, Labels, Selectors & Annotations

### 1. Deep-Dive Theory & Mental Model

#### Labels vs Annotations
Both are key-value maps inside `metadata`, but they serve completely different purposes in Kubernetes:

| Feature | `metadata.labels` | `metadata.annotations` |
| :--- | :--- | :--- |
| **Purpose** | **Identifying metadata** used to group, select, and route to objects | **Non-identifying metadata** read by tools, controllers, or humans |
| **Can Services / Deployments / NetworkPolicies select on them?** | **YES** (`-l app=web`, `matchLabels`, `matchExpressions`) | **NO** (cannot select objects by annotation) |
| **Size / Character constraints** | Strict (`<= 63` chars for value, alphanumeric + `-_.`) | Flexible (can store JSON blobs, URLs, build SHAs, Ingress rules, `change-cause`) |
| **Imperative CLI** | `k label pod web env=prod --overwrite` / `k label pod web env-` | `k annotate pod web owner="team-a"` / `k annotate pod web owner-` |

#### Equality-Based vs Set-Based Label Selectors
- **Equality-based (`=`, `==`, `!=`):**
  ```bash
  k get pods -l env=prod,tier!=frontend
  ```
- **Set-based (`in`, `notin`, `exists`):**
  ```bash
  k get pods -l 'env in (prod, staging),tier notin (legacy),release'
  ```
  *(Above matches Pods where `env` is `prod` or `staging`, `tier` is not `legacy`, and the label key `release` exists with any value!)*

#### Using `kubectl explain` Like a Cheat Code
Every resource schema in the API server is self-documented. You never need to guess whether a field is an array (`[]Object`) or an object (`Object`) or a string (`string`):
- `<[]Object>` means **YAML list** (needs a leading `- ` dash!).
- `<Object>` means **nested YAML map** (indent by 2 spaces, **no** `- ` dash!).

```bash
k explain pod.spec.volumes                        # Shows every volume type supported by the cluster!
k explain pod.spec.containers --recursive         # Prints the full tree of container fields
k explain deploy.spec.strategy.rollingUpdate      # Shows maxSurge & maxUnavailable docs
```

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 2A: Live Label Filtering, Bulk Operations & Annotations
```bash
k create ns day2-lab
k run app-1 -n day2-lab --image=nginx:1.25-alpine --labels="app=store,env=prod,tier=web"
k run app-2 -n day2-lab --image=nginx:1.25-alpine --labels="app=store,env=staging,tier=web"
k run db-1  -n day2-lab --image=redis:7.2-alpine  --labels="app=store,env=prod,tier=db"

# 1. Query only prod web pods
k get pods -n day2-lab -l env=prod,tier=web --show-labels

# 2. Query using set-based selector
k get pods -n day2-lab -l 'env in (prod,staging),tier!=db'

# 3. Overwrite an existing label and add an annotation to all 'tier=web' pods at once!
k label pods -n day2-lab -l tier=web tier=frontend --overwrite
k annotate pods -n day2-lab -l tier=frontend commit-sha="a81f92c"

# 4. Remove the 'commit-sha' annotation using the trailing minus sign (-)
k annotate pods -n day2-lab -l tier=frontend commit-sha-
k delete ns day2-lab
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 2.1: Finish the Warm-Up Lab (Tasks 7–12) + Baseline MCQ Quiz
1. Complete **Tasks 7–12** of [`warmup-exam/CKAD_WARMUP_EXAM.html`](warmup-exam/CKAD_WARMUP_EXAM.html) and delete `ckad-warmup`.
2. Open [`quiz/CKAD_QUIZ_MCQ.html`](quiz/CKAD_QUIZ_MCQ.html) in your browser and complete all 25 questions to establish your baseline score.

#### Scenario 2.2: Label Surgery & Bulk Selection (Exam Weight: 4%)
In namespace `inventory`:
1. Create 3 Pods (`item-a`, `item-b`, `item-c` using `busybox:1.36` running `sleep 3600`):
   - `item-a`: `region=us,env=prod,version=v1`
   - `item-b`: `region=eu,env=prod,version=v2`
   - `item-c`: `region=us,env=dev,version=v1`
2. Using a single `kubectl` command, add label `audited=true` to all Pods in `inventory` that are in `env=prod` AND have `region` in `(us, eu)`.
3. Delete `version` label from `item-c`, and add annotation `reviewed-by="sec-ops"` to `item-a`.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 2.2)</strong></summary>

```bash
k create ns inventory
k run item-a -n inventory --image=busybox:1.36 --labels="region=us,env=prod,version=v1" -- sleep 3600
k run item-b -n inventory --image=busybox:1.36 --labels="region=eu,env=prod,version=v2" -- sleep 3600
k run item-c -n inventory --image=busybox:1.36 --labels="region=us,env=dev,version=v1"  -- sleep 3600

# Bulk label matching pods
k label pods -n inventory -l 'env=prod,region in (us,eu)' audited=true

# Remove label 'version' from item-c and annotate item-a
k label pod item-c -n inventory version-
k annotate pod item-a -n inventory reviewed-by="sec-ops"

# Verify
k get pods -n inventory --show-labels
k describe pod item-a -n inventory | grep Annotations -A2
k delete ns inventory
```
</details>

---

## Day 3 (Wed Nov 18): Container Images, Multi-Stage Dockerfiles, `command`/`args` & Lifecycle Hooks

### 1. Deep-Dive Theory & Mental Model

#### How Dockerfile `ENTRYPOINT` & `CMD` Interact with Kubernetes `command` & `args`
This is one of the most frequently tested concepts on CKAD. Memorize this table:

```text
+----------------------------------+----------------------------------+
| Dockerfile Instruction           | Kubernetes Pod YAML Equivalent   |
+----------------------------------+----------------------------------+
| ENTRYPOINT ["python3", "app.py"] | command: ["python3", "app.py"]   |
| CMD ["--port", "8080"]           | args: ["--port", "8080"]         |
+----------------------------------+----------------------------------+
```

**The 4 Combination Rules:**
1. **Neither `command` nor `args` in Pod YAML:** Runs `ENTRYPOINT` + `CMD` from the Dockerfile (`python3 app.py --port 8080`).
2. **Only `args` defined in Pod YAML (`args: ["--port", "9090"]`):** Keeps Dockerfile `ENTRYPOINT` and replaces `CMD` (`python3 app.py --port 9090`).
3. **Only `command` defined in Pod YAML (`command: ["sh", "-c", "date"]`):** Replaces **BOTH** Dockerfile `ENTRYPOINT` and `CMD`! (`CMD` from the Dockerfile is completely ignored).
4. **Both `command` and `args` defined in Pod YAML:** Runs `command` followed by `args`.

#### Container Lifecycle Hooks & Graceful Termination Sequence
When a Pod is deleted (`kubectl delete pod`), Kubernetes executes this exact sequence in parallel:

```text
[kubectl delete pod]
        │
        ├──► 1. Control Plane marks Pod "Terminating" -> Kube-proxy removes Pod IP from Service Endpoints
        │
        └──► 2. Kubelet starts terminationGracePeriodSeconds timer (default 30s):
                ├── Step 2a: Runs container `lifecycle.preStop` hook synchronously
                ├── Step 2b: Sends SIGTERM (PID 1 in container)
                └── Step 2c: If still running when timer hits 0 -> Sends SIGKILL (immediate kill)
```

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 3A: Building a Multi-Stage Container Image & Exporting to `.tar`
Inspect the existing [`probes/Dockerfile`](probes/Dockerfile) in your repo, then build a custom Go utility image from scratch:

```bash
mkdir -p /tmp/opencode/day3-build && cd /tmp/opencode/day3-build

cat << 'EOF' > main.go
package main
import (
    "fmt"
    "os"
)
func main() {
    args := os.Args[1:]
    fmt.Printf("Binary running with args: %v\n", args)
}
EOF

cat << 'EOF' > Dockerfile
# Stage 1: Build binary
FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY main.go .
RUN CGO_ENABLED=0 go build -o /out/runner main.go

# Stage 2: Minimal runtime image
FROM alpine:3.20
RUN adduser -D -u 10001 appuser
USER 10001
COPY --from=builder /out/runner /usr/local/bin/runner
ENTRYPOINT ["/usr/local/bin/runner"]
CMD ["--default-mode"]
EOF

# Build, tag, and save as OCI/Docker tar archive (classic CKAD task!)
docker build -t runner:1.0 .
docker tag runner:1.0 registry.example.com/tools/runner:1.0
docker save -o /tmp/opencode/runner-1.0.tar runner:1.0
ls -lh /tmp/opencode/runner-1.0.tar
cd /home/a00876142/ckad-session/ckad
```

#### Lab 3B: Testing `command`, `args`, `postStart` & `preStop` Live
```yaml
# Save as /tmp/opencode/hooks-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: hooks-pod
spec:
  terminationGracePeriodSeconds: 20
  restartPolicy: Never
  containers:
  - name: web
    image: nginx:1.25-alpine
    lifecycle:
      postStart:
        exec:
          command: ["sh", "-c", "echo 'Started at $(date)' > /usr/share/nginx/html/status.txt"]
      preStop:
        exec:
          command: ["sh", "-c", "nginx -s quit; sleep 3"]
```
```bash
k apply -f /tmp/opencode/hooks-pod.yaml
k wait --for=condition=Ready pod/hooks-pod --timeout=30s
k exec hooks-pod -- cat /usr/share/nginx/html/status.txt
# Expected Output: Started at ...
k delete pod hooks-pod
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 3.1: Overriding `ENTRYPOINT` vs `CMD` (Exam Weight: 6%)
In namespace `default`, create two Pods using `busybox:1.36` (`restartPolicy: Never`):
1. `pod-cmd-only`: Uses `command: ["sh", "-c"]` and `args: ["echo 'Hello from args' && sleep 120"]`.
2. `pod-imperative`: Generated purely via `kubectl run pod-imperative --image=busybox:1.36 --restart=Never --command -- sh -c "echo 'Imperative command' && sleep 120"`.
Compare `k get pod pod-cmd-only -o yaml` and `k get pod pod-imperative -o yaml` to see how `--command` vs no `--command` populates `command:` vs `args:`!

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 3.1)</strong></summary>

- When you run `k run mypod --image=busybox -- sh -c "sleep 10"` (**without** `--command`), `kubectl` sets `args: ["sh", "-c", "sleep 10"]`.
- When you run `k run mypod --image=busybox --command -- sh -c "sleep 10"` (**with** `--command`), `kubectl` sets `command: ["sh", "-c", "sleep 10"]`.
```bash
k delete pod pod-cmd-only pod-imperative $now
```
</details>

#### Scenario 3.2: Graceful Drain Hook (Exam Weight: 6%)
Create a Pod named `graceful-worker` in namespace `default` (`image: busybox:1.36`) with:
- Main command: `["sh", "-c", "echo Working... && sleep 3600"]`
- `terminationGracePeriodSeconds: 15`
- A `preStop` exec hook that writes `"Draining queues"` to `/tmp/drain.log` and sleeps for `4` seconds.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 3.2)</strong></summary>

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: graceful-worker
spec:
  terminationGracePeriodSeconds: 15
  containers:
  - name: worker
    image: busybox:1.36
    command: ["sh", "-c", "echo Working... && sleep 3600"]
    lifecycle:
      preStop:
        exec:
          command: ["sh", "-c", "echo 'Draining queues' > /tmp/drain.log && sleep 4"]
```
```bash
k apply -f graceful-worker.yaml
k describe pod graceful-worker
k delete pod graceful-worker $now
```
</details>

---

## Day 4 (Thu Nov 19): Multi-Container Pod Patterns (Init, Sidecar, Adapter, Ambassador & Native Sidecars)

### 1. Deep-Dive Theory & Mental Model

#### Execution Order Inside a Pod
```text
 [Pod Scheduled to Node]
          │
          ▼
 +-------------------------------------------------------------------------+
 | 1. initContainers[0] (Normal Init)                                      |
 |    Runs to completion (Exit 0) before next initContainer starts.        |
 +-------------------------------------------------------------------------+
          │
          ▼
 +-------------------------------------------------------------------------+
 | 2. initContainers[1] (Native Sidecar -> has `restartPolicy: Always`)    |
 |    Starts before main containers, stays running alongside them!         |
 +-------------------------------------------------------------------------+
          │
          ▼
 +-------------------------------------------------------------------------+
 | 3. containers[0..N] (Main App + Traditional Sidecars/Adapters/Proxies)  |
 |    All start in parallel. Share network (localhost) & mounted Volumes.  |
 +-------------------------------------------------------------------------+
```

#### Comparing the 4 Multi-Container Patterns

| Pattern | Role | Classic CKAD Exam Scenario |
| :--- | :--- | :--- |
| **InitContainer** | Pre-boot setup / gating | Wait for a DB service via `nc -z` or `nslookup`, or clone/generate a config file into a shared `emptyDir` before the main web server starts. |
| **Sidecar** | Log shipping / file sync | Main app writes logs to `/var/log/app.log` on an `emptyDir`; Sidecar runs `tail -n+1 -F /var/log/app.log` so `kubectl logs <pod> -c sidecar` streams them. |
| **Adapter** | Output normalization | Main app writes legacy metrics format to `/tmp/stats.txt`; Adapter reads `/tmp/stats.txt`, reformats to Prometheus syntax, and serves on port `9100`. |
| **Ambassador** | Outbound network proxy | Main app connects to `localhost:6379`; Ambassador container listens on `6379` and proxies/encrypts traffic to an external Redis cluster. |

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 4A: Building an InitContainer + Main App + Adapter + Ambassador Pod
Let's build a 3-stage Pod demonstrating **InitContainer** (pre-check), **Main App** (writes raw status), and **Adapter** (transforms and serves over HTTP):

```yaml
# Save as /tmp/opencode/adapter-demo.yaml
apiVersion: v1
kind: Pod
metadata:
  name: adapter-demo
spec:
  initContainers:
  - name: init-dir
    image: busybox:1.36
    command: ["sh", "-c", "mkdir -p /data && echo 'INIT_OK' > /data/init.status"]
    volumeMounts:
    - name: shared-data
      mountPath: /data

  containers:
  # Main container: writes legacy pipe-delimited metrics every 3 seconds
  - name: legacy-app
    image: busybox:1.36
    command:
    - sh
    - -c
    - |
      i=1
      while true; do
        echo "requests_total|$i" > /data/raw.metric
        i=$((i+1))
        sleep 3
      done
    volumeMounts:
    - name: shared-data
      mountPath: /data

  # Adapter container: converts 'requests_total|N' into Prometheus 'requests_total N'
  - name: prom-adapter
    image: busybox:1.36
    command:
    - sh
    - -c
    - |
      while true; do
        if [ -f /data/raw.metric ]; then
          tr '|' ' ' < /data/raw.metric > /data/metrics.prom
        fi
        sleep 3
      done
    volumeMounts:
    - name: shared-data
      mountPath: /data

  volumes:
  - name: shared-data
    emptyDir: {}
```

Apply and verify how the three containers interact through `/data`:
```bash
k apply -f /tmp/opencode/adapter-demo.yaml
sleep 5
k exec adapter-demo -c prom-adapter -- cat /data/init.status
# Output: INIT_OK
k exec adapter-demo -c prom-adapter -- cat /data/raw.metric
# Output: requests_total|2
k exec adapter-demo -c prom-adapter -- cat /data/metrics.prom
# Output: requests_total 2
k delete pod adapter-demo $now
```

#### Lab 4B: Break It & Fix It — InitContainer Blocking Pod Startup (`Init:0/1`)
If an `initContainer` fails or waits forever on a service that doesn't exist yet, the Pod status shows `Init:0/1` or `Init:CrashLoopBackOff` and main containers **never start**:
```bash
k run blocked-pod --image=nginx:1.25-alpine $do > /tmp/opencode/blocked.yaml
# Add an initContainer waiting for a Service named 'db-ready-svc'
cat << 'EOF' > /tmp/opencode/blocked.yaml
apiVersion: v1
kind: Pod
metadata:
  name: blocked-pod
spec:
  initContainers:
  - name: wait-for-db
    image: busybox:1.36
    command: ["sh", "-c", "until nslookup db-ready-svc; do echo waiting for db; sleep 2; done"]
  containers:
  - name: web
    image: nginx:1.25-alpine
EOF

k apply -f /tmp/opencode/blocked.yaml
k get pod blocked-pod
# STATUS: Init:0/1
k logs blocked-pod -c wait-for-db
# Now create the missing Service and watch blocked-pod automatically transition to Running!
k create service clusterip db-ready-svc --tcp=5432:5432
k get pod blocked-pod -w
k delete pod blocked-pod $now && k delete svc db-ready-svc
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 4.1: Retrofitting a Logging Sidecar onto an Existing Deployment (Exam Weight: 8%)
1. Create a Deployment named `legacy-payment` (`replicas: 1`, `image: busybox:1.36`) whose container `payment` writes `"$(date) TX_OK"` every 4 seconds into `/var/log/payment/tx.log`.
2. Update the `legacy-payment` Deployment to add:
   - An `emptyDir` volume named `log-storage` mounted at `/var/log/payment` in both containers.
   - A sidecar container named `log-forwarder` (`image: busybox:1.36`) that runs `sh -c "tail -n+1 -F /var/log/payment/tx.log"`.
3. Verify `k logs deploy/legacy-payment -c log-forwarder` outputs the transaction lines.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 4.1)</strong></summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: legacy-payment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: legacy-payment
  template:
    metadata:
      labels:
        app: legacy-payment
    spec:
      containers:
      - name: payment
        image: busybox:1.36
        command: ["sh", "-c", "while true; do echo \"$(date) TX_OK\" >> /var/log/payment/tx.log; sleep 4; done"]
        volumeMounts:
        - name: log-storage
          mountPath: /var/log/payment
      - name: log-forwarder
        image: busybox:1.36
        command: ["sh", "-c", "tail -n+1 -F /var/log/payment/tx.log"]
        volumeMounts:
        - name: log-storage
          mountPath: /var/log/payment
      volumes:
      - name: log-storage
        emptyDir: {}
```
```bash
k apply -f legacy-payment.yaml
k rollout status deploy/legacy-payment
k logs deploy/legacy-payment -c log-forwarder
k delete deploy legacy-payment
```
</details>

#### Scenario 4.2: Ambassador Localhost Proxy Pod (Exam Weight: 7%)
Create a Pod named `ambassador-pod` in namespace `default`:
- Main container `client` (`busybox:1.36`) curls `http://localhost:8080` every 5 seconds and prints the output to stdout.
- Ambassador container `proxy` (`nginx:1.25-alpine`) listens on port `8080` (mounted via a ConfigMap `nginx-proxy-conf` at `/etc/nginx/conf.d/default.conf` that returns `200 "Proxied via Ambassador\n"` on port `8080`).

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 4.2)</strong></summary>

```bash
cat << 'EOF' > /tmp/opencode/proxy.conf
server {
    listen 8080;
    location / {
        return 200 "Proxied via Ambassador\n";
    }
}
EOF
k create configmap nginx-proxy-conf --from-file=default.conf=/tmp/opencode/proxy.conf
```
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: ambassador-pod
spec:
  containers:
  - name: client
    image: busybox:1.36
    command: ["sh", "-c", "while true; do wget -qO- http://localhost:8080; sleep 5; done"]
  - name: proxy
    image: nginx:1.25-alpine
    volumeMounts:
    - name: conf
      mountPath: /etc/nginx/conf.d
  volumes:
  - name: conf
    configMap:
      name: nginx-proxy-conf
```
```bash
k apply -f ambassador-pod.yaml
sleep 6
k logs ambassador-pod -c client
# Output: Proxied via Ambassador
k delete pod ambassador-pod $now && k delete cm nginx-proxy-conf
```
</details>

---

## Day 5 (Fri Nov 20): Batch Workloads — Jobs & CronJobs In-Depth

### 1. Deep-Dive Theory & Mental Model

#### Annotated CronJob & Job Nesting Hierarchy
The #1 mistake people make with `CronJob` YAML is getting lost in the **3 levels of `.spec`**:
1. `CronJob.spec` -> Schedule & concurrency rules (`schedule`, `concurrencyPolicy`, `startingDeadlineSeconds`, `successfulJobsHistoryLimit`, `failedJobsHistoryLimit`, `suspend`, `jobTemplate`).
2. `CronJob.spec.jobTemplate.spec` -> Job execution rules (`completions`, `parallelism`, `backoffLimit`, `activeDeadlineSeconds`, `ttlSecondsAfterFinished`, `template`).
3. `CronJob.spec.jobTemplate.spec.template.spec` -> Pod specification (`restartPolicy: OnFailure|Never`, `containers`).

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: report-generator
spec:                                      # <--- SPEC #1: CronJob Spec
  schedule: "30 2 * * 1-5"                 # 02:30 AM Monday through Friday
  concurrencyPolicy: Forbid                # Allow (default) | Forbid | Replace
  startingDeadlineSeconds: 120             # Max delay if controller missed schedule
  successfulJobsHistoryLimit: 3            # Default: 3
  failedJobsHistoryLimit: 1                # Default: 1
  suspend: false                           # Set true to pause scheduling
  jobTemplate:
    spec:                                  # <--- SPEC #2: Job Spec
      completions: 3                       # 3 Pods must exit 0
      parallelism: 2                       # Up to 2 Pods run at the same time
      backoffLimit: 4                      # Max retries before marking Job Failed
      activeDeadlineSeconds: 180           # Kill whole Job if running > 180s
      ttlSecondsAfterFinished: 600         # Auto-delete 10m after completion
      template:
        spec:                              # <--- SPEC #3: Pod Spec
          restartPolicy: OnFailure         # REQUIRED: Never or OnFailure (cannot be Always!)
          containers:
          - name: reporter
            image: busybox:1.36
            command: ["sh", "-c", "echo Generating report && sleep 2"]
```

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 5A: Watching `parallelism`, `completions`, and `backoffLimit` in Real Time
```bash
# 1. Generate a Job via $do and customize it
k create job worker-queue --image=busybox:1.36 $do -- sh -c "echo Processed item on \$HOSTNAME && sleep 3" > /tmp/opencode/job-queue.yaml
```
Edit `/tmp/opencode/job-queue.yaml` to set `completions: 6` and `parallelism: 2`:
```bash
sed -i '/metadata:/a \  labels:\n    job-group: batch' /tmp/opencode/job-queue.yaml
```
 Or apply directly:
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: worker-queue
spec:
  completions: 6
  parallelism: 2
  ttlSecondsAfterFinished: 120
  template:
    spec:
      restartPolicy: Never
      containers:
      - name: worker
        image: busybox:1.36
        command: ["sh", "-c", "echo Processed item on $HOSTNAME && sleep 2"]
```
```bash
k apply -f /tmp/opencode/job-queue.yaml
k get job worker-queue -w
# Watch COMPLETIONS progress: 0/6 -> 2/6 -> 4/6 -> 6/6!
k delete job worker-queue
```

#### Lab 5B: Break It & Fix It — `activeDeadlineSeconds` Timeout vs `backoffLimit`
Let's create a Job that sleeps 30 seconds, but has `activeDeadlineSeconds: 5`:
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: slow-job
spec:
  activeDeadlineSeconds: 5
  backoffLimit: 3
  template:
    spec:
      restartPolicy: Never
      containers:
      - name: slow
        image: busybox:1.36
        command: ["sleep", "30"]
```
```bash
k apply -f slow-job.yaml
sleep 7
k describe job slow-job | tail -n 10
# Expected Event: Warning  DeadlineExceeded  Job was active longer than specified deadline
k delete job slow-job
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 5.1: Parallel Calculation Job (Exam Weight: 6%)
In namespace `default`, create a Job named `pi-calc` using image `perl:5.34-slim` that computes $\pi$ to 1000 places (`perl -Mbignum=bpi -wle 'print bpi(1000)'`):
- It must complete **3 times** in total (`completions: 3`), running **3 Pods in parallel** (`parallelism: 3`).
- It must allow at most **2 retries** (`backoffLimit: 2`) and automatically delete itself **90 seconds** after finishing (`ttlSecondsAfterFinished: 90`).

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 5.1)</strong></summary>

```bash
k create job pi-calc --image=perl:5.34-slim $do -- perl -Mbignum=bpi -wle 'print bpi(1000)' > /tmp/opencode/pi.yaml
```
Add `completions: 3`, `parallelism: 3`, `backoffLimit: 2`, `ttlSecondsAfterFinished: 90` under `spec:`:
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pi-calc
spec:
  completions: 3
  parallelism: 3
  backoffLimit: 2
  ttlSecondsAfterFinished: 90
  template:
    spec:
      restartPolicy: Never
      containers:
      - name: pi-calc
        image: perl:5.34-slim
        command: ["perl", "-Mbignum=bpi", "-wle", "print bpi(1000)"]
```
```bash
k apply -f /tmp/opencode/pi.yaml
k wait --for=condition=complete job/pi-calc --timeout=90s
k logs job/pi-calc
k delete job pi-calc
```
</details>

#### Scenario 5.2: CronJob with Overlap Protection & Manual Trigger (Exam Weight: 7%)
Create a CronJob named `db-vacuum` in namespace `default`:
- Schedule: Every day at `03:15` (`15 3 * * *`)
- `concurrencyPolicy: Replace`
- `startingDeadlineSeconds: 60`
- `successfulJobsHistoryLimit: 4`, `failedJobsHistoryLimit: 2`
- Each Job must have `activeDeadlineSeconds: 45` and run `busybox:1.36` with command `["sh", "-c", "echo Vacuuming DB && sleep 2"]`.
- Immediately trigger a manual Job named `db-vacuum-manual` from `db-vacuum` to verify it completes.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 5.2)</strong></summary>

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: db-vacuum
spec:
  schedule: "15 3 * * *"
  concurrencyPolicy: Replace
  startingDeadlineSeconds: 60
  successfulJobsHistoryLimit: 4
  failedJobsHistoryLimit: 2
  jobTemplate:
    spec:
      activeDeadlineSeconds: 45
      template:
        spec:
          restartPolicy: OnFailure
          containers:
          - name: vacuum
              image: busybox:1.36
              command: ["sh", "-c", "echo Vacuuming DB && sleep 2"]
```
```bash
k apply -f db-vacuum.yaml
k create job db-vacuum-manual --from=cronjob/db-vacuum
k wait --for=condition=complete job/db-vacuum-manual --timeout=30s
k logs job/db-vacuum-manual
k delete cronjob db-vacuum && k delete job db-vacuum-manual
```
</details>

---

## Day 6 (Sat Nov 21 — Weekend Deep Lab): Volumes, PV, PVC, StorageClasses & Projected Volumes

### 1. Deep-Dive Theory & Mental Model

#### The Kubernetes Persistent Storage Lifecycle
```text
+-----------------------+       Matches on:        +-----------------------+
| PersistentVolume (PV) | ◄──────────────────────► | PersistentVolumeClaim |
| (Cluster-Scoped)      |  1. storageClassName     | (Namespaced)          |
| - capacity: 2Gi       |  2. accessModes          | - requests: 1Gi       |
| - accessModes: RWO    |  3. capacity >= request  | - accessModes: RWO    |
| - reclaimPolicy       |  4. label selector (opt) | - storageClassName    |
+-----------------------+                          +-----------+-----------+
                                                               │
                                                               │ claimName
                                                               ▼
                                                   +-----------------------+
                                                   | Pod                   |
                                                   | volumes[].pvc         |
                                                   | volumeMounts[]        |
                                                   +-----------------------+
```

#### Access Modes & Reclaim Policies Reference

| Access Mode | CLI Short | Meaning |
| :--- | :---: | :--- |
| `ReadWriteOnce` | `RWO` | Mounted read-write by a **single Node** (multiple Pods on that *same* Node can use it). |
| `ReadOnlyMany` | `ROX` | Mounted read-only by **many Nodes**. |
| `ReadWriteMany` | `RWX` | Mounted read-write by **many Nodes** simultaneously. |
| `ReadWriteOncePod` | `RWOP` | Mounted read-write by a **single Pod** across the entire cluster. |

| Reclaim Policy | What happens when the PVC is deleted? |
| :--- | :--- |
| `Retain` | PV moves to **`Released`** state. Data stays intact on disk, but **a `Released` PV cannot be bound by a new PVC** until an admin clears its `spec.claimRef`! |
| `Delete` | Both the PV object and the underlying backing storage volume are automatically deleted. |

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 6A: Static Provisioning (`PV` -> `PVC` -> `Pod`) & The `Released` PV Recovery Trick
This is a classic weekend-worthy deep dive that frequently trips people up on CKAD/CKA:

```yaml
# Save as /tmp/opencode/pv-lifecycle.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: finance-pv
spec:
  capacity:
    storage: 1Gi
  accessModes:
  - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  storageClassName: local-fast
  hostPath:
    path: /tmp/finance-data
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: finance-pvc
spec:
  accessModes:
  - ReadWriteOnce
  storageClassName: local-fast
  resources:
    requests:
      storage: 500Mi
```
```bash
k apply -f /tmp/opencode/pv-lifecycle.yaml
k get pv,pvc
# Both show STATUS: Bound!

# Now delete the PVC and see what happens to the PV with Retain policy:
k delete pvc finance-pvc
k get pv finance-pv
# STATUS is now 'Released' (NOT 'Available'!)

# If you recreate finance-pvc right now, it will stay Pending because finance-pv still holds the old claimRef!
# Fix: clear spec.claimRef on the PV so it returns to 'Available':
k patch pv finance-pv --type=json -p='[{"op": "remove", "path": "/spec/claimRef"}]'
k get pv finance-pv
# STATUS is back to 'Available'!
k delete pv finance-pv
```

#### Lab 6B: `projected` Volumes (Combining Secret + ConfigMap + DownwardAPI in One Mount)
```yaml
# Save as /tmp/opencode/projected-pod.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-info
data:
  region: "eu-west-1"
---
apiVersion: v1
kind: Secret
metadata:
  name: app-key
stringData:
  token: "xyz-987-secret"
---
apiVersion: v1
kind: Pod
metadata:
  name: projected-demo
  labels:
    tier: backend
spec:
  containers:
  - name: app
    image: busybox:1.36
    command: ["sleep", "3600"]
    volumeMounts:
    - name: all-in-one
      mountPath: /etc/bundle
      readOnly: true
  volumes:
  - name: all-in-one
    projected:
      sources:
      - configMap:
          name: app-info
      - secret:
          name: app-key
      - downwardAPI:
          items:
          - path: "pod_labels"
            fieldRef:
              fieldPath: metadata.labels
          - path: "cpu_limit"
            resourceFieldRef:
              containerName: app
              resource: limits.cpu
```
```bash
k apply -f /tmp/opencode/projected-pod.yaml
k exec projected-demo -- ls -la /etc/bundle
k exec projected-demo -- head -n 20 /etc/bundle/region /etc/bundle/token /etc/bundle/pod_labels
k delete -f /tmp/opencode/projected-pod.yaml
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 6.1: Why Is My PVC Stuck `Pending`? (Diagnostic Scenario — Weight: 7%)
Apply this broken manifest and figure out the **two** reasons `broken-pvc` refuses to bind to `avail-pv`, then fix `broken-pvc` so it reaches `Bound`:
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: avail-pv
spec:
  capacity:
    storage: 500Mi
  accessModes: ["ReadWriteOnce"]
  storageClassName: gold-ssd
  hostPath:
    path: /tmp/avail-pv
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: broken-pvc
spec:
  accessModes: ["ReadWriteMany"]
  storageClassName: silver-hdd
  resources:
    requests:
      storage: 1Gi
```

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 6.1)</strong></summary>

There are actually **three** mismatches between `broken-pvc` and `avail-pv`:
1. `storageClassName`: `silver-hdd` != `gold-ssd`
2. `accessModes`: `ReadWriteMany` != `ReadWriteOnce`
3. `requests.storage`: `1Gi` > PV's `500Mi` capacity!

Recreate `broken-pvc` with `storageClassName: gold-ssd`, `accessModes: ["ReadWriteOnce"]`, and `storage: 400Mi`:
```bash
k delete pvc broken-pvc
cat << 'EOF' | k apply -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: broken-pvc
spec:
  accessModes: ["ReadWriteOnce"]
  storageClassName: gold-ssd
  resources:
    requests:
      storage: 400Mi
EOF
k get pv,pvc
k delete pvc broken-pvc && k delete pv avail-pv
```
</details>

---

## Day 7 (Sun Nov 22 — Weekend Boss Lab): Week 1 Synthesis Arena (5 Multi-Part Scenarios)

Treat today like a real **60-minute mini-exam** covering everything from Week 1. Create namespace `week1-arena` and solve all 5 tasks without looking at solutions until the end.

```bash
k create ns week1-arena
```

### Arena Task 1 (15 pts): Custom Image & Lifecycle Pod
In `week1-arena`, create a Pod `lifecycle-runner` (`image: busybox:1.36`) with labels `domain=design,week=1`:
- Override `command` to `["sh", "-c"]` and `args` to `["while true; do date >> /var/log/clock.log; sleep 5; done"]`.
- Add a `postStart` hook that creates `/tmp/healthy` (`touch /tmp/healthy`).
- Set `terminationGracePeriodSeconds: 10`.

### Arena Task 2 (25 pts): Init-Gated Multi-Container Web Server
In `week1-arena`, create a Pod `gated-web`:
- Shared `emptyDir` volume `web-vol`.
- `initContainer` #1 `wait-svc` (`busybox:1.36`): runs `until nslookup auth-backend.week1-arena.svc.cluster.local; do sleep 2; done`.
- `initContainer` #2 `write-page` (`busybox:1.36`): writes `"<h1>Authenticated Portal</h1>"` to `/html/index.html` on `web-vol`.
- Main container `web` (`nginx:1.25-alpine`, port `80`) mounts `web-vol` at `/usr/share/nginx/html`.
- Sidecar container `access-tail` (`busybox:1.36`) runs `tail -n+1 -F /var/log/nginx/access.log` (share a second `emptyDir` `nginx-logs` mounted at `/var/log/nginx` in both `web` and `access-tail`!).
- Also create a dummy ClusterIP Service `auth-backend` (port `8080`) in `week1-arena` so `wait-svc` succeeds and `gated-web` transitions to `Running (2/2)`!

### Arena Task 3 (20 pts): Batch Job with Native Sidecar
In `week1-arena`, create a Job `export-ledger`:
- `completions: 2`, `parallelism: 1`, `backoffLimit: 2`, `ttlSecondsAfterFinished: 300`.
- A **native sidecar** (`initContainers` with `restartPolicy: Always`) named `monitor` (`busybox:1.36`) running `sh -c "while true; do echo monitoring; sleep 3; done"`.
- Main container `exporter` (`busybox:1.36`) running `sh -c "echo Ledger exported && sleep 2"`.
- Confirm the Job reaches `2/2` completions cleanly.

### Arena Task 4 (20 pts): CronJob with Deadline & History Limits
In `week1-arena`, create a CronJob `audit-sweep`:
- Schedule: every 15 minutes (`*/15 * * * *`)
- `concurrencyPolicy: Forbid`
- `successfulJobsHistoryLimit: 2`, `failedJobsHistoryLimit: 2`
- Job `activeDeadlineSeconds: 20`
- Container `auditor` (`busybox:1.36`) running `date`.
- Manually trigger a Job named `audit-run-1` from `audit-sweep`.

### Arena Task 5 (20 pts): Persistent Storage + DownwardAPI
In `week1-arena`:
- Create a PV `arena-pv` (`capacity: 1Gi`, `accessModes: [ReadWriteOnce]`, `storageClassName: arena-sc`, `hostPath: /tmp/arena-pv`).
- Create a PVC `arena-pvc` (`requests.storage: 500Mi`, `accessModes: [ReadWriteOnce]`, `storageClassName: arena-sc`).
- Create a Pod `storage-inspector` (`busybox:1.36`, sleeping `3600`, `limits: { cpu: 200m, memory: 128Mi }`) mounting `arena-pvc` at `/mnt/persist` and a `downwardAPI` volume at `/etc/podmeta` exposing `metadata.namespace` as file `ns` and container `limits.memory` as file `mem_limit`.

<details>
<summary><strong>Click for Complete Week 1 Arena Solutions & Verification Script</strong></summary>

```yaml
# Save as /tmp/opencode/week1-arena-solutions.yaml
apiVersion: v1
kind: Pod
metadata:
  name: lifecycle-runner
  namespace: week1-arena
  labels: { domain: design, week: "1" }
spec:
  terminationGracePeriodSeconds: 10
  containers:
  - name: runner
    image: busybox:1.36
    command: ["sh", "-c"]
    args: ["while true; do date >> /var/log/clock.log; sleep 5; done"]
    lifecycle:
      postStart:
        exec:
          command: ["touch", "/tmp/healthy"]
---
apiVersion: v1
kind: Service
metadata:
  name: auth-backend
  namespace: week1-arena
spec:
  ports: [{ port: 8080, targetPort: 8080 }]
---
apiVersion: v1
kind: Pod
metadata:
  name: gated-web
  namespace: week1-arena
spec:
  initContainers:
  - name: wait-svc
    image: busybox:1.36
    command: ["sh", "-c", "until nslookup auth-backend.week1-arena.svc.cluster.local; do sleep 2; done"]
  - name: write-page
    image: busybox:1.36
    command: ["sh", "-c", "echo '<h1>Authenticated Portal</h1>' > /html/index.html"]
    volumeMounts: [{ name: web-vol, mountPath: /html }]
  containers:
  - name: web
    image: nginx:1.25-alpine
    ports: [{ containerPort: 80 }]
    volumeMounts:
    - { name: web-vol, mountPath: /usr/share/nginx/html }
    - { name: nginx-logs, mountPath: /var/log/nginx }
  - name: access-tail
    image: busybox:1.36
    command: ["sh", "-c", "touch /var/log/nginx/access.log && tail -n+1 -F /var/log/nginx/access.log"]
    volumeMounts:
    - { name: nginx-logs, mountPath: /var/log/nginx }
  volumes:
  - { name: web-vol, emptyDir: {} }
  - { name: nginx-logs, emptyDir: {} }
---
apiVersion: batch/v1
kind: Job
metadata:
  name: export-ledger
  namespace: week1-arena
spec:
  completions: 2
  parallelism: 1
  backoffLimit: 2
  ttlSecondsAfterFinished: 300
  template:
    spec:
      restartPolicy: Never
      initContainers:
      - name: monitor
        image: busybox:1.36
        restartPolicy: Always
        command: ["sh", "-c", "while true; do echo monitoring; sleep 3; done"]
      containers:
      - name: exporter
        image: busybox:1.36
        command: ["sh", "-c", "echo Ledger exported && sleep 2"]
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: audit-sweep
  namespace: week1-arena
spec:
  schedule: "*/15 * * * *"
  concurrencyPolicy: Forbid
  successfulJobsHistoryLimit: 2
  failedJobsHistoryLimit: 2
  jobTemplate:
    spec:
      activeDeadlineSeconds: 20
      template:
        spec:
          restartPolicy: OnFailure
          containers:
          - name: auditor
            image: busybox:1.36
            command: ["date"]
---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: arena-pv
spec:
  capacity: { storage: 1Gi }
  accessModes: ["ReadWriteOnce"]
  storageClassName: arena-sc
  hostPath: { path: /tmp/arena-pv }
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: arena-pvc
  namespace: week1-arena
spec:
  accessModes: ["ReadWriteOnce"]
  storageClassName: arena-sc
  resources:
    requests: { storage: 500Mi }
---
apiVersion: v1
kind: Pod
metadata:
  name: storage-inspector
  namespace: week1-arena
spec:
  containers:
  - name: inspector
    image: busybox:1.36
    command: ["sleep", "3600"]
    resources:
      limits: { cpu: 200m, memory: 128Mi }
    volumeMounts:
    - { name: pvol, mountPath: /mnt/persist }
    - { name: meta, mountPath: /etc/podmeta }
  volumes:
  - name: pvol
    persistentVolumeClaim: { claimName: arena-pvc }
  - name: meta
    downwardAPI:
      items:
      - path: "ns"
        fieldRef: { fieldPath: metadata.namespace }
      - path: "mem_limit"
        resourceFieldRef:
          containerName: inspector
          resource: limits.memory
```
```bash
k apply -f /tmp/opencode/week1-arena-solutions.yaml
k create job audit-run-1 -n week1-arena --from=cronjob/audit-sweep
k get pods,job,cronjob,pvc -n week1-arena
k delete ns week1-arena && k delete pv arena-pv
```
</details>


---

# Week 2 (Nov 23–29): Application Deployment, Scheduling & Helm

**Weekly Objective:** Master **Application Deployment (20% of CKAD)** plus advanced workload scheduling: Deployments, RollingUpdates, Rollbacks, Blue-Green & Canary strategies, Node/Pod Affinity, Taints/Tolerations, Horizontal Pod Autoscaling (`autoscaling/v2`), Helm v3, StatefulSets, and DaemonSets.

---

## Day 8 (Mon Nov 23): Deployments, ReplicaSets, RollingUpdates & Rollbacks

### 1. Deep-Dive Theory & Mental Model

#### How a Deployment Controls ReplicaSets During a Rollout
A **Deployment** never manages Pods directly—it creates and scales **ReplicaSets**. Whenever you change *anything* inside `spec.template` (`PodTemplateSpec`), the Deployment controller creates a new ReplicaSet (Revision $N+1$) and gradually scales it up while scaling the old ReplicaSet down:

```text
Deployment: checkout (replicas: 4, maxSurge: 25% [1], maxUnavailable: 25% [1])
  │
  ├──► ReplicaSetcheckout-7b9f8 (Revision 1: nginx:1.25)  [4 -> 3 -> 2 -> 1 -> 0 Pods]
  │         (Kept at 0 replicas for rollback history!)
  │
  └──► ReplicaSetcheckout-5d4c2 (Revision 2: nginx:1.26)  [0 -> 2 -> 3 -> 4 Pods]
```

#### Calculating `maxSurge` and `maxUnavailable` Math
Suppose `replicas: 4`, `maxSurge: 25%`, `maxUnavailable: 25%`:
- **Minimum Available Pods at any moment:** $4 - (4 \times 0.25) = \mathbf{3}$ Ready Pods.
- **Maximum Total Pods at any moment:** $4 + (4 \times 0.25) = \mathbf{5}$ Pods.

#### Annotated Deployment YAML Anatomy
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: checkout
spec:
  replicas: 4
  revisionHistoryLimit: 10             # Number of old ReplicaSets to retain (default 10)
  progressDeadlineSeconds: 600         # Marks deployment ProgressDeadlineExceeded if stuck > 600s
  strategy:
    type: RollingUpdate                # RollingUpdate (default) | Recreate
    rollingUpdate:
      maxSurge: 25%                    # Int or % above desired replicas
      maxUnavailable: 0                # 0 guarantees capacity never dips below 4 during update!
  selector:
    matchLabels:
      app: checkout                    # MUST match spec.template.metadata.labels!
  template:
    metadata:
      labels:
        app: checkout
    spec:
      containers:
      - name: web                      # Remember this container name for `k set image`!
        image: nginx:1.25-alpine
```

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 8A: Multi-Revision Rollouts, Pausing/Batching & Rolling Back
```bash
k create ns day8-lab

# 1. Create Revision 1
k create deploy storefront -n day8-lab --image=nginx:1.24-alpine --replicas=3
k annotate deploy/storefront -n day8-lab kubernetes.io/change-cause="Rev 1: nginx 1.24"
k rollout status deploy/storefront -n day8-lab

# 2. Pause rollout so we can batch TWO spec.template changes into ONE single revision (Rev 2)!
k rollout pause deploy/storefront -n day8-lab
k set image deploy/storefront -n day8-lab nginx=nginx:1.25-alpine
k set resources deploy/storefront -n day8-lab -c=nginx --requests=cpu=50m,memory=64Mi --limits=cpu=100m,memory=128Mi
k annotate deploy/storefront -n day8-lab kubernetes.io/change-cause="Rev 2: nginx 1.25 + resources"

# Verify no new ReplicaSet rolled out while paused:
k get rs -n day8-lab

# Now resume and watch Revision 2 roll out in one go:
k rollout resume deploy/storefront -n day8-lab
k rollout status deploy/storefront -n day8-lab

# 3. Push a broken Revision 3
k set image deploy/storefront -n day8-lab nginx=nginx:broken-tag-999
k annotate deploy/storefront -n day8-lab kubernetes.io/change-cause="Rev 3: broken image"

# 4. Inspect rollout history & ReplicaSets
k rollout history deploy/storefront -n day8-lab
k rollout history deploy/storefront -n day8-lab --revision=2
k get rs -n day8-lab
```
**Expected Output of `k rollout history`:**
```text
deployment.apps/storefront
REVISION  CHANGE-CAUSE
1         Rev 1: nginx 1.24
2         Rev 2: nginx 1.25 + resources
3         Rev 3: broken image
```
Roll back to Revision 2:
```bash
k rollout undo deploy/storefront -n day8-lab --to-revision=2
k rollout status deploy/storefront -n day8-lab
k delete ns day8-lab
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 8.1: Zero-Degradation Rolling Update (Exam Weight: 6%)
In namespace `default`, create a Deployment `order-engine` (`image: nginx:1.25-alpine`, `replicas: 5`):
- Configure `RollingUpdate` so **at least 5 Pods are always available** during updates (`maxUnavailable: 0`) and **up to 2 extra Pods** can be created (`maxSurge: 2`).
- Set `revisionHistoryLimit: 5`.
- Update the image to `nginx:1.26-alpine` and verify all 5 replicas are running the new image.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 8.1)</strong></summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: order-engine
spec:
  replicas: 5
  revisionHistoryLimit: 5
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 2
      maxUnavailable: 0
  selector:
    matchLabels: { app: order-engine }
  template:
    metadata:
      labels: { app: order-engine }
    spec:
      containers:
      - name: nginx
        image: nginx:1.25-alpine
```
```bash
k apply -f order-engine.yaml
k set image deploy/order-engine nginx=nginx:1.26-alpine
k rollout status deploy/order-engine
k get rs -l app=order-engine
k delete deploy order-engine
```
</details>

---

## Day 9 (Tue Nov 24): Advanced Release Strategies — Recreate, Blue-Green & Canary

### 1. Deep-Dive Theory & Mental Model

#### Visual Architecture of Blue-Green vs Canary

```text
BLUE-GREEN DEPLOYMENT (Instant 100% Cutover via Service Selector)
=================================================================
  [Service: pay-svc]
  selector: { app: pay, version: green }   <-- Flipped from 'blue' to 'green'!
                     │
        +------------+------------+
        │ (0% traffic)            │ (100% traffic)
        ▼                         ▼
 [Deploy: pay-blue]        [Deploy: pay-green]
 labels: app=pay           labels: app=pay
         version=blue              version=green
 (3 Pods Ready - Standby)  (3 Pods Ready - Active)


CANARY DEPLOYMENT (Proportional Traffic Split via Shared Label)
===============================================================
  [Service: pay-svc]
  selector: { app: pay }                   <-- Matches BOTH v1 and v2!
                     │
        +------------+------------+
        │ (80% traffic = 4/5)     │ (20% traffic = 1/5)
        ▼                         ▼
 [Deploy: pay-v1]          [Deploy: pay-v2-canary]
 selector: app=pay,v=v1    selector: app=pay,v=v2
 (4 Pods Ready)            (1 Pod Ready)
```

> **Critical Canary Rule:**
> In a Canary deployment, the **Service** selects only `app: pay`, while each **Deployment** MUST select **both** `app: pay` AND `v: v1` (or `v: v2`). If `pay-v1`'s `spec.selector` only had `app: pay`, the two Deployments would fight over each other's Pods!

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 9A: Live Blue-Green Cutover & Canary Traffic Split Verification
```yaml
# Save as /tmp/opencode/release-strategies.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: release-lab
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-v1
  namespace: release-lab
spec:
  replicas: 4
  selector:
    matchLabels: { app:greeter, track: stable }
  template:
    metadata:
      labels: { app: greeter, track: stable }
    spec:
      containers:
      - name: web
        image: hashicorp/http-echo:0.2.3
        args: ["-listen=:8080", "-text=VERSION_1_STABLE"]
        ports: [{ containerPort: 8080 }]
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-v2
  namespace: release-lab
spec:
  replicas: 1
  selector:
    matchLabels: { app: greeter, track: canary }
  template:
    metadata:
      labels: { app: greeter, track: canary }
    spec:
      containers:
      - name: web
        image: hashicorp/http-echo:0.2.3
        args: ["-listen=:8080", "-text=VERSION_2_CANARY"]
        ports: [{ containerPort: 8080 }]
---
apiVersion: v1
kind: Service
metadata:
  name: greeter-svc
  namespace: release-lab
spec:
  # Initially Blue-Green mode: points 100% to track=stable
  selector:
    app: greeter
    track: stable
  ports:
  - port: 80
    targetPort: 8080
```

Apply and test **Blue-Green** first, then **Canary**:
```bash
k apply -f /tmp/opencode/release-strategies.yaml
k wait --for=condition=Available deploy/app-v1 deploy/app-v2 -n release-lab --timeout=60s

# 1. Test Blue-Green (100% hits VERSION_1_STABLE)
k run tester -n release-lab --image=busybox:1.36 --rm -it --restart=Never -- \
  sh -c 'for i in $(seq 1 10); do wget -qO- http://greeter-svc; done'

# 2. Switch Service to Canary mode (remove 'track' from Service selector so 4 v1 + 1 v2 = 80%/20% split!)
k patch svc greeter-svc -n release-lab --type=json -p='[{"op":"remove","path":"/spec/selector/track"}]'

# 3. Send 10 requests and observe ~8 VERSION_1_STABLE and ~2 VERSION_2_CANARY!
k run tester -n release-lab --image=busybox:1.36 --rm -it --restart=Never -- \
  sh -c 'for i in $(seq 1 10); do wget -qO- http://greeter-svc; done'

# 4. Promote Canary to 100% (point Service to track: canary and scale app-v2 to 4)
k scale deploy app-v2 -n release-lab --replicas=4
k patch svc greeter-svc -n release-lab -p '{"spec":{"selector":{"track":"canary"}}}'
k delete ns release-lab
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 9.1: Implementing a 10% Canary Release (Exam Weight: 8%)
In namespace `canary-prod` (create it):
1. Deploy `frontend-v1` (`nginx:1.25-alpine`, `replicas: 9`, labels `app: frontend, version: v1`).
2. Expose it via ClusterIP Service `frontend-svc` on port `80` selecting **only** `app: frontend`.
3. Deploy `frontend-v2` (`nginx:1.26-alpine`, labels `app: frontend, version: v2`) scaled so that **exactly 10%** of total Pods backing `frontend-svc` run `v2`.
4. Verify `k get endpoints frontend-svc -n canary-prod` shows 10 Pod IPs.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 9.1)</strong></summary>

9 `v1` Pods + 1 `v2` Pod = 10 total Pods (1/10 = 10%):
```bash
k create ns canary-prod
k create deploy frontend-v1 -n canary-prod --image=nginx:1.25-alpine --replicas=9 $do > /tmp/opencode/f1.yaml
# Set labels & matchLabels to { app: frontend, version: v1 }
sed -i 's/app: frontend-v1/app: frontend\n        version: v1/g' /tmp/opencode/f1.yaml
k apply -f /tmp/opencode/f1.yaml

# Copy f1.yaml to f2.yaml with replicas: 1, version: v2, image: nginx:1.26-alpine
sed -e 's/frontend-v1/frontend-v2/g' -e 's/replicas: 9/replicas: 1/' -e 's/version: v1/version: v2/g' -e 's/1.25-alpine/1.26-alpine/' /tmp/opencode/f1.yaml > /tmp/opencode/f2.yaml
k apply -f /tmp/opencode/f2.yaml

k create service clusterip frontend-svc -n canary-prod --tcp=80:80 $do | sed 's/app: frontend-svc/app: frontend/' | k apply -f -
k get endpoints frontend-svc -n canary-prod
k delete ns canary-prod
```
</details>

---

## Day 10 (Wed Nov 25): Pod Scheduling — `nodeSelector`, Node/Pod Affinity & Taints/Tolerations

### 1. Deep-Dive Theory & Mental Model

#### Attraction (Affinity) vs Repulsion (Taints & Tolerations)
```text
NODE AFFINITY (Pod attracts to Node)           TAINTS & TOLERATIONS (Node repels Pods)
====================================           =======================================
[Pod: "I require disk=nvme"]                   [Node: Tainted with dedicated=ml:NoSchedule]
            │                                              │
            ▼                                              ├──► Normal Pod -> REJECTED (Pending)
   [Node: disk=nvme]                                       └──► Pod with Toleration -> ALLOWED
                                                                (Note: Can still land on other
                                                                 untainted nodes unless paired
                                                                 with nodeSelector/nodeAffinity!)
```

#### Taint Effects Cheat Table

| Taint Effect | Effect on New Pods (Scheduling) | Effect on Already Running Pods (Execution) |
| :--- | :--- | :--- |
| `NoSchedule` | Will **not** schedule onto the Node unless Pod has matching toleration. | Already running Pods on the Node are **not** evicted. |
| `PreferNoSchedule` | Scheduler tries to avoid the Node, but **will** schedule there if no other Node has capacity. | Already running Pods are **not** evicted. |
| `NoExecute` | Will **not** schedule onto the Node without matching toleration. | **Immediately evicts** running Pods lacking the toleration (or evicts after `tolerationSeconds` if specified!). |

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 10A: Combining Taints, Tolerations (`NoExecute` + `tolerationSeconds`) & NodeAffinity
```bash
# Pick the second node in your minikube cluster (or first node if single-node)
TARGET_NODE=$(k get nodes -o jsonpath='{.items[-1].metadata.name}')
echo "Target node: $TARGET_NODE"

# 1. Label and taint the node
k label node $TARGET_NODE accelerator=nvidia-t4 --overwrite
k taint node $TARGET_NODE workload=ai:NoSchedule --overwrite

# 2. Deploy a Pod that both tolerates the taint AND requires the label via nodeAffinity
cat << 'EOF' | k apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: ai-trainer
spec:
  tolerations:
  - key: "workload"
    operator: "Equal"
    value: "ai"
    effect: "NoSchedule"
  - key: "workload"
    operator: "Equal"
    value: "ai"
    effect: "NoExecute"
    tolerationSeconds: 30
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: accelerator
            operator: In
            values:
            - nvidia-t4
            - nvidia-a100
  containers:
  - name: trainer
    image: busybox:1.36
    command: ["sleep", "3600"]
EOF

k get pod ai-trainer -o wide
# Clean up taint & label immediately after verifying!
k delete pod ai-trainer $now
k taint node $TARGET_NODE workload=ai:NoSchedule-
k label node $TARGET_NODE accelerator-
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 10.1: High-Availability Pod Anti-Affinity (Exam Weight: 7%)
Create a Deployment `redis-ha` (`image: redis:7.2-alpine`, `replicas: 2`) in namespace `default` that uses **hard Pod Anti-Affinity** (`requiredDuringSchedulingIgnoredDuringExecution`) with `topologyKey: kubernetes.io/hostname` so that **no two replicas of `redis-ha` can ever be scheduled onto the same Node**.
*(Bonus test: scale `redis-ha` to `3` replicas on a 2-node cluster and verify the 3rd Pod stays `Pending` because of the anti-affinity rule!)*

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 10.1)</strong></summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis-ha
spec:
  replicas: 2
  selector:
    matchLabels: { app: redis-ha }
  template:
    metadata:
      labels: { app: redis-ha }
    spec:
      affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values: ["redis-ha"]
            topologyKey: kubernetes.io/hostname
      containers:
      - name: redis
        image: redis:7.2-alpine
```
```bash
k apply -f redis-ha.yaml
k get pods -l app=redis-ha -o wide
k delete deploy redis-ha
```
</details>

---

## Day 11 (Thu Nov 26): Horizontal Pod Autoscaling (`autoscaling/v2`) & Metrics

### 1. Deep-Dive Theory & Mental Model

#### Annotated `autoscaling/v2` HPA Manifest (CPU + Memory + Custom Behavior)
While `kubectl autoscale` only generates a simple CPU utilization target, CKAD can ask you to add a **memory metric** or customize the **scaleDown stabilization window**:

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: payment-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: payment-api
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization          # Utilization (%) -> compares against pod's requests.cpu
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: AverageValue         # AverageValue -> raw value like 100Mi (or Utilization %)
        averageValue: 100Mi
  behavior:
    scaleDown:
      stabilizationWindowSeconds: 180   # Default is 300s (5m)
      policies:
      - type: Pods
        value: 1
        periodSeconds: 60               # Remove at most 1 Pod per minute
    scaleUp:
      stabilizationWindowSeconds: 0
```

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 11A: Generating an `autoscaling/v2` Skeleton & Extending It
Instead of memorizing the `autoscaling/v2` YAML from scratch, let `kubectl autoscale ... $do` generate the structure, then edit it!
```bash
k autoscale deploy web --min=2 --max=8 --cpu-percent=65 $do
```
**Output:**
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  creationTimestamp: null
  name: web
spec:
  maxReplicas: 8
  metrics:
  - resource:
      name: cpu
      target:
        averageUtilization: 65
        type: Utilization
    type: Resource
  minReplicas: 2
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: web
status:
  currentMetrics: null
  desiredReplicas: 0
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 11.1: Dual-Metric HPA with Stabilization Window (Exam Weight: 7%)
In namespace `autoscale-ns`:
1. Create Deployment `worker-api` (`nginx:1.25-alpine`, `replicas: 2`) with container resource requests `cpu: 100m, memory: 128Mi`.
2. Create an `autoscaling/v2` HorizontalPodAutoscaler named `worker-hpa` targeting `worker-api` with:
   - `minReplicas: 2`, `maxReplicas: 6`
   - Metric 1: CPU `Utilization` at `60%`
   - Metric 2: Memory `Utilization` at `80%`
   - `behavior.scaleDown.stabilizationWindowSeconds: 90`

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 11.1)</strong></summary>

```bash
k create ns autoscale-ns
k create deploy worker-api -n autoscale-ns --image=nginx:1.25-alpine --replicas=2
k set resources deploy/worker-api -n autoscale-ns --requests=cpu=100m,memory=128Mi

cat << 'EOF' | k apply -f -
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: worker-hpa
  namespace: autoscale-ns
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: worker-api
  minReplicas: 2
  maxReplicas: 6
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 60
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
  behavior:
    scaleDown:
      stabilizationWindowSeconds: 90
EOF
k describe hpa worker-hpa -n autoscale-ns
k delete ns autoscale-ns
```
</details>

---

## Day 12 (Fri Nov 27): Helm Package Manager — Charts, Values, Releases & Rollbacks

### 1. Deep-Dive Theory & Mental Model

#### Helm Value Precedence Order (Lowest to Highest Priority)
When installing or upgrading a chart, values merge in this exact order (rightmost wins):
```text
Chart's default `values.yaml`  <  `-f custom-values.yaml`  <  `--set key=val`
```

#### Complete Helm Exam Command Matrix

| Exam Task | Exact Command |
| :--- | :--- |
| List configured chart repos | `helm repo list` |
| Add & refresh a chart repo | `helm repo add bitnami https://charts.bitnami.com/bitnami && helm repo update` |
| Search all versions of a chart | `helm search repo bitnami/nginx --versions` |
| Download/extract a chart locally to edit templates | `helm pull bitnami/nginx --untar --untardir /tmp/opencode` |
| View configurable chart parameters | `helm show values bitnami/nginx` |
| Install specific chart version in namespace | `helm install my-rel bitnami/nginx --version 18.1.0 -n web --create-namespace` |
| Upgrade release keeping existing values & changing one | `helm upgrade my-rel bitnami/nginx -n web --reuse-values --set replicaCount=3` |
| Find broken / pending releases across all namespaces | `helm list -a -A` (or `helm list --pending -A`) |
| Roll back release to revision 1 | `helm rollback my-rel 1 -n web` |
| Uninstall release | `helm uninstall my-rel -n web` |

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 12A: Creating, Installing, Modifying, and Rolling Back a Local Helm Chart
On the CKAD exam, you might be given a **local chart directory** on disk (e.g., `/opt/charts/webapp`) instead of a remote repo! Let's practice creating and working with a local chart:

```bash
mkdir -p /tmp/opencode/helm-local && cd /tmp/opencode/helm-local

# 1. Scaffold a standard Helm chart
helm create mychart
tree mychart || ls -R mychart

# 2. Inspect default values and render the templates locally
helm show values ./mychart | head -n 25
helm template test-render ./mychart --set replicaCount=3 | grep -E "kind:|replicas:"

# 3. Install the local chart into namespace 'chart-ns'
helm install local-app ./mychart -n chart-ns --create-namespace --set replicaCount=2
k get all -n chart-ns

# 4. Upgrade with --reuse-values and override service.port
helm upgrade local-app ./mychart -n chart-ns --reuse-values --set service.port=8080
helm get values local-app -n chart-ns
helm history local-app -n chart-ns

# 5. Rollback to revision 1 and clean up
helm rollback local-app 1 -n chart-ns
helm uninstall local-app -n chart-ns && k delete ns chart-ns
cd /home/a00876142/ckad-session/ckad
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 12.1: Cleaning Up a Pending Release & Deploying a Custom Local Chart (Exam Weight: 8%)
1. Using `/tmp/opencode/helm-local/mychart`, modify `mychart/values.yaml` so `replicaCount` is `2` and `image.tag` is `"1.25-alpine"`.
2. Install it as release `portal-v1` in namespace `portal` (`--create-namespace`).
3. Upgrade `portal-v1` using `--set image.pullPolicy=Always` while preserving your previous values.
4. Verify `helm get values portal-v1 -n portal` shows `pullPolicy: Always`, then uninstall and delete `portal`.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 12.1)</strong></summary>

```bash
helm install portal-v1 /tmp/opencode/helm-local/mychart -n portal --create-namespace \
  --set replicaCount=2 --set image.tag=1.25-alpine
helm upgrade portal-v1 /tmp/opencode/helm-local/mychart -n portal --reuse-values --set image.pullPolicy=Always
helm get values portal-v1 -n portal
helm uninstall portal-v1 -n portal && k delete ns portal
```
</details>

---

## Day 13 (Sat Nov 28 — Weekend Deep Lab): StatefulSets, Headless Services & DaemonSets

### 1. Deep-Dive Theory & Mental Model

#### Deployment vs StatefulSet vs DaemonSet Comparison

| Dimension | `Deployment` | `StatefulSet` | `DaemonSet` |
| :--- | :--- | :--- | :--- |
| **Pod Naming** | Random hash (`web-7d8f9-x2k9p`) | Ordered ordinal (`db-0`, `db-1`, `db-2`) | Node-bound (`agent-k8x2m`) |
| **Storage** | All replicas share the same PVC (if any) | Each replica gets its **own dedicated PVC** via `volumeClaimTemplates` (`data-db-0`, `data-db-1`) | Usually `hostPath` or `emptyDir` |
| **Network Identity** | Single load-balanced ClusterIP | Paired with **Headless Service (`clusterIP: None`)** -> dedicated DNS record per Pod (`db-0.db-svc.ns.svc.cluster.local`) | Often `hostNetwork: true` or node-local |
| **Scaling** | Controlled by `replicas` (parallel startup) | Controlled by `replicas` (ordered `0 -> 1 -> 2` startup, `2 -> 1 -> 0` termination by default) | Automatically `1` Pod per matching Node |

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 13A: StatefulSet + Headless Service + Per-Pod Storage Persistence Test
```yaml
# Save as /tmp/opencode/sts-full.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: sts-lab
---
apiVersion: v1
kind: Service
metadata:
  name: redis-headless
  namespace: sts-lab
spec:
  clusterIP: None                  # Headless Service!
  selector:
    app: redis-cluster
  ports:
  - name: redis
    port: 6379
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: redis-node
  namespace: sts-lab
spec:
  serviceName: redis-headless      # REQUIRED: Must match Headless Service name!
  replicas: 3
  selector:
    matchLabels:
      app: redis-cluster
  template:
    metadata:
      labels:
        app: redis-cluster
    spec:
      containers:
      - name: redis
        image: redis:7.2-alpine
        ports:
        - containerPort: 6379
        volumeMounts:
        - name: redis-data
          mountPath: /data
  volumeClaimTemplates:
  - metadata:
      name: redis-data
    spec:
      accessModes: ["ReadWriteOnce"]
      resources:
        requests:
          storage: 100Mi
```

Apply and test the 3 defining properties of a StatefulSet:
```bash
k apply -f /tmp/opencode/sts-full.yaml
k rollout status sts/redis-node -n sts-lab

# 1. Check predictable Pod names & dedicated PVCs:
k get pods,pvc -n sts-lab
# Pods: redis-node-0, redis-node-1, redis-node-2
# PVCs: redis-data-redis-node-0, redis-data-redis-node-1, redis-data-redis-node-2

# 2. Test individual Pod DNS records via Headless Service:
k run dns-check -n sts-lab --image=busybox:1.36 --rm -it --restart=Never -- \
  nslookup redis-node-1.redis-headless.sts-lab.svc.cluster.local

# 3. Write a file to redis-node-1's volume, delete redis-node-1, and prove the new redis-node-1 reattaches to the SAME PVC:
k exec -n sts-lab redis-node-1 -- sh -c "echo 'NODE_1_STATE' > /data/id.txt"
k delete pod redis-node-1 -n sts-lab
k wait --for=condition=Ready pod/redis-node-1 -n sts-lab --timeout=60s
k exec -n sts-lab redis-node-1 -- cat /data/id.txt
# Output: NODE_1_STATE
k delete ns sts-lab
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 13.1: Converting a Deployment Spec into a DaemonSet (Exam Weight: 6%)
Pro-Tip: `kubectl create` does **not** have a `kubectl create daemonset` subcommand! The fastest way to create a DaemonSet on the exam is:
1. Run `k create deploy node-exporter --image=prom/node-exporter:v1.7.0 $do > ds.yaml`
2. Change `kind: Deployment` to `kind: DaemonSet`, and delete the `replicas:`, `strategy:`, and `status:` lines!
Try this right now in namespace `monitoring`: create a DaemonSet named `log-agent` (`image: busybox:1.36`, command `sleep 3600`) using the Deployment-to-DaemonSet shortcut in under 90 seconds.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 13.1)</strong></summary>

```bash
k create ns monitoring
k create deploy log-agent -n monitoring --image=busybox:1.36 $do -- sleep 3600 > /tmp/opencode/ds.yaml
# Edit /tmp/opencode/ds.yaml: change kind: DaemonSet and remove replicas: 1 and strategy: {}
cat << 'EOF' > /tmp/opencode/ds.yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    app: log-agent
  name: log-agent
  namespace: monitoring
spec:
  selector:
    matchLabels:
      app: log-agent
  template:
    metadata:
      labels:
        app: log-agent
    spec:
      containers:
      - command: ["sleep", "3600"]
        image: busybox:1.36
        name: busybox
EOF
k apply -f /tmp/opencode/ds.yaml
k get ds,pods -n monitoring -o wide
k delete ns monitoring
```
</details>

---

## Day 14 (Sun Nov 29 — Weekend Boss Lab): Week 2 End-to-End Release Engineering Lab

Set a **60-minute timer**, create namespace `week2-arena`, and complete all 5 tasks without notes:

```bash
k create ns week2-arena
```

1. **Task 1 (20 pts — Rollouts):** Create Deployment `catalog` (`nginx:1.24-alpine`, `replicas: 4`, `maxSurge: 25%`, `maxUnavailable: 25%`). Update it to `nginx:1.25-alpine`, then `nginx:1.26-alpine`, then roll back to revision `2` (`nginx:1.25-alpine`).
2. **Task 2 (20 pts — Blue-Green):** Deploy `auth-blue` (`hashicorp/http-echo:0.2.3`, args `["-listen=:8080", "-text=BLUE"]`, 2 replicas, labels `app: auth, slot: blue`) and `auth-green` (same with `-text=GREEN` and `slot: green`). Create Service `auth-svc` (port `80` -> `8080`) initially routing to `blue`, then patch it to route 100% to `green`.
3. **Task 3 (20 pts — Scheduling):** Label one node `disktype=nvme`. Create Pod `fast-db` (`redis:7.2-alpine`) in `week2-arena` with `nodeAffinity` (`requiredDuringSchedulingIgnoredDuringExecution`) requiring `disktype In [nvme, ssd]`.
4. **Task 4 (20 pts — HPA):** Add resource requests (`cpu: 100m, memory: 64Mi`) to Deployment `catalog`, and create an `autoscaling/v2` HPA `catalog-hpa` (`minReplicas: 2`, `maxReplicas: 8`, target CPU `65%`, `scaleDown.stabilizationWindowSeconds: 60`).
5. **Task 5 (20 pts — StatefulSet):** Create Headless Service `zk-svc` (`port: 2181`) and StatefulSet `zk` (`busybox:1.36`, sleeping `3600`, `replicas: 2`, `serviceName: zk-svc`) with a `100Mi` `ReadWriteOnce` `volumeClaimTemplates` entry named `zk-data` mounted at `/var/lib/zk`.

<details>
<summary><strong>Click for Complete Week 2 Arena Verification & Cleanup</strong></summary>

```bash
# Verify each task:
k get deploy catalog -n week2-arena -o jsonpath='{.spec.template.spec.containers[0].image}{"\n"}'  # nginx:1.25-alpine
k run t -n week2-arena --image=busybox:1.36 --rm -it --restart=Never -- wget -qO- http://auth-svc  # GREEN
k get pod fast-db -n week2-arena -o wide
k get hpa catalog-hpa -n week2-arena
k get sts,pvc -n week2-arena

# Cleanup
k delete ns week2-arena
NODE=$(k get nodes -l disktype=nvme -o jsonpath='{.items[0].metadata.name}')
[ -n "$NODE" ] && k label node $NODE disktype-
```
</details>


---

# Week 3 (Nov 30–Dec 6): Configuration, Security & Observability (40% of Exam)

**Weekly Objective:** Master the two highest-ROI domains on the exam: **Application Environment, Configuration and Security (25%)** and **Application Observability and Maintenance (15%)**.

---

## Day 15 (Mon Nov 30): ConfigMaps & Secrets — Env Vars, Volume Mounts, `subPath` & TLS/Registry Secrets

### 1. Deep-Dive Theory & Mental Model

#### The 4 Ways Pods Consume ConfigMaps & Secrets

```text
+---------------------------------------------------------------------------------+
|Method                     | Auto-Updates on Change? | Best Use Case             |
+---------------------------+-------------------------+---------------------------+
|1. envFrom (All keys)      | NO (Requires Pod restart)| Injecting 5+ env vars     |
|2. env.valueFrom (1 key)   | NO (Requires Pod restart)| Renaming a key to an env  |
|3. Volume Mount (Full/items)| YES (~60s kubelet sync) | Config files, TLS certs   |
|4. Volume Mount + subPath  | NO (Frozen at boot!)    | Mounting 1 file into an   |
|                           |                         | existing dir (/etc/nginx) |
+---------------------------------------------------------------------------------+
```

#### Annotated Side-by-Side YAML: All 4 Consumption Patterns in One Pod
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: config-master
spec:
  imagePullSecrets:
  - name: my-registry-cred               # Type: kubernetes.io/dockerconfigjson
  containers:
  - name: app
    image: busybox:1.36
    command: ["sleep", "3600"]
    # Pattern 1: Import ALL keys from a ConfigMap & Secret as Env Vars
    envFrom:
    - configMapRef:
        name: global-cm
    - secretRef:
        name: global-sec
    # Pattern 2: Import a SINGLE key and map it to a custom Env Var name
    env:
    - name: DB_PASSWORD
      valueFrom:
        secretKeyRef:
          name: db-secret
          key: password
          optional: false                # Set true so Pod still starts if secret is missing!
    # Pattern 3 & 4: Volume Mounts
    volumeMounts:
    - name: cm-vol
      mountPath: /etc/app-config         # Auto-updates when ConfigMap changes!
      readOnly: true
    - name: sec-vol
      mountPath: /etc/nginx/secret.key   # Uses subPath -> DOES NOT auto-update!
      subPath: secret.key
      readOnly: true
  volumes:
  - name: cm-vol
    configMap:
      name: global-cm
      defaultMode: 0440                  # Octal file permissions (0440 = r--r-----)
      items:                             # Optional: project only specific keys
      - key: app.json
        path: settings.json
  - name: sec-vol
    secret:
      secretName: global-sec
      defaultMode: 0400                  # Read-only by owner (0400 = r--------)
```

> **Exam Trap — Octal Permissions in YAML (`0400` vs `256`):**
> In Kubernetes YAML, if you write `defaultMode: 0400` (with a leading zero) **without quotes**, some YAML parsers treat it as octal (`256` decimal) while `400` without a leading zero is treated as **decimal `400`** (which is `0620` octal!). Either write `defaultMode: 0400` carefully and verify with `k exec -- stat -c %a <file>`, or use decimal (`0400` octal = `256` decimal; `0644` octal = `420` decimal).

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 15A: Imperative Creation of All 3 Secret Types + Live Permission & Update Test
```bash
k create ns day15-lab

# 1. Generic Secret from literals and files
echo -n "super-secret-api-key" > /tmp/opencode/api.key
k create secret generic app-sec -n day15-lab \
  --from-literal=DB_USER=postgres \
  --from-file=API_KEY=/tmp/opencode/api.key

# Inspect base64-encoded secret and decode a value with base64 -d:
k get secret app-sec -n day15-lab -o jsonpath='{.data.DB_USER}' | base64 -d ; echo

# 2. TLS Secret (generate self-signed cert + key in 1 command)
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout /tmp/opencode/tls.key -out /tmp/opencode/tls.crt -subj "/CN=pay.example.com"
k create secret tls pay-tls -n day15-lab --cert=/tmp/opencode/tls.crt --key=/tmp/opencode/tls.key

# 3. Docker Registry Secret
k create secret docker-registry ecr-cred -n day15-lab \
  --docker-server=123456.dkr.ecr.eu-west-1.amazonaws.com \
  --docker-username=AWS --docker-password=my-token --docker-email=ops@example.com

k get secrets -n day15-lab
k delete ns day15-lab
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 15.1: Immutable Secret + Mode `0400` Mount + Single Env Key (Exam Weight: 8%)
In namespace `vault-ns`:
1. Create an **immutable** Secret (`immutable: true`) named `payment-creds` with keys `STRIPE_KEY=sk_live_999` and `WEBHOOK_SECRET=whsec_111`.
2. Create a Pod `payment-worker` (`busybox:1.36`, sleeping `3600`) that:
   - Injects `STRIPE_KEY` as environment variable `PROVIDER_KEY`.
   - Mounts `payment-creds` at `/var/run/creds` with `defaultMode: 0400` (`readOnly: true`).
3. Verify that attempting to `kubectl edit secret payment-creds -n vault-ns` to change `STRIPE_KEY` is rejected by Kubernetes because `immutable: true` is set!

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 15.1)</strong></summary>

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: vault-ns
---
apiVersion: v1
kind: Secret
metadata:
  name: payment-creds
  namespace: vault-ns
immutable: true
stringData:
  STRIPE_KEY: "sk_live_999"
  WEBHOOK_SECRET: "whsec_111"
---
apiVersion: v1
kind: Pod
metadata:
  name: payment-worker
  namespace: vault-ns
spec:
  containers:
  - name: worker
    image: busybox:1.36
    command: ["sleep", "3600"]
    env:
    - name: PROVIDER_KEY
      valueFrom:
        secretKeyRef:
          name: payment-creds
          key: STRIPE_KEY
    volumeMounts:
    - name: creds-vol
      mountPath: /var/run/creds
      readOnly: true
  volumes:
  - name: creds-vol
    secret:
      secretName: payment-creds
      defaultMode: 0400
```
```bash
k apply -f vault-scenario.yaml
k exec -n vault-ns payment-worker -- printenv PROVIDER_KEY
k exec -n vault-ns payment-worker -- ls -la /var/run/creds/..data
k delete ns vault-ns
```
</details>

---

## Day 16 (Tue Dec 1): Resource Management — Requests/Limits, QoS, ResourceQuota, LimitRange & PriorityClass

### 1. Deep-Dive Theory & Mental Model

#### How Kubernetes Scheduler, Kubelet & Eviction Manager Use Resource Controls

```text
+----------------------------------------------------------------------------+
| Stage 1: ADMISSION CONTROL (Namespace Policy Gatekeepers)                  |
|  - LimitRange: Injects default request/limit & checks per-Pod/Container    |
|                min/max constraints.                                        |
|  - ResourceQuota: Checks if new Pod pushes Namespace aggregate total       |
|                   over the namespace budget (CPU, RAM, PVC, Pod count).    |
+----------------------------------------------------------------------------+
                                      │
                                      ▼
+----------------------------------------------------------------------------+
| Stage 2: SCHEDULING (kube-scheduler)                                       |
|  - Looks ONLY at `resources.requests` (ignores `limits`!) + `PriorityClass`|
|  - Places Pod on a Node with sufficient unallocated capacity.              |
|  - If no Node fits a high-priority Pod -> Preempts lower-priority Pods!    |
+----------------------------------------------------------------------------+
                                      │
                                      ▼
+----------------------------------------------------------------------------+
| Stage 3: RUNTIME ENFORCEMENT & EVICTION (Node Kubelet + Linux cgroups)     |
|  - Container exceeds `limits.cpu`    -> Throttled via CFS quota (slows down|
|  - Container exceeds `limits.memory` -> Kernel OOM Killer -> `OOMKilled`   |
|  - Node runs low on RAM/Disk         -> Evicts Pods in order:              |
|      1. Lowest PriorityClass -> 2. BestEffort -> 3. Burstable -> 4. Guaranteed
+----------------------------------------------------------------------------+
```

#### QoS Class Decision Matrix
Inspect [`manifests/resource-management/`](manifests/resource-management/) in your repo:
- **`Guaranteed`:** **Every** container in the Pod sets both `cpu` and `memory` `requests` AND `limits`, and `request == limit` for each. (Note: if you specify *only* `limits` without `requests` in YAML, Kubernetes automatically copies `limits` into `requests`—which also qualifies as `Guaranteed`!)
- **`Burstable`:** At least one container sets a `cpu` or `memory` request or limit, and the Pod does not meet `Guaranteed` criteria.
- **`BestEffort`:** Zero requests and zero limits on **all** containers in the Pod.

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 16A: Testing `ResourceQuota` + `LimitRange` Admission Control Live
Inspect [`manifests/resource-quota/resource-quota.yaml`](manifests/resource-quota/resource-quota.yaml) and [`manifests/limitrange/limitrange.yaml`](manifests/limitrange/limitrange.yaml), then run this experiment:

```bash
k create ns day16-lab

# 1. Create a ResourceQuota imperatively or via YAML:
k create quota ns-budget -n day16-lab \
  --hard=pods=3,requests.cpu=400m,requests.memory=256Mi,limits.cpu=800m,limits.memory=512Mi,configmaps=5

k describe quota ns-budget -n day16-lab

# 2. Try creating a Pod WITHOUT resource requests/limits:
k run bare-pod -n day16-lab --image=busybox:1.36 -- sleep 3600
# Expected Error: Error from server (Forbidden): pods "bare-pod" is forbidden: failed quota: ns-budget: must specify limits.cpu for: bare-pod; limits.memory for: bare-pod; requests.cpu for: bare-pod; requests.memory for: bare-pod

# 3. Apply a LimitRange that injects default requests/limits AND enforces min/max:
cat << 'EOF' | k apply -f -
apiVersion: v1
kind: LimitRange
metadata:
  name: container-guard
  namespace: day16-lab
spec:
  limits:
  - type: Container
    defaultRequest:
      cpu: 100m
      memory: 64Mi
    default:
      cpu: 200m
      memory: 128Mi
    min:
      cpu: 50m
      memory: 32Mi
    max:
      cpu: 500m
      memory: 256Mi
EOF

# 4. Retry the exact same bare Pod command — now it succeeds!
k run bare-pod -n day16-lab --image=busybox:1.36 -- sleep 3600
k get pod bare-pod -n day16-lab -o jsonpath='QoS: {.status.qosClass}, CPU Req: {.spec.containers[0].resources.requests.cpu}{"\n"}'
k describe quota ns-budget -n day16-lab
k delete ns day16-lab
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 16.1: Multi-Container QoS Puzzle (Exam Weight: 6%)
In namespace `qos-test`, a 2-container Pod `mixed-pod` currently has:
- Container `c1` (`nginx:1.25-alpine`): `requests: { cpu: 100m, memory: 128Mi }`, `limits: { cpu: 100m, memory: 128Mi }`
- Container `c2` (`busybox:1.36`, sleeping): `requests: { cpu: 50m }`, `limits: { cpu: 50m }` (no memory set!)

1. Without applying it yet, predict what QoS class `mixed-pod` will get (`Guaranteed` or `Burstable`?) and why.
2. Fix `mixed-pod` so it achieves `Guaranteed` QoS and has a `PriorityClass` named `high-tier` (`value: 50000`, `preemptionPolicy: PreemptLowerPriority`).

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 16.1)</strong></summary>

Because `c2` is missing `memory` requests/limits, `mixed-pod` gets **`Burstable`**! To achieve **`Guaranteed`**, *every* container must specify *both* CPU and memory with `request == limit`.
```yaml
apiVersion: scheduling.k8s.io/v1
kind: PriorityClass
metadata:
  name: high-tier
value: 50000
preemptionPolicy: PreemptLowerPriority
---
apiVersion: v1
kind: Pod
metadata:
  name: mixed-pod
spec:
  priorityClassName: high-tier
  containers:
  - name: c1
    image: nginx:1.25-alpine
    resources:
      requests: { cpu: 100m, memory: 128Mi }
      limits: { cpu: 100m, memory: 128Mi }
  - name: c2
    image: busybox:1.36
    command: ["sleep", "3600"]
    resources:
      requests: { cpu: 50m, memory: 64Mi }
      limits: { cpu: 50m, memory: 64Mi }
```
```bash
k apply -f mixed-pod.yaml
k get pod mixed-pod -o jsonpath='{.status.qosClass}{"\n"}'  # Guaranteed
k delete pod mixed-pod $now && k delete priorityclass high-tier
```
</details>

---

## Day 17 (Wed Dec 2): Container & Pod SecurityContexts, Linux Capabilities, `fsGroup` & PSS

### 1. Deep-Dive Theory & Mental Model

#### Pod-Level vs Container-Level `securityContext` Placement Matrix
One of the easiest ways to get a YAML validation error on CKAD is placing a container-only security field at the Pod level (or vice versa). Keep this table burned into memory:

| Field | Pod-Level (`spec.securityContext`) | Container-Level (`spec.containers[].securityContext`) | Notes |
| :--- | :---: | :---: | :--- |
| `runAsUser` | YES | YES (overrides Pod) | UID for container process (`1000+` = non-root) |
| `runAsGroup` | YES | YES (overrides Pod) | Primary GID for container process |
| `runAsNonRoot` | YES | YES (overrides Pod) | Blocks start if effective UID is `0` |
| `seccompProfile` | YES | YES (overrides Pod) | `type: RuntimeDefault` or `Unconfined` or `Localhost` |
| **`fsGroup`** | **YES (ONLY!)** | **NO (Invalid!)** | Group owner applied to mounted volumes |
| **`supplementalGroups`** | **YES (ONLY!)** | **NO (Invalid!)** | Additional GIDs added to the user |
| **`allowPrivilegeEscalation`**| **NO (Invalid!)** | **YES (ONLY!)** | Blocks `setuid` elevation (`false` required by Restricted PSS) |
| **`readOnlyRootFilesystem`** | **NO (Invalid!)** | **YES (ONLY!)** | Mounts container `/` read-only |
| **`privileged`** | **NO (Invalid!)** | **YES (ONLY!)** | Full root host device access |
| **`capabilities`** | **NO (Invalid!)** | **YES (ONLY!)** | `drop: ["ALL"]`, `add: ["NET_BIND_SERVICE", "SYS_TIME"]` |

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 17A: Enforcing Pod Security Standards (PSS) on a Namespace & Fixing Violations
Kubernetes has a built-in Pod Security Admission controller configured via namespace labels:
- `pod-security.kubernetes.io/enforce: privileged | baseline | restricted`
- `pod-security.kubernetes.io/warn: restricted`

Let's see what happens when a namespace enforces `restricted` and you try to deploy a normal Pod:

```bash
k create ns secure-zone
k label ns secure-zone pod-security.kubernetes.io/enforce=restricted

# 1. Try creating a normal nginx Pod:
k run naive-pod -n secure-zone --image=nginx:1.25-alpine
```
**Expected Admission Rejection Error:**
```text
Error from server (Forbidden): pods "naive-pod" is forbidden: violates PodSecurity "restricted:latest": allowPrivilegeEscalation != false (container "naive-pod" must set securityContext.allowPrivilegeEscalation=false), unrestricted capabilities (container "naive-pod" must set securityContext.capabilities.drop=["ALL"]), runAsNonRoot != true (pod or container "naive-pod" must set securityContext.runAsNonRoot=true), seccompProfile (pod or container "naive-pod" must set securityContext.seccompProfile.type to "RuntimeDefault" or "Localhost")
```
Look at that error message—**Kubernetes literally tells you the exact 4 fields required to pass `restricted` PSS!** Now apply a compliant Pod:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: compliant-pod
  namespace: secure-zone
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 1000
    runAsGroup: 3000
    fsGroup: 2000
    seccompProfile:
      type: RuntimeDefault
  containers:
  - name: app
    image: busybox:1.36
    command: ["sleep", "3600"]
    securityContext:
      allowPrivilegeEscalation: false
      readOnlyRootFilesystem: true
      capabilities:
        drop: ["ALL"]
```
```bash
k apply -f compliant-pod.yaml
k get pods -n secure-zone
# STATUS: Running!
k delete ns secure-zone
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 17.1: Capabilities + `fsGroup` + Read-Only Root Filesystem (Exam Weight: 8%)
In namespace `sec-ops`, create a Deployment `net-auditor` (`replicas: 1`, `image: busybox:1.36`, command `["sleep", "3600"]`):
- Pod-level: `runAsUser: 2000`, `runAsGroup: 2000`, `fsGroup: 4000`, `runAsNonRoot: true`.
- Container-level: `readOnlyRootFilesystem: true`, `allowPrivilegeEscalation: false`, drop `ALL` capabilities, and add `NET_ADMIN` and `NET_RAW`.
- Mount an `emptyDir` volume named `tmp-vol` at `/tmp`.
- Verify via `k exec` that `/tmp/report.txt` can be created and has group ID `4000`, while `/etc/test` fails with `Read-only file system`.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 17.1)</strong></summary>

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: sec-ops
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: net-auditor
  namespace: sec-ops
spec:
  replicas: 1
  selector:
    matchLabels: { app: net-auditor }
  template:
    metadata:
      labels: { app: net-auditor }
    spec:
      securityContext:
        runAsUser: 2000
        runAsGroup: 2000
        fsGroup: 4000
        runAsNonRoot: true
      containers:
      - name: auditor
        image: busybox:1.36
        command: ["sleep", "3600"]
        securityContext:
          readOnlyRootFilesystem: true
          allowPrivilegeEscalation: false
          capabilities:
            drop: ["ALL"]
            add: ["NET_ADMIN", "NET_RAW"]
        volumeMounts:
        - name: tmp-vol
          mountPath: /tmp
      volumes:
      - name: tmp-vol
        emptyDir: {}
```
```bash
k apply -f net-auditor.yaml
k exec -n sec-ops deploy/net-auditor -- sh -c "touch /tmp/report.txt && ls -ln /tmp/report.txt"
k delete ns sec-ops
```
</details>

---

## Day 18 (Thu Dec 3): ServiceAccounts & Role-Based Access Control (RBAC)

### 1. Deep-Dive Theory & Mental Model

#### The RBAC Matrix: `Role` vs `ClusterRole` & `RoleBinding` vs `ClusterRoleBinding`

```text
+---------------------------------------------------------------------------------+
|                   RBAC PERMISSION & BINDING COMBINATIONS                        |
+------------------------+---------------------------+----------------------------+
| Role / ClusterRole     | Bound via                 | Effective Scope            |
+------------------------+---------------------------+----------------------------+
| 1. Role (Namespaced)   | RoleBinding (Namespaced)  | Single Namespace ONLY      |
| 2. ClusterRole         | ClusterRoleBinding        | ALL Namespaces + Cluster   |
| 3. ClusterRole         | RoleBinding (Namespaced)  | Single Namespace ONLY!     |
|                        |                           | (Great for reusing a       |
|                        |                           |  ClusterRole per-team!)    |
+------------------------+---------------------------+----------------------------+
  *(Note: A ClusterRoleBinding CANNOT reference a namespaced Role!)*
```

#### Common `apiGroups` Lookup (Or How to Avoid Looking Them Up!)
When you use **`kubectl create role` or `kubectl create clusterrole`**, `kubectl` automatically fills in the right `apiGroups` for you when you pass `--resource=<resource>`:
```bash
k create role MultiRole --verb=get,list,watch --resource=pods,deployments.apps,jobs.batch,ingresses.networking.k8s.io $do
```
Look at what that generates:
```yaml
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["apps"]
  resources: ["deployments"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["batch"]
  resources: ["jobs"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get", "list", "watch"]
```
**You never have to hand-write `apiGroups` again!**

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 18A: Diagnosing a Pod Failing `kubectl` API Calls & Fixing Its RBAC
On CKAD, a common scenario is: *"Pod `kubectl-monitor` in namespace `ops` is failing with `403 Forbidden` in its logs. Inspect the logs and grant its ServiceAccount only the exact permissions it needs."*

```bash
k create ns ops-rbac
k create sa monitor-sa -n ops-rbac

# Launch a Pod using bitnami/kubectl that tries to list deployments and services every 5s:
cat << 'EOF' | k apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: k8s-monitor
  namespace: ops-rbac
spec:
  serviceAccountName: monitor-sa
  containers:
  - name: kubectl
    image: bitnami/kubectl:latest
    command:
    - sh
    - -c
    - |
      while true; do
        kubectl get deployments,services -n ops-rbac
        sleep 5
      done
EOF

sleep 8
k logs k8s-monitor -n ops-rbac
```
**Expected Log Output:**
```text
Error from server (Forbidden): deployments.apps is forbidden: User "system:serviceaccount:ops-rbac:monitor-sa" cannot list resource "deployments" in API group "apps" in the namespace "ops-rbac"
Error from server (Forbidden): services is forbidden: User "system:serviceaccount:ops-rbac:monitor-sa" cannot list resource "services" in API group "" in the namespace "ops-rbac"
```
Look at the error: it tells you the exact `User`, `verb` (`list`), `resource` (`deployments`, `services`), and `API group` (`apps`, `""`)! Now fix it in two commands:
```bash
k create role monitor-role -n ops-rbac --verb=get,list,watch --resource=deployments.apps,services
k create rolebinding monitor-rb -n ops-rbac --role=monitor-role --serviceaccount=ops-rbac:monitor-sa

sleep 6
k logs k8s-monitor -n ops-rbac | tail -n 5
# Now prints the deployments and services cleanly!
k delete ns ops-rbac
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 18.1: Scoping a `ClusterRole` to a Single Namespace via `RoleBinding` (Exam Weight: 8%)
1. Create namespaces `team-alpha` and `team-beta`.
2. Create a **ClusterRole** named `secret-inspector` that allows `get`, `list`, `watch` on `secrets` and `configmaps`.
3. Create ServiceAccount `alpha-bot` in `team-alpha`.
4. Bind `ClusterRole/secret-inspector` to `alpha-bot` **only inside `team-alpha`** using a `RoleBinding` named `alpha-secret-bind`.
5. Verify with `k auth can-i` that `alpha-bot` can list secrets in `team-alpha`, but **cannot** list secrets in `team-beta`!

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 18.1)</strong></summary>

```bash
k create ns team-alpha && k create ns team-beta
k create clusterrole secret-inspector --verb=get,list,watch --resource=secrets,configmaps
k create sa alpha-bot -n team-alpha
k create rolebinding alpha-secret-bind -n team-alpha \
  --clusterrole=secret-inspector \
  --serviceaccount=team-alpha:alpha-bot

# Verify:
k auth can-i list secrets -n team-alpha --as=system:serviceaccount:team-alpha:alpha-bot   # yes
k auth can-i list secrets -n team-beta  --as=system:serviceaccount:team-alpha:alpha-bot   # no!
k delete ns team-alpha team-beta && k delete clusterrole secret-inspector
```
</details>

---

## Day 19 (Fri Dec 4): Container Probes (`startup`, `readiness`, `liveness`) & Live Debugging (`logs`, `top`, `debug`)

### 1. Deep-Dive Theory & Mental Model

#### Probe Timeline & Interaction Diagram

```text
Container Starts (t = 0s)
   │
   ▼
+-------------------------------------------------------------------+
| 1. startupProbe (if configured)                                   |
|    - Runs every `periodSeconds` up to `failureThreshold` times.   |
|    - Liveness & Readiness are DISABLED until startupProbe passes! |
|    - Maximum boot budget = failureThreshold * periodSeconds.      |
+-------------------------------------------------------------------+
   │ (startupProbe succeeds once!)
   ▼
+---------------------------------+---------------------------------+
| 2. livenessProbe (Continuous)   | 3. readinessProbe (Continuous)  |
|    - Fails -> Kubelet RESTARTS  |    - Fails -> Pod IP removed    |
|      the container!             |      from Service Endpoints     |
|                                 |      (Container NOT restarted!) |
+---------------------------------+---------------------------------+
```

#### All 4 Probe Action Handlers
```yaml
# 1. HTTP GET (200-399 = OK)
httpGet:
  path: /healthz
  port: 8080
  httpHeaders:
  - name: Custom-Header
    value: Awesome

# 2. TCP Socket (Handshake succeeds = OK)
tcpSocket:
  port: 3306

# 3. Exec Command (Exit code 0 = OK)
exec:
  command: ["sh", "-c", "pg_isready -U postgres"]

# 4. gRPC (Implements gRPC Health Checking Protocol)
grpc:
  port: 9000
```

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 19A: Fixing a Misconfigured Probe Causing `CrashLoopBackOff`
On CKAD, you will frequently encounter a Deployment whose Pods keep restarting every 30 seconds or show `0/1 Ready`. Let's deploy one and debug it:

```yaml
# Save as /tmp/opencode/broken-probes.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: flaky-web
spec:
  replicas: 1
  selector:
    matchLabels: { app: flaky-web }
  template:
    metadata:
      labels: { app: flaky-web }
    spec:
      containers:
      - name: web
        image: nginx:1.25-alpine
        ports: [{ containerPort: 80 }]
        livenessProbe:
          httpGet:
            path: /healthz       # <-- nginx default doesn't have /healthz (returns 404)!
            port: 80
          initialDelaySeconds: 2
          periodSeconds: 3
          failureThreshold: 2
        readinessProbe:
          tcpSocket:
            port: 8080           # <-- nginx listens on 80, not 8080!
          periodSeconds: 3
```
```bash
k apply -f /tmp/opencode/broken-probes.yaml
sleep 12
k describe pod -l app=flaky-web | tail -n 15
```
**Expected Events Output:**
```text
Warning  Unhealthy  Pod/flaky-web-...  Readiness probe failed: dial tcp 10.244.1.40:8080: connect: connection refused
Warning  Unhealthy  Pod/flaky-web-...  Liveness probe failed: HTTP probe failed with statuscode: 404
Normal   Killing    Pod/flaky-web-...  Container web failed liveness probe, will be restarted
```
Notice how `k describe pod` tells you **both** bugs immediately (`statuscode: 404` on Liveness and `8080: connection refused` on Readiness)! Fix both (`path: /` and `port: 80`) and apply:
```bash
sed -i -e 's|path: /healthz|path: /|' -e 's|port: 8080|port: 80|' /tmp/opencode/broken-probes.yaml
k apply -f /tmp/opencode/broken-probes.yaml
k rollout status deploy/flaky-web
k delete deploy flaky-web
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 19.1: Protecting a Legacy Slow-Boot App with `startupProbe` (Exam Weight: 7%)
Create a Pod `legacy-erp` (`image: nginx:1.25-alpine`, port `80`) that:
- Can take up to **90 seconds** to start up, checked every **5 seconds** via HTTP GET `/` on port `80` (`startupProbe`: what must `failureThreshold` be so `failureThreshold * 5s = 90s`?).
- Once started, checks `livenessProbe` via HTTP GET `/` on port `80` every `10` seconds (`timeoutSeconds: 2`, `failureThreshold: 3`).
- Checks `readinessProbe` via `exec` command `["test", "-f", "/usr/share/nginx/html/index.html"]` every `5` seconds.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 19.1)</strong></summary>

`90s / 5s = 18`, so `startupProbe.failureThreshold: 18` and `periodSeconds: 5`:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: legacy-erp
spec:
  containers:
  - name: erp
    image: nginx:1.25-alpine
    ports: [{ containerPort: 80 }]
    startupProbe:
      httpGet: { path: /, port: 80 }
      periodSeconds: 5
      failureThreshold: 18
    livenessProbe:
      httpGet: { path: /, port: 80 }
      periodSeconds: 10
      timeoutSeconds: 2
      failureThreshold: 3
    readinessProbe:
      exec:
        command: ["test", "-f", "/usr/share/nginx/html/index.html"]
      periodSeconds: 5
```
```bash
k apply -f legacy-erp.yaml
k wait --for=condition=Ready pod/legacy-erp --timeout=30s
k delete pod legacy-erp $now
```
</details>

---

## Day 20 (Sat Dec 5 — Weekend Clinic): The 7-Scenario Broken Cluster Triage Clinic + CRDs & API Deprecations

### Part 1: The 7-Scenario Broken Cluster Triage Clinic
Apply this complete broken environment into namespace `clinic-20`, then diagnose and fix **all 7 broken workloads** so that every Pod is `Running (1/1 Ready)` and `clinic-svc` responds to `wget`:

```yaml
# Save as /tmp/opencode/clinic-20.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: clinic-20
---
# Bug 1: Why is pod-alpha stuck in CreateContainerConfigError?
apiVersion: v1
kind: Pod
metadata: { name: pod-alpha, namespace: clinic-20 }
spec:
  containers:
  - name: app
    image: busybox:1.36
    command: ["sleep", "3600"]
    env:
    - name: SECRET_KEY
      valueFrom:
        secretKeyRef: { name: alpha-sec, key: api-token }
---
apiVersion: v1
kind: Secret
metadata: { name: alpha-sec, namespace: clinic-20 }
stringData:
  token: "12345"                     # Hint: compare key name with pod-alpha!
---
# Bug 2: Why is pod-beta stuck in CrashLoopBackOff?
apiVersion: v1
kind: Pod
metadata: { name: pod-beta, namespace: clinic-20 }
spec:
  containers:
  - name: app
    image: busybox:1.36
    command: ["sh", "-c", "ech0 'Starting worker' && sleep 3600"]
---
# Bug 3: Why is pod-gamma rejected / failing security checks?
apiVersion: v1
kind: Pod
metadata: { name: pod-gamma, namespace: clinic-20 }
spec:
  securityContext:
    runAsNonRoot: true
  containers:
  - name: app
    image: busybox:1.36              # Hint: busybox runs as UID 0 by default!
    command: ["sleep", "3600"]
---
# Bug 4: Why is pod-delta stuck in Init:0/1 forever?
apiVersion: v1
kind: Pod
metadata: { name: pod-delta, namespace: clinic-20 }
spec:
  initContainers:
  - name: check-file
    image: busybox:1.36
    command: ["sh", "-c", "until [ -f /shared/ready.flag ]; do sleep 2; done"]
    volumeMounts: [{ name: v, mountPath: /shared }]
  containers:
  - name: app
    image: busybox:1.36
    command: ["sleep", "3600"]
    volumeMounts: [{ name: v, mountPath: /shared }]
  volumes:
  - { name: v, emptyDir: {} }
---
# Bug 5, 6, 7: Why does `wget http://clinic-svc` fail in clinic-20? (3 bugs in deploy + svc!)
apiVersion: apps/v1
kind: Deployment
metadata: { name: clinic-web, namespace: clinic-20 }
spec:
  replicas: 2
  selector:
    matchLabels: { app: clinic, role: web }
  template:
    metadata:
      labels: { app: clinic, role: web }
    spec:
      containers:
      - name: nginx
        image: nginx:1.25-apline     # Bug 5: image typo!
        ports: [{ containerPort: 80 }]
        readinessProbe:
          httpGet: { path: /, port: 81 } # Bug 6: wrong readiness port!
          periodSeconds: 2
---
apiVersion: v1
kind: Service
metadata: { name: clinic-svc, namespace: clinic-20 }
spec:
  selector: { app: clinic, role: frontend } # Bug 7: selector mismatch!
  ports: [{ port: 80, targetPort: 80 }]
```

<details>
<summary><strong>Click for Complete Triage Clinic Diagnosis & Fixes</strong></summary>

1. **`pod-alpha` (`CreateContainerConfigError`):** `k describe pod pod-alpha -n clinic-20` shows `couldn't find key api-token in Secret clinic-20/alpha-sec`. Fix: change `key: api-token` to `key: token` (or add `api-token` to `alpha-sec`).
2. **`pod-beta` (`CrashLoopBackOff`):** `k logs pod-beta -n clinic-20` shows `sh: ech0: not found`. Fix typo `ech0` -> `echo` and `k replace --force -f`.
3. **`pod-gamma` (`CreateContainerConfigError`):** `k describe pod pod-gamma -n clinic-20` shows `container has runAsNonRoot and image will run as root`. Fix: add `runAsUser: 1000` under `securityContext`.
4. **`pod-delta` (`Init:0/1`):** InitContainer waits for `/shared/ready.flag` before starting main container, so nothing ever creates `/shared/ready.flag`! Fix: change the `initContainer` command to `touch /shared/ready.flag`.
5. **`clinic-web` & `clinic-svc`:**
   - Bug 5: `nginx:1.25-apline` -> `nginx:1.25-alpine`
   - Bug 6: `readinessProbe` port `81` -> `80`
   - Bug 7: `clinic-svc` selector `role: frontend` -> `role: web`
```bash
k delete ns clinic-20
```
</details>

### Part 2: Custom Resource Definitions (CRDs), Operators & API Deprecations Lab
```yaml
# Save as /tmp/opencode/crd-backup.yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: backups.stable.example.com      # MUST be <plural>.<group>
spec:
  group: stable.example.com
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
              schedule:
                type: string
              retentionDays:
                type: integer
  scope: Namespaced
  names:
    plural: backups
    singular: backup
    kind: Backup
    shortNames:
    - bkp
---
apiVersion: stable.example.com/v1
kind: Backup
metadata:
  name: nightly-pg-backup
spec:
  schedule: "0 1 * * *"
  retentionDays: 14
```
```bash
k apply -f /tmp/opencode/crd-backup.yaml
k get crds | grep example
k get bkp                               # Uses the custom shortName 'bkp'!
k explain backup.spec                   # Your CRD now works with `kubectl explain`!
k delete -f /tmp/opencode/crd-backup.yaml
```

---

## Day 21 (Sun Dec 6 — Weekend Boss Lab): Week 3 Hardened Microservice Architecture Exam

1. **Part A (20 mins):** Retake [`quiz/CKAD_QUIZ_MCQ.html`](quiz/CKAD_QUIZ_MCQ.html) — aim for **23+/25 (92%+)**.
2. **Part B (60 mins — Hardened Microservice Exam):**
   Create namespace `pci-zone` and build the following interconnected resources without notes:
   - **Task 1 (Quota & Limits):** `ResourceQuota` `pci-quota` (`pods: 5, requests.cpu: 1, requests.memory: 1Gi, limits.cpu: 2, limits.memory: 2Gi`) and `LimitRange` `pci-limits` (`type: Container`, `defaultRequest: { cpu: 100m, memory: 128Mi }`, `default: { cpu: 200m, memory: 256Mi }`).
   - **Task 2 (Config & Secret):** Immutable Secret `pci-tls-key` (`KEY=rsa-4096-xyz`) and ConfigMap `pci-nginx-conf` containing an NGINX config listening on port **`8080`** (so it can run as non-root!).
   - **Task 3 (RBAC):** ServiceAccount `pci-runner`, Role `pci-reader` (`get,list,watch` on `pods,configmaps,events`), and RoleBinding `pci-bind`.
   - **Task 4 (Hardened Deployment + Probes):** Deployment `pci-gateway` (`replicas: 2`, `image: nginx:1.25-alpine`) using `pci-runner`:
     - Pod `securityContext`: `runAsNonRoot: true`, `runAsUser: 101`, `runAsGroup: 101`, `fsGroup: 101`.
     - Container `securityContext`: `readOnlyRootFilesystem: true`, `allowPrivilegeEscalation: false`, `capabilities: { drop: ["ALL"] }`.
     - Mount `pci-nginx-conf` at `/etc/nginx/conf.d`, plus `emptyDir` volumes at `/var/cache/nginx`, `/var/run`, and `/tmp`.
     - Inject `KEY` from `pci-tls-key` as env var `TLS_MASTER_KEY`.
     - Configure `startupProbe`, `readinessProbe`, and `livenessProbe` on HTTP `/` port `8080`.
     - Set explicit resources so the Pods get **`Guaranteed`** QoS (`cpu: 200m, memory: 128Mi` for both requests & limits).

<details>
<summary><strong>Click for Complete Day 21 Boss Lab Solution & Verification</strong></summary>

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: pci-zone
---
apiVersion: v1
kind: ResourceQuota
metadata: { name: pci-quota, namespace: pci-zone }
spec:
  hard:
    pods: "5"
    requests.cpu: "1"
    requests.memory: 1Gi
    limits.cpu: "2"
    limits.memory: 2Gi
---
apiVersion: v1
kind: LimitRange
metadata: { name: pci-limits, namespace: pci-zone }
spec:
  limits:
  - type: Container
    defaultRequest: { cpu: 100m, memory: 128Mi }
    default: { cpu: 200m, memory: 256Mi }
---
apiVersion: v1
kind: Secret
metadata: { name: pci-tls-key, namespace: pci-zone }
immutable: true
stringData:
  KEY: "rsa-4096-xyz"
---
apiVersion: v1
kind: ConfigMap
metadata: { name: pci-nginx-conf, namespace: pci-zone }
data:
  default.conf: |
    server {
        listen 8080;
        location / {
            return 200 "PCI Gateway OK\n";
        }
    }
---
apiVersion: v1
kind: ServiceAccount
metadata: { name: pci-runner, namespace: pci-zone }
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata: { name: pci-reader, namespace: pci-zone }
rules:
- apiGroups: [""]
  resources: ["pods", "configmaps", "events"]
  verbs: ["get", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata: { name: pci-bind, namespace: pci-zone }
subjects:
- kind: ServiceAccount
  name: pci-runner
  namespace: pci-zone
roleRef:
  kind: Role
  name: pci-reader
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: apps/v1
kind: Deployment
metadata: { name: pci-gateway, namespace: pci-zone }
spec:
  replicas: 2
  selector:
    matchLabels: { app: pci-gateway }
  template:
    metadata:
      labels: { app: pci-gateway }
    spec:
      serviceAccountName: pci-runner
      securityContext:
        runAsNonRoot: true
        runAsUser: 101
        runAsGroup: 101
        fsGroup: 101
      containers:
      - name: nginx
        image: nginx:1.25-alpine
        ports: [{ containerPort: 8080 }]
        env:
        - name: TLS_MASTER_KEY
          valueFrom:
            secretKeyRef: { name: pci-tls-key, key: KEY }
        resources:
          requests: { cpu: 200m, memory: 128Mi }
          limits: { cpu: 200m, memory: 128Mi }
        securityContext:
          readOnlyRootFilesystem: true
          allowPrivilegeEscalation: false
          capabilities:
            drop: ["ALL"]
        startupProbe:
          httpGet: { path: /, port: 8080 }
          periodSeconds: 3
          failureThreshold: 10
        readinessProbe:
          httpGet: { path: /, port: 8080 }
          periodSeconds: 5
        livenessProbe:
          httpGet: { path: /, port: 8080 }
          periodSeconds: 10
        volumeMounts:
        - { name: conf, mountPath: /etc/nginx/conf.d }
        - { name: cache, mountPath: /var/cache/nginx }
        - { name: run, mountPath: /var/run }
        - { name: tmp, mountPath: /tmp }
      volumes:
      - name: conf
        configMap: { name: pci-nginx-conf }
      - { name: cache, emptyDir: {} }
      - { name: run, emptyDir: {} }
      - { name: tmp, emptyDir: {} }
```
```bash
k apply -f pci-zone.yaml
k rollout status deploy/pci-gateway -n pci-zone
k get pods -n pci-zone -o custom-columns=NAME:.metadata.name,STATUS:.status.phase,QOS:.status.qosClass
k auth can-i list configmaps -n pci-zone --as=system:serviceaccount:pci-zone:pci-runner  # yes
k delete ns pci-zone
```
</details>


---

# Week 4 (Dec 7–13): Services, Networking & Full Mock Exam #1

**Weekly Objective:** Conquer **Services and Networking (20% of CKAD)**, master fast `kubectl` data extraction (`jsonpath` / `custom-columns`), and complete **Full 2-Hour Simulated CKAD Exam #1**.

---

## Day 22 (Mon Dec 7): Kubernetes Networking Model, CoreDNS & Service Types (`ClusterIP`, `NodePort`, `LoadBalancer`, `ExternalName`)

### 1. Deep-Dive Theory & Mental Model

#### End-to-End Packet Flow: From Client Pod to Backend Pod via Service & `kube-proxy`

```text
+-----------------------+
| Client Pod            |
| `curl http://api:80`  |
+-----------+-----------+
            │ 1. DNS Query `api.default.svc.cluster.local` -> CoreDNS (10.96.0.10:53)
            │    Returns Service Virtual ClusterIP: 10.104.22.88
            ▼
+-------------------------------------------------------------------------+
| Node Kernel (iptables / IPVS programmed by `kube-proxy`)                |
|  - Intercepts destination 10.104.22.88:80                               |
|  - Picks one Ready backend from Service's `Endpoints` / `EndpointSlice` |
|  - Rewrites DNAT to Pod IP `10.244.2.19:8080` (`targetPort`)            |
+-----------+-------------------------------------------------------------+
            │
            ▼
+-------------------------------------------------------------------------+
| Backend Pod (IP: 10.244.2.19, Container listening on `containerPort: 8080`)
+-------------------------------------------------------------------------+
```

#### Annotated Service YAML with Named Ports & NodePort
Using **named ports** (e.g., `name: http-api` on the Pod's `containerPort` and `targetPort: http-api` on the Service) decouples the Service from port number changes on the Pod!

```yaml
apiVersion: v1
kind: Service
metadata:
  name: payment-svc
  namespace: prod
spec:
  type: NodePort                   # ClusterIP (default) | NodePort | LoadBalancer | ExternalName
  selector:
    app: payment                   # Selects Ready Pods matching ALL listed labels
    tier: backend
  ports:
  - name: http
    protocol: TCP
    port: 80                       # Port the Service listens on internally (ClusterIP:80)
    targetPort: http-api           # Port (or named port!) on the Pod (e.g., 8080)
    nodePort: 30080                # Static port opened on EVERY Node IP (30000-32767)
```

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 22A: Named Ports, Cross-Namespace DNS, `ExternalName` & Endpoint Debugging
```bash
k create ns net-front && k create ns net-back

# 1. Create a backend Deployment in 'net-back' with a named containerPort 'web-port' (8080)
cat << 'EOF' | k apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: catalog-api
  namespace: net-back
spec:
  replicas: 2
  selector:
    matchLabels: { app: catalog }
  template:
    metadata:
      labels: { app: catalog }
    spec:
      containers:
      - name: api
        image: hashicorp/http-echo:0.2.3
        args: ["-listen=:8080", "-text=CATALOG_OK"]
        ports:
        - name: web-port
          containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: catalog-svc
  namespace: net-back
spec:
  selector: { app: catalog }
  ports:
  - port: 80
    targetPort: web-port
EOF

k wait --for=condition=Available deploy/catalog-api -n net-back --timeout=60s

# 2. Create an ExternalName Service in 'net-front' pointing to 'catalog-svc.net-back.svc.cluster.local'
k create service externalname local-catalog -n net-front \
  --external-name=catalog-svc.net-back.svc.cluster.local

# 3. Curl 'http://local-catalog' from a Pod inside 'net-front' — it resolves across namespaces via CNAME!
k run client -n net-front --image=busybox:1.36 --rm -it --restart=Never -- \
  wget -qO- http://local-catalog
# Expected Output: CATALOG_OK

k delete ns net-front net-back
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 22.1: Multi-Port Service & Broken Endpoint Triage (Exam Weight: 7%)
In namespace `telecom`:
1. Create a Pod `dual-server` (`nginx:1.25-alpine`, labels `app=telecom,role=edge`) exposing container ports `80` (named `http`) and `443` (named `https`).
2. Create a `NodePort` Service named `edge-svc` in `telecom` that exposes:
   - Service port `8080` -> `targetPort: http` (`nodePort: 30180`)
   - Service port `8443` -> `targetPort: https` (`nodePort: 30443`)
3. Verify `k get endpoints edge-svc -n telecom` lists both ports (`:80` and `:443`) for `dual-server`'s IP.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 22.1)</strong></summary>

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: telecom
---
apiVersion: v1
kind: Pod
metadata:
  name: dual-server
  namespace: telecom
  labels: { app: telecom, role: edge }
spec:
  containers:
  - name: nginx
    image: nginx:1.25-alpine
    ports:
    - { name: http, containerPort: 80 }
    - { name: https, containerPort: 443 }
---
apiVersion: v1
kind: Service
metadata:
  name: edge-svc
  namespace: telecom
spec:
  type: NodePort
  selector: { app: telecom, role: edge }
  ports:
  - name: http
    port: 8080
    targetPort: http
    nodePort: 30180
  - name: https
    port: 8443
    targetPort: https
    nodePort: 30443
```
```bash
k apply -f telecom.yaml
k get endpoints edge-svc -n telecom
k delete ns telecom
```
</details>

---

## Day 23 (Tue Dec 8): Ingress Controllers, `IngressClass`, Path/Host Routing & HTTPS/TLS Termination

### 1. Deep-Dive Theory & Mental Model

#### How Ingress Routes External HTTP/HTTPS Traffic into Services

```text
                       [External Client]
            https://shop.example.com/api/orders
                              │
                              ▼
             +---------------------------------+
             | Ingress Controller Pod (NGINX)  |
             |  - Terminates TLS (`shop-tls`)  |
             |  - Evaluates Ingress Rules      |
             +----------------+----------------+
                              │
        +---------------------+---------------------+
        │ Host: shop.example.com                    │ Host: admin.example.com
        ├── Path: /api (Prefix)                     └── Path: / (Exact)
        ▼                                           ▼
 [Service: api-svc:8080]                    [Service: admin-svc:9000]
        │                                           │
        ▼                                           ▼
   [API Pods]                                 [Admin Pods]
```

#### Annotated `networking.k8s.io/v1` Ingress YAML
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ecommerce-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /   # Common exam annotation!
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  ingressClassName: nginx                           # Or via annotation kubernetes.io/ingress.class
  defaultBackend:                                   # Fallback when no host/path matches
    service:
      name: fallback-svc
      port:
        number: 80
  tls:
  - hosts:
    - shop.example.com
    secretName: shop-tls-secret                     # Type: kubernetes.io/tls
  rules:
  - host: shop.example.com
    http:
      paths:
      - path: /api
        pathType: Prefix                            # Prefix | Exact | ImplementationSpecific
        backend:
          service:
            name: api-svc
            port:
              number: 8080
```

#### The Fastest Way to Generate Ingress YAML on the Exam (`k create ingress`)
Never type the deep `rules[].http.paths[].backend.service.port.number` hierarchy by hand! Use `kubectl create ingress`:
```bash
# Note: Adding '*' after the path sets pathType=Prefix; omitting '*' sets pathType=Exact!
k create ingress my-ing \
  --class=nginx \
  --default-backend=fallback-svc:80 \
  --rule="shop.example.com/api*=api-svc:8080,tls=shop-tls" \
  --rule="shop.example.com/login=auth-svc:9000,tls=shop-tls" \
  --annotation="nginx.ingress.kubernetes.io/rewrite-target=/" \
  $do
```

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 23A: End-to-End Live Ingress Routing + Path Rewrite + HTTPS Test
```bash
k create ns ing-lab

# 1. Create two echo backend Deployments + Services (video-svc and pay-svc)
k create deploy video -n ing-lab --image=hashicorp/http-echo:0.2.3 -- -listen=:5678 -text="VIDEO_STREAM_SERVICE"
k expose deploy video -n ing-lab --name=video-svc --port=80 --target-port=5678

k create deploy pay -n ing-lab --image=hashicorp/http-echo:0.2.3 -- -listen=:5678 -text="PAYMENT_GATEWAY_SERVICE"
k expose deploy pay -n ing-lab --name=pay-svc --port=80 --target-port=5678

# 2. Generate TLS cert & Secret
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout /tmp/opencode/ing.key -out /tmp/opencode/ing.crt -subj "/CN=media.local"
k create secret tls media-tls -n ing-lab --cert=/tmp/opencode/ing.crt --key=/tmp/opencode/ing.key

# 3. Create the Ingress imperatively!
k create ingress media-ing -n ing-lab \
  --class=nginx \
  --annotation="nginx.ingress.kubernetes.io/rewrite-target=/" \
  --rule="media.local/watch*=video-svc:80,tls=media-tls" \
  --rule="media.local/pay*=pay-svc:80,tls=media-tls"

k describe ingress media-ing -n ing-lab
k delete ns ing-lab
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 23.1: Multi-Host & Multi-Path Ingress with Rewrite (Exam Weight: 8%)
In namespace `gateway-ns`, create an Ingress named `corp-ingress` (`ingressClassName: nginx`, annotation `nginx.ingress.kubernetes.io/ssl-redirect: "false"`) that routes:
- Host `hr.corp.internal`:
  - `/payroll` (`pathType: Prefix`) -> Service `payroll-svc` port `8080`
  - `/benefits` (`pathType: Exact`) -> Service `benefits-svc` port `80`
- Host `status.corp.internal`:
  - `/` (`pathType: Prefix`) -> Service `status-svc` port `3000`
- Configure TLS for `hr.corp.internal` using Secret `hr-tls-cert`.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 23.1)</strong></summary>

```bash
k create ns gateway-ns
k create ingress corp-ingress -n gateway-ns \
  --class=nginx \
  --annotation='nginx.ingress.kubernetes.io/ssl-redirect=false' \
  --rule="hr.corp.internal/payroll*=payroll-svc:8080,tls=hr-tls-cert" \
  --rule="hr.corp.internal/benefits=benefits-svc:80,tls=hr-tls-cert" \
  --rule="status.corp.internal/*=status-svc:3000" \
  $do > /tmp/opencode/corp-ing.yaml

k apply -f /tmp/opencode/corp-ing.yaml
k describe ingress corp-ingress -n gateway-ns
k delete ns gateway-ns
```
</details>

---

## Day 24 (Wed Dec 9): NetworkPolicies Part 1 — Default Isolation, Ingress Rules & The `AND` vs `OR` Selector Trap

### 1. Deep-Dive Theory & Mental Model

#### Visual Guide to the Infamous `AND` vs `OR` YAML Dash Trap

```text
CASE 1: SINGLE DASH (-) = LOGICAL AND
=====================================
  ingress:
  - from:
    - namespaceSelector:           <-- ONE dash starts a single rule element
        matchLabels:
          env: prod
      podSelector:                 <-- NO dash! Same rule element = AND!
        matchLabels:
          role: api

  Result: Allows traffic ONLY from Pods labeled `role=api` that live INSIDE
          namespaces labeled `env=prod`.


CASE 2: TWO DASHES (- ... -) = LOGICAL OR
=========================================
  ingress:
  - from:
    - namespaceSelector:           <-- Dash #1: Rule element 1
        matchLabels:
          env: prod
    - podSelector:                 <-- Dash #2: Separate rule element = OR!
        matchLabels:
          role: api

  Result: Allows traffic from ANY Pod in namespace `env=prod`
          OR from Pods labeled `role=api` in the NetworkPolicy's OWN namespace!
```

#### The 4 Default Isolation Patterns

| Goal | `podSelector` | `policyTypes` | `ingress` / `egress` field |
| :--- | :--- | :--- | :--- |
| **Deny ALL Ingress** in namespace | `{}` | `["Ingress"]` | Omit `ingress:` (or set `ingress: []`) |
| **Allow ALL Ingress** in namespace | `{}` | `["Ingress"]` | `ingress: [{}]` |
| **Deny ALL Egress** in namespace | `{}` | `["Egress"]` | Omit `egress:` (or set `egress: []`) |
| **Deny ALL Ingress + Egress** | `{}` | `["Ingress", "Egress"]` | Omit both `ingress:` and `egress:` |

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 24A: Proving `AND` vs `OR` Live with Calico
```bash
k create ns np-target && k create ns np-trusted && k create ns np-untrusted
k label ns np-trusted access=granted
k label ns np-untrusted access=denied

# 1. Deploy target nginx server in 'np-target'
k run vault-server -n np-target --image=nginx:1.25-alpine --labels="app=vault" --port=80
k expose pod vault-server -n np-target --port=80

# 2. Apply an AND NetworkPolicy: ONLY Pods with role=auditor INSIDE namespaces with access=granted!
cat << 'EOF' | k apply -f -
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: vault-and-policy
  namespace: np-target
spec:
  podSelector:
    matchLabels:
      app: vault
  policyTypes:
  - Ingress
  ingress:
  - from:
    - namespaceSelector:
        matchLabels:
          access: granted
      podSelector:
        matchLabels:
          role: auditor
    ports:
    - protocol: TCP
      port: 80
EOF

# Test 1: Pod in 'np-trusted' WITH label role=auditor -> SUCCEEDS!
k run ok-pod -n np-trusted --image=busybox:1.36 --labels="role=auditor" --rm -it --restart=Never -- \
  wget -qO- --timeout=3 http://vault-server.np-target.svc.cluster.local | head -n 3

# Test 2: Pod in 'np-trusted' WITHOUT label role=auditor -> BLOCKED (Times out)!
k run blocked-1 -n np-trusted --image=busybox:1.36 --labels="role=guest" --rm -it --restart=Never -- \
  wget -qO- --timeout=2 http://vault-server.np-target.svc.cluster.local || echo "Blocked as expected!"

# Test 3: Pod in 'np-untrusted' WITH label role=auditor -> BLOCKED (Times out)!
k run blocked-2 -n np-untrusted --image=busybox:1.36 --labels="role=auditor" --rm -it --restart=Never -- \
  wget -qO- --timeout=2 http://vault-server.np-target.svc.cluster.local || echo "Blocked as expected!"

k delete ns np-target np-trusted np-untrusted
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 24.1: Multi-Rule Ingress NetworkPolicy (Exam Weight: 8%)
In namespace `database-ns`, Pod `postgres` has label `tier=db` and listens on TCP port `5432`. Create a NetworkPolicy `db-access` in `database-ns` that allows incoming TCP traffic on port `5432` to `tier=db` from:
1. Any Pod with label `tier=backend` in the **same** namespace (`database-ns`), **OR**
2. Pods with label `role=migrator` running inside any namespace labeled `env=ci`.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 24.1)</strong></summary>

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: db-access
  namespace: database-ns
spec:
  podSelector:
    matchLabels:
      tier: db
  policyTypes:
  - Ingress
  ingress:
  - from:
    # Rule item 1: tier=backend in database-ns
    - podSelector:
        matchLabels:
          tier: backend
    # Rule item 2 (separate dash = OR): role=migrator AND namespace env=ci
    - namespaceSelector:
        matchLabels:
          env: ci
      podSelector:
        matchLabels:
          role: migrator
    ports:
    - protocol: TCP
      port: 5432
```
</details>

---

## Day 25 (Thu Dec 10): NetworkPolicies Part 2 — Egress Rules, `ipBlock` CIDRs & The CoreDNS Port 53 Trap

### 1. Deep-Dive Theory & Mental Model

#### Why Egress Policies Break Service Calls (The Port 53 Trap)
When a Pod runs `curl http://payment-svc:8080`, two network connections happen in sequence:
1. **DNS Lookup (`UDP/TCP port 53`):** Pod queries CoreDNS in `kube-system` to resolve `payment-svc` -> IP.
2. **HTTP Request (`TCP port 8080`):** Pod connects to the resolved Pod IP on port `8080`.

If your NetworkPolicy lists `policyTypes: ["Egress"]` and only allows port `8080`, **Step 1 (DNS) is dropped by the firewall**, and `curl http://payment-svc:8080` fails with `Could not resolve host`!

```yaml
# Always include this Egress rule when restricting Egress on Pods that use DNS names:
  egress:
  - ports:
    - protocol: UDP
      port: 53
    - protocol: TCP
      port: 53
```
*(Notice: if you specify `ports: [53]` without a `to:` block, it allows port 53 to any destination IP, which is the fastest and safest way to allow cluster DNS on the exam unless the prompt specifically restricts DNS to `kube-system`!)*

---

### 2. Guided Step-by-Step Walkthrough

#### Lab 25A: Combined Ingress + Egress NetworkPolicy with `ipBlock`
```yaml
# Save as /tmp/opencode/full-netpol.yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: api-zero-trust
  namespace: default
spec:
  podSelector:
    matchLabels:
      app: order-api
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: api-gateway
    ports:
    - protocol: TCP
      port: 8080
  egress:
  # 1. Allow DNS resolution
  - ports:
    - protocol: UDP
      port: 53
    - protocol: TCP
      port: 53
  # 2. Allow talking to internal Postgres Pods on 5432
  - to:
    - podSelector:
        matchLabels:
          app: postgres
    ports:
    - protocol: TCP
      port: 5432
  # 3. Allow talking to external payment gateway subnet except 10.100.5.0/24
  - to:
    - ipBlock:
        cidr: 10.100.0.0/16
        except:
        - 10.100.5.0/24
    ports:
    - protocol: TCP
      port: 443
```
```bash
k apply -f /tmp/opencode/full-netpol.yaml
k describe netpol api-zero-trust
k delete netpol api-zero-trust
```

---

### 3. Exam-Style Scenarios (Your Turn)

#### Scenario 25.1: Fixing a Broken Microservice Chain Using Labels Instead of Editing the NetworkPolicy (Exam Weight: 8%)
Sometimes an exam prompt says:
> *"Do **NOT** modify the existing NetworkPolicy ` NP` in namespace `secure-app`. Instead, adjust the labels on Pods `frontend` and `backend` so `frontend` can reach `backend` on port `80`."*

Practice this right now:
```bash
k create ns secure-app
k run backend -n secure-app --image=nginx:1.25-alpine --port=80 --labels="tier=api"
k expose pod backend -n secure-app --port=80
k run frontend -n secure-app --image=busybox:1.36 --labels="tier=ui" -- sleep 3600

cat << 'EOF' | k apply -f -
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: locked-backend
  namespace: secure-app
spec:
  podSelector:
    matchLabels:
      tier: api
  policyTypes: ["Ingress", "Egress"]
  ingress:
  - from:
    - podSelector:
        matchLabels:
          access-backend: "true"
    ports: [{ protocol: TCP, port: 80 }]
---
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: locked-frontend
  namespace: secure-app
spec:
  podSelector:
    matchLabels:
      tier: ui
  policyTypes: ["Egress"]
  egress:
  - ports: [{ protocol: UDP, port: 53 }, { protocol: TCP, port: 53 }]
  - to:
    - podSelector:
        matchLabels:
          accept-ui: "true"
    ports: [{ protocol: TCP, port: 80 }]
EOF
```
Without editing either NetworkPolicy, add the required labels to `frontend` and `backend` so `k exec -n secure-app frontend -- wget -qO- --timeout=2 http://backend` succeeds!

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 25.1)</strong></summary>

- `locked-backend` requires incoming Pods to have `access-backend=true` -> label `frontend`!
- `locked-frontend` requires destination Pods to have `accept-ui=true` -> label `backend`!
```bash
k label pod frontend -n secure-app access-backend=true
k label pod backend -n secure-app accept-ui=true
k exec -n secure-app frontend -- wget -qO- --timeout=2 http://backend | head -n 4
k delete ns secure-app
```
</details>

---

## Day 26 (Fri Dec 11): Advanced `kubectl` Extraction — JSONPath, Custom Columns, Field Selectors & Sorting

### 1. Deep-Dive Theory & Mental Model

#### The Essential JSONPath & Custom-Columns Cookbook
| Task | Exact Command |
| :--- | :--- |
| Custom Columns (Pod Name + Node + QoS) | `k get pods -A -o custom-columns=POD:.metadata.name,NODE:.spec.nodeName,QOS:.status.qosClass` |
| All Container Images across all Pods | `k get pods -A -o jsonpath='{.items[*].spec.containers[*].image}{"\n"}'` |
| Newline-separated list using `range` | `k get pods -A -o jsonpath='{range .items[*]}{.metadata.namespace}{"/"}{.metadata.name}{"\n"}{end}'` |
| Filter array by condition `[?(@.key=="val")]` | `k get nodes -o jsonpath='{.items[*].status.addresses[?(@.type=="InternalIP")].address}{"\n"}'` |
| Sort Pods by restart count | `k get pods -A --sort-by='.status.containerStatuses[0].restartCount'` |
| Sort PVCs by storage capacity | `k get pvc -A --sort-by='.spec.resources.requests.storage'` |
| Field selector (non-Running Pods only) | `k get pods -A --field-selector=status.phase!=Running` |

---

### 2. Guided Step-by-Step Walkthrough & Scenario

Run each command in the table above against your cluster, then solve this exam scenario:
- **Scenario 26.1:** Find all Pods across all namespaces that use `serviceAccountName: default`, sorted by `.metadata.creationTimestamp`, and output `NAMESPACE` and `NAME` using `-o custom-columns` to `/tmp/opencode/default-sa-pods.txt`.

<details>
<summary><strong>Click for Step-by-Step Solution & Verification (Scenario 26.1)</strong></summary>

```bash
k get pods -A --field-selector=spec.serviceAccountName=default \
  --sort-by=.metadata.creationTimestamp \
  -o custom-columns=NAMESPACE:.metadata.namespace,NAME:.metadata.name > /tmp/opencode/default-sa-pods.txt
cat /tmp/opencode/default-sa-pods.txt
```
</details>

---

## Day 27 (Sat Dec 12 — Weekend Mock #1): Full 15-Task Simulated CKAD Exam #1 (Built-In) + Killer.sh Session 1

> **If you have already registered for CKAD:** You can use today for **Killer.sh Session 1**, OR run this **complete, built-in 15-task simulated CKAD exam** right on your local cluster first!
> **Strict Rules:** Set a **120-minute timer**. Close `CKAD_NOTES.md`. Use only `kubectl explain` and `https://kubernetes.io/docs/`. Passing score: **66/100 points**.

### Step 0: Apply the Mock Exam #1 Cluster Environment
```yaml
# Save as /tmp/opencode/mock1-setup.yaml and run: k apply -f /tmp/opencode/mock1-setup.yaml
apiVersion: v1
kind: Namespace
metadata: { name: mock1-alpha }
---
apiVersion: v1
kind: Namespace
metadata: { name: mock1-beta, labels: { trusted: "yes" } }
---
apiVersion: apps/v1
kind: Deployment
metadata: { name: legacy-orders, namespace: mock1-alpha }
spec:
  replicas: 2
  selector: { matchLabels: { app: orders } }
  template:
    metadata: { labels: { app: orders } }
    spec:
      containers:
      - name: orders
        image: nginx:1.24-alpine
        ports: [{ containerPort: 80 }]
---
apiVersion: v1
kind: Service
metadata: { name: orders-svc, namespace: mock1-alpha }
spec:
  selector: { app: order-wrong }
  ports: [{ port: 80, targetPort: 80 }]
---
apiVersion: apps/v1
kind: Deployment
metadata: { name: crashing-worker, namespace: mock1-alpha }
spec:
  replicas: 1
  selector: { matchLabels: { app: crasher } }
  template:
    metadata: { labels: { app: crasher } }
    spec:
      securityContext:
        runAsNonRoot: true
      containers:
      - name: worker
        image: busybox:1.36
        command: ["sh", "-c", "sleep 3600"]
```

---

### The 15 Exam Tasks (Total: 100 Points)

- **Question 1 (5 pts — Pod & Resources):** In `mock1-alpha`, create a Pod `cache-pod` (`redis:7.2-alpine`) with label `tier=cache` and `Guaranteed` QoS using `cpu: 100m` and `memory: 128Mi`.
- **Question 2 (7 pts — Multi-Container Sidecar):** In `mock1-alpha`, create a Pod `web-logger` with main container `app` (`busybox:1.36`, writing `"$(date) OK"` every 3s to `/var/log/app/out.log`) and sidecar `shipper` (`busybox:1.36`, running `tail -n+1 -F /var/log/app/out.log`), sharing an `emptyDir` volume `log-dir`.
- **Question 3 (6 pts — CronJob):** In `mock1-alpha`, create a CronJob `backup-job` (`schedule: "0 */4 * * *"`, `concurrencyPolicy: Forbid`, `successfulJobsHistoryLimit: 2`, `failedJobsHistoryLimit: 1`) running `busybox:1.36` with command `["sh", "-c", "echo backup done"]`. Trigger a manual Job named `backup-test` from it.
- **Question 4 (8 pts — PV & PVC):** Create a cluster PV `mock1-pv` (`500Mi`, `ReadWriteOnce`, `storageClassName: mock1-sc`, `Retain`, `hostPath: /tmp/mock1-pv`). In `mock1-alpha`, create PVC `mock1-pvc` (`300Mi`, `ReadWriteOnce`, `storageClassName: mock1-sc`) and mount it at `/data` in Pod `data-writer` (`busybox:1.36`, sleeping `3600`).
- **Question 5 (6 pts — Rolling Update & Rollback):** In `mock1-alpha`, update Deployment `legacy-orders` so `maxSurge: 1` and `maxUnavailable: 0`, update its image to `nginx:1.25-alpine`, and scale it to `3` replicas.
- **Question 6 (8 pts — Canary Deployment):** In `mock1-alpha`, fix Service `orders-svc` so it selects `app: orders`. Then create a Canary Deployment `legacy-orders-canary` (`nginx:1.26-alpine`, `replicas: 1`, labels `app: orders, track: canary`) so that 25% of requests to `orders-svc` (1 out of 4 total Pods) hit the canary.
- **Question 7 (6 pts — Pod Scheduling):** Label one of your cluster nodes `env=staging`. In `mock1-alpha`, create a Pod `stage-runner` (`nginx:1.25-alpine`) with `nodeAffinity` (`requiredDuringSchedulingIgnoredDuringExecution`) requiring `env In [staging, prod]`.
- **Question 8 (6 pts — HPA):** In `mock1-alpha`, add CPU requests (`100m`) to `legacy-orders` and create an `autoscaling/v2` HPA `orders-hpa` targeting `legacy-orders` (`minReplicas: 2`, `maxReplicas: 6`, CPU utilization `75%`).
- **Question 9 (7 pts — ConfigMap & Secret):** In `mock1-alpha`, create ConfigMap `cm-settings` (`PORT=8080`, `settings.ini="debug=true"`) and Secret `sec-token` (`TOKEN=secret-abc`). Create Pod `cfg-pod` (`busybox:1.36`, sleeping `3600`) exposing `PORT` and `TOKEN` as env vars, and mounting only `settings.ini` into `/etc/cfg/settings.ini`.
- **Question 10 (6 pts — Quota & LimitRange):** In `mock1-beta`, create a `LimitRange` `beta-limits` (`Container` default `cpu: 200m, memory: 128Mi`, defaultRequest `cpu: 100m, memory: 64Mi`) and a `ResourceQuota` `beta-quota` (`requests.cpu: 1, limits.memory: 1Gi, pods: 5`).
- **Question 11 (7 pts — Fix Broken SecurityContext):** In `mock1-alpha`, Deployment `crashing-worker` is failing to create containers due to its `securityContext`. Fix `crashing-worker` so it runs as UID `1000`, GID `2000`, with `readOnlyRootFilesystem: true`, `allowPrivilegeEscalation: false`, and drops `ALL` capabilities.
- **Question 12 (7 pts — RBAC):** In `mock1-alpha`, create ServiceAccount ` job-runner`, Role `job-manager` (`get,list,watch,create,delete` on `jobs` in `batch` and `pods` in core `""`), and RoleBinding `job-runner-bind`.
- **Question 13 (7 pts — Probes):** In `mock1-alpha`, add an `httpGet` `readinessProbe` (`path: /`, `port: 80`, `periodSeconds: 5`) and a `tcpSocket` `livenessProbe` (`port: 80`, `initialDelaySeconds: 5`, `periodSeconds: 10`) to Deployment `legacy-orders`.
- **Question 14 (7 pts — Ingress):** In `mock1-alpha`, create Ingress `alpha-ing` (`ingressClassName: nginx`) routing `http://alpha.exam.local/orders` (`Prefix`) to `orders-svc:80`.
- **Question 15 (7 pts — NetworkPolicy):** In `mock1-alpha`, create a NetworkPolicy `orders-netpol` selecting `app: orders` (`policyTypes: ["Ingress"]`) that allows TCP port `80` **only** from Pods in `mock1-alpha` labeled `role=client` **OR** from Pods in namespaces labeled `trusted=yes` that also have Pod label `role=auditor`.

<details>
<summary><strong>Click for Mock Exam #1 Grading Rubric & Verification Script</strong></summary>

Run this verification script to grade your Mock Exam #1:
```bash
echo "=== Q1: Guaranteed QoS ==="
k get pod cache-pod -n mock1-alpha -o jsonpath='{.status.qosClass}{"\n"}'

echo "=== Q2: Sidecar Logs ==="
k logs web-logger -n mock1-alpha -c shipper | head -n 2

echo "=== Q3: CronJob + Manual Job ==="
k get cronjob backup-job -n mock1-alpha && k get job backup-test -n mock1-alpha

echo "=== Q4: Bound PVC ==="
k get pvc mock1-pvc -n mock1-alpha

echo "=== Q5 & Q6: Canary Endpoints (should show 4 IPs) ==="
k get endpoints orders-svc -n mock1-alpha

echo "=== Q7: Node Affinity ==="
k get pod stage-runner -n mock1-alpha -o wide

echo "=== Q8: HPA ==="
k get hpa orders-hpa -n mock1-alpha

echo "=== Q9: ConfigMap & Secret Pod ==="
k exec cfg-pod -n mock1-alpha -- sh -c 'printenv PORT TOKEN && cat /etc/cfg/settings.ini'

echo "=== Q10: Quota & LimitRange ==="
k get quota,limitrange -n mock1-beta

echo "=== Q11: Fixed Crashing Worker ==="
k get deploy crashing-worker -n mock1-alpha

echo "=== Q12: RBAC Check (should print yes) ==="
k auth can-i create jobs.batch -n mock1-alpha --as=system:serviceaccount:mock1-alpha:job-runner

echo "=== Q13: Probes on legacy-orders ==="
k describe deploy legacy-orders -n mock1-alpha | grep -E "Liveness|Readiness"

echo "=== Q14: Ingress ==="
k get ingress alpha-ing -n mock1-alpha

echo "=== Q15: NetworkPolicy ==="
k describe netpol orders-netpol -n mock1-alpha
```
Clean up after Day 28 review: `k delete ns mock1-alpha mock1-beta && k delete pv mock1-pv`
</details>

---

## Day 28 (Sun Dec 13 — Weekend Post-Mortem): Mock #1 Deep Re-Solve & Gap Elimination

1. **Step 1 (Score Analysis):** Go through every task from Day 27 (or Killer.sh Session 1) where you lost points or spent more than **6 minutes**.
2. **Step 2 (The "3x Rule"):** For each missed/slow task, delete the resource and recreate it **3 times in a row from a blank terminal** until you can finish it in under **3.5 minutes** without checking notes.
3. **Step 3:** Record your top lessons in the [Personal Mistake & Speed Tracker](#appendix-personal-mistake--speed-tracker).


---

# Week 5 (Dec 14–20): High-Speed Clinics, Mock Exam #2 & Exam Day

**Weekly Objective:** Eliminate hesitation on the hardest YAML structures (NetworkPolicies, Ingress, RBAC, SecurityContext, Multi-Container Storage), complete **Mock Exam #2**, and walk into **Sunday, December 20** well-rested and razor-sharp.

---

## Day 29 (Mon Dec 14): High-Speed Clinic #1 — NetworkPolicy + Ingress + Service Gauntlet

Set a **40-minute timer** and complete all 4 networking tasks below in namespace `clinic-net` ( target: < 8 minutes per task).

```bash
k create ns clinic-net
k create ns external-clients && k label ns external-clients partner=verified
```

### Gauntlet 29.1: Default-Deny + Precision Ingress & Egress NetworkPolicy
In `clinic-net`:
1. Create a **default-deny-all** NetworkPolicy `deny-all` that blocks all Ingress and Egress traffic for all Pods in `clinic-net`.
2. Create a second NetworkPolicy `allow-checkout` selecting Pods with `app=checkout` (`policyTypes: ["Ingress", "Egress"]`) that:
   - Allows **Ingress** on TCP `8443` from Pods in namespaces labeled `partner=verified` **AND** having Pod label `client=api-gw`.
   - Allows **Egress** on UDP/TCP `53` (to anywhere for DNS) and TCP `6379` to Pods labeled `app=redis` in `clinic-net`.

### Gauntlet 29.2: Multi-Host TLS Ingress with Custom Default Backend
In `clinic-net`, create an Ingress `edge-router` (`ingressClassName: nginx`):
- Default backend: Service `error-page-svc` port `8080`.
- TLS for `api.bank.local` and `portal.bank.local` using Secret `bank-tls`.
- Host `api.bank.local` path `/v2` (`Prefix`) -> `v2-svc:8080`.
- Host `portal.bank.local` path `/login` (`Exact`) -> `auth-svc:9000`.

### Gauntlet 29.3: Fixing a Broken 3-Tier Service Chain
Apply this broken 2-tier setup in `clinic-net` and fix the **2 bugs** preventing `ui-pod` from reaching `backend-svc` on port `80`:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: backend-pod
  namespace: clinic-net
  labels: { app: backend, env: prod }
spec:
  containers:
  - name: web
    image: hashicorp/http-echo:0.2.3
    args: ["-listen=:9090", "-text=BACKEND_ALIVE"]
    ports: [{ name: http-alt, containerPort: 9090 }]
---
apiVersion: v1
kind: Service
metadata:
  name: backend-svc
  namespace: clinic-net
spec:
  selector: { app: backend, env: staging }   # Bug A
  ports:
  - port: 80
    targetPort: 8080                         # Bug B
```

<details>
<summary><strong>Click for Complete Day 29 Gauntlet Solutions</strong></summary>

```yaml
# Gauntlet 29.1 Solution
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata: { name: deny-all, namespace: clinic-net }
spec:
  podSelector: {}
  policyTypes: ["Ingress", "Egress"]
---
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata: { name: allow-checkout, namespace: clinic-net }
spec:
  podSelector:
    matchLabels: { app: checkout }
  policyTypes: ["Ingress", "Egress"]
  ingress:
  - from:
    - namespaceSelector:
        matchLabels: { partner: verified }
      podSelector:
        matchLabels: { client: api-gw }
    ports: [{ protocol: TCP, port: 8443 }]
  egress:
  - ports: [{ protocol: UDP, port: 53 }, { protocol: TCP, port: 53 }]
  - to:
    - podSelector:
        matchLabels: { app: redis }
    ports: [{ protocol: TCP, port: 6379 }]
---
# Gauntlet 29.2 Solution
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata: { name: edge-router, namespace: clinic-net }
spec:
  ingressClassName: nginx
  defaultBackend:
    service: { name: error-page-svc, port: { number: 8080 } }
  tls:
  - hosts: ["api.bank.local", "portal.bank.local"]
    secretName: bank-tls
  rules:
  - host: api.bank.local
    http:
      paths:
      - path: /v2
        pathType: Prefix
        backend:
          service: { name: v2-svc, port: { number: 8080 } }
  - host: portal.bank.local
    http:
      paths:
      - path: /login
        pathType: Exact
        backend:
          service: { name: auth-svc, port: { number: 9000 } }
```
For **Gauntlet 29.3**, patch `backend-svc` selector to `env: prod` and `targetPort: 9090` (or `targetPort: http-alt`):
```bash
k patch svc backend-svc -n clinic-net -p '{"spec":{"selector":{"env":"prod"},"ports":[{"port":80,"targetPort":9090}]}}'
k delete ns clinic-net external-clients
```
</details>

---

## Day 30 (Tue Dec 15): High-Speed Clinic #2 — RBAC + SecurityContext + Volumes/PVC Gauntlet

Set a **40-minute timer** and complete all 3 high-density tasks in namespace `clinic-sec`:

```bash
k create ns clinic-sec
```

### Gauntlet 30.1: Least-Privilege RBAC across API Groups
In `clinic-sec`:
1. Create ServiceAccount `release-bot`.
2. Create Role `release-role` granting:
   - `get, list, watch` on `pods` and `services` (core group `""`)
   - `get, list, update, patch` on `deployments` and `statefulsets` (`apps` group)
   - `create, get` on `jobs` (`batch` group)
3. Bind `release-role` to `release-bot` via RoleBinding `release-bind`.
4. Verify with `k auth can-i patch statefulsets.apps -n clinic-sec --as=system:serviceaccount:clinic-sec:release-bot`.

### Gauntlet 30.2: Hardened Multi-Container Pod with PVC & Projected Volume
In `clinic-sec`:
1. Create Secret `api-creds` (`TOKEN=tok_999`) and ConfigMap `env-cfg` (`REGION=us-east`).
2. Create a Pod `secure-pipeline` using `serviceAccountName: release-bot` (`automountServiceAccountToken: false`):
   - Pod `securityContext`: `runAsNonRoot: true`, `runAsUser: 3000`, `runAsGroup: 3000`, `fsGroup: 3000`.
   - A `projected` volume `bundle` combining `api-creds` and `env-cfg` mounted read-only at `/etc/projected`.
   - An `emptyDir` volume `work` mounted at `/work`.
   - `initContainer` `prep` (`busybox:1.36`): copies `/etc/projected/REGION` to `/work/region.txt`.
   - Main container `runner` (`busybox:1.36`, `sleep 3600`): `readOnlyRootFilesystem: true`, `allowPrivilegeEscalation: false`, `capabilities: { drop: ["ALL"] }`, mounting `work` at `/work`.

<details>
<summary><strong>Click for Complete Day 30 Gauntlet Solutions</strong></summary>

```bash
# Gauntlet 30.1 Solution (100% imperative!)
k create sa release-bot -n clinic-sec
k create role release-role -n clinic-sec \
  --verb=get,list,watch,update,patch,create \
  --resource=pods,services,deployments.apps,statefulsets.apps,jobs.batch $do > /tmp/opencode/r.yaml
```
Adjust `/tmp/opencode/r.yaml` so each rule has its exact requested verbs:
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata: { name: release-role, namespace: clinic-sec }
rules:
- apiGroups: [""]
  resources: ["pods", "services"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["apps"]
  resources: ["deployments", "statefulsets"]
  verbs: ["get", "list", "update", "patch"]
- apiGroups: ["batch"]
  resources: ["jobs"]
  verbs: ["create", "get"]
```
```bash
k apply -f /tmp/opencode/r.yaml
k create rolebinding release-bind -n clinic-sec --role=release-role --serviceaccount=clinic-sec:release-bot

# Gauntlet 30.2 Solution
k create secret generic api-creds -n clinic-sec --from-literal=TOKEN=tok_999
k create configmap env-cfg -n clinic-sec --from-literal=REGION=us-east
cat << 'EOF' | k apply -f -
apiVersion: v1
kind: Pod
metadata: { name: secure-pipeline, namespace: clinic-sec }
spec:
  serviceAccountName: release-bot
  automountServiceAccountToken: false
  securityContext:
    runAsNonRoot: true
    runAsUser: 3000
    runAsGroup: 3000
    fsGroup: 3000
  initContainers:
  - name: prep
    image: busybox:1.36
    command: ["sh", "-c", "cp /etc/projected/REGION /work/region.txt"]
    volumeMounts:
    - { name: bundle, mountPath: /etc/projected, readOnly: true }
    - { name: work, mountPath: /work }
  containers:
  - name: runner
    image: busybox:1.36
    command: ["sleep", "3600"]
    securityContext:
      readOnlyRootFilesystem: true
      allowPrivilegeEscalation: false
      capabilities: { drop: ["ALL"] }
    volumeMounts:
    - { name: work, mountPath: /work }
  volumes:
  - name: work
    emptyDir: {}
  - name: bundle
    projected:
      sources:
      - secret: { name: api-creds }
      - configMap: { name: env-cfg }
EOF
k exec -n clinic-sec secure-pipeline -- cat /work/region.txt
k delete ns clinic-sec
```
</details>

---

## Day 31 (Wed Dec 16): Full Mock Exam #2 (Killer.sh Session 2 or Split 1h+1h)

Today is your **final dress rehearsal** under exam conditions:
- **Option A:** Take **Killer.sh Session 2** in a single 2-hour block (or 60 minutes today + 60 minutes tomorrow if weekday time is tight).
- **Option B:** Re-run **Simulated CKAD Exam #1** from Day 27 with a strict **75-minute timer** (aiming for **95%+** with zero errors).

**Focus during Mock #2:**
1. Paste the context/namespace command before touching anything.
2. Never spend more than 6 minutes on Pass 1 before flagging a question.
3. Run the 15-second verification check (`k get`, `k describe`, `k get endpoints`) before moving to the next question.

---

## Day 32 (Thu Dec 17): Mock #2 Post-Mortem & Surgical Fixes

Review every task from Mock #2 and run through this **Rapid Diagnostic Triage Table** so every failure pattern is instant on Sunday:

| Symptom Observed | First Command to Run | Root Cause 95% of the Time |
| :--- | :--- | :--- |
| Pod `Pending` | `k describe pod <name>` | Node taint, `nodeAffinity` mismatch, insufficient CPU/RAM requests, or `PVC` still `Pending`. |
| Pod `CreateContainerConfigError` | `k describe pod <name>` | Missing `ConfigMap`/`Secret` key, or `runAsNonRoot: true` on a root image without `runAsUser`. |
| Pod `CrashLoopBackOff` | `k logs <name> --previous` | Typo in `command`/`args`, app exiting `0` immediately, or failing `livenessProbe`. |
| Pod `0/1 Running` (Not Ready) | `k describe pod <name>` | `readinessProbe` wrong port or path! |
| Service `curl` connection refused / timeout | `k get endpoints <svc>` | Service `selector` doesn't match Pod `labels`, or `targetPort` doesn't match `containerPort`. |
| `wget http://svc` says `bad address` | `k get netpol` | An `Egress` NetworkPolicy is blocking UDP/TCP port `53` to CoreDNS! |

---

## Day 33 (Fri Dec 18): The 15-Command Muscle-Memory Speed Run

Run all 15 imperative generators below in namespace `speed-run` in **under 20 minutes**—this is your final finger-speed warm-up before the weekend:

```bash
k create ns speed-run

# 1. Pod with port, labels, and env
k run p1 -n speed-run --image=nginx:1.25-alpine --port=80 --labels="app=web,env=prod" --env="TIER=front"

# 2. Pod with custom command & args
k run p2 -n speed-run --image=busybox:1.36 --command -- sh -c "echo fast && sleep 3600"

# 3. Deployment with 3 replicas
k create deploy d1 -n speed-run --image=nginx:1.24-alpine --replicas=3

# 4. Set container image & check rollout status
k set image deploy/d1 -n speed-run nginx=nginx:1.25-alpine && k rollout status deploy/d1 -n speed-run

# 5. Set container resource requests & limits imperatively!
k set resources deploy/d1 -n speed-run -c=nginx --requests=cpu=100m,memory=64Mi --limits=cpu=200m,memory=128Mi

# 6. Rollback deployment to revision 1
k rollout undo deploy/d1 -n speed-run --to-revision=1

# 7. Expose Deployment as NodePort Service
k expose deploy d1 -n speed-run --name=d1-svc --type=NodePort --port=8080 --target-port=80

# 8. Create HPA imperatively
k autoscale deploy d1 -n speed-run --min=2 --max=6 --cpu-percent=70

# 9. Create Job & CronJob imperatively
k create job j1 -n speed-run --image=busybox:1.36 -- echo "done"
k create cronjob cj1 -n speed-run --image=busybox:1.36 --schedule="*/5 * * * *" -- date

# 10. Trigger manual Job from CronJob
k create job cj1-manual -n speed-run --from=cronjob/cj1

# 11. ConfigMap & Secret from literals
k create cm cm1 -n speed-run --from-literal=K1=V1 --from-literal=K2=V2
k create secret generic sec1 -n speed-run --from-literal=PASS=s3cr3t

# 12. ServiceAccount + Role + RoleBinding + auth can-i check
k create sa sa1 -n speed-run
k create role r1 -n speed-run --verb=get,list,watch --resource=pods,deployments.apps
k create rolebinding rb1 -n speed-run --role=r1 --serviceaccount=speed-run:sa1
k auth can-i list deployments.apps -n speed-run --as=system:serviceaccount:speed-run:sa1

# 13. ResourceQuota imperatively
k create quota q1 -n speed-run --hard=pods=10,requests.cpu=2,limits.memory=2Gi

# 14. Ingress with TLS and Prefix rule imperatively
k create ingress ing1 -n speed-run --class=nginx --rule="app.local/api*=d1-svc:8080,tls=app-tls"

# 15. Clean up
k delete ns speed-run
```

---

## Day 34 (Sat Dec 19): Final Dress Rehearsal, Logistics & Mental Rest

- **Morning (45 mins max):**
  - Skim [`CKAD_NOTES.md`](CKAD_NOTES.md) Section 6 (`kubectl` Cheat Sheet) and your [Personal Mistake & Speed Tracker](#appendix-personal-mistake--speed-tracker).
  - Do **not** start any new heavy mock exams today—save your mental energy for tomorrow!
- **Afternoon (15 mins — Exam Room & System Check):**
  - Run the **PSI Secure Browser compatibility check**.
  - Clear your desk completely (no papers, second monitors unplugged, phones out of reach, clear glass of water only).
  - Have your **unexpired government-issued photo ID** ready on the desk.
- **Evening:** Eat a good dinner and get a full 8 hours of sleep.

---

## Day 35 (Sun Dec 20): CKAD EXAM DAY Game Plan

### Minute 0–1: Terminal Setup inside the Exam Environment
As soon as your exam terminal opens:
```bash
alias k=kubectl
export do="--dry-run=client -o yaml"
export now="--force --grace-period=0"
```
Check `~/.vimrc` (`cat ~/.vimrc`) and add `set tabstop=2 shiftwidth=2 expandtab` if not already set.

### Minutes 1–120: The Winning Loop on Every Question
1. **Context First:** Copy and paste the `kubectl config use-context ...` box at the top of the question.
2. **Namespace Check:** Read the prompt carefully for the target namespace (`-n <ns>`), exact object name, image tag, and port numbers.
3. **Generate, Don't Hand-Type:** Use `$do` (`k run`, `k create deploy/job/cronjob/cm/secret/role/rolebinding/ingress/quota`, `k expose`, `k set resources/image/env`).
4. **The 6-Minute Cutoff:** If a question is fighting you at the 6-minute mark, **flag it and move on**. Bank all the fast points first!
5. **15-Second Verification:** Run `k get pods,svc,endpoints -n <ns>` before clicking the next question.

You have put in 35 days of real, hands-on engineering practice. Trust your preparation and go crush it!

---

## Appendix: Personal Mistake & Speed Tracker

Use this table throughout Weeks 1–5 whenever a command trips you up or takes longer than 5 minutes:

| Date | Week / Day | Topic / Resource | What Tripped Me Up | The Fix / Command to Remember | Verified 3x? |
| :--- | :--- | :--- | :--- | :--- | :---: |
| *Nov 19* | *W1 / D4* | *Multi-Container Job* | *Standard sidecar kept Job running forever* | *Put sidecar under `initContainers` with `restartPolicy: Always`* | [x] |
| *Nov 30* | *W3 / D15* | *ConfigMap Mount* | *Used `subPath` when prompt asked for live config updates* | *Use `items:` or full dir mount (never `subPath`) if auto-update is needed* | [x] |
| *Dec 9* | *W4 / D24* | *NetworkPolicy* | *Put a second `-` before `podSelector`, turning `AND` into `OR`* | *1 dash = `AND` (same list item); 2 dashes = `OR` (separate list items)* | [x] |
| | | | | | [ ] |
| | | | | | [ ] |
| | | | | | [ ] |
| | | | | | [ ] |

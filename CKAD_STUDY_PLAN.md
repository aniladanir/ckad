# 5-Week CKAD Study Plan (Nov 16 – Dec 20, 2026)

- **Start Date:** Monday, November 16, 2026
- **Target Exam Date:** Sunday, December 20, 2026 (**35 days total**)
- **Daily Commitment:**
  - **Weekdays (Mon–Fri):** **1–1.5 hours** core session (never exceeding your 2-hour cap).
  - **Weekends (Sat–Sun):** **2–3 hours** focused lab/mock blocks (comfortably inside your 3–4 hour window, leaving room for breaks or optional extra drills without burnout).
- **Your Starting Profile:** You studied the theory ~1 year ago, so concepts will come back quickly—but you haven't touched `kubectl` or written YAML manifests in a year.

> **Core Strategy — 20% Reading / 80% Typing:**
> Do **not** spend weeks passively reading docs. Skim theory in [`CKAD_NOTES.md`](CKAD_NOTES.md) for 10–15 minutes, then spend the remaining 45+ minutes in your terminal creating, breaking, debugging, and applying manifests.

---

## Before You Start (Exam-Speed Habits from Day 1)

Set these up in your practice terminal and use them **every single day** so they become automatic by December 20:

```bash
alias k=kubectl
export do="--dry-run=client -o yaml"
export now="--force --grace-period=0"
```

Add clean YAML indentation to `~/.vimrc` (takes 30 seconds and saves minutes of YAML pain):
```vim
set tabstop=2
set shiftwidth=2
set expandtab
```

**Three golden CKAD habits:**
1. **Never write YAML from scratch** when `k run ... $do`, `k create ... $do`, or `k expose ... $do` can generate 80% of it in 2 seconds.
2. **Use `kubectl explain`** (e.g., `k explain pod.spec.containers.livenessProbe`) before opening browser docs.
3. **Always verify `-n <namespace>`**—the #1 silent point-loser on CKAD is creating a resource in `default` instead of the requested namespace.

---

## Recommended Practice Environment

- **Local Cluster:** `kind`, `minikube`, `k3s`, or Docker Desktop Kubernetes (with **Metrics Server** enabled for HPA/top, plus **Helm v3** installed).
- **Repo Materials:**
  - Notes: [`CKAD_NOTES.md`](CKAD_NOTES.md)
  - Concept Check: [`quiz/CKAD_QUIZ_MCQ.html`](quiz/CKAD_QUIZ_MCQ.html)
  - Gentle Warm-Up Lab: [`warmup-exam/CKAD_WARMUP_EXAM.html`](warmup-exam/CKAD_WARMUP_EXAM.html) + [`warmup-exam/setup.yaml`](warmup-exam/setup.yaml)
  - Reference Manifests: [`manifests/`](manifests/) and [`probes/`](probes/)
- **Included with CKAD Registration:** You get **2 free Killer.sh simulator sessions** (harder than the real exam—amazing for Weeks 4–5).

---

## Week 1 (Nov 16 – Nov 22): Muscle Memory Reboot & Core Workloads

**Goal:** Get comfortable back in the terminal, generate YAML fast, and lock down **Application Design & Build (20%)**.

- [ ] **Mon Nov 16 (Day 1) — Terminal Reboot & Baseline** *(60 min)*
  - Start your local cluster, set up `k`, `$do`, and `~/.vimrc`.
  - Apply `warmup-exam/setup.yaml` and complete **Tasks 1–6** of [`warmup-exam/CKAD_WARMUP_EXAM.html`](warmup-exam/CKAD_WARMUP_EXAM.html).
  - *Bonus (+30m):* Take [`quiz/CKAD_QUIZ_MCQ.html`](quiz/CKAD_QUIZ_MCQ.html) once as a baseline (don't worry about the score—use it to spot what faded over the year).

- [ ] **Tue Nov 17 (Day 2) — Finish the Warm-Up Exam** *(60 min)*
  - Complete **Tasks 7–12** of [`warmup-exam/CKAD_WARMUP_EXAM.html`](warmup-exam/CKAD_WARMUP_EXAM.html) (Probes, Jobs/CronJobs, fixing broken Service, SecurityContext, RBAC, NetworkPolicy).
  - Run `k delete namespace ckad-warmup` when done.

- [ ] **Wed Nov 18 (Day 3) — Container Images & Pod Lifecycle** *(60 min)*
  - Brush up on Section 1 of [`CKAD_NOTES.md`](CKAD_NOTES.md) (Dockerfiles, `ENTRYPOINT`/`CMD` vs `command`/`args`, `imagePullPolicy`, `restartPolicy`, `preStop`/`postStart`).
  - **Hands-on:** Inspect [`probes/Dockerfile`](probes/Dockerfile); build a tiny image locally (or practice `docker`/`podman build`, `tag`, `save`); launch Pods overriding `command` vs `args` and verify with `k logs`.

- [ ] **Thu Nov 19 (Day 4) — Multi-Container Pod Patterns** *(60 min)*
  - Review Sidecar, Ambassador, Adapter, InitContainers, and Native Sidecars (`restartPolicy: Always` inside `initContainers`).
  - **Hands-on:** Write a Pod with an `initContainer` that writes `index.html` into an `emptyDir`, served by an `nginx` main container, plus a sidecar tailing logs—test with `k exec` and `k logs -c`.

- [ ] **Fri Nov 20 (Day 5) — Jobs & CronJobs** *(60 min)*
  - Review `completions`, `parallelism`, `backoffLimit`, `activeDeadlineSeconds`, `ttlSecondsAfterFinished`, `concurrencyPolicy` (`Allow`/`Forbid`/`Replace`).
  - **Hands-on:** Create a parallel Job (`completions: 4`, `parallelism: 2`), a failing Job with `backoffLimit: 2`, and a CronJob that you manually trigger with `k create job manual --from=cronjob/<name>`.

- [ ] **Sat Nov 21 (Day 6) — Storage Deep Dive + Week 1 Synthesis** *(2–2.5 hrs)*
  - Review `emptyDir`, `hostPath`, `PV`/`PVC`/`StorageClass`, Access Modes (`RWO` vs `RWOP` vs `RWX`), `reclaimPolicy`, `volumeBindingMode`.
  - **Hands-on:** Inspect and apply [`manifests/pvc/`](manifests/pvc/); create a PVC + Pod, write a file, delete the Pod, recreate it, and confirm data persists.
  - **Synthesis Lab:** Build a 2-container Pod (InitContainer + main container) mounted to a PVC, overriding `command`/`args`.

- [ ] **Sun Nov 22 (Day 7) — Week 1 Timed Drill & MCQ Baseline** *(2–2.5 hrs)*
  - Reset and redo 6 random tasks from `CKAD_WARMUP_EXAM` **without looking at hints or solutions**, aiming for < 5 min per task.
  - Take [`quiz/CKAD_QUIZ_MCQ.html`](quiz/CKAD_QUIZ_MCQ.html) once to identify which theory areas faded most over the year.
  - Review any weak Week 1 topics in `CKAD_NOTES.md`.

---

## Week 2 (Nov 23 – Nov 29): Application Deployment, Scheduling & Helm

**Goal:** Master **Application Deployment (20%)**—rollouts, strategies, scheduling constraints, HPA, and Helm.

- [ ] **Mon Nov 23 (Day 8) — Deployments, Rolling Updates & Rollbacks** *(60 min)*
  - Review `maxSurge`, `maxUnavailable`, revisions, `change-cause`.
  - **Hands-on:** Create a Deployment via `$do`, edit `strategy.rollingUpdate`, perform 3 image updates, inspect `k rollout history --revision`, practice `rollout pause/resume` and `rollout undo --to-revision=1`.

- [ ] **Tue Nov 24 (Day 9) — Blue-Green & Canary Deployments** *(60 min)*
  - Review Deployment Strategies table in `CKAD_NOTES.md`.
  - **Hands-on:** Build a **Blue-Green** setup (`v1` and `v2` Deployments + 1 Service) and switch traffic via `k patch svc` selector; then build a **Canary** setup (`9` replicas `v1` + `1` replica `v2` behind one Service).

- [ ] **Wed Nov 25 (Day 10) — Labels, Affinity, Taints & Tolerations** *(60 min)*
  - Review `nodeSelector`, `nodeAffinity` (`required` vs `preferred`), `podAntiAffinity`, and Taints/Tolerations (`NoSchedule`, `PreferNoSchedule`, `NoExecute`).
  - **Hands-on:** Label a node (`k label node ... disk=ssd`), schedule Pods with `nodeAffinity`, taint a node (`k taint nodes ...`), verify Pods stay `Pending`, then add a matching toleration.

- [ ] **Thu Nov 26 (Day 11) — Horizontal Pod Autoscaling (HPA)** *(60 min)*
  - Review HPA control loop, stabilization windows, and `autoscaling/v2` (`manifests/autoscaler/horizontal.yaml`).
  - **Hands-on:** Ensure Metrics Server is running in your local cluster; deploy an app with `requests.cpu: 100m`, create an HPA with `k autoscale` and via YAML, generate CPU load with a busybox `while true` loop, and watch `k get hpa -w`.

- [ ] **Fri Nov 27 (Day 12) — Helm Package Manager** *(60 min)*
  - Review Helm concepts (`Chart`, `Values`, `Release`) and commands in `CKAD_NOTES.md`.
  - **Hands-on:** `helm repo add bitnami ...`, `helm search repo`, `helm show values`, `helm install` with `--set` and `-f values.yaml`, `helm upgrade`, `helm history`, `helm rollback`, `helm template`, and `helm uninstall`.

- [ ] **Sat Nov 28 (Day 13) — StatefulSets, DaemonSets & ConfigMaps Intro** *(2–2.5 hrs)*
  - Review StatefulSet (ordered naming, Headless Service, `volumeClaimTemplates`) and DaemonSet.
  - **Hands-on:** Deploy a 3-replica StatefulSet with a Headless Service (`clusterIP: None`), test DNS lookup for `<pod-0>.<svc>` from a busybox Pod, and deploy a DaemonSet.
  - *Optional (+45m if energetic):* Preview [`manifests/configmap/`](manifests/configmap/) to get a head start on Week 3.

- [ ] **Sun Nov 29 (Day 14) — Week 2 Timed Mini-Lab & Catch-Up** *(2–2.5 hrs)*
  - Give yourself **60 minutes** to complete an end-to-end chain: *Helm install -> Blue-Green switch -> Rolling update & undo -> Taint/Toleration Pod -> HPA manifest*.
  - Spend the second hour catching up on any weekday topic that felt shaky and tightening your YAML editing speed.

---

## Week 3 (Nov 30 – Dec 6): Configuration, Security & Observability

**Goal:** Conquer the two heaviest/trickiest areas: **Configuration & Security (25%)** + **Observability & Maintenance (15%)** = **40% of the exam**.

- [ ] **Mon Nov 30 (Day 15) — ConfigMaps & Secrets** *(60 min)*
  - Review ConfigMaps & Secrets in `CKAD_NOTES.md` (especially `envFrom` vs `volumeMounts` vs `subPath` caveat).
  - **Hands-on:** Practice [`manifests/configmap/`](manifests/configmap/) and [`manifests/secret/`](manifests/secret/); create `generic`, `tls` (with `openssl req`), and `docker-registry` Secrets; mount specific keys using `items` and verify permissions/updates.

- [ ] **Tue Dec 1 (Day 16) — Resource Management, Quotas, LimitRanges, QoS & PriorityClass** *(60 min)*
  - Review CPU throttling vs `OOMKilled`, `ResourceQuota`, `LimitRange`, QoS classes, `PriorityClass`.
  - **Hands-on:** Apply [`manifests/resource-quota/`](manifests/resource-quota/), [`manifests/limitrange/`](manifests/limitrange/), [`manifests/resource-management/`](manifests/resource-management/), and [`manifests/priority-class/`](manifests/priority-class/). Observe how a Quota rejects Pods without resources until a `LimitRange` is added.

- [ ] **Wed Dec 2 (Day 17) — SecurityContext & Linux Capabilities** *(60 min)*
  - Review `runAsUser`, `runAsGroup`, `fsGroup`, `runAsNonRoot`, `allowPrivilegeEscalation`, `capabilities`, `readOnlyRootFilesystem`.
  - **Hands-on:** Apply and test [`manifests/security-context/security-context.yaml`](manifests/security-context/security-context.yaml); deliberately trigger a `runAsNonRoot` conflict with an `nginx` image, then fix it with `runAsUser` + `emptyDir` for writable paths (`/var/cache/nginx`, `/var/run`).

- [ ] **Thu Dec 3 (Day 18) — ServiceAccounts & RBAC** *(60 min)*
  - Review ServiceAccounts (`automountServiceAccountToken`, projected tokens) and RBAC (`Role`, `ClusterRole`, `RoleBinding`, `ClusterRoleBinding`).
  - **Hands-on:** Apply [`manifests/service-account/`](manifests/service-account/); build Roles/RoleBindings from scratch via `k create role/rolebinding $do`; test permissions with `k auth can-i ... --as=system:serviceaccount:<ns>:<sa>`.

- [ ] **Fri Dec 4 (Day 19) — Probes & Application Observability** *(60 min)*
  - Review `startupProbe`, `readinessProbe`, `livenessProbe`, probe handlers, and timing fields.
  - **Hands-on:** Inspect [`probes/deployment.yaml`](probes/deployment.yaml) and [`probes/main.go`](probes/main.go); create Pods with `httpGet`, `tcpSocket`, and `exec` probes; practice `k top`, `k logs -f`, `k logs --previous`, `k port-forward`, and `k debug -it`.

- [ ] **Sat Dec 5 (Day 20) — Troubleshooting Clinic, CRDs & API Deprecations** *(2–3 hrs)*
  - Review **API Deprecations** and **CRDs & Operators** in `CKAD_NOTES.md`.
  - **Hands-on:** Practice `k api-resources`, `k api-versions`, `k get crds`. Deliberately break 5 Pods (`bad-image`, `bad-configmap`, `oom-limit`, `bad-command`, `bad-probe`) and diagnose each using only `k get events`, `k describe`, `k debug`, and `k logs --previous`.

- [ ] **Sun Dec 6 (Day 21) — Week 3 Checkpoint & Cumulative Mini-Mock** *(2.5–3 hrs)*
  - Retake [`quiz/CKAD_QUIZ_MCQ.html`](quiz/CKAD_QUIZ_MCQ.html) — target **22+/25 (88%+)**.
  - Do a **75-minute timed mini-mock** combining Secret + SecurityContext (`readOnlyRootFilesystem: true` + `emptyDir`) + RBAC + ResourceQuota/LimitRange + Probes, then spend 30–45 min reviewing mistakes.

---

## Week 4 (Dec 7 – Dec 13): Services, Networking & Full Mock #1

**Goal:** Master **Services & Networking (20%)**, build CLI query speed, and complete your first full 2-hour timed mock exam.

- [ ] **Mon Dec 7 (Day 22) — Services & CoreDNS** *(60 min)*
  - Review `ClusterIP`, `NodePort`, `LoadBalancer`, `ExternalName`, Headless Services, and `Endpoints`.
  - **Hands-on:** Inspect [`manifests/service/nodeport.yaml`](manifests/service/nodeport.yaml); expose Deployments with `k expose`, test cross-namespace DNS (`<svc>.<ns>.svc.cluster.local`) from a temporary `busybox` Pod using `wget` / `nslookup`.

- [ ] **Tue Dec 8 (Day 23) — Ingress Routing & TLS** *(60 min)*
  - Review Ingress `pathType`, `ingressClassName`, host/path routing, and TLS termination.
  - **Hands-on:** Inspect [`manifests/ingress/`](manifests/ingress/); write an Ingress YAML from memory (or `k create ingress ... $do`) with 2 paths (`/api`, `/auth`), 2 hosts, and a `tls` block referencing a `kubernetes.io/tls` Secret. *(Install NGINX Ingress Controller on `kind`/`minikube` if you want to test live `curl`!)*

- [ ] **Wed Dec 9 (Day 24) — NetworkPolicies Part 1 (Ingress & Selectors)** *(60 min)*
  - Review default allow vs isolation, `podSelector`, `namespaceSelector`, and the **`AND` vs `OR` rule trap**.
  - **Hands-on:** Write from scratch: (1) default-deny-all-ingress policy, (2) allow ingress from `podSelector` **AND** `namespaceSelector`, (3) allow ingress from `podSelector` **OR** `namespaceSelector`. Double-check every `-`!

- [ ] **Thu Dec 10 (Day 25) — NetworkPolicies Part 2 (Egress, Ports & DNS)** *(60 min)*
  - Review `Egress`, `ipBlock`, and allowing UDP/TCP port `53` to `kube-system` for DNS resolution.
  - **Hands-on:** Write an `Egress` policy allowing traffic only to a target Pod label on port `8080` + DNS port `53`. Inspect [`manifests/generic/network-policy.yaml`](manifests/generic/network-policy.yaml).

- [ ] **Fri Dec 11 (Day 26) — `kubectl` Speed, JSONPath & Docs Navigation** *(60 min)*
  - Practice extracting data fast: `-o custom-columns`, `-o jsonpath`, `--sort-by=.metadata.creationTimestamp`, `k explain --recursive | grep -i <field>`.
  - Practice finding exact YAML snippets on `kubernetes.io/docs` (e.g., NetworkPolicy, Ingress, SecurityContext, PersistentVolumeClaim) in < 20 seconds.

- [ ] **Sat Dec 12 (Day 27) — FULL MOCK EXAM #1 + Initial Review** *(2.5–3 hrs)*
  - Run **Killer.sh Session 1** (or complete the full [Medium CKAD Practice Questions](https://medium.com/bb-tutorials-and-thoughts/practice-enough-with-these-questions-for-the-ckad-exam-2f42d1228552)) in one strict **2-hour** timed block.
  - Rules: No notes, only `kubernetes.io/docs` and `kubectl explain`. Flag/skip any question taking > 6–7 minutes.
  - Spend 30–60 minutes after a short break listing every missed or slow topic.

- [ ] **Sun Dec 13 (Day 28) — Mock Exam #1 Deep Review & Re-Solve** *(2–3 hrs)*
  - Re-solve every question you missed, skipped, or took > 6 minutes on from scratch until each takes < 4 minutes.
  - Write down your top 2 weakest domains to target on Mon–Tue.

---

## Week 5 (Dec 14 – Dec 20): Final Polishing, Mock #2 & Exam Day

**Goal:** Turn weak areas into automatic muscle memory, pass Mock #2 comfortably, and arrive rested on Sunday.

- [ ] **Mon Dec 14 (Day 29) — Weak-Spot Clinic #1** *(60 min)*
  - Focus exclusively on your **lowest-scoring domain** from Mock #1. Repeat 4–5 variations of those tasks cleanly.

- [ ] **Tue Dec 15 (Day 30) — Weak-Spot Clinic #2** *(60 min)*
  - Focus on your **second-weakest domain** plus high-frequency fast points (Labels, Jobs/CronJobs, Rollbacks, ResourceQuotas, NetworkPolicies).

- [ ] **Wed Dec 16 (Day 31) — FULL MOCK EXAM #2** *(2 hrs max; or split: 1h Wed + 1h Thu)*
  - Run **Killer.sh Session 2** (or a fresh timed mock exam). If you only have 1 hour on Wednesday, do half the questions in 60 min and the other half in 60 min on Thursday.
  - Focus on pacing, reading every requirement carefully (namespace, names, ports, labels), and verifying every created resource before moving on.

- [ ] **Thu Dec 17 (Day 32) — Mock Exam #2 Review** *(60–90 min)*
  - Redo any missed tasks from Mock #2. Ensure you have zero hesitation on `kubectl` flags or YAML structure.

- [ ] **Fri Dec 18 (Day 33) — Rapid-Fire Imperative & YAML Drill** *(60 min)*
  - Do 12–15 quick 3-minute drills covering: `run`, `create deploy/job/cronjob/cm/secret/sa/role/rolebinding/ingress`, `expose`, `set image`, `rollout undo`, `scale`, `auth can-i`, `explain`.

- [ ] **Sat Dec 19 (Day 34) — Light Confidence Check & Rest** *(60–90 min max, then stop!)*
  - Do a relaxed 45-min sweep of any remaining shaky tasks (e.g., 1 NetworkPolicy, 1 RBAC, 1 Multi-container Pod) and skim Section 6 (`kubectl` Cheat Sheet) of [`CKAD_NOTES.md`](CKAD_NOTES.md).
  - Check exam room, desk, government ID, webcam/mic, and PSI browser requirements.
  - **Stop studying by early afternoon and get a full night's sleep.**

- [ ] **Sun Dec 20 (Day 35) — CKAD EXAM DAY!**
  - **First 30 seconds of exam:** Confirm aliases (`alias k=kubectl`, `export do="--dry-run=client -o yaml"`) if not preconfigured.
  - **During exam:**
    - Always check the context/namespace command at the top of each question!
    - Don't get trapped on a single hard 4% question—flag it and grab all the easy 6–8% points first.
    - Before moving to the next question, run `k get ...` / `k describe ...` to confirm Pods are `Running`/`Ready` and Services have `Endpoints`.

---

## Quick Adaptation Guide (For Busy vs. High-Energy Days)

- **Only have ~1 hour on a weekday?** Skip the reading—pick **2 hands-on tasks** for that day's topic and do them in the terminal. Move any remaining drills to the weekend block.
- **Have 2 hours on a weekday or 3–4 hours on a weekend?** Complete the core lab first, then build one extra unguided variation from scratch without `CKAD_NOTES.md` (e.g., combine `readinessProbe` + `SecurityContext` + `NetworkPolicy` on a Deployment) before taking a break.

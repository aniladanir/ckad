# CKAD EXAM NOTES
&nbsp;
## Topics
### Multi-Container Patterns
1. **Sidecar:** A sidecar container runs alongside the main application container to provide supporting functionality. Examples include logging, monitoring, and security.</br></br>
*Problem*: You have a web application container that serves your website. You want to collect logs and metrics from this web server without bloating the web server's code.</br>
*Solution*: You add a sidecar container to the same pod that reads the log file from the shared volume and forwards the logs to a centralized logging system.

2. **Ambassador:** An ambassador container acts as a proxy for the main container, handling all outgoing network traffic. This can be useful for tasks like service discovery, routing, and circuit breaking, without having to build that logic into the main application.</br></br>
*Problem:* Your application needs to connect to a sharded database. The application logic for figuring out which shard to connect to for which piece of data is complex.</br>
*Solution:*  A proxy (like envoy or a custom-built one) that listens on localhost:5432. When the main application connects, the ambassador inspects the request, applies the sharding logic, and forwards the request to the correct database shard out on the network.

3. **Adapter:** An adapter container is used to standardize or normalize the output of the main container. For example, it might
reformat logs to a specific format required by a centralized logging service.</br></br>
*Problem:* You have an older, legacy application that you can't modify. It exposes monitoring data in an outdated, non-standard XML format over HTTP. Your company's standard monitoring system, however, only understands the Prometheus metrics format.</br>
*Solution:* A small web service you write that queries the main container's XML endpoint, converts the data into the Prometheus text-based format, and exposes it on a different port. Your monitoring system can now scrape the adapter.

4. **InitContainer:** An init container runs and completes before the main containers in the pod are started. This is useful for performing setup tasks that the main application relies on.</br></br>
*Problem:* A container that runs a script. The script makes an API call to a remote configuration server, fetches the data, and writes it to a config.json file on a shared volume. This container runs and exits successfully.

** Note: Ambassador and adapter patterns are more specific implementations of the sidecar pattern.

## Definitions
**Pod:** A Pod is the smallest deployable unit in Kubernetes that runs one or more containers sharing network and storage, scheduled onto a single node.

**ReplicaSet:** A ReplicaSet is a Kubernetes object that ensures a specified number of pod "replicas" are running at any given time. Its main job is to maintain a stable set of pods.</br>
- No automatic update strategy.
- No rollback capability

**Deployment:** A deployment manages replicasets to ensure a desired number of identical pods are running and kept in the desired state.
- Provides automated, rolling updates.
- Provides easy rollbacks.

**emptyDir:** A volume type that provides ephemeral(temporary) shared storage for containers in the same Pod.
- Created when pod starts and deleted when pod dies.
- Shared between containers
- Lives on the node
- Safe and isolated to the pod

**hostPath:** A volume type that mounts a file or directory from the host node's filesystem into your Pod.
- Accessing files/directories on the host node can expose security risks.
- Need to monitor the disk usage manually.
- Pods using hostPath may behave differently on different nodes depending on the host's filesystem.</br> 
**Note: *Due to these reasons, hostPath usage is discouraged in production.*

**PersistentVolume (PV):** A PersistentVolume (PV) is a piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned using Storage Classes. It is a resource in the cluster just like a node is a cluster resource.

**Storage Class:** A template or recipe for creating storage. It does not represent any actual storage itself. It is used to define a "class" or "tier" of storage.

**PersistentVolumeClaim (PVC):** A request for storage by a user. It is similar to a Pod. Pods consume node resources and PVCs consume PV resources. PVCs allow applications to request persistent storage without knowing how or where that storage is implemented. </br> 
Access Modes:
- *ReadWriteOnce*: can be mounted by  a single node, multiple pods
- *ReadWriteMany*: can be mounted by multiple nodes
- *ReadWriteOncePod*: can be mounted by one pod
- *ReadOnlyMany*: can be mounted by multiple nodes as read-only

**ConfigMap:** A ConfigMap is an API object used to store non-confidential data in key-value pairs. Pods can consume ConfigMaps as environment variables, command-line arguments, or as configuration files in a volume.</br>
Configmaps are stored in cluster's etcd.

**NodePort:** Exposes the Service on each Node's IP at a static port (the NodePort). To make the node port available, Kubernetes sets up a cluster IP address, the same as if you had requested a Service of type: ClusterIP.

**Ingress:**  an API object that manages external access to the services within a cluster, typically for HTTP and HTTPS traffic.</br>
Other ways to expose a service are Nodeport and LoadBalancer. But they act as a one-to-one mapping between a public IP and a service.</br>
Ingress acts as a single entry point that can route traffic to many different services based on the incoming request's hostname or URL path.

```
Real-World Ingress Example: An E-commerce Website

Imagine you have an e-commerce website running as microservices in Kubernetes:
* A frontend-service that serves the main user interface.
* An api-service that handles product data and business logic.
* An auth-service that manages user login and authentication.

You want to expose these services to the public at the domain my-cool-shop.com.

Without Ingress, you might have to expose each service with its own LoadBalancer, leading to multiple public IPs and higher costs:
* 203.0.113.10 -> frontend-service
* 203.0.113.11 -> api-service
* 203.0.113.12 -> auth-service

With Ingress, you can use a single public IP and route traffic intelligently.

Here's how you'd do it with path-based routing:

* http://my-cool-shop.com -> frontend-service (serves the website)
* http://my-cool-shop.com/api/ -> api-service (handles API requests)
* http://my-cool-shop.com/auth/ -> auth-service (handles logins)

You can also handle host-based routing:

* http://my-cool-shop.com -> frontend-service
* http://api.my-cool-shop.com -> api-service
```

**Network Policies:** NetworkPolicies are Kubernetes resources that control how Pods are allowed to communicate with each other and with other network endpoints. 
They use label selectors to define which Pods a policy applies to, and specify allowed ingress (incoming) and egress (outgoing) traffic rules.

**Probes:** Probes are diagnostic checks performed by the kubelet (the agent that runs on each node) to determine the health and readiness of containers running within a Pod. They are crucial for ensuring the reliability and self-healing capabilities of your applications.

Three main types of probes;
- *Startup:*  Determines if a container application has started successfully. If a startup probe is configured, it disables liveness and readiness checks until it succeeds.
- *Readiness:*  Determines if a container is ready to serve traffic. If a readiness probe fails, Kubernetes will remove the Pod's IP address from the endpoints of all Services that match the Pod.
- *Liveness:*  Determines if a container is running and healthy. If a liveness probe fails, Kubernetes assumes the container is in a broken state and will restart the container.

</br>
There are also three main types of probe handlers;

- *HTTP probe:* Sends an HTTP GET request to a specified path on a specified portUseful when app exposes HTTP endpoints. 2xx-3xx status codes are considered success. 4xx-5xx codes and other network errors are considered failure. Useful for apps that expose HTTP endpoints.
- *TCP probe:*  Attempts to open a TCP socket on a specified port. Considered successful if TCP connection is established. Useful for apps that listen on a TCP port such as databases, message queues etc...
- *Exec prob:* Kubelet executes a specified command inside the container. If command returns 0, it is considered success. If returns non-zero code or times out, it is considered a failure.

</br>
Timing fields for probes;

- *initialDelaySeconds:* Number of seconds to wait after the container has started before the very first probe is performed.
- *failureThreshold:* How often (in seconds) to perform the probe.
- *timeoutSeconds:* Number of seconds after which the probe times out.
- *failureThreshold:* Number of consecutive times a probe must fail for Kubernetes to consider the check failed.

**Request:** The minimum amount of resources Kubernetes guarantees to a container. Kubernetes Scheduler uses this value to decide where to place a Pod. If no such Node exists, the Pod remains in a Pending state. If no requests are specified, Kubernetes treats them as 0.

**Limit:** The maximum amount of resources a container is allowed to use. Kubelet (the agent running on each Node) enforces this limit using Linux cgroups.

- *CPU*: If a container tries to use more CPU than its limit, the Linux kernel throttles the container's processes. The application will run slower, but it won't be terminated.
- *Memory:* If a container tries to allocate more memory than its limit, Linux kernel triggers an Out-Of-Memory (OOM) kill for the container's processes. Kubelet observes that the container has stopped and reports its status as <ins>OOMKilled</ins> to the Kubernetes control plane.

**Quality of Service (QoS) Classes:** Kubernetes assigns every Pod a QoS class based on the resource requests and limits of its component Containers. QoS classes are used by Kubernetes to decide which Pods to evict from a Node running low on resources like memory and cpu.</br>
There are three QoS classes, from highest to lowest priority:
   1. Guaranteed
   2. Burstable
   3. BestEffort
   - **Guaranteed:**  Every container in the Pod must have both a memory-cpu limit and a memory-cpu request defined, and they must be the same value.</br>
   *BestFor*: Critical workloads that cannot tolerate downtime or performance degradation, such as databases, message queues or stateful services.
   - **Burstable:** The Pod does not meet the criteria for Guaranteed, but at least one container in the Pod has a CPU or memory request defined. The most common pattern is setting a request lower than a limit.</br>
   *BestFor:* The vast majority of applications such as Web servers, API backends, and stateless microservices.
   - **BestEffort:** Assigned when no requests or limits is set. They are first to be killed.</br>
   *BestFor:* Low-priority tasks that can be interrupted and are not critical like batch jobs, development and test containers.

**Priority Class:** A PriorityClass is an object that defines a priority level for Pods. Its primary purpose is to influence the scheduling and preemption of Pods, ensuring that more critical workloads are given precedence over less important ones.

1. *Scheduling Order:* When the Kubernetes scheduler has multiple Pods waiting to be scheduled, it will prioritize scheduling the Pods with a higher priority value first.
2. *Preemption Order:*  If a high-priority Pod cannot be scheduled because there are not enough resources, the scheduler can evict (terminate and remove) lower-priority Pods from a node to make room for the high-priority Pod. The evicted Pods may be rescheduled on other nodes if resources are available.
3. *Node Pressure Eviction:* When pods are evicted due to node running low on resources, kubelet determines if Pods are using more resources than they requested. Then it groups all candidate Pods by their priority. It will always evict Pods from the lowest priority group first. If multiple Pods exist at the same lowest priority level, the kubelet then uses their QoS class to decide the order of eviction.

**Service Accounts:** Service Account provides an identity for processes that run inside a Pod. Each service account is bound to a Kubernetes namespace.

- When a pod needs to talk to the Kubernetes API, it authenticates itself using the identity of a service account.
- Kubernetes automatically creates a secret that holds an authentication token for the Service Account. This token is then mounted into the pod's filesystem (at /var/run/secrets kubernetes.io/serviceaccount/token).
- Every namespace has a default Service Account. If you don't assign a specific Service Account to a pod when you create it, it automatically uses the default one.
- A Service Account on its own has no permissions. To grant it the ability to do things, you use Kubernetes' Role-Based Access Control (RBAC). You define a Role (for permissions within a namespace) or a ClusterRole (for cluster-wide permissions) and then bind that role to the Service Account with a RoleBinding or ClusterRoleBinding.


**Security Context:** Security Context is a feature that allows you to define privilege and access control settings for a Pod or an individual Container.

1. User Control
    - *runAsUser:* Specifies the User ID (UID) that the container process will run as.
    - *runAsGroup:* Specifies the Group ID (GID) that the container process will run as.
    - *runAsNonRoot:*  A simple boolean (true/false). If set to true, the Kubelet will validate that the container does not
     run as UID 0 (root) before starting it.
2. Privilege Escalation Controls
    - *allowPrivilegeEscalation:* Controls whether a process can gain more privileges than its parent. Setting this to
    false prevents a child process from using mechanisms like setuid to elevate its permissions.
    - *privileged:* A boolean (true/false). Running a container in privileged mode gives it access to all devices on the
    host and disables nearly all security mechanisms. This is extremely dangerous and should be avoided unless
    absolutely necessary.
3. Linux Capabilities</br>
    Instead of giving a container all-or-nothing root access, Linux capabilities allow you to grant specific kernel-level
    privileges. This is a much more granular and secure approach.
    - *capabilities*
        - *drop: ["ALL"]*: A common practice to drop all default capabilities.
        - *add: ["NET_BIND_SERVICE"]*: Then, you can add back only the specific ones you need, like allowing a process to bind to a port below 1024 without running as root.
4. Filesystem Controls
    - *readOnlyRootFilesystem*: A boolean (true/false). If true, the container's root filesystem is mounted as read-only.
    This is an excellent security measure to prevent attackers from modifying application binaries or configuration
    files.

## Kubectl Commands
```
- To quickly run a pod:
kubectl run <pod_name> --image=<image> (--command -- sh -c "sleep 3600")

- To apply a manifest file:
kubectl apply -f pod.yaml


- To view pod specs and state:
kubectl describe pod <pod_name>


- To edit a resource manifest:
kubectl edit <resource> <name>


- To print logs:
kubectl logs <pod_name>|job/<job_name> (--previous)


- To open a interactive shell inside pod:
kubectl exec -it <pod_name> -- sh


- To add label to a resource:
kubectl label <resource> <name> <label>=<value>


- To remove label from a resource:
kubectl label <resource> <name> <label>-


- To filter resources by label:
kubectl get <resource>s -l <label>=<value> 


- To slace replicaset of a deployment:
kubectl scale deployment <name> --replicas=<number>


- To set container image in a resource such as a pod or a deployment:
kubectl set image <resource> <resource_name> <container_name>=<image>


- To view rollout status:
kubectl rollout status deployment <name>


- To rollback a rollout:
kubectl rollout undo deployment <name>


-View rollout history of a deployment:
kubectl rollout history deployment <deployment_name>


- To create a job:
kubectl create job <name> --image=<image> -- <command>


- To create a cronjob:
kubectl create cronjob <name> --image=<image> --schedule="<schedule>" -- <command>


- To create configmap from file:
kubectl create configmap <name> --from-file=<file_key>=<file_name>

- To create configmap from literals:
kubectl create configmap <name> --from-literal=<key1>=<value1> --from-literal=<key2>=<value2> ...

- To quickly create a temporary interactive pod:
kubectl run temp --image=nginx --image-pull-policy=IfNotPresent --rm -it -- sh

- To retrive a custom value from a resource
kubectl get <resource> <resource_name> -o custom-columns=<Column_Name>:<path> 
Example: kubectl get pods web -o custom-columns=IMAGE:.spec.containers[].image

```

## Notes

- **All containers in a pod shares the same network namespace.** </br>
That means one ip address per pod, one loopback interface(localhost) and same port space.
</br></br>
- **Deployment prioritizes availability over exact replica count during updates.** </br>
Temporary over- or under-provisioning is normal. For example, when maxSurge is set to %25 for a rolling update and replica count is 4, rollout process can create an extra pod just to keep availability at a desired level.
</br></br>
- **Jobs treat each Pod execution as an immutable attempt.**<br>
    - Clear success/failure tracking
    - Accurate retry counting
    - Clean seperation of attempts
</br></br>
- **Kubernetes DNS creates a DNS record for each Service name that resolves to its ClusterIP.**</br>
Service load-balances to matching Pod IPs.
</br></br>
- **A service selects Pods using selectors and forwards traffic to their IPs.**
</br></br>

### Storage
- **Pods use PVCs to request persistent storage without knowing the underlying storage implementation.**
</br></br>
- **The first pod that successfully mounts a ReadWriteOnce PVC will cause its underlying PV to be attached to the node where that pod is scheduled.**
</br></br>
- **Deployments that use a single ReadWriteOnce PVC may not scale across multiple nodes.**
</br></br>
- **Use env vars for static startup config; use ConfigMap volumes for dynamic or file-based config. [(see deployment file)](/manifests/configmap/deployment.yaml)**</br>
When building a manifest file, use envFrom if configs are static, and mount file-based ConfigMap if configs are dynamic.
    - ConfigMap volumes can update while the Pod is running
    - Environment variables cannot

- **Why prefer volume mounted configmaps (file-based)**
    1. Large or structured config (yaml, json)
    2. Config may change while Pod is running
    3. App expects filesystem-based config

- **Why “encrypted ConfigMap ≠ Secret”**
    1. *Kubernetes treats it as normal config*. It maybe logged, exposed in debug output, mounted with broader permissions etc...
    2. *RBAC (role base access control) is different*. kubernetes lets you restrict access to secrets separately. If you store secrets in ConfigMaps, anyone with get configmaps can read them.
    3. *Secrets have special handling.* 
        - Can be mounted as tmpfs (in-memory). 
        - Can be excluded from logs and debug dumps. 
        - Can integrate with 'Encryption at rest' (etcd encryption) or external secret managers (Vault, AWS Secrets Manager, etc.)
    4. *Secrets communicate intent.* “This data is sensitive.”

### Networking
- **Pod IPs are cluster-wide routable. Every Pod can reach every other Pod IP directly, without NAT.**

- **Pod IPs are ephemeral (temporary). They are assigned new IPs on restart.**

- **One Pod = one network namespace = one IP**

- **Kubernetes networking is Pod-to-Pod, flat, and IP-based — Services are just virtual IPs on top.**

- **What problem does a Service solve?**
    - Stable Ip
    - Stable DNS name
    - LoadBalancing

- **ClusterIP pod-to-pod traffic route:**
     ``` 
    Pod → web (DNS-CoreDNS)
        ↓
    ClusterIP (internal virtual IP)
        ↓
    kube-proxy (runs on node)
        ↓
    Endpoint (Pod IP)
- **kube-proxy turns a virtual Service IP into real Pod IP traffic by programming node-level networking rules.**

- **NodePort ≠ pod exposure**</br>
    Nodeport exposes the service, not pods.

- **A NodePort Service is still a normal ClusterIP Service internally. NodePort just adds one more way to enter that same Service.**

- **When traffic arrives at a node via a NodePort, kube-proxy on that node may forward the request to any ready Pod selected by the Service, <ins>regardless of whether that Pod is running on the same node or on a different node within the cluster.<ins>**

- **About network policies:**
    - If no NetworkPolicy selects a Pod, all traffic is allowed.
    - If NetworkPolicy has no ingress rules, all incoming traffic is denied to selected Pods.
    - A pod is isolated for egress(outgoing traffic) if there is any NetworkPolicy that both selects the pod and has "Egress" in its policyTypes.
    - NetworkPolicies are namespaced. Once a NetworkPolicy applies to a Pod, traffic from other namespaces is blocked unless the policy explicitly allows those namespaces or Pods.
    - Network policies are additive.

- **Liveness probe restarts a pod on failure, readiness probe cuts traffic to a pod. They complement each other.**

- **Startup protects boot, readiness protects traffic, liveness protects uptime.**

### Resource Management

- **Limits are ignored by the scheduler. Requests are what matters.**

- **Guaranteed pods are killed last because their requested resources are guaranteed and accounted for by the scheduler(request=limit), so evicting them would break scheduling guarantees.**

- **Requests decide placement. Limits decide enforcement. QoS decides eviction.**

- **Memory and CPU units:**</br>-For memory units, Mi stands for Mebibyte. This is a binary unit, where 1Mi = 1024 KiB = 1024 * 1024 bytes. (Megabyte is 1000 * 1000 bytes).There is also Gibibyte(Gi) and Kibibyte(Ki).</br>
-For cpu unit, m stands for millicores.  1000m is equivalent to 1 CPU core. So, 250m means 0.25 (one-quarter) of a CPU core, and 500m means 0.5 (half) of a CPU core.


### Security

- **Assigning user/group IDs**
    - 0 : Is the root user/group id.
    - 1-999 : This range is reserved for system services and background processes (daemons).
    - 1000+ : This is the standard range for regular, human login accounts. When you install a Linux desktop, the first user you create is almost always assigned UID 1000 and GID 1000.

- **SetUID Bit**</br>
The setuid bit (short for "Set User ID") is a special type of permission bit that can be set on executable files in
Unix-like operating systems (like Linux).  
*Normally, when a user executes a program, the resulting process runs with the effective user ID (EUID) of the user who
executed the program.</br> 
*However, when an executable file has the setuid bit enabled, and a user executes that file:
<ins>The resulting process will run with the effective user ID (EUID) of the file's owner, instead of the user who
executed it.</ins>
- **Privilege Escalation Real World Example (Linux Security):**</br>

    ```
    The passwd command in Linux

    First, let's understand why this mechanism exists for legitimate reasons. Think about how you change your password on
    a Linux system.

    1. The Problem: Your user information, including your encrypted password, is stored in the /etc/shadow file. For
        security, this file is highly protected and can only be read and written to by the `root` user.
    2. The Question: So, how can you, a normal user (let's say your user is developer), change your own password if you
        don't have permission to write to /etc/shadow?
    3. The Solution (`setuid`): The passwd command (located at /usr/bin/passwd) is the solution. This program file is
        owned by root, and it has a special setuid permission bit set on it.

    When you, the developer user, run the passwd command:
    * The passwd program starts as a child process of your shell.
    * Because the setuid bit is active on the file, the kernel does something special: it runs this new process not as
        you (developer), but as the owner of the file, which is root.
    * Now this passwd process, running as root, has the necessary privileges to modify the /etc/shadow file.
    * The program is carefully written to only change the password for the original user (developer) and then exit.

    This is a classic case of privilege escalation: your shell process, running as developer, spawned a child process that
    is running as root.

- **If an image runs as root and *runAsNonRoot* is set to 'true', pod creation will fail due to conflict. To prevent this, user id must be specified using *runAsUser* security context field (e.g. 'runAsUser: 1001').**

-  **If readOnlyRootFilesystem is set to true, then no one has write access to the container's root filesystem, not even
the runAsUser and not even the root user (UID 0).**

## DOCS

- Practice questions: https://medium.com/bb-tutorials-and-thoughts/practice-enough-with-these-questions-for-the-ckad-exam-2f42d1228552
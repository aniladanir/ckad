# CKAD EXAM NOTES
&nbsp;
## Definitions
**Pod:** A Pod is the smallest deployable unit in Kubernetes that runs one or more containers sharing network and storage, scheduled onto a single node.

**Deployment:** A deployment manages replicasets to ensure a desired number of identical pods are running and kept in the desired state.

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
   - **Guaranteed:**  Every container in the Pod must have both a memory-cpu limit and a memory-cpu request defined, and they must be the same value.
   *BestFor*: Critical workloads that cannot tolerate downtime or performance degradation, such as databases, message queues or stateful services.
   - **Burstable:** The Pod does not meet the criteria for Guaranteed, but at least one container in the Pod has a CPU or memory request defined. The most common pattern is setting a request lower than a limit.</br>
   *BestFor:* The vast majority of applications such as Web servers, API backends, and stateless microservices.
   - **BestEffort:** Assigned when no requests or limits is set. They are first to be killed.</br>
   *BestFor:* Low-priority tasks that can be interrupted and are not critical like batch jobs, development and test containers.

**Priority Class:** Ayni qos class'ta bulunan iki farkli poddan once hangisinin evict edilecegine karar verilirken bakilir. lowest priority first. ONEMLI. eger priority classlari ayniysa ya da belirtilmemisse, memory requestini en cok asan pod terminate edilir.

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
kubectl logs <pod_name>|job/<job_name>


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


- To create a job:
kubectl create job <name> --image=<image> -- <command>


- To create a cronjob:
kubectl create cronjob <name> --image=<image> --schedule="<schedule>" -- <command>


- To create configmap from file:
kubectl create configmap <name> --from-file=<file_key>=<file_name>

- To create configmap from literals:
kubectl create configmap <name> --from-literal=<key1>=<value1> --from-literal=<key2>=<value2> ...

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

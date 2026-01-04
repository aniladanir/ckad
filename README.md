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

- **Liveness probe restarts a pod on failure, readiness probe cuts traffic to a pod. They complement each other.**

- **Startup protects boot, readiness protects traffic, liveness protects uptime.**
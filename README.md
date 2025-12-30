# CKAD EXAM NOTES
&nbsp;
## Definitions
**Pod:** A Pod is the smallest deployable unit in Kubernetes that runs one or more containers sharing network and storage, scheduled onto a single node.

**Deployment:** A deployment manages replicasets to ensure a desired number of identical pods are running and kept in the desired state.

**emptyDir:** A volume type that provides ephemeral(temporary) shared storage for containers in the same Pod.
- Created when pod starts and deleted when pod dies.
- Shared between containers
- Lives on the node

**hostPath:** A volume type that mounts a file or directory from the host node's filesystem into your Pod.
- Access to the host filesystem can expose security risks.
- Need to monitor the disk usage manually.
- Identical pods may behave differently due to different files on the nodes. (node-specific behaviour)</br> 
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

#### Storage
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

#### Networking
- **Pod IPs are cluster-wide routable. Every Pod can reach every other Pod IP directly, without NAT.**

- **Pod IPs are ephemeral (temporary). They are assigned new IPs on restart.**

- **One Pod = one network namespace = one IP**

- **Kubernetes networking is Pod-to-Pod, flat, and IP-based — Services are just virtual IPs on top.**
                
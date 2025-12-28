# CKAD EXAM NOTES
&nbsp;
## Definitions
**Pod:** A Pod is the smallest deployable unit in Kubernetes that runs one or more containers sharing network and storage, scheduled onto a single node.

**emptyDir:** A volume type that provides ephemeral(temporary) shared storage for containers in the same Pod.
- Created when pod starts and deleted when pod dies.
- Shared between containers
- Lives on the node

**Deployment:** A deployment manages replicasets to ensure a desired number of identical pods are running and kept in the desired state.

## Kubectl Commands
```
kubectl apply -f pod.yaml (apply a manifest file)

kubectl describe pod <pod_name> (view pod specs and state)

kubectl edit pod <pod_name> (edit pod manifest)

kubectl logs <pod_name> (print logs of a pod)

kubectl exec -it <pod_name> -- sh (open a interactive shell)

kubectl label pod <pod_name> <label>=<value> (add label to a pod)

kubectl label pod <pod_name> <label>- (remove label from a pod)

kubectl get pods -l <label>=<value> (filter pods by label)

kubectl scale deployment <name> --replicas=<number> (slaces replicaset of a deployment)

kubectl set image <resource> <resource_name> <container_name>=<image> (sets container image in a resource such as a pod or a deployment )

kubectl rollout status deployment <name> (watch rollout)

kubectl rollout undo deployment <name> (rollback a rollout)

```

## Notes

- **All containers in a pod shares the same network namespace** </br>
That means one ip address per pod, one loopback interface(localhost) and same port space.

- **Deployment prioritizes availability over exact replica count during updates.** </br>
Temporary over- or under-provisioning is normal. For example, when maxSurge is set to %25 for a rolling update and replica count is 4, rollout process can create an extra pod just to keep availability at a desired level.


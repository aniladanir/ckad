alias k=kubectl
source <(kubectl completion bash)
complete -o default -F __start_kubectl k

export DO="--dry-run=client -o yaml"
export NOW="--force --grace-period=0"

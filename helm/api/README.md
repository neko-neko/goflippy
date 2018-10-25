# goflippy API server Helm chart
## How to render k8s resources
Run this:
```sh
# dry-run
$ helm template . --name production | kubectl apply -f - --dry-run

# apply
$ helm template . --name production | kubectl apply -f -
```

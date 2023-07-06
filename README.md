# Kube-ingress-dashboard

*Like a tiny Kubernetes Dashboard, but just for ingress resources.*

**Intended for use behind an authentication proxy or in a private network - do not expose this dashboard with a fully public ingress.**

Built with **[Go](https://go.dev/)**, **[Pico v2](https://v2.picocss.com/docs/v2)** and **[Alpine.js](https://alpinejs.dev/)**. Container image is approx ~<50 MB.

## Installation

```bash
# 0. Set desired namespace
export NAMESPACE="yournamespace"

# 1. create deployment and supporting resources
kubectl apply -f https://raw.githubusercontent.com/Antvirf/kube-ingress-dashboard/main/manifests/deploy.yaml -n $NAMESPACE

# 2. create clusterrolebinding
kubectl create clusterrolebinding kube-ingress-dashboard --clusterrole=kube-ingress-dashboard --serviceaccount=$NAMESPACE:kube-ingress-dasbhoard
```

## Functionality

The dashboard fetches Ingresses for all namespaces the calling pod has access to, and caches the results for 15 seconds. Any requests within that period are served the same information, after which a new request is made to the API server. Refreshing the cache is done on request only, so if the dashboard is not accessed, no requests are made to the API server.

![dark](/docs/darkmode.png)

### Search by name or URL

![search](/docs/filter_namespace_and_search.png)

### Filter by namespace

![namespace](/docs/filter_namespace.png)

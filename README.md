# Kube-ingress-dashboard

*Like a tiny Kubernetes Dashboard, but just for ingress resources.*

**Intended for use behind an authentication proxy or in a private network - do not expose this dashboard with a fully public ingress.**

Built with **[Go](https://go.dev/)**, **[Pico v2](https://v2.picocss.com/docs/v2)** and **[Alpine.js](https://alpinejs.dev/)**.

## Installation

```bash
export NAMESPACE="yournamespace"
kubectl apply -f https://raw.githubusercontent.com/antvirf/kube-ingress-dashboard/main/deploy.yaml -n $NAMESPACE
```

## Functionality

The dashboard fetches Ingresses for all namespaces the calling pod has access to, and caches the results for 15 seconds. Any requests within that period are served the same information, after which a new request is made to the API server. Refreshing the cache is done on request only, so if the dashboard is not accessed, no requests are made to the API server.

### Search by name or URL

### Filter by namespace

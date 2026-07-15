# application-mesh Helm chart

This chart deploys Istio service mesh configuration for your application: DestinationRules, VirtualServices, Gateways, Egress rules and EnvoyFilters.

## Usage

Add this chart as dependency to your application chart:
```yaml
dependencies:
  - name: application-mesh
    version: 0.0.0
    repository: "https://matic-insurance.github.io/helm-charts"
    alias: "mesh"
```

### Destinations

Define mesh destinations (one per component) with traffic policies:

```yaml
mesh:
  destinations:
    - component: webserver
    - component: sidekiq
```

### Ingress

Route external/internal traffic to destinations:

```yaml
mesh:
  ingress:
    - host: api.matic.com
      gateway: external
      destination: webserver
      deny_locations:
        - /admin

    - host: api-internal.matic.link
      gateway: internal
      destination: webserver
```

### Egress

Allow outbound traffic to external services:

```yaml
mesh:
  egress:
    - name: db
      type: postgres
      host: myapp.aws.db
    - name: api
      type: https
      host: external-api.com
```

## Sticky Sessions

Sticky sessions route requests from the same client to the same pod using Istio's [consistent hash load balancing](https://istio.io/latest/docs/reference/config/networking/destination-rule/#LoadBalancerSettings-ConsistentHashLB). When enabled, the default `LEAST_REQUEST` simple load balancer is replaced with a `consistentHash` configuration.

### Configuration

Add `stickySession` to a destination entry:

```yaml
mesh:
  destinations:
    - component: webserver
      stickySession:
        enabled: true
        type: sourceIp | httpCookie | httpHeaderName
```

| Parameter | Description | Default |
|---|---|---|
| `stickySession.enabled` | Enable sticky sessions for this destination | `false` |
| `stickySession.type` | Hash method: `sourceIp`, `httpCookie`, or `httpHeaderName` | `sourceIp` |
| `stickySession.httpCookie.name` | Cookie name (when type is `httpCookie`) | `istio-session` |
| `stickySession.httpCookie.ttl` | Cookie lifetime as duration (when type is `httpCookie`) | `3600s` |
| `stickySession.httpHeaderName` | Header name to hash on (when type is `httpHeaderName`) | `x-session-id` |

### Examples

**Source IP** — simplest option, routes by client IP address:

```yaml
mesh:
  destinations:
    - component: webserver
      stickySession:
        enabled: true
        type: sourceIp
```

Produces:
```yaml
trafficPolicy:
  loadBalancer:
    consistentHash:
      useSourceIp: true
```

**HTTP Cookie** — Istio generates and manages a session cookie automatically:

```yaml
mesh:
  destinations:
    - component: webserver
      stickySession:
        enabled: true
        type: httpCookie
        httpCookie:
          name: my-app-session
          ttl: 7200s
```

Produces:
```yaml
trafficPolicy:
  loadBalancer:
    consistentHash:
      httpCookie:
        name: "my-app-session"
        ttl: 7200s
```

**HTTP Header** — route based on a request header (e.g. for streamable HTTP sessions):

```yaml
mesh:
  destinations:
    - component: webserver
      stickySession:
        enabled: true
        type: httpHeaderName
        httpHeaderName: x-session-id
```

Produces:
```yaml
trafficPolicy:
  loadBalancer:
    consistentHash:
      httpHeaderName: "x-session-id"
```

> **Note:** When using `httpHeaderName`, requests without the header will not be sticky. This is useful when only specific routes (e.g. streamable HTTP endpoints) include the header, while regular traffic is distributed normally.

### Choosing a method

| Method | Use when | Limitations |
|---|---|---|
| `sourceIp` | Simple affinity needed, clients have stable IPs | Breaks behind shared proxies/NAT |
| `httpCookie` | Browser-based clients, need reliable affinity | Requires cookie support in client |
| `httpHeaderName` | Only specific routes need affinity (e.g. streaming) | Header must be set by the client |

### Why Istio DestinationRule (not Kubernetes Service sessionAffinity)

Sticky sessions are implemented at the Istio `DestinationRule` level using `consistentHash` rather than Kubernetes Service `sessionAffinity: ClientIP`. The reasons:

1. **More hash methods.** Kubernetes Service only supports `ClientIP`. Istio adds `httpCookie` (managed automatically by the sidecar) and `httpHeaderName`, which are essential for scenarios like streamable HTTP where only specific routes need affinity.

2. **Mesh-aware routing.** All traffic in our stack flows through the Istio sidecar proxy. Kubernetes Service `sessionAffinity` operates at kube-proxy/iptables level, which is bypassed when the mesh intercepts traffic — making it effectively a no-op for mesh-enabled services.

3. **Per-destination control.** The `DestinationRule` is already the resource that governs traffic policy per component (load balancer algorithm, connection pool, outlier detection). Sticky sessions are a load balancing concern and belong in the same place.

4. **Consistent architecture.** The `application-mesh` chart owns all traffic management configuration. Adding session affinity to the `application-component` Service template would split traffic policy across two charts and create a confusing precedence situation when mesh is enabled.

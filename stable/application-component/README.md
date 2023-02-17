# application-component Helm chart
This chart is designed to deploy single component of the application like webserver or sidekiq 
and configure all additional resources: service account, service, ingress, autoscaling

When deploying full application - this chart should be used several times for every applicaiton, 
as well we should use other charts to handle migration and additional configs and secrets

## Usage

### Example for deploying webserver

```yaml
product: olb
component: webserver
environment: production
version: release-1.2.3

image:
  repository: matic/olb

ingress:
  enabled: true
  list:
    - host: olb.matic.link
      type: internal
      tls: matic.link-tls

    - host: olb.matic.com
      type: external
      deny_locations:
        - /admin
      tls: matic.com-tls

deployment:
  command: "bundle exec puma -t 8:8"
  replicas: 2
  port: 9292
  resources:
    requests:
      memory: 200Mi
      cpu: 100m
    limits:
      memory: 700Mi
      cpu: 300m
  startupProbe:
    enabled: true
  readinessProbe:
    enabled: true
  livenessProbe:
    enabled: true

service:
  enabled: true

serviceAccount:
  enabled: true

autoscaling:
  enabled: true
  minReplicas: 2
  maxReplicas: 4
  metrics:
    - name: olb-requests-per-second
      type: Datadog
      query: ewma_10(avg:nginx_ingress.controller.requests{service:olb-webserver}.as_count().rollup(avg, 60))
      target:
        type: AverageValue
        averageValue: 12

datadog:
  enabled: true
```
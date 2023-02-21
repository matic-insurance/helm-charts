# application-component Helm chart
This chart is designed to deploy single component of the application like webserver or sidekiq 
and configure all additional resources: service account, service, ingress, autoscaling

When deploying full application - this chart should be used several times for every component, 
as well we should use other charts to handle migration and additional configs and secrets

## Usage

Add this chart as dependency to your application chart:
```yaml
dependencies:
  - name: application-component
    version: 0.0.0
    repository: "https://matic-insurance.github.io/helm-charts"
    alias: "webserver"
  - name: application-component
    version: 0.0.0
    repository: "https://matic-insurance.github.io/helm-charts"
    alias: "sidekiq"
  - name: application-component
    version: 0.0.0
    repository: "https://matic-insurance.github.io/helm-charts"
    alias: "kafka"
```

Chart using common global values that control application deployment:
```yaml
global:
  application:
    product: olb
    environment: "staging"
    version: "1.2.3"

  applicationImage:
    repository: "maticinsurance/olb"
```

Additional deployment components:
- [application-config](../application-config) - allows to define environment variables and config files for all components
- [application-migration](../application-migration) - allows to define job that runs before deployment (aka migration)

Define components that should be run:

```yaml
webserver:
  command: "puma -p 9292"
  deployment:
    replicas: 2
    port: 9292

  # Enable service for the app
  service:
    enabled: true

  # Define ingress configuration for webserver
  ingress:
    enabled: true
    list:
      - host: olb.matic.com
        type: external
        tls: matic.com-tls
    
  # Create service account in one of the components
  serviceAccount:
    create: true

sidekiq:
  command: "sidekiq -q default"

kafka:
  command: "karafka server"
```

## Application monitoring

Chart support configuration of additional monitoring information for Datadog and Sentry:
```yaml
global:
  applicationMonitoring:
    datadog: true
    sentry: true
```

Additional documentation can be found in Confluence:
- [Datadog](https://maticinsurance.atlassian.net/wiki/spaces/DOPS/pages/3652649000/Datadog+Standard+Configuration)
- [Sentry](https://maticinsurance.atlassian.net/wiki/spaces/DOPS/pages/3652649045/Sentry+Configuration)

## Full application example

Below is example configuration for regular app at matic 
```yaml
global:
  application:
    product: olb
    component: webserver
    environment: production
    version: release-1.2.3

  applicationImage:
    repository: matic/olb
    
  applicationMonitoring:
    datadog: true
    sentry: true

migrations:
  command: "bundle exec rails db:migrate"

webserver:
  component: "webserver"
  command: "bundle exec puma -t 8:8"

  deployment:
    replicas: 2
    port: 9292
    resources:
      requests:
        memory: 500Mi
        cpu: 200m
      limits:
        memory: 1500Mi
        cpu: 600m
    startupProbe:
      enabled: true
    readinessProbe:
      enabled: true
    livenessProbe:
      enabled: true

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


  service:
    enabled: true

  serviceAccount:
    create: true

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

sidekiq:
  component: "sidekiq"
  command: "bundle exec sidekiq --concurrency 10"
  deployment:
    resources:
      requests:
        memory: 200Mi
        cpu: 100m
      limits:
        memory: 700Mi
        cpu: 300m
    startupProbe:
      enabled: true
      command: /app/ops/scripts/check_sidekiq.sh
    readinessProbe:
      enabled: true
      command: /app/ops/scripts/check_sidekiq.sh
  
  autoscaling:
    enabled: true
    minReplicas: 1
    maxReplicas: 3
    metrics:
      - name: olb-sidekiq-utilization
        type: Datadog
        query: ewma_10(avg:sidekiq.process.utilization{product:olb}.rollup(avg, 60))
        target:
          type: Value
          value: 75

karakfka:
  component: "karafka"
  command: "bundle exec karafka server"
  deployment:
    replicas: 1
    resources:
      requests:
        memory: 200Mi
        cpu: 100m
      limits:
        memory: 700Mi
        cpu: 300m
```
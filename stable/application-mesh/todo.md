# Implement custom gateways after cert-manager can generate certificates

Add support for `gateway: external` and `gateway: internal` in the ingress list

# Support custom traffic configurations for `DestinationRule`

- `.Values.destinations[].connectionPool` 
- `.Values.destinations[].loadBalancer`

Use defaults when not specified

# Add common labels to all components

```yaml
labels:
  {{- include "application-helpers.labels" . | nindent 4 }}
```

# Review and think if we need outlier configuration in `DestinationRule`
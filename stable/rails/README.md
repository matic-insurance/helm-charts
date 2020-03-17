# rails Helm chart

The Helm chart perform Ruby on Rails application deployment to the Kubernetes cluster. With default parameters simple Puma web-server will be installed and available via ClusterIP within Kubernetes CNI network.

## Requirements

* In case when custom configuration is required to pass, it must be created in advance as Secret. The Secret name should meet following pattern: `<product>-<configuration-file-path>`.

For example we have following structure with configuration files:
```
configs/
└── staging
    └── app
        ├── config.yml
        └── settings.yml
```

In order to convert them to Secrets, use the following chart:

```yaml
{{- if .Values.rails.enabled -}}

{{- range $config := $.Values.rails.custom_configs.files }}
---
apiVersion: v1
kind: Secret
metadata:
  name: {{ $.Values.rails.product}}{{ $config | replace "/" "-" | replace "." "-" }}
type: Opaque
data:
  {{- $path := printf "configs/%s%s" $.Values.rails.environment $config}}
  {{ ($.Files.Glob $path).AsSecrets | indent 2 }}
{{- end }}

{{- end -}}
```
As a result we'll get this rendered part:
```yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: myproduct-app-settings-yml
type: Opaque
data:
    settings.yml: ....
---
apiVersion: v1
kind: Secret
metadata:
  name: myproduct-app-config-yml
type: Opaque
data:
    config.yml: ....
```

The Deployment will be looking for the Secret with corresponding naming pattern and mount it to the Pods.

NOTE: `custom_configs.files` values should match configuration files path. Otherwise Chart will not be able to populate Secrets with proper data.

## Usage
In order to install chart: `helm install -n rails ./rails --namespace mynamespace`

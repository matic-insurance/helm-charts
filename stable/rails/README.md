# Rails Helm chart

The Helm chart perform Ruby on Rails application deployment into the Kubernetes cluster. With default parameters application will be served by Puma web-server and available towards ClusterIP service within Kubernetes CNI network.

## Custom configuration

With this chart, you may pass customized configuration files required by an application.

All custom configuration files should be created in advance and available as k8s Secrets.

The chart will be looking for a Secret name with the following pattern:

`<product>-<configuration-file-path>`

If you'd like to use database migrations option, you must create a separate Secret with this naming format:

`<product>-migrations-<configuration-file-path>`

Let's take an example next configuration structure:

```bash
configs/
└── staging
    └── app
        ├── config.yml
        └── settings.yml
```

In order to make a Secret with needed pattern, use following template

```yaml
{{- if .Values.rails.enabled -}}

{{- range $config := $.Values.rails.custom_configs.files }}
---
apiVersion: v1
kind: Secret
metadata:
  name: {{ $.Values.rails.product}}{{ $config | replace "/" "-" | replace "." "-" | replace "_" "-"}}
  annotations:
    "helm.sh/hook": pre-install,pre-upgrade
    "helm.sh/hook-delete-policy": before-hook-creation
type: Opaque
data:
  {{- $path := printf "configs/%s%s" $.Values.rails.environment $config}}
  {{ ($.Files.Glob $path).AsSecrets | indent 2 }}
{{- end }}

{{- end -}}
```

The rendered part will be looking like this:

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

**NOTE**: you should keep consistency with your configuration file path and `Values.custom_configs.files` parameters. Otherwise, the chart will not be able to update Secrets with proper data.

## Usage

In order to install chart: `helm install -n rails ./rails --namespace mynamespace`

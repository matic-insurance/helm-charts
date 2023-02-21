{{/*
Expand the name of the chart.
*/}}
{{- define "application-helpers.name" -}}
{{- required ".Values.global.application.product required for deployment" .Values.global.application.product | trunc 63 | trimSuffix "-" -}}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "application-helpers.fullname" -}}
{{- printf "%s-%s" (include "application-helpers.name" .) (required ".Values.component required for component deployment" .Values.component) | trunc 63 | trimSuffix "-" -}}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "application-helpers.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
All Common Labels
*/}}
{{- define "application-helpers.labels" -}}
{{ include "application-helpers.selectorLabels" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/component: {{ .Values.component }}
app.kubernetes.io/part-of: {{ .Values.global.application.product }}
app.kubernetes.io/version: {{ (required ".Values.global.application.version is required for applicaiton deployment" .Values.global.application.version) | quote }}
helm.sh/chart: {{ include "application-helpers.chart" . -}}
{{/* Just checing for environment configuraiton */}}
{{- if (required ".Values.global.application.environment" .Values.global.application.environment) -}}{{- end -}}
{{- end }}

{{/*
Datadog unfied tags: https://docs.datadoghq.com/getting_started/tagging/unified_service_tagging/?tab=kubernetes
*/}}
{{- define "application-helpers.monitoring.datadog.labels" -}}
tags.datadoghq.com/env: {{ .Values.global.application.environment }}
tags.datadoghq.com/service: {{ .Values.global.application.product }}
tags.datadoghq.com/version: {{ .Values.global.application.version | quote }}
{{- end }}

{{/*
Datadog environment variables to be used by datadog libraries
*/}}
{{- define "application-helpers.monitoring.datadog.env" -}}
- name: DD_AGENT_HOST
  valueFrom:
    fieldRef:
      fieldPath: status.hostIP
- name: DD_TRACE_AGENT_PORT
  value: "8126"
- name: DD_DOGSTATSD_PORT
  value: "8125"
- name: DD_ENV
  valueFrom:
    fieldRef:
      fieldPath: metadata.labels['tags.datadoghq.com/env']
- name: DD_SERVICE
  valueFrom:
    fieldRef:
      fieldPath: metadata.labels['tags.datadoghq.com/service']
- name: DD_VERSION
  valueFrom:
    fieldRef:
      fieldPath: metadata.labels['tags.datadoghq.com/version']
- name: DD_TAGS
  value: "product:{{ .Values.global.application.product}},component:{{ .Values.component }}"
{{- end }}

{{/*
Sentry environment variables to be used by libraries
*/}}
{{- define "application-helpers.monitoring.sentry.env" -}}
- name: SENTRY_RELEASE
  value: "{{ include "application-helpers.name" . }}-{{ .Values.global.application.version }}"
- name: SENTRY_ENVIRONMENT
  value: {{ .Values.global.application.environment }}
{{- end }}

{{/*
Common Selector labels
*/}}
{{- define "application-helpers.selectorLabels" -}}
app.kubernetes.io/name: {{ include "application-helpers.name" . }}
app.kubernetes.io/instance: {{ include "application-helpers.fullname" . }}
{{- end }}

{{/*
Return name of the secret (volume) for specific configuration (env vars, migration configs, runtime configs)
Usage: `{{ include "application-helpers.config-files.volume-name" (merge (dict "deployment" "environment") $) }}`
Returns: String
*/}}
{{- define "application-helpers.config-files.volume-name" -}}
{{ include "application-helpers.name" . }}-configs-{{ .deployment }}
{{- end }}

{{/*
Returns name of environment variables secret/volume
*/}}
{{- define "application-helpers.configs.environment-variables.volume-name" -}}
{{ include "application-helpers.name" . }}-configs-environment
{{- end }}

{{/*
Returns name of runtime config files secret/volume
*/}}
{{- define "application-helpers.configs.files-runtime.volume-name" -}}
{{ include "application-helpers.name" . }}-configs-runtime
{{- end }}

{{/*
Returns name of migration config files secret/volume
*/}}
{{- define "application-helpers.configs.files-migration.volume-name" -}}
{{ include "application-helpers.name" . }}-configs-migration
{{- end }}

{{/*
Define the name of the service account to use
*/}}
{{- define "application-helpers.serviceAccountName" -}}
{{- default (include "application-helpers.name" .) .Values.serviceAccount.name }}
{{- end }}

{{/*
Construct full docker image
*/}}
{{- define "application-helpers.docker-image" }}
{{- $repository := required ".Values.global.applicationImage.repository is required for application deployment" .Values.global.applicationImage.repository -}}
{{- $tag := default .Values.global.application.version .Values.global.applicationImage.tag -}}
{{ printf "%s:%s" $repository $tag }}
{{- end }}



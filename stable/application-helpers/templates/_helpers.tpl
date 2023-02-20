{{/*
Expand the name of the chart.
*/}}
{{- define "application-helpers.name" -}}
{{- .Values.global.application.product | trunc 63 | trimSuffix "-" -}}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "application-helpers.fullname" -}}
{{- printf "%s-%s" .Values.global.application.product .Values.application.component | trunc 63 | trimSuffix "-" -}}
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
app.kubernetes.io/component: {{ .Values.application.component }}
app.kubernetes.io/part-of: {{ .Values.global.application.product }}
app.kubernetes.io/version: {{ .Values.global.application.version | quote }}
helm.sh/chart: {{ include "application-helpers.chart" . }}
{{- end }}

{{/*
Datadog unfied tags: https://docs.datadoghq.com/getting_started/tagging/unified_service_tagging/?tab=kubernetes
*/}}
{{- define "application-helpers.datadog-labels" -}}
tags.datadoghq.com/env: {{ .Values.global.application.environment }}
tags.datadoghq.com/service: {{ .Values.global.application.product }}
tags.datadoghq.com/version: {{ .Values.global.application.version | quote }}
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



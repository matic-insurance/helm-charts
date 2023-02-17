{{/*
Expand the name of the chart.
*/}}
{{- define "application-component.name" -}}
{{- .Values.product | trunc 63 | trimSuffix "-" -}}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "application-component.fullname" -}}
{{- printf "%s-%s"  .Values.product .Values.component | trunc 63 | trimSuffix "-" -}}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "application-component.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "application-component.labels" -}}
{{ include "application-component.selectorLabels" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/component: {{ .Values.component }}
app.kubernetes.io/part-of: {{ .Values.product }}
app.kubernetes.io/version: {{ .Values.version | quote }}
helm.sh/chart: {{ include "application-component.chart" . }}
{{- end }}

{{/*
Datadog unfied tags: https://docs.datadoghq.com/getting_started/tagging/unified_service_tagging/?tab=kubernetes
*/}}
{{- define "application-component.datadog-labels" -}}
tags.datadoghq.com/env: {{ .Values.environment }}
tags.datadoghq.com/service: {{ .Values.product }}
tags.datadoghq.com/version: {{ .Values.version | quote }}
{{- end }}
{{/*
Selector labels
*/}}
{{- define "application-component.selectorLabels" -}}
app.kubernetes.io/name: {{ include "application-component.name" . }}
app.kubernetes.io/instance: {{ include "application-component.fullname" . }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "application-component.serviceAccountName" -}}
{{- default (include "application-component.name" .) .Values.serviceAccount.name }}
{{- end }}

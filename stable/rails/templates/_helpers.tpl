{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "rails.name" -}}
{{- printf "%s-%s"  .Values.product .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "rails.environment" -}}
{{- .Values.environment | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "rails.product" -}}
{{- .Values.product | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "rails.version" -}}
{{- .Values.image.repository | regexFind ":.*$" | trimPrefix ":" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "rails.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "rails.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

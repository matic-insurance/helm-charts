{{/*
Constructs FQDN domain name for particular component passed as .component attribute
*/}}
{{- define "application-mesh.service-host" -}}
{{ include "application-mesh.component-fullname" . }}.{{ .Release.Namespace }}.svc.cluster.local
{{- end }}

{{/*
Generate component full name for particular component passed as .component attribute
*/}}
{{- define "application-mesh.component-fullname" -}}
{{ include "application-helpers.fullname" (merge (dict "Values" (dict "component" .component)) $) }}
{{- end }}

{{/*
Generate version label for currently deployed pods
*/}}
{{- define "application-mesh.current-version-label" -}}
app.kubernetes.io/version: {{ (required ".Values.global.application.version is required for applicaiton deployment" .Values.global.application.version) | quote }}
{{- end }}

{{/*
Add header at virtual service for datadog proxy tracing
*/}}
{{- define "application-mesh.proxy-tracing-headers" -}}
{{- if .Values.global.applicationMonitoring.datadog -}}
headers:
  request:
    add:
      "x-request-start": "t=%START_TIME(%s.%3f)%"
{{- end }}
{{- end }}

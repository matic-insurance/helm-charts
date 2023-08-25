{{/*
Constructs FQDN domain name for particular component
*/}}
{{- define "application-mesh.service-host" -}}
{{ include "application-helpers.name" . }}-{{ .component }}.{{ .Release.Namespace }}.svc.cluster.local
{{- end }}

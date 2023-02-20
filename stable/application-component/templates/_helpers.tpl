{{/*
Create the name of the service account to use
*/}}
{{- define "application-components.serviceAccountName" -}}
{{- default (include "application-helpers.name" .) .Values.serviceAccount.name }}
{{- end }}

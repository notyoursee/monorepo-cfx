{{/*
Return chart name
*/}}
{{- define "node-app.name" -}}
{{- .Chart.Name -}}
{{- end -}}

{{/*
Return full name
*/}}
{{- define "node-app.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

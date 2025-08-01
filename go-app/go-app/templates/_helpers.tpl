{{/*
Return the name of the chart
*/}}
{{- define "go-app.name" -}}
{{- .Chart.Name -}}
{{- end -}}

{{/*
Return the full name of the release
*/}}
{{- define "go-app.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

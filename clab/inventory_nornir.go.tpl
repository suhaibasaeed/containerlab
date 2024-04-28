{{- range .Nodes }}
{{.LongName}}:
  hostname: {{.MgmtIPv4Address}}
  groups:
    - nokia_srlinux
{{- end }}

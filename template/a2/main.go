package main

// template 换行, 去掉空行
import (
	"bytes"
	"fmt"
	"text/template"
)

type Data struct {
	Name   string
	Alerts Alerts
}

type Alerts []Alert

type Alert struct {
	Labels Labels
}

type Labels map[string]string

func main() {
	//	text := `[告警:{{ len .Alerts }}] {{ .Name }}
	//{{ range .Alerts }}
	//{{ range $key, $value := .Labels -}}
	//{{- if eq $key "hostname" -}}
	//主机名: {{ $value }}
	//{{- else if eq $key "instance" -}}
	//告警主机: {{ $value }} {{- end -}}
	//{{- end }}
	//a:b
	//--------
	//{{ end }}
	//`

	text := `[告警:{{ len .Alerts }}] {{ .Name }}
{{ range .Alerts -}}
{{- if .Labels.hostname }}
主机名: {{ .Labels.hostname }} 
{{- end }} 

{{- if .Labels.instance }} 
告警主机: {{ .Labels.instance }} 
{{- end }}
-----------

{{- end }}
	`
	t := template.New("").Option("missingkey=zero")

	tmpl, err := t.Parse(text)
	if err != nil {
		fmt.Printf("err1: %s", err)
		return
	}

	data := Data{
		Name: "test",
		Alerts: Alerts{
			Alert{
				Labels: map[string]string{
					"hostname": "11",
					"instance": "111",
				},
			},
			Alert{
				Labels: map[string]string{
					"hostname": "22",
					"instance": "222",
				},
			},
		},
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		fmt.Printf("err2: %s", err)
		return
	}

	fmt.Println(buf.String())
}

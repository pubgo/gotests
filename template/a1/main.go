package main

import (
	"fmt"
	"github.com/flosch/pongo2"
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
		text, err := pongo2.FromString(`[告警:{{ len(alerts) }}] {{ name }}
	{% for alert in alerts -%}
	{%- if alert.Labels.hostname %}
主机名: {{ alert.Labels.hostname }}
{%- endif %}
{%- if alert.Labels.instance%} 
告警主机: {{ alert.Labels.instance }}
{%- endif %}
--------
{%- endfor %}
	`)
	text1, err := pongo2.FromString(`Standard whitespace control:
{% if true %}
Standard whitespace control
{% endif %}

Full Trim whitespace control:
{% if true -%}
Full Trim whitespace control
{%- endif %}

Useful with logic:
{%- if false %}
1st choice
{%- elif false %}
2nd choice
{%- elif false %}
3rd choice
{%- endif %}

Cycle without whitespace control:
{% for i in simple %}
{{ i }}
{% endfor %}

Cycle with whitespace control:
{% for i in simple -%}
{%- if i==2 %}
1st choice
{%- elif i==3 %}
2nd choice
{%- elif i==4 %}
3rd choice
{%- endif %}
{%- endfor %}

Trim everything:
{% for i in simple -%}
{{ i }}
{%- endfor %}
`)

	_ = text1

	if err != nil {
		panic(err)
	}

	//text := text1

	out, err := text.Execute(pongo2.Context{
		"simple": []int{1, 2, 3, 4, 5, 6, 7, 8},
		"name":   "test",
		"alerts": Alerts{
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
		}})
	if err != nil {
		panic(err)
	}
	fmt.Println(out) // Output: Hello Florian!
}

package models

import "github.com/waldhalf/gotmpl/pkg/forms"

// TemplateData holds data from handlers to template
type TemplateData struct {
	StringMap 	map[string]string
	StringInt 	map[int]int
	StringFloat map[float32]float32
	Data 		map[string]interface{}
	CSRFToken	string
	Flash		string
	Warning		string
	Error		string
	Form		*forms.Form
}
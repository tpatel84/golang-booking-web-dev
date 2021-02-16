package models

import "github.com/tpatel84/golang-booking-web-dev/internal/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	intMap    map[string]int
	FloatMap  map[string]float32
	// string as key and any type as value (interface{}(empty interface) represents any type)
	Data      map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
	Form *forms.Form
}


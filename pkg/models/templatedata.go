package models

//* TemplateData holds data send to handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int64
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

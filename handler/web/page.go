package web

import (
	"html/template"
)

// Page holds page information that is sent to all webpages rendered
// by perseus.
type Page struct {
	SafeList []template.HTML
	List     []string
	Error    []string
	Success  []string
	Notice   []string
	Version  string
}

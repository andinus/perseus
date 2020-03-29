package web

import "html/template"

// Page holds page information.
type Page struct {
	SafeList []template.HTML
	List     []string
	Error    []string
	Success  []string
	Notice   []string
	Version  string
}

package functions

import "html/template"

func FuncMap() template.FuncMap {

	return template.FuncMap{
		"title":  Title,
		"split":  Split,
		"domain": GetDomainName,
	}

}

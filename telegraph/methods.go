package telegraph

import (
	"html"
	"regexp"

	"github.com/anaskhan96/soup"
)

func GetMethods(doc soup.Root) []Method {
	methods := make([]Method, 0)
	for _, methodName := range Methods {
		method := Method{}
		method.Name = methodName
		m := doc.Find("h4", "id", methodName)
		m = m.FindNextSibling()
		method.Description = m.FullText()
		m = m.FindNextSibling()
		method.Parameters = GetMethodParams(m)
		m = m.FindNextSibling()
		method.SampleRequest = m.Find("a").FullText()
		methods = append(methods, method)
	}
	return methods
}

func GetMethodParams(m soup.Root) []Parameter {
	regex := regexp.MustCompile(HTML_REGEX_PATTERN)
	params := make([]Parameter, 0)
	for _, val := range m.FindAll("li") {
		param := Parameter{}
		children := val.Children()
		param.Name = val.Find("strong").FullText()
		param.DataType = GetDataType(children)
		param.Required = IsMethodRequired(children)
		param.Description = html.UnescapeString(regex.ReplaceAllString(GetDescription(children), ""))
		params = append(params, param)
	}
	return params
}

func IsMethodRequired(children []soup.Root) bool {
	for _, child := range children {
		if child.NodeValue == "em" {
			if child.FullText() == "Required" {
				return true
			}
		}
	}
	return false
}

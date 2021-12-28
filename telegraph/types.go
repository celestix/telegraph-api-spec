package telegraph

import (
	"html"
	"regexp"

	"github.com/anaskhan96/soup"
)

func GetTypes(doc soup.Root) []Type {
	types := make([]Type, 0)
	for _, typeName := range Types {
		object := Type{}
		object.Name = typeName
		m := doc.Find("h4", "id", typeName)
		m = m.FindNextSibling()
		object.Description = m.FullText()
		m = m.FindNextSibling()
		object.Attributes = GetTypeAttr(m)
		types = append(types, object)
	}
	return types
}

func GetTypeAttr(m soup.Root) []Attrubute {
	regex := regexp.MustCompile(HTML_REGEX_PATTERN)
	attrs := make([]Attrubute, 0)
	for _, val := range m.FindAll("li") {
		attr := Attrubute{}
		children := val.Children()
		attr.Name = val.Find("strong").FullText()
		attr.DataType = GetDataType(children)
		attr.Optional = IsTypeOptional(children)
		attr.Description = html.UnescapeString(regex.ReplaceAllString(GetDescription(children), ""))
		attrs = append(attrs, attr)
	}
	return attrs
}

func IsTypeOptional(children []soup.Root) bool {
	for _, child := range children {
		if child.NodeValue == "em" {
			if child.FullText() == "Optional" {
				return true
			}
		}
	}
	return false
}

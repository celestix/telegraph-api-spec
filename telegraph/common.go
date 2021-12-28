package telegraph

import (
	"strings"

	"github.com/anaskhan96/soup"
)

func MakeResult(doc soup.Root) Result {
	result := Result{}
	result.Methods = GetMethods(doc)
	result.Types = GetTypes(doc)
	return result
}

func FindMethodsAndTypes(doc soup.Root) {
	for _, vals := range doc.Find("h4", "id", "1.-Methods").FindNextSibling().FindAll("li") {
		Methods = append(Methods, vals.Find("a").FullText())
	}
	for _, vals := range doc.Find("h4", "id", "2.-Types").FindNextSibling().FindAll("li") {
		Types = append(Types, vals.Find("a").FullText())
	}
}

func GetDataType(children []soup.Root) DataTypeInfo {
	dataType := DataTypeInfo{}
	for _, child := range children {
		if child.NodeValue == "em" {
			for _, coreType := range CoreDataTypes {
				if strings.Contains(child.FullText(), coreType) {
					dataType.Name = coreType
					if strings.Contains(child.FullText(), "Array") {
						dataType.IsArray = true
					}
					return dataType
				}
			}
			for _, Type := range Types {
				if child.FullText() == "Array of " {
					dataType.IsArray = true
					child = child.FindNextSibling()
					if child.NodeValue == "a" {
						dataType.Name = child.Find("em").FullText()
					}
					return dataType
				}
				if strings.Contains(child.FullText(), Type) {
					dataType.Name = Type
					if strings.Contains(child.FullText(), "Array") {
						dataType.IsArray = true
					}
					return dataType
				}
			}
		}
	}
	return dataType
}

func GetDescription(children []soup.Root) string {
	gotBreak := false
	text := ""
	for _, child := range children {
		if gotBreak {
			text += child.HTML()
		}
		if child.NodeValue == "br" || (child.NodeValue == "strong" && child.Find("br").NodeValue == "br") {
			gotBreak = true
		}
	}
	return text
}

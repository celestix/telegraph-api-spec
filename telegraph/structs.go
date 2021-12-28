package telegraph

type Result struct {
	Methods []Method `json:"methods"`
	Types   []Type   `json:"types"`
}

type DataTypeInfo struct {
	Name    string `json:"name"`
	IsArray bool   `json:"is_array"`
}

// ****************************************

type Type struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Attributes  []Attrubute `json:"attributes"`
}

type Attrubute struct {
	Name        string       `json:"name"`
	DataType    DataTypeInfo `json:"data_type"`
	Optional    bool         `json:"optional"`
	Description string       `json:"description"`
}

// ****************************************

type Method struct {
	Name          string      `json:"name"`
	Returns       string      `json:"returns"`
	Description   string      `json:"description"`
	SampleRequest string      `json:"sample_request"`
	Parameters    []Parameter `json:"parameters"`
}

type Parameter struct {
	Name        string       `json:"name"`
	DataType    DataTypeInfo `json:"data_type"`
	Required    bool         `json:"required"`
	Description string       `json:"description"`
}

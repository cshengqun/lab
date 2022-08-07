package value_object


type BdQuestion struct {
	Data questionData `json:"data"`
}

type questionData struct {
	Items []item `json:"items"`
}

type item struct {
	Basic basic `json:"basic"`
}

type basic struct {
	Name []name `json:"_name"`
}
type name struct {
	LanguageCode int64 `json:"language_code"`
	Text string `json:"text"`
}

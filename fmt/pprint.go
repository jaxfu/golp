package fmt

import "encoding/json"

type Struct interface {
	~struct{}
}

func PrettifyJSON[T Struct](content T) string {
	marsh, _ := json.MarshalIndent(
		content,
		"",
		" ",
	)
	return string(marsh)
}

package param

type ParamOrder struct {
	Fields []*OrderField `json:"fields"`
}
type OrderField struct {
	Which int `json:"which"`
	Count int `json:"count"`
}
type ParamMenu struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Amount  int64 `json:"amount"`
	Options []int  `json:"options"`
}

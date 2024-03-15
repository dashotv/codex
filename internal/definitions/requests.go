package definitions

type Request struct {
	ID    string `json:"id" query:"id"`
	Limit int    `json:"limit" query:"limit"`
	Skip  int    `json:"skip" query:"skip"`
}

type IndexRequest struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

type KeyRequest struct {
	Key string `json:"id"`
}

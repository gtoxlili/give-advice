package domain

type ChatDto struct {
	Topic   string `json:"topic"`
	Records []struct {
		Q string `json:"q"`
		A string `json:"a"`
	} `json:"records"`
}

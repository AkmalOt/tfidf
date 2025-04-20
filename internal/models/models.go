package models

type Response struct {
	Key   string  `json:"key"`
	Tf    float64 `json:"tf"`
	Idf   float64 `json:"idf"`
	TfIdf float64 `json:"tf_idf"`
}

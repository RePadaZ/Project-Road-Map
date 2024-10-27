package model

type Request struct {
	Meaning     float64 `json:"Meaning"`
	ConvertFrom string  `json:"convertFrom"`
	ConvertTo   string  `json:"convertTo"`
	Flag        string  `json:"flag"`
}

type Response struct {
	Result float64 `json:"result"`
}

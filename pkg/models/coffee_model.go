package models

type Coffee struct {
	ID    string `json:"id" readonly:"true"`
	Water int    `json:"water" readonly:"true"`
	Milk  int    `json:"milk" readonly:"true"`
	Beans int    `json:"beans" readonly:"true"`
	Cost  int    `json:"cost" readonly:"true"`
}

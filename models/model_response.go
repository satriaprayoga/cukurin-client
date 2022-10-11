package models

type ResponseList struct {
	Page      int         `json:"page"`
	Total     int         `json:"total"`
	Last      int         `json:"last"`
	Size      string      `json:"size"`
	Column    string      `json:"column"`
	AllColumn string      `json:"all_column"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

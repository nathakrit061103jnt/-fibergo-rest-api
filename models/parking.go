package models

type Tbl_parkings struct {
	Pk_id    int    `json:"pk_id"`
	Pk_name  string `json:"pk_name"`
	St_id    int    `json:"st_id"`
	St_name  string `json:"st_name"`
	St_color string `json:"st_color"`
}

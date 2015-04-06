package model

type Scheudle struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Mon  int    `json:"mon"`
	Tue  int    `json:"tue"`
	Wed  int    `json:"wed"`
	Thu  int    `json:"thu"`
	Fri  int    `json:"fri"`
	Sat  int    `json:"sat"`
	Sun  int    `json:"sun"`
}

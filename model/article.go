package model

type Article struct {
	// Id
	// in: string
	Id string `json:"Id"`
	// Title
	// in: string
	Title string `json:"title"`
	// Desc
	// in: string
	Desc string `json:"desc"`
	// Content
	// in: string
	Content string `json:"content"`
}

var Articles []Article

package main

type Message struct {
	Message string `json:"message"`
}

type ListItem struct {
	ID   int64  `json:"id"`
	Item string `json:"item"`
}

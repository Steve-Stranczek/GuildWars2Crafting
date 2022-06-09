package main

type itemRequest struct {
	Bearer string `json:"bearer"`
	Item   int    `json:"item"`
}
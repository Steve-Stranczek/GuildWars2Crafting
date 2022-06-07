package main

type materialStorage []struct {
	ID       int `json:"id"`
	Category int `json:"category"`
	Count    int `json:"count"`
}
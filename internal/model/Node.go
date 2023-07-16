package model

type Node struct {
	NodeId   int64
	NodeCode string
	Type     int
	SubType  int // how to find actor by type?
	Config   map[string]string //what is key?
}

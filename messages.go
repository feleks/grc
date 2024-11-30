package main

type message struct {
	Type string `json:"type"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type MouseMoveMessage struct {
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	SelectMode bool    `json:"select_mode"`
}

type KeypressMessage struct {
	Value string `json:"value"`
}

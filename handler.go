package main

import (
	"encoding/json"
	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
	"log"
	"math"
)

type point struct {
	x float64
	y float64
}

type handler struct {
	mouseDown          bool
	mouseMove          bool
	moveLast           point
	scrollAccumulation point
}

func newHandler() *handler {
	return &handler{}
}

func sign(f float64) float64 {
	if f < 0 {
		return -1
	} else if f > 0 {
		return 1
	} else {
		return 0
	}
}

func (h *handler) handle(conn *websocket.Conn, msgRaw []byte) error {
	msg := &message{}
	err := json.Unmarshal(msgRaw, msg)
	if err != nil {
		return err
	}

	switch msg.Type {
	case "mousemove":
		mousemoveMsg := &MouseMoveMessage{}
		err = json.Unmarshal(msgRaw, mousemoveMsg)
		if err != nil {
			return err
		}

		x := mousemoveMsg.X * 3.3
		y := mousemoveMsg.Y * 3.3

		if !h.mouseMove {
			h.mouseMove = true
			h.moveLast = point{}
			log.Printf("start move")
		}

		log.Printf("mousemove message: x=%f (%f), y=%f (%f)", x, h.moveLast.x, y, h.moveLast.y)

		if math.Abs(x-h.moveLast.x) > 3.5 {
			s := sign(x - h.moveLast.x)
			x = h.moveLast.x + (s * 3.5)
		}
		if math.Abs(y-h.moveLast.y) > 3.5 {
			s := sign(y - h.moveLast.y)
			y = h.moveLast.y + (s * 3.5)
		}

		log.Printf("sets to: x=%f, y=%f", x, y)

		h.moveLast = point{
			x: x,
			y: y,
		}

		if mousemoveMsg.SelectMode && !h.mouseDown {
			robotgo.Toggle("left", "down")
			h.mouseDown = true
		}

		robotgo.MoveRelative(int(x), int(y))
	case "end_mousemove":
		h.mouseMove = false
		log.Printf("end move")
		h.moveLast = point{}
	case "end_select":
		robotgo.Toggle("left", "up")
		h.mouseDown = false
	case "click":
		robotgo.Click()
	case "shift+click":
		robotgo.KeyDown("shift")
		robotgo.Click()
		robotgo.KeyUp("shift")
	case "ctrl+click":
		robotgo.KeyDown("ctrl")
		robotgo.Click()
		robotgo.KeyUp("ctrl")
	case "rclick":
		robotgo.Click("right")
	case "scroll":
		mousemoveMsg := &MouseMoveMessage{}
		err = json.Unmarshal(msgRaw, mousemoveMsg)
		if err != nil {
			return err
		}

		h.scrollAccumulation.x += mousemoveMsg.X / 5
		h.scrollAccumulation.y += mousemoveMsg.Y / 5

		log.Printf("scroll message: x=%f (%f), y=%f (%f)", mousemoveMsg.X, h.scrollAccumulation.x, mousemoveMsg.Y, h.scrollAccumulation.y)

		x := 0
		y := 0

		if math.Abs(h.scrollAccumulation.x) >= 1 {
			s := sign(h.scrollAccumulation.x)
			x = int(s * math.Floor(math.Abs(h.scrollAccumulation.x)))
			h.scrollAccumulation.x = math.Mod(h.scrollAccumulation.x, 1.0)
		}
		if math.Abs(h.scrollAccumulation.y) >= 1 {
			s := sign(h.scrollAccumulation.y)
			y = int(s * math.Floor(math.Abs(h.scrollAccumulation.y)))
			h.scrollAccumulation.y = math.Mod(h.scrollAccumulation.y, 1.0)
		}

		//log.Printf("scroll, x=%d, y=%d", x, y)гггггhello

		if x != 0 || y != 0 {
			robotgo.Scroll(x, y)
		}
	case "keypress":
		keypressMsg := &KeypressMessage{}
		err = json.Unmarshal(msgRaw, keypressMsg)
		if err != nil {
			return err
		}

		log.Printf("Keypress msg: %s", keypressMsg.Value)

		if keypressMsg.Value == "$backspace" {
			robotgo.KeyPress("backspace")
		} else {
			robotgo.TypeStr(keypressMsg.Value)
		}
	}

	return nil
}

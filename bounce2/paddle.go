package main

import (
	"strings"
)

type Paddle struct {
	width int 
	height int
	X int
	Y int 
	Xspeed int
}



func (p *Paddle) Display() string {
	return strings.Repeat(" ", p.width)
}


func (p *Paddle) MoveLeft() {

	if p.X > 0 {
		p.X -= p.Xspeed
	}
}

func (p *Paddle) MoveRight(width int) {
	if p.X <  width - p.width {
		p.X += p.Xspeed
	}
}
package main

type Ball struct {
	X int
	Y int
	Xspeed int
	Yspeed int
	bounces int
}

func (b *Ball) Display() string {
	return "\u25CF"
}

func (b *Ball) Update() {
	b.X += b.Xspeed
	b.Y += b.Yspeed
}

func (b *Ball) checkEdges(mw int, mh int) {
	if b.X <= 0 || b.X > mw{
		b.reverseX()
		b.bounces++

	}

	if b.Y<=0 {
		b.reverseY()
		b.bounces++
	}
}

func (b *Ball) Intersects(p Paddle) bool {
	return b.Y == p.Y && b.X >= p.X && b.X <= p.X + p.width
}

func (b *Ball) reverseX() {
	b.Xspeed *= -1
}

func (b *Ball) reverseY() {
	b.Yspeed *= -1
}

func (b *Ball) reset( x int, y int, Xspeed int, Yspeed int) {
	b.X = x
	b.Y =  y
	b.Xspeed = Xspeed
	b.Yspeed = Yspeed
}

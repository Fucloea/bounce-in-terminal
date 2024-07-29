package main

import (
	"github.com/gdamore/tcell/v2"
	"time"
	 "strconv"
)

type Game struct {
	Screen tcell.Screen
	Ball Ball
	Player Player
}


func (g *Game) GameOver() bool {
	return g.Player.Score > 20
}


func (g *Game) Run() {
	s := g.Screen
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	for {

		s.Clear()


		width, height := s.Size()


		if g.GameOver() {
			drawSprite(s, (width/2)-4, 7, (width/2)+5, 7, defStyle, "YOU LOST!")
			s.Show()
		}


		if g.Ball.Intersects(g.Player.Paddle)  {
			g.Ball.reverseX()
			g.Ball.reverseY()
		}


		if g.Ball.X <= 0 {
			g.Ball.reset(width/2, height/2, -1, 1)
		}

		if g.Ball.X >= width {
			g.Ball.reset(width/2, height/2, 1, 1)
		}

		if g.Ball.Y > height {
			g.Player.Score++
			g.Ball.reset(width/2, height/2, -1, -1)
		}


		pc := strconv.Itoa(g.Player.Score)
		for i,r := range []rune(pc) {
			s.SetContent((width/2)+i,(height/2), r, nil, defStyle)
		}


		g.Ball.checkEdges(width, height)
		g.Ball.Update()

		drawSprite(s,
			g.Ball.X,
			g.Ball.Y,
			g.Ball.X,
			g.Ball.Y,
			defStyle,
			g.Ball.Display())


		paddleStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
		drawSprite(s,
			g.Player.Paddle.X,
			g.Player.Paddle.Y,
			g.Player.Paddle.X+g.Player.Paddle.width,
			g.Player.Paddle.Y+g.Player.Paddle.height,
			paddleStyle,
			g.Player.Paddle.Display())

	


		time.Sleep(40 * time.Millisecond)
		s.Show()
	}

}

func drawSprite(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1

	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}
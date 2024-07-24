package main

import (
	"github.com/gdamore/tcell/v2"
	"time"
)

type Game struct {
	Screen tcell.Screen
	Ball Ball
}


func (g *Game) Run() () {
	

	defStyle := tcell.StyleDefault.Background(tcell.ColorAqua).Foreground(tcell.ColorRed)
	g.Screen.SetStyle(defStyle)
	

	for {

		width, height := g.Screen.Size()
		g.Screen.Clear()
        g.Ball.Update()

        g.Screen.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, defStyle)
        g.Ball.checkEdges(width,height)
        g.Screen.Show()

        time.Sleep(40*time.Millisecond)

    }
}

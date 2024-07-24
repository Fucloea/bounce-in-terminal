package main

import (
	"github.com/gdamore/tcell/v2"
	"time"
	 "strconv"
)

type Game struct {
	Screen tcell.Screen
	Ball Ball
}


func (g *Game) Run() () {
	

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorRed)
	g.Screen.SetStyle(defStyle)
	

	for {

		width, height := g.Screen.Size()
		g.Screen.Clear()
        g.Ball.Update()
        g.Screen.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, defStyle)
        display_bounce_count(g)
        g.Ball.checkEdges(width,height)
        time.Sleep(40*time.Millisecond)
        g.Screen.Show()

    }
}

func display_bounce_count(g *Game) {

	width, height := g.Screen.Size()
	defS := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorAqua)
	str := strconv.Itoa(g.Ball.bounces)
    runes := []rune(str)

    for i, r := range runes {
            g.Screen.SetContent((width/2)+i,height/2, r, nil, defS) 
    }
}
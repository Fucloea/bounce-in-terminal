package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
)


func main() {

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}


	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	width, height := s.Size()

	ball := Ball{
		X:      width / 2,
		Y:      10,
		Xspeed: 1,
		Yspeed: 1,
	}

	player1 := Player{
		Score: 0,
		Paddle: Paddle{
			width:  15,
			height: 1,
			Y:      height - 5,
			X:      width/2,
			Xspeed: 3,
		},
	}
	game := Game{
		Screen:  s,
		Ball:    ball,
		Player: player1,
	}

	go game.Run()

	for {

		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyRight {
				game.Player.Paddle.MoveRight(width)
			} else if event.Key() == tcell.KeyLeft {
				game.Player.Paddle.MoveLeft()

			}
		}
	}

}
package main

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// handle user trying to close the window
func windowClosingHandler() {
	DB.Close()
	//time.Sleep(time.Second * 1)
}

func inputEventHandler(screen *ebiten.Image, mouseInput bool, kbdInput bool, buttons []*Button, textboxes []*TextBox) {
	if mouseInput {
		cursorX, cursorY := ebiten.CursorPosition()
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			for _, button := range buttons {
				button.Click(screen, cursorX, cursorY)
			}
			wg.Done()
		}()
		go func() {
			for _, textbox := range textboxes {
				if withinBounds(cursorX, cursorY, textbox.posX, textbox.posY, textbox.width, textbox.height) {
					// #+TODO: implement textbox input
					vector.DrawFilledRect(screen, float32(textbox.posX), float32(textbox.posY), float32(textbox.width), float32(textbox.height), Teal, true)
				}
			}
			wg.Done()
		}()
		wg.Wait()
	}
	if kbdInput {
		cursorX := KBDC.X
		cursorY := KBDC.Y
		for _, textbox := range textboxes {
			if withinBounds(cursorX, cursorY, textbox.posX, textbox.posY, textbox.width, textbox.height) {
				// #+TODO: implement textbox input
				// #+TODO: implement vim motions
				vector.DrawFilledRect(screen, float32(textbox.posX), float32(textbox.posY), float32(textbox.width), float32(textbox.height), Teal, true)
			}
		}
	}
}

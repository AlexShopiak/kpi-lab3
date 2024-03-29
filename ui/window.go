package ui

import (
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/imageutil"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"

	"image"
	"image/color"
	"log"
)

type Visualizer struct {
	Width         int
	Height        int
	Title         string
	Debug         bool
	OnScreenReady func(s screen.Screen)

	w    screen.Window
	tx   chan screen.Texture
	done chan struct{}

	sz  size.Event
	mouseX int
	mouseY int
}

func (pw *Visualizer) Main() {
	pw.tx = make(chan screen.Texture)
	pw.done = make(chan struct{})
	pw.mouseX = pw.Width/2
	pw.mouseY = pw.Height/2
    driver.Main(pw.run)
}

func (pw *Visualizer) Update(t screen.Texture) {
	pw.tx <- t
}

func (pw *Visualizer) run(s screen.Screen) {
	

	w, err := s.NewWindow(&screen.NewWindowOptions{
		Width:  pw.Width,
		Height:  pw.Height,
		Title:  pw.Title,
	})

	if err != nil {
		log.Fatal("Failed to initialize the app window:", err)
	}
	
	defer func() {
		w.Release()
		close(pw.done)
	}()

	if pw.OnScreenReady != nil {
		pw.OnScreenReady(s)
	}

	pw.w = w

	events := make(chan any)
	go func() {
		for {
			e := w.NextEvent()
			if pw.Debug {
				log.Printf("new event: %v", e)
			}
			if detectTerminate(e) {
				close(events)
				break
			}
			events <- e
		}
	}()

	var t screen.Texture

	for {
		select {
		case e, ok := <-events:
			if !ok {
				return
			}
			pw.handleEvent(e, t)

		case t = <-pw.tx:
			w.Send(paint.Event{})
		}
	}
}

func detectTerminate(e any) bool {
	switch e := e.(type) {
	case lifecycle.Event:
		if e.To == lifecycle.StageDead {
			return true // Window destroy initiated.
		}
	case key.Event:
		if e.Code == key.CodeEscape {
			return true // Esc pressed.
		}
	}
	return false
}

func (pw *Visualizer) handleEvent(e any, t screen.Texture) {
	switch e := e.(type) {

		case size.Event: // Оновлення даних про розмір вікна.
			pw.sz = e
	
		case error:
			log.Printf("ERROR: %s", e)
	
		case mouse.Event:
			if t == nil {
				if e.Button == mouse.ButtonLeft && e.Direction == mouse.DirPress {
					pw.mouseX = int(e.X)
					pw.mouseY = int(e.Y)
					pw.w.Send(paint.Event{})
				}
			}
	
		case paint.Event:
			// Малювання контенту вікна.
			if t == nil {
				pw.drawDefaultUI()
			} else {
				// Використання текстури отриманої через виклик Update.
				pw.w.Scale(pw.sz.Bounds(), t, t.Bounds(), draw.Src, nil)
			}
			pw.w.Publish()
		}
}

func (pw *Visualizer) drawDefaultUI() {
	pw.w.Fill(pw.sz.Bounds(), color.White, draw.Src) // Фон.

	w4 := pw.Width/4 
	w8 := pw.Width/8 
	x, y := pw.mouseX, pw.mouseY
	c := color.RGBA{R: 225, G: 225, B: 0, A: 1}
	
	pw.w.Fill(image.Rect(x - w4, y - w4, x + w4, y), c, draw.Src)
	pw.w.Fill(image.Rect(x - w8, y, x + w8, y + w4), c, draw.Src)

	// Малювання білої рамки.
	for _, br := range imageutil.Border(pw.sz.Bounds(), 10) {
		pw.w.Fill(br, color.White, draw.Src)
	}
}

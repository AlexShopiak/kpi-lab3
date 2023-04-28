package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

func Reset(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, screen.Src) //todo
}

type BgRect struct{
	x1, y1, x2, y2 int 
}
func (b *BgRect) Do(t screen.Texture) bool {
	c := color.Black
	t.Fill(image.Rect(b.x1, b.y1, b.x2, b.y2), c, screen.Src)
	return false
}

type Figure struct{
	x, y int //todo
}
func (f Figure) Do(t screen.Texture) bool {
	return false //todo
}

type Move struct{
	x, y int //todo
}
func (m Move) Do(t screen.Texture) bool {
	return false //todo
}
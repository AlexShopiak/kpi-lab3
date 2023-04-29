package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

type Operation interface { // Operation змінює вхідну текстуру.
	Do(t screen.Texture) (ready bool) // Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
}

type OperationList []Operation // OperationList групує список операції в одну.
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
	c := color.White
	t.Fill(t.Bounds(), c, screen.Src)
}
// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	c := color.RGBA{G: 0xff, A: 0xff}
	t.Fill(t.Bounds(), c, screen.Src)
}
func Reset(t screen.Texture) {
	c := color.Black
	t.Fill(t.Bounds(), c, screen.Src)
}

type BgRect struct{
	x1, y1, x2, y2 int 
}
func (b *BgRect) Do(t screen.Texture) bool {
	c := color.Black
	t.Fill(image.Rect(b.x1, b.y1, b.x2, b.y2), c, screen.Src)
	return false
}
var figures []Figure //todo
type Figure struct{
	x, y int
}
func (f Figure) Do(t screen.Texture) bool {
	c := color.RGBA{R: 225, G: 225, B: 0, A: 1}
	t.Fill(image.Rect(f.x-200, f.y-200, f.x+200, f.y), c, screen.Src)
	t.Fill(image.Rect(f.x-100, f.y-200, f.x+100, f.y+200), c, screen.Src)
	return false 
}

type Move struct{
	x, y int
}
func (m Move) Do(t screen.Texture) bool {
	for _, figure := range figures {
		figure.x += m.x
		figure.y += m.y
		//todo
	}
	return false
}
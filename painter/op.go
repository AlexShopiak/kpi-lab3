package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

//Зберігає поточний стан текстури
type TextureState struct {
	BgClr color.Color
	BgRect Operation
	Figures []Figure
}
var state = TextureState{}

//Змінює вхідну текстуру
//Повертаючи true, якщо текстура готова для відображення
type Operation interface { 
	Do(t screen.Texture) (ready bool) 
}

//Групує список операції в одну
type OperationList []Operation 
func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

//Сигналізує, що текстура готова
type updateOp struct{}
func (op updateOp) Do(t screen.Texture) bool { return true }
var UpdateOp = updateOp{}

//Перетворює функцію оновлення текстури в Operation
type OperationFunc func(t screen.Texture)
func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

//Зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	state.BgClr = color.White
	t.Fill(t.Bounds(), state.BgClr, screen.Src)
}
//Зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	state.BgClr = color.RGBA{G: 0xff, A: 0xff}
	t.Fill(t.Bounds(), state.BgClr , screen.Src)
}


type BgRect struct{
	x1, y1, x2, y2 int 
}
func (r *BgRect) Do(t screen.Texture) bool {
	state.BgRect = r
	c := color.Black
	t.Fill(image.Rect(r.x1, r.y1, r.x2, r.y2), c, screen.Src)
	return false
}


type Figure struct{
	x, y int
}
func (f *Figure) Do(t screen.Texture) bool {
	state.Figures = append(state.Figures, *f)
	c := color.RGBA{R: 225, G: 225, B: 0, A: 1}
	t.Fill(image.Rect(f.x-200, f.y-200, f.x+200, f.y), c, screen.Src)
	t.Fill(image.Rect(f.x-100, f.y-200, f.x+100, f.y+200), c, screen.Src)
	return false 
}


type Move struct{
	x, y int
}
func (m *Move) Do(t screen.Texture) bool {
	for _, f := range state.Figures {
		f.x += m.x
		f.y += m.y
		//todo
	}
	return false
}


func Reset(t screen.Texture) {	
	state.BgClr = color.Black
	state.BgRect = nil
	state.Figures = state.Figures[:0] //TODO can nil be here

	t.Fill(t.Bounds(), state.BgClr, screen.Src)
}
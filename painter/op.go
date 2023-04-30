package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

//Зберігає поточний стан текстури
var state = TextureState{}
type TextureState struct {
	BgRect Operation
	Figures []FigureOp
}

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
type UpdateOp struct{}
func (op UpdateOp) Do(t screen.Texture) bool { return true }

//Перетворює функцію оновлення текстури в Operation
type OperationFunc func(t screen.Texture)
func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

//Зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}
//Зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff} , screen.Src)
}
//Очищує текстуру та зафарбовує у чорний колір. Може бути викоистана як Operation через OperationFunc(Reset).
func Reset(t screen.Texture) {	
	state.BgRect = nil
	state.Figures = state.Figures[:0] //TODO can nil be here
	t.Fill(t.Bounds(), color.Black, screen.Src)
}



type BgrectOp struct{
	X1, Y1, X2, Y2 int 
}
func (r BgrectOp) Do(t screen.Texture) bool {
	state.BgRect = r
	c := color.Black
	t.Fill(image.Rect(r.X1, r.Y1, r.X2, r.Y2), c, screen.Src) //neeedd
	return false
}



type FigureOp struct{
	X, Y int
}
func (f FigureOp) Do(t screen.Texture) bool {
	state.Figures = append(state.Figures, f)
	c := color.RGBA{R: 225, G: 225, B: 0, A: 1}
	t.Fill(image.Rect(f.X-200, f.Y-200, f.X+200, f.Y), c, screen.Src)
	t.Fill(image.Rect(f.X-100, f.Y, f.X+100, f.Y+200), c, screen.Src)
	return false 
}


type MoveOp struct{
	X, Y int
}
func (m MoveOp) Do(t screen.Texture) bool {
	for _, f := range state.Figures {
		f.X += m.X
		f.Y += m.Y
		//todo
	}
	return false
}
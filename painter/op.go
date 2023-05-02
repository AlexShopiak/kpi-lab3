package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

//Зберігає поточний стан текстури
var state = TextureState{
	color.White,
	false,
	BgrectOp{},
	false,
    []FigureOp{},
}

type TextureState struct {
	BgClr      color.Color
	HasRect    bool
	BgRect     BgrectOp
	HasFigures bool
	Figures    []FigureOp
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

//Поетапно малює 3 шари
//Власне малювання відбувається один раз в кінці, аби зберегти ресурси та не малювати в холосту
//Сигналізує, що текстура готова
type UpdateOp struct{}
func (op UpdateOp) Do(t screen.Texture) bool {
	//1st layer
	t.Fill(t.Bounds(), state.BgClr, screen.Src)
	//2nd layer
	if state.HasRect {
		r := state.BgRect
		t.Fill(image.Rect(r.X1, r.Y1, r.X2, r.Y2), color.Black, screen.Src) 
	}
    //3rd layer
	if state.HasFigures {
		for _, f := range state.Figures {
			c := color.RGBA{R: 225, G: 225, B: 0, A: 1}
			t.Fill(image.Rect(f.X-100, f.Y-100, f.X+100, f.Y), c, screen.Src)
			t.Fill(image.Rect(f.X-50, f.Y, f.X+50, f.Y+100), c, screen.Src)
		}
	}
	return true
}

//Перетворює функцію оновлення текстури в Operation
type OperationFunc func(t screen.Texture)
func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

//Зафарбовує тестуру у білий колір
//Може бути викоистана як Operation через OperationFunc(WhiteFill)
func WhiteFill(t screen.Texture) {
	state.BgClr = color.White                          
}

//Зафарбовує тестуру у зелений колір
//Може бути викоистана як Operation через OperationFunc(GreenFill)
func GreenFill(t screen.Texture) {
	state.BgClr = color.RGBA{R: 0, G: 255, B: 0, A: 1}
}

//Очищує текстуру та зафарбовує у чорний колір
//Може бути викоистана як Operation через OperationFunc(Reset)
func Reset(t screen.Texture) {	
	state.BgClr = color.Black 
	state.HasRect = false
	state.BgRect = BgrectOp{}
	state.HasFigures = false
	state.Figures = []FigureOp{}
}

//Тестова операція для використання у сриптах. Не бажана для використання
type CustomFill struct{
	R,G,B,A int 
}
func (c CustomFill) Do(t screen.Texture) bool {
	state.BgClr = color.RGBA{R: uint8(c.R), G: uint8(c.G), B: uint8(c.B), A: uint8(c.A)}
	return false
}


type BgrectOp struct{ X1, Y1, X2, Y2 int }
func (r BgrectOp) Do(t screen.Texture) bool {
	state.BgRect = BgrectOp{r.X1, r.Y1, r.X2, r.Y2}
	state.HasRect = true 
	return false
}


type FigureOp struct{ X, Y int }
func (f FigureOp) Do(t screen.Texture) bool {	
    state.Figures = append(state.Figures, FigureOp{f.X, f.Y})
	state.HasFigures = true 
	return false 
}

type MoveOp struct{ X, Y int }
func (m MoveOp) Do(t screen.Texture) bool {
	for i := range state.Figures {
		state.Figures[i].X += m.X
		state.Figures[i].Y += m.Y
	}
	return false
}
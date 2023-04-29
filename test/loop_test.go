package test

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"github.com/AlexShopiak/kpi-lab3/painter"
	"golang.org/x/exp/shiny/screen"
)
func TestLoop_Post(t *testing.T) {
	var (
		l painter.Loop
		tr TestReceiver
	)
	l.Receiver = &tr

	l.Start(mockScreen{})

	l.Post(painter.OperationFunc(painter.WhiteFill)) //#1
	l.Post(painter.OperationFunc(painter.GreenFill)) //#2
	l.Post(painter.UpdateOp)

	if tr.LastTexture != nil {
		t.Fatal("Receiver got the texture too early")
	}

	l.StopAndWait()

	tx, ok := tr.LastTexture.(*mockTexture)
	if !ok {
		t.Fatal("Receiver still nasn't texture")
	}
	if tx.FillCnt != 2 {
		t.Fatal("Unexpected num of Fill calls")
	}
}

//Mocks
type TestReceiver struct {
	LastTexture screen.Texture	
}
func (tr *TestReceiver) Update(t screen.Texture) {
	tr.LastTexture = t
}



type mockScreen struct {}
func (m mockScreen) NewBuffer(size image.Point) (screen.Buffer, error) {
	panic("Implement me")
}
func (m mockScreen) NewTexture(size image.Point) (screen.Texture, error) {
	return new(mockTexture), nil
}
func (m mockScreen) NewWindow(opts *screen.NewWindowOptions) (screen.Window, error) {
	panic("Implement me")
}



type mockTexture struct {
	FillCnt int
}
func (m *mockTexture) Release() {}
func (m *mockTexture) Size() image.Point { return image.Pt(400, 400) }
func (m *mockTexture) Bounds() image.Rectangle { return image.Rectangle{Max: image.Pt(400, 400)} }
func (m *mockTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	panic("Implement me")
}
func (m *mockTexture) Fill(dp image.Rectangle, src color.Color, op draw.Op) {
	m.FillCnt++
}
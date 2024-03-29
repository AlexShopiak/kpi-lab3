package test

import (
	"image"
	"image/color"
	"image/draw"
	"reflect"
	"testing"

	"github.com/AlexShopiak/kpi-lab3/painter"
	"golang.org/x/exp/shiny/screen"
)
func TestLoop_EarlyTexture(t *testing.T) {
	var l painter.Loop
	var	tr TestReceiver
	l.Receiver = &tr
	l.Start(mockScreen{})

	l.Post(painter.OperationFunc(painter.Reset)) 
	l.Post(painter.OperationFunc(painter.WhiteFill)) 
	l.Post(painter.OperationFunc(painter.GreenFill)) 
	l.Post(painter.UpdateOp{})

	if tr.LastTexture != nil {
		t.Fatal("Receiver got the texture too early")
	}

	l.StopAndWait()

	if tr.LastTexture == nil {
		t.Fatal("Receiver still nasn't texture")
	}
}

func TestLoop_Optimisation(t *testing.T) {
	var l painter.Loop
	var	tr TestReceiver
	l.Receiver = &tr
	l.Start(mockScreen{})

	l.Post(painter.OperationFunc(painter.GreenFill)) 
	l.Post(painter.OperationFunc(painter.WhiteFill)) 
	l.Post(painter.OperationFunc(painter.GreenFill)) 
	l.Post(painter.UpdateOp{})

	if l.MQ.Length() != 4 {
		t.Fatal("Bad number of operations")
	}

	l.StopAndWait()

	tx, _ := tr.LastTexture.(*mockTexture)
	if tx.FillCnt != 1 {
		t.Fatal("Too much usless Fill calls")
	}

	if l.MQ.Length() != 0 {
		t.Fatal("Bad number of operations")
	}
}

func TestLoop_Post3(t *testing.T) {
	var (
		l painter.Loop
		tr TestReceiver
	)
	l.Receiver = &tr
	l.Start(mockScreen{})
	var testOps []string

	l.Post(painter.OperationFunc(func(t screen.Texture) {
		testOps = append(testOps, "op1")

		l.Post(painter.OperationFunc(func(t screen.Texture) {
			testOps = append(testOps, "op3")
		}))
	}))

	l.Post(painter.OperationFunc(func(t screen.Texture) {
		testOps = append(testOps, "op2")
	}))

	l.StopAndWait()

	if !reflect.DeepEqual(testOps, []string{"op1","op2","op3"}) {
		t.Fatal("Bad order of operations")
	}
}

//=================================Mocks===================================
//=========================================================================
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
func (m *mockTexture) Size() image.Point {
	return image.Pt(400, 400)
}
func (m *mockTexture) Bounds() image.Rectangle {
	return image.Rectangle{Max: image.Pt(400, 400)}
}
func (m *mockTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	panic("Implement me")
}
func (m *mockTexture) Fill(dp image.Rectangle, src color.Color, op draw.Op) {
	m.FillCnt++
}
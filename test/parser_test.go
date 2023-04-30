package test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/AlexShopiak/kpi-lab3/painter"
	"github.com/AlexShopiak/kpi-lab3/painter/lang"
)

func TestParser_Type1(t *testing.T) {
	p := lang.Parser{}
	
	expected := []painter.Operation{
		painter.OperationFunc(painter.WhiteFill),
		painter.BgrectOp{X1: 1,Y1: 2,X2: 3,Y2: 4},
		painter.FigureOp{X: 1,Y: 2},
		painter.OperationFunc(painter.GreenFill),
		painter.FigureOp{X: 3,Y: 4},
		painter.UpdateOp{},
	}

	r := strings.NewReader("white\n bgrect 1 2 3 4\n figure 1 2\n green\n figure 3 4\n update")
	res, err := p.Parse(r)
	if err != nil {
		t.Fatal("Bad parsing")
	}

	for i := range expected {
		if !reflect.DeepEqual(reflect.TypeOf(expected[i]), reflect.TypeOf(res[i])) {
			t.Fatal("Bad operation order")
		}
	}
}

//=================================Mocks===================================
//=========================================================================
/*
type mockReader struct {}

func (r mockReader) Read(bytes []byte) (int, error) {
	b := []byte("green \nwhite \nbgrect 1 2 3 4 \nupdate")
	for i := range bytes {
		bytes[i] = b[i]
	}
	return len(bytes), nil
}


type MyReader struct {
	data []byte
	readIndex int64
}

func NewReader(toRead string) *MyReader {
	return &MyReader{data: []byte(toRead)}
}


func (r *MyReader) Read(p []byte) (n int, err error) {
	if r.readIndex >= int64(len(r.data)) {
		err = io.EOF
		return
	}
	
	n = copy(p, r.data[r.readIndex:])
	r.readIndex += int64(n)
	return
}*/


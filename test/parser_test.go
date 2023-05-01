package test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/AlexShopiak/kpi-lab3/painter"
	"github.com/AlexShopiak/kpi-lab3/painter/lang"
)

func TestParser_Type(t *testing.T) {
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

func TestParser_InvalitOperation(t *testing.T) {
	p := lang.Parser{}
	r := strings.NewReader("upgrade")
	_, err := p.Parse(r)

	if  err != lang.InvOprErr {
		t.Fatal("Doesnt catch InvOprErr")
	}

}

func TestParser_InvalidParams(t *testing.T) {
	p := lang.Parser{}
	r := strings.NewReader("bgrect 100 Alex 300 Andrew")
	_, err := p.Parse(r)

	if  err != lang.InvParErr {
		t.Fatal("Doesnt catch InvParErr")
	}

}

func TestParser_LittleParams(t *testing.T) {
	p := lang.Parser{}
	r := strings.NewReader("bgrect 100 200")
	_, err := p.Parse(r)

	if  err != lang.LitParErr {
		t.Fatal("Doesnt catch LitParErr")
	}

}
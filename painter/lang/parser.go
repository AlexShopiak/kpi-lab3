package lang

import (
    "bufio"
    "io"
    "strconv"
    "strings"

    "github.com/AlexShopiak/kpi-lab3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
    scanner := bufio.NewScanner(in)
    scanner.Split(bufio.ScanLines)
    var res []painter.Operation 

    for scanner.Scan() {     
        commandLine := scanner.Text() 
        op, err := parseLine(commandLine) // parse the line to get Operation  
		if err != nil {
			return nil, err
		}
        res = append(res, op) 
    }
    return res, nil 
}

func parseLine(line string) (painter.Operation, error) {
	cmd, params, err := getCmdAndParams(line)
	if err != nil {
		return nil, err
	}
	switch cmd {
	case "fill":
		return painter.CustomFill{
			R:params[0],
			G:params[1],
			B:params[2],
			A:params[3],
		}, nil
	case "white":
		return painter.OperationFunc(painter.WhiteFill), nil
	case "green":
		return painter.OperationFunc(painter.GreenFill), nil
	case "update":
		return painter.UpdateOp{}, nil
	case "bgrect":
		return painter.BgrectOp{
			X1:params[0],
			Y1:params[1],
			X2:params[2],
			Y2:params[3],
		}, nil
	case "figure":
		return painter.FigureOp{
			X:params[0],
			Y:params[1],
		}, nil
	case "move":
		return painter.MoveOp{
			X:params[0],
			Y:params[1],
		}, nil
	case "reset":
		return painter.OperationFunc(painter.Reset), nil
	} 
	return nil, InvOprErr
}

func getCmdAndParams(cl string) (string, []int, error) {
	fields := strings.Fields(cl)
	cmd := fields[0]
	params := []int{}
	
	for _,field := range fields[1:] {
		res, err := strconv.ParseFloat(field, 32)
		if err != nil {
			return "", nil, InvParErr
		}
		params = append(params, int(res))
	}

	switch cmd {
	case "bgrect","fill":
		if len(params) < 4 {
			return "", nil, LitParErr
		}
	case "figure","move":
		if len(params) < 2 {
			return "", nil, LitParErr
		}
	}

	return cmd, params, nil
}
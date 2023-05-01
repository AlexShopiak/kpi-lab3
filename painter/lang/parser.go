package lang

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/AlexShopiak/kpi-lab3/painter"
)

type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	var res []painter.Operation  
  for scanner.Scan() {     
		commandLine := scanner.Text() 
    op, _ := parseLine(commandLine) // parse the line to get Operation  
    res = append(res, op) 
  }
	return res, nil }

func parseLine(line string) (painter.Operation, error) {
  cmd, params, err := getCmdAndParams(line)
  if err != nil {
    return nil, err
  }
  switch cmd {
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
  return nil, errors.New("Invalid operation")
}

func getCmdAndParams(cl string) (string, []int, error) {
  fields := strings.Fields(cl)
  if len(fields) == 0 {
    return "", nil, errors.New("Empty string")
  }
  cmd := fields[0]
  params := []int{}
  
  for _,field := range fields[1:] {
    res, err := strconv.ParseFloat(field, 32)
    if err != nil {
      return "", nil, errors.New("Invalid params")
    }
    params = append(params, int(res))
  }

  switch cmd {
  case "bgrect":
    if len(params) < 4 {
      return "", nil, errors.New("Little params")
    }
  case "figure","move":
    if len(params) < 2 {
      return "", nil, errors.New("Little params")
    }
  }

  return cmd, params, nil
}
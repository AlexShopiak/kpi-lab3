package lang

import "errors"

var (
	InvOprErr = errors.New("Invalid operation")
	InvParErr = errors.New("Invalid params")
	LitParErr = errors.New("Little params")
)
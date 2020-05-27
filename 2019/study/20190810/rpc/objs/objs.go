package objs

import (
	"fmt"
)

type Request struct {
	Left  int
	Right int
}

type Response struct {
	Result int
}

type Calc struct{}

func (c *Calc) Sum(r *Request, rp *Response) error {
	fmt.Println("Sum: ", r.Right, r.Left)
	rp.Result = r.Left + r.Right
	return nil
}

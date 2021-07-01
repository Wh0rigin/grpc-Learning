package services

import (
	context "context"
	"fmt"
)

type BlackBoxServices struct{}

func (this *BlackBoxServices) SayHello(ctx context.Context, in *Request) (*Respond, error) {
	sendmsg := "hello:" + fmt.Sprint(in.GetId())
	return &Respond{Msg: sendmsg}, nil
}

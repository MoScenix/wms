package service

import (
	"context"
	user "github.com/MoScenix/wms/rpc_gen/kitex_gen/user"
	"testing"
)

func TestListUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListUserService(ctx)
	// init req and assert value

	req := &user.ListUserReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

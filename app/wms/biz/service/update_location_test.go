package service

import (
	"context"
	"testing"
	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
)

func TestUpdateLocation_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateLocationService(ctx)
	// init req and assert value

	req := &wms.UpdateLocationReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

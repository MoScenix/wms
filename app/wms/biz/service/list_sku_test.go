package service

import (
	"context"
	"testing"
	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
)

func TestListSku_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListSkuService(ctx)
	// init req and assert value

	req := &wms.ListSkuReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

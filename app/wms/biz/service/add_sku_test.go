package service

import (
	"context"
	"testing"
	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
)

func TestAddSku_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddSkuService(ctx)
	// init req and assert value

	req := &wms.AddSkuReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

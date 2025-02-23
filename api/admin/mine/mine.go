// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package mine

import (
	"context"

	"github.com/greper/ulogin/api/admin/mine/v1"
)

type IMineV1 interface {
	Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error)
}

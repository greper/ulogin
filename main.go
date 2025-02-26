package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/greper/ulogin/internal/boot"
	_ "github.com/greper/ulogin/internal/packed"

	"github.com/greper/ulogin/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	ctx := gctx.GetInitCtx()
	boot.Init(ctx)
	cmd.Main.Run(ctx)
}

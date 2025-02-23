package main

import (
	_ "github.com/greper/ulogin/internal/packed"
	"github.com/greper/ulogin/utility"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/greper/ulogin/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	utility.DbUp()
	cmd.Main.Run(gctx.GetInitCtx())
}

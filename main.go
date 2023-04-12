package main

import (
	_ "learn_goframe01/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"learn_goframe01/internal/cmd"

	_ "learn_goframe01/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.New())
}

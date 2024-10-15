package action

import (
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/goctl-swagger/power"
)

func PowerGenerate(ctx *cli.Context) error {
	fileName := "policy.csv"
	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	return power.Do(fileName, p)
}

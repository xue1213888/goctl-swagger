package power

import (
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"strconv"
)

func applyGenerate(in *plugin.Plugin) [][]string {
	var res [][]string
	for _, group := range in.Api.Service.Groups {
		for _, route := range group.Routes {
			path := group.GetAnnotation("prefix") + route.Path
			if path[0] != '/' {
				path = "/" + path
			}
			if path == "" {
				continue
			}
			power := ""
			if route.AtDoc.Properties != nil {
				power, _ = strconv.Unquote(route.AtDoc.Properties["x_power"])
			}
			if power == "" {
				power = group.GetAnnotation("x_power")
			}
			if power == "" {
				continue
			}
			power = "sys_" + power
			res = append(res, []string{
				"p",
				power,
				path,
			})
		}
	}
	return res
}

package iris

import (
	"github.com/kataras/iris/v12"
	"github.com/yangguifnag/simplex2go/consul"
)

type Iris struct {
	Port            int    `json:"port"`
	Host            string `json:"host"`
	HasConsul       bool   `json:"hasConsul"`
	ConsulCheckName string `json:"consulCheckName"`
}

func (i *Iris) Init() *iris.Application {
	app := iris.New()

	if i.HasConsul {
		if i.ConsulCheckName == "" {
			i.ConsulCheckName = consul.DEFAULT_CHECK_NAME
		}
		app.Get(`/`+i.ConsulCheckName, func(ctx iris.Context) {
			ctx.JSON(iris.Map{"status": "UP"})
		})
	}
	return app

}

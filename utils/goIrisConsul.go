package utils

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/kataras/iris/v12"
	"github.com/yangguifnag/simplex2go/consul"
	iris_ "github.com/yangguifnag/simplex2go/iris"
)

type GoIrisConsul struct {
	consulConfig consul.BaseConfig
	irisConfig   iris_.Iris
}

func (conf *GoIrisConsul) Init() (*iris.Application, *consulapi.Client, error) {
	consulApp, err := conf.consulConfig.RegisterService()
	if err != nil {
		return nil, nil, err
	}

	irisApp := conf.irisConfig.Init()
	err = irisApp.Run(iris.Addr(fmt.Sprintf(":%d", conf.consulConfig.Port)))
	if err != nil {
		return nil, nil, err
	}

	return irisApp, consulApp, nil
}

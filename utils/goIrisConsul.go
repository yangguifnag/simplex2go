package utils

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/kataras/iris/v12"
	"github.com/yangguifnag/simplex2go/consul"
	iris_ "github.com/yangguifnag/simplex2go/iris"
)

type GoIrisConsul struct {
	ConsulConfig consul.BaseConfig
	IrisConfig   iris_.Iris
	irisApp      *iris.Application
	consulApp    *consulapi.Client
}

func (conf *GoIrisConsul) Init() (*iris.Application, *consulapi.Client, error) {
	consulApp, err := conf.ConsulConfig.RegisterService()
	if err != nil {
		return nil, nil, err
	}

	irisApp := conf.IrisConfig.Init()
	//err = irisApp.Run(iris.Addr(fmt.Sprintf(":%d", conf.ConsulConfig.Port)))
	//if err != nil {
	//	return nil, nil, err
	//}

	return irisApp, consulApp, nil
}

func (conf *GoIrisConsul) IrisRun() (*iris.Application, error) {
	irisApp := conf.irisApp
	err := irisApp.Run(iris.Addr(fmt.Sprintf(":%d", conf.ConsulConfig.Port)))
	if err != nil {
		return nil, err
	}

	return irisApp, nil

}

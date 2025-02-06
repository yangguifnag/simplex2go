package consul

import consulapi "github.com/hashicorp/consul/api"

var DEFAULT_CHECK_NAME = "consul00000health"

type CheckConfig struct {
	HTTP      string `json:"http"`
	Interval  string `json:"interval"`
	Timeout   string `json:"timeout"`
	CheckName string `json:"checkName"`
}

type BaseConfig struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Tags    []string `json:"tags"`
	Address string   `json:"address"`
	Port    int      `json:"port"`
	Check   CheckConfig
}

func (conf *BaseConfig) RegisterService() (*consulapi.Client, error) {
	if conf.Check.CheckName == "" {
		conf.Check.CheckName = DEFAULT_CHECK_NAME
	}
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}
	registration := &consulapi.AgentServiceRegistration{
		ID:      conf.ID,
		Name:    conf.Name,
		Tags:    conf.Tags,
		Address: conf.Address,
		Port:    conf.Port,
		Check: &consulapi.AgentServiceCheck{
			HTTP:     conf.Check.HTTP + "/" + conf.Check.CheckName,
			Interval: conf.Check.Interval,
			Timeout:  conf.Check.Timeout,
		},
	}
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		return client, err
	}
	return client, nil

}

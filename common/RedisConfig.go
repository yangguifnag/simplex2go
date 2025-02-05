package common

type RedisConfigModule struct {
	Pass string `json:"pass"`
	Port string `json:"port"`
	Host string `json:"host"`
	DB   int    `json:"db"`
}

package common

type HttpCodeConfig struct {
	Success map[string]int `json:"success"`
	Fail    map[string]int `json:"fail"`
}

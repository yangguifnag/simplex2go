package utils

import common "github.com/yangguifnag/simplex2go/common"

func GetDSN(c common.MysqlConfigModule) string {
	dsn := c.User + ":" + c.Pass + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}

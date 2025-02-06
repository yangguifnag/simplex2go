package utils

import (
	"github.com/mitchellh/mapstructure"
	"github.com/yangguifnag/simplex2go/common"
	"reflect"
	"time"
)

func InitUpdateAuditField[T any](entity T, name string) T {
	typeO := reflect.TypeOf(entity)
	typeStr := typeO.Kind().String()
	if typeStr == "slice" {
		var mp []map[string]interface{}
		mapstructure.WeakDecode(entity, &mp)
		for _, val := range mp {
			val[`UpdateDate`] = time.Now()
			val[`UpdateTime`] = time.Now()
			val[`UpdateBy`] = name
		}
		mapstructure.WeakDecode(mp, &entity)
		return entity
	} else {
		var mp map[string]interface{}
		mapstructure.WeakDecode(entity, &mp)
		mp[`UpdateDate`] = time.Now()
		mp[`UpdateTime`] = time.Now()
		mp[`UpdateBy`] = name
		mapstructure.WeakDecode(mp, &entity)
		return entity
	}
}

func InitUpdateAuditFieldWhitJwt[T any](entity T, jwt common.JWTSession) T {
	return InitUpdateAuditField(entity, jwt.Account)
}

func InitAuditField[T any](entity T, name string) T {
	typeO := reflect.TypeOf(entity)
	typeStr := typeO.Kind().String()
	if typeStr == "slice" {
		var mp []map[string]interface{}
		mapstructure.WeakDecode(entity, &mp)
		for _, val := range mp {
			val[`CreateDate`] = time.Now()
			val[`CreateTime`] = time.Now()
			val[`UpdateDate`] = time.Now()
			val[`UpdateTime`] = time.Now()
			val[`UpdateBy`] = name
			val[`CreateBy`] = name
		}
		mapstructure.WeakDecode(mp, &entity)
		return entity
	} else {
		var mp map[string]interface{}
		mapstructure.WeakDecode(entity, &mp)
		mp[`CreateDate`] = time.Now()
		mp[`CreateTime`] = time.Now()
		mp[`UpdateDate`] = time.Now()
		mp[`UpdateTime`] = time.Now()
		mp[`UpdateBy`] = name
		mp[`CreateBy`] = name
		mapstructure.WeakDecode(mp, &entity)
		return entity
	}
}

func InitAuditFieldWhitJwt[T any](entity T, jwt common.JWTSession) T {
	return InitAuditField(entity, jwt.Account)
}

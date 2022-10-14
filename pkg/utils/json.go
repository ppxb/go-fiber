package utils

import (
	"encoding/json"
	"github.com/ppxb/go-fiber/pkg/log"
)

func Struct2Json(obj interface{}) string {
	str, err := json.Marshal(obj)
	if err != nil {
		log.Error("[struct2json]can not convert: %v", err)
	}
	return string(str)
}

func Json2Struct(str string, obj interface{}) {
	err := json.Unmarshal([]byte(str), obj)
	if err != nil {
		log.Error("[json2struct]can not convert: %v", err)
	}
}

func Struct2StructByJson(struct1 interface{}, struct2 interface{}) {
	jsonStr := Struct2Json(struct1)
	Json2Struct(jsonStr, struct2)
}

func JsonWithSort(str string) string {
	m := make(map[string]interface{})
	Json2Struct(str, &m)
	return Struct2Json(m)
}

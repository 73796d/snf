package message

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"reflect"
)

// json消息转化为字节
func MsgToBuf(msg interface{}) ([]byte, error) {
	b, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// 字节转化为json
func BufToJson(buf []byte) (*simplejson.Json, error) {
	return simplejson.NewJson(buf)
}

// json转化为Struct 通过反射
func JsonToStruct(j *simplejson.Json, i interface{}) {
	dataStruct := reflect.Indirect(reflect.ValueOf(i))
	dataStructType := dataStruct.Type()
	dataFieldNumber := dataStructType.NumField()
	for i := 0; i < dataFieldNumber; i++ {
		field := dataStructType.Field(i)
		fieldV := dataStruct.Field(i)
		if field.Type.Kind() == reflect.Int {
			n, _ := j.Get(field.Name).Int()
			fieldV.SetInt(int64(n))
		} else if field.Type.Kind() == reflect.Int32 {
			n, _ := j.Get(field.Name).Int()
			fieldV.SetInt(int64(n))
		} else if field.Type.Kind() == reflect.Uint32 {
			n, _ := j.Get(field.Name).Int()
			fieldV.SetUint(uint64(n))
		} else if field.Type.Kind() == reflect.Uint64 {
			n, _ := j.Get(field.Name).Int64()
			fieldV.SetUint(uint64(n))
		} else if field.Type.Kind() == reflect.Int64 {
			n, _ := j.Get(field.Name).Int64()
			fieldV.SetInt(n)
		} else if field.Type.Kind() == reflect.String {
			s, _ := j.Get(field.Name).String()
			fieldV.SetString(s)
		} else if field.Type.Kind() == reflect.Bool {
			b, _ := j.Get(field.Name).Bool()
			fieldV.SetBool(b)
		} else if field.Type.Kind() == reflect.Interface {
			JsonToStruct(j.Get(field.Name), fieldV)
		}


	}

}

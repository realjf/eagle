package utils

import (
	. "eagle/service/common"
	"reflect"
)

//把类似slice的map转为slice
func MapToMapEntrySet(input interface{}) []EntrySet {
	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Map {
		return nil
	}
	keys := v.MapKeys()
	output := []EntrySet{}
	for i, l := 0, v.Len(); i < l; i++ {
		tmp := EntrySet{}
		tmp.SetKey(keys[i].Interface())
		tmp.SetValue(v.MapIndex(keys[i]).Interface())
		output = append(output, tmp)
	}
	return output
}

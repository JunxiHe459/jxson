package encode

import (
	"errors"
	"reflect"
	"strconv"
)

func Marshal(v any) ([]byte, error) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.AppendInt(nil, reflect.ValueOf(v).Int(), 10), nil
	case reflect.Struct:
	default:
		return nil, errors.New("unsupported type.\n前面的区域，以后再来探索吧。")
	}
	return nil, nil
}

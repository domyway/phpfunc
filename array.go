package phpfunc

import "reflect"

func ArrayMap(callback interface{}, array interface{}) interface{} {
	f := reflect.ValueOf(callback)
	a := reflect.ValueOf(array)
	in := make([]reflect.Value, reflect.TypeOf(callback).NumIn())

	var newArray interface{}

	switch reflect.TypeOf(array).Kind() {
	case reflect.Array, reflect.Slice:
		tmpArray := make([]interface{}, a.Len())

		for i := 0; i < a.Len(); i++ {
			in[0] = a.Index(i)
			tmpArray[i] = f.Call(in)[0].Interface().(string)
		}

		newArray = tmpArray

	case reflect.Map:
		tmpMap := make(map[string]interface{})
		for _, k := range a.MapKeys() {
			v := a.MapIndex(k)
			in[0] = v
			tmpMap[k.String()] = f.Call(in)[0].Interface()
		}

		newArray = tmpMap
	}

	return newArray
}

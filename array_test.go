package phpfunc

import (
	"fmt"
	"testing"
)

func TestArrayMap(t *testing.T) {
	// array := [...]string{"1", "2", "3"}
	array := make([]string, 3)
	array[0] = "1"
	array[1] = "2"
	array[2] = "3"

	ay := ArrayMap(Hw, array)
	ay2 := ay.([]interface{})
	fmt.Println(Implode("|", ay2))

	array2 := map[string]string{"username": "weizhao", "age": "18"}
	ay3 := ArrayMap(Hw, array2)
	fmt.Println(Implode("|", ay3.(map[string]interface{})))
	// fmt.Println(Implode("|", ay.([]interface{})))
	// fmt.Println(Implode("|", ay2.([]string)))

}

func Hw(str string) string {
	// return str
	return "[" + str + "]"
}

package slice

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSliceToString(t *testing.T) {

	var a = []int{1,2,3,4}
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	var result = string(b)
	fmt.Println(result[])
}


func SliceToString(s []string) string {
	var str string
	if len(s) > 0 {
		for _, v := range s {
			str = fmt.Sprintf("%s,%s", str, v)
		}

		return str[1:]
	}
	return ""
}
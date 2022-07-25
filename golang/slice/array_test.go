package slice

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestSliceToString(t *testing.T) {

	var a = []int{1, 2, 3, 4}
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	var result = string(b)
	fmt.Println(result)
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

func TestSlice(t *testing.T) {
	leader := "10.16.1.11"
	backupEndpointsString := "10.16.1.11,10.16.1.14"

	backupEndpoints := strings.Split(backupEndpointsString, ",")

	for _, endpoint := range backupEndpoints {
		if endpoint != leader {
			fmt.Println(endpoint)
		}
	}
}

func TestGetSlice(t *testing.T) {
	slice := []string{"a", "b", "b"}
	fmt.Println(slice[:3]) //[a b c d]
	fmt.Println(len(slice))

	if len(slice) > 2 {
		fmt.Println(slice[:2])
	}
	fmt.Println(slice)
	slice = slice[:2]
	fmt.Println(slice)
}

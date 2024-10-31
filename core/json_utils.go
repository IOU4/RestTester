package core

import (
	"log"
	"reflect"
)

func IsSameJSON(a, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		log.Println("not same type")
		return false
	}
	switch x := a.(type) {
	case map[string]interface{}:
		y := b.(map[string]interface{})

		if len(x) != len(y) {
			return false
		}

		for k, v := range x {
			val2 := y[k]

			if (v == nil) != (val2 == nil) {
				return false
			}

			if !IsSameJSON(v, val2) {
				return false
			}
		}
		return true
	case []interface{}:
		y := b.([]interface{})

		if len(x) != len(y) {
			return false
		}
		var matches int
		flagged := make([]bool, len(y))
		for _, v := range x {
			for i, v2 := range y {
				if IsSameJSON(v, v2) && !flagged[i] {
					matches++
					flagged[i] = true

					break
				}
			}
		}
		return matches == len(x)
	default:
		return a == b
	}

}

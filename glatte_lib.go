package glatte_lib
import (
	"reflect"
	
)

func IsArray(object interface{}) (bool) {
	return Type(object) == "string"
}
func IsMap(object interface{}) (bool) {
	return Type(object) == "map"
}

func Type(object interface{}) (string) {
	reflectVal := reflect.ValueOf(object)
	switch(reflectVal.Kind()) {
		case reflect.Array , reflect.Slice:
			return "string"
		break;
		case reflect.Map:
			return "map"
		break;
	} 
	return "nil"
}
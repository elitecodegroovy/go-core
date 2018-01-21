package encapsulation

import (
	"fmt"
	"reflect"
)

type Helper struct {}

func (h *Helper) doHelp(who string){
	fmt.Println("Help "+ who)
}

func (h *Helper)DoTask(x interface{}){
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Bool:
		h.doHelp(fmt.Sprintf("%v", v.Bool()))
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		h.doHelp(fmt.Sprintf("%v", v.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		h.doHelp(fmt.Sprintf("%v", v.Uint()))
	case reflect.Float32, reflect.Float64:
		h.doHelp(fmt.Sprintf("%v", v.Float()))
	case reflect.String:
		h.doHelp(fmt.Sprintf("%v", v.String()))
	case reflect.Slice:
		h.doHelp(fmt.Sprintf("len=%d, %v", v.Len(), v.Interface()))
	case reflect.Map:
		h.doHelp(fmt.Sprintf("%v", v.Interface()))
	case reflect.Chan:
		h.doHelp(fmt.Sprintf("%v\n", v.Interface()))
	default:
		h.doHelp(fmt.Sprint(x))
	}
}

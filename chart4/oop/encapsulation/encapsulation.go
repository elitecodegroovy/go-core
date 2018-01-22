package encapsulation

import (
	"fmt"
	"reflect"
	"strconv"
)

type Helper struct {
	Id		int64
	Name 	string
	status  bool
}

func (h *Helper) SetStatus(b bool){
	h.status = b
}

func (h *Helper) GetStatus() bool{
	return h.status
}

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

func (h *Helper)String() string{
	return "Id :"+ strconv.FormatInt(h.Id, 10)+
		", name:"+ h.Name + ", status:"+ strconv.FormatBool(h.status)
}
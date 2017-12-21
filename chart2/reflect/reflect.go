package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"util"
)

func doTypeNValue() {
	s := "reflect"
	f := float32(8.88)
	fmt.Printf("s : %v, %v\n", reflect.TypeOf(s), s)
	fmt.Printf("f : %v, %v\n", reflect.TypeOf(f), f)
	fmt.Println()

	v1 := reflect.ValueOf(s)
	v2 := reflect.ValueOf(f)
	fmt.Printf("s value %v\n", v1)
	fmt.Printf("f value %v\n", v2)
}

type User struct {
	UserName string `tag_name:"tag 1"`
	NickName string `tag_name:"tag 2"`
	Age      int    ` tag_name:"tag 3"`
}

func (u *User) printFields() {
	val := reflect.ValueOf(u).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n",
			typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
}
func doReflectStructFields() {
	u := &User{
		UserName: "梦放飞",
		NickName: "梦子",
		Age:      30,
	}
	u.printFields()
}

func getAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
		// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

//print value of the reflect value
func printValue(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			printValue(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			printValue(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			printValue(fmt.Sprintf("%s[%s]", path,
				getAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			printValue(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			printValue(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, getAtom(v))
	}
}

func testPrintValue() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Musics          []string
	}
	//!-movie
	goodMovie := Movie{
		Title:    "芳华",
		Subtitle: "The Youth",
		Color:    true,
		Year:     2017,
		Actor: map[string]string{
			"刘峰":  "黄轩",
			"何小萍": "苗苗",
			"萧穗子": "钟楚曦",
			"林丁丁": "杨采钰",
			"郝淑雯": "李晓峰",
			`陈灿`:  "王天辰",
		},

		Musics: []string{
			"那些花儿",
			"想把你留在这里",
			"美好生活",
			"绒花",
		},
	}

	printValue("goodMovie", reflect.ValueOf(goodMovie))

	changeValue()
}

func changeValue() {
	a := "C programming"
	d := reflect.ValueOf(&a).Elem()      //变量d，拥有地址
	pa := d.Addr().Interface().(*string) //获取string指针
	*pa = "Go Programming"               //赋值给指针的值
	fmt.Println(a)

	isAddr()
}

func isAddr() {
	x := "immutable"                  // 不是常量
	a := reflect.ValueOf("immutable") // 不是常量
	b := reflect.ValueOf(x)           // 不是常量
	c := reflect.ValueOf(&x)          // 不是常量
	d := c.Elem()                     // 是常量

	fmt.Println("a是常量：", a.CanAddr()) // "false"
	fmt.Println("b是常量：", b.CanAddr()) // "false"
	fmt.Println("c是常量：", c.CanAddr()) // "false"
	fmt.Println("d是常量：", d.CanAddr()) // "true"
}

func CallFunc(m interface{}, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m)
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of parameters is not right, please check again!")
		return
	}
	ps := make([]reflect.Value, len(params))
	for i, param := range params {
		ps[i] = reflect.ValueOf(param)
	}
	result = f.Call(ps)
	return
}

func callReflectFunc() {
	s := "I LOVE YOU"
	fmt.Println("swap case s:", util.SwapCase(s))

	if result, err := CallFunc(util.SwapCase, s); err != nil {
		fmt.Errorf("error %s", err.Error())
	} else {
		fmt.Printf("call reflect func, result :%s \n", result[0].String())
	}
}

type Test struct{}

func (t *Test) PrintInfo(i int, s string) string {
	fmt.Println("call method PrintInfo i", i, ",s :", s)
	return s + strconv.Itoa(i)
}

func (t *Test) ShowMsg() string {
	fmt.Println("\nshow msg input 'call reflect'")
	return "ShowMsg"
}

func callReflect(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	if v := reflect.ValueOf(any).MethodByName(name); v.String() == "<invalid Value>" {
		return nil
	} else {
		return v.Call(inputs)
	}

}

func callReflectMethod() {
	fmt.Printf("\n callReflectMethod PrintInfo :%s", callReflect(&Test{}, "PrintInfo", 10, "TestMethod")[0].String())
	fmt.Printf("\n callReflectMethod ShowMsg  %s", callReflect(&Test{}, "ShowMsg")[0].String())

	//<invalid Value> case
	if result := callReflect(&Test{}, "ShowMs"); result != nil {
		fmt.Printf("\n callReflectMethod ShowMs %s", result[0].String())
	} else {
		fmt.Println("\n callReflectMethod ShowMs didn't run ")
	}
	fmt.Println("\n reflect all ")
}

func main() {
	doTypeNValue()

	//print the struct field
	doReflectStructFields()

	//print the struct tree
	testPrintValue()

	callReflectFunc()

	callReflectMethod()

	//附近内容，见目录mapstructure
}

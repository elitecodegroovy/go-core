
package json2struct

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"gopkg.in/yaml.v2"
)

var ForceFloats bool

// 常用的缩略词集合
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"NTP":   true,
	"DB":    true,
}

// 数值的字母集合
var intToWordMap = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

type Parser func(io.Reader) (interface{}, error)

//解析JSON数据，以interface{}类型格式输出
func ParseJson(input io.Reader) (interface{}, error) {
	var result interface{}
	if err := json.NewDecoder(input).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
//解析YAML数据，以interface{}类型格式输出
func ParseYaml(input io.Reader) (interface{}, error) {
	var result interface{}
	b, err := readFile(input)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}

//读取文件input的内容，以[]byte输出
func readFile(input io.Reader) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, input)
	if err != nil {
		return []byte{}, nil
	}
	return buf.Bytes(), nil
}

//产生一个结构定义：
//实现JSON数据转化为结构定义
func GenerateJSON(input io.Reader, parser Parser, structName, pkgName string, tags []string, subStruct bool, convertFloats bool) ([]byte, error) {
	var subStructMap map[string]string
	if subStruct {
		subStructMap = make(map[string]string)
	}

	var result map[string]interface{}

	iresult, err := parser(input)
	if err != nil {
		return nil, err
	}
	//json数据映射为Go结构有三种数据类型：
	// key为interface{}类型、key为字符串类型、key为泛化的数组类型[]interface{}
	//处理步骤：
	//步骤一，处理数组类型[]interface{}
	//步骤二，统一处理key为interface{}或者字符串类型的情况
	switch iresult := iresult.(type) {
	case []interface{}:
		//步骤一,处理数组类型[]interface{}类型的数据
		src := fmt.Sprintf("package %s\n\ntype %s %s\n",
			pkgName,
			structName,
			typeForValue(iresult, structName, tags, subStructMap, convertFloats))
		formatted, err := format.Source([]byte(src))
		if err != nil {
			err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, src)
		}
		return formatted, err
	case map[interface{}]interface{}:
		//转化过程：统一key为strin值为interface{}的map类型
		result = convertKeysToStrings(iresult)
	case map[string]interface{}:
		result = iresult
	default:
		return nil, fmt.Errorf("unexpected type: %T", iresult)
	}
	///步骤二,统一处理的数据类型类型--key为strin值为interface{}的map类型
	src := fmt.Sprintf("package %s\ntype %s %s}",
		pkgName,
		structName,
		generateTypes(result, structName, tags, 0, subStructMap, convertFloats))

	keys := make([]string, 0, len(subStructMap))
	for key := range subStructMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, k := range keys {
		src = fmt.Sprintf("%v\n\ntype %v %v", src, subStructMap[k], k)
	}

	formatted, err := format.Source([]byte(src))
	if err != nil {
		err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, src)
	}
	return formatted, err
}

func convertKeysToStrings(obj map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})

	for k, v := range obj {
		res[fmt.Sprintf("%v", k)] = v
	}

	return res
}

// map[string]interface{} 类型的数据生成对应的结构
func generateTypes(obj map[string]interface{}, structName string, tags []string, depth int,
					subStructMap map[string]string, convertFloats bool) string {
	structure := "struct {"

	keys := make([]string, 0, len(obj))
	for key := range obj {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	//构建字段序列
	for _, key := range keys {
		value := obj[key]
		valueType := typeForValue(value, structName, tags, subStructMap, convertFloats)
		//value = mergeElements(value)
		//If a nested value, recurse
		switch value := value.(type) {
		case []interface{}:
			if len(value) > 0 {
				sub := ""
				if v, ok := value[0].(map[interface{}]interface{}); ok {
					sub = generateTypes(convertKeysToStrings(v), structName, tags, depth+1, subStructMap, convertFloats) + "}"
				} else if v, ok := value[0].(map[string]interface{}); ok {
					sub = generateTypes(v, structName, tags, depth+1, subStructMap, convertFloats) + "}"
				}

				if sub != "" {
					subName := sub

					if subStructMap != nil {
						if val, ok := subStructMap[sub]; ok {
							subName = val
						} else {
							subName = fmt.Sprintf("%v_sub%v", structName, len(subStructMap)+1)

							subStructMap[sub] = subName
						}
					}

					valueType = "[]" + subName
				}
			}
		case map[interface{}]interface{}:
			sub := generateTypes(convertKeysToStrings(value), structName, tags, depth+1, subStructMap, convertFloats) + "}"
			subName := sub

			if subStructMap != nil {
				if val, ok := subStructMap[sub]; ok {
					subName = val
				} else {
					subName = fmt.Sprintf("%v_sub%v", structName, len(subStructMap)+1)

					subStructMap[sub] = subName
				}
			}
			valueType = subName
		case map[string]interface{}:
			sub := generateTypes(value, structName, tags, depth+1, subStructMap, convertFloats) + "}"
			subName := sub

			if subStructMap != nil {
				if val, ok := subStructMap[sub]; ok {
					subName = val
				} else {
					subName = fmt.Sprintf("%v_sub%v", structName, len(subStructMap)+1)

					subStructMap[sub] = subName
				}
			}

			valueType = subName
		}
		//格式化字段
		fieldName := FmtFieldName(key)

		tagList := make([]string, 0)
		for _, t := range tags {
			tagList = append(tagList, fmt.Sprintf("%s:\"%s\"", t, key))
		}

		structure += fmt.Sprintf("\n%s %s `%s`",
			fieldName,
			valueType,
			strings.Join(tagList, " "))
	}
	return structure
}

// 格式化结构字段的呈现形式。
// e.g.
// 	FmtFieldName("goo_id")
//  Output: GooId
func FmtFieldName(s string) string {
	runes := []rune(s)
	//case 1: 如果字段开头字符不是数字和字母的话，去除掉。
	for len(runes) > 0 && !unicode.IsLetter(runes[0]) && !unicode.IsDigit(runes[0]) {
		runes = runes[1:]
	}
	//case 1: 如果字段开头字符为空字符（没有任何字符），以“_”代替。
	if len(runes) == 0 {
		return "_"
	}

	s = stringifyFirstChar(string(runes))
	//首字母大写算法
	name := lintFieldName(s)
	runes = []rune(name)
	for i, c := range runes {
		ok := unicode.IsLetter(c) || unicode.IsDigit(c)
		if i == 0 {
			ok = unicode.IsLetter(c)
		}
		if !ok {
			runes[i] = '_'
		}
	}
	s = string(runes)
	s = strings.Trim(s, "_")
	if len(s) == 0 {
		return "_"
	}
	return s
}

//首字母大写算法，包含一下规则：
//	1. 保留特殊缩略词
//  2. 去掉下划线，并且排除下划线在两个数字之间的情况
func lintFieldName(name string) string {
	if name == "_" {
		return name
	}

	allLower := true
	for _, r := range name {
		if !unicode.IsLower(r) {
			allLower = false
			break
		}
	}
	//所有字符都是小写字母
	if allLower {
		runes := []rune(name)
		if u := strings.ToUpper(name); commonInitialisms[u] {
			copy(runes[0:], []rune(u))
		} else {
			runes[0] = unicode.ToUpper(runes[0])
		}
		return string(runes)
	}
	//所有字符都是大写字母，并且没有下划线
	allUpperWithUnderscore := true
	for _, r := range name {
		if !unicode.IsUpper(r) && r != '_' {
			allUpperWithUnderscore = false
			break
		}
	}
	if allUpperWithUnderscore {
		name = strings.ToLower(name)
	}
	// 驼峰命名转化，并且分离下划线
	// Check each word for common initialisms.
	runes := []rune(name)
	w, i := 0, 0
	for i+1 <= len(runes) {
		eow := false 															// 是否到达单词结尾

		if i+1 == len(runes) {
			eow = true
		} else if runes[i+1] == '_' {
			// underscore; shift the remainder forward over any run of underscores
			eow = true
			n := 1
			for i+n+1 < len(runes) && runes[i+n+1] == '_' {
				n++
			}

			// 排除下划线在两个数字之间的情况
			if i+n+1 < len(runes) && unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+n+1]) {
				n--
			}

			copy(runes[i+1:], runes[i+n+1:])
			runes = runes[:len(runes)-n]
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i) is a word.
		word := string(runes[w:i])
		if u := strings.ToUpper(word); commonInitialisms[u] {                 //一般首字母缩略词处理
			copy(runes[w:], []rune(u))
		} else if strings.ToLower(word) == word {                            //已经都是小写字母，将首字母大写
			runes[w] = unicode.ToUpper(runes[w])
		}
		w = i
	}
	return string(runes)
}

// 数值的数据类型
func typeForValue(value interface{}, structName string, tags []string, subStructMap map[string]string, convertFloats bool) string {
	//JSON数组类型，对应Go类是[]interface{}
	if objects, ok := value.([]interface{}); ok {
		types := make(map[reflect.Type]bool, 0)
		for _, o := range objects {
			types[reflect.TypeOf(o)] = true
		}
		if len(types) == 1 {
			return "[]" + typeForValue(mergeElements(objects).([]interface{})[0], structName, tags, subStructMap, convertFloats)
		}
		return "[]interface{}"
	} else if object, ok := value.(map[interface{}]interface{}); ok {
		return generateTypes(convertKeysToStrings(object), structName, tags, 0, subStructMap, convertFloats) + "}"
	} else if object, ok := value.(map[string]interface{}); ok {
		return generateTypes(object, structName, tags, 0, subStructMap, convertFloats) + "}"
	} else if reflect.TypeOf(value) == nil {
		return "interface{}"
	}
	v := reflect.TypeOf(value).Name()
	if v == "float64" && convertFloats {
		v = disambiguateFloatInt(value)
	}
	return v
}

// 以float64读取所有的数值，利用了Go的反射机制
func disambiguateFloatInt(value interface{}) string {
	const epsilon = .0001
	vfloat := value.(float64)
	if !ForceFloats && math.Abs(vfloat-math.Floor(vfloat+epsilon)) < epsilon {
		var tmp int64
		return reflect.TypeOf(tmp).Name()
	}
	return reflect.TypeOf(value).Name()
}

// 转化第一个字符int到字符串
func stringifyFirstChar(str string) string {
	first := str[:1]

	i, err := strconv.ParseInt(first, 10, 8)

	if err != nil {
		return str
	}

	return intToWordMap[i] + "_" + str[1:]
}

func mergeElements(i interface{}) interface{} {
	switch i := i.(type) {
	default:
		return i
	case []interface{}:
		l := len(i)
		if l == 0 {
			return i
		}
		for j := 1; j < l; j++ {
			i[0] = mergeObjects(i[0], i[j])
		}
		return i[0:1]
	}
}

func mergeObjects(o1, o2 interface{}) interface{} {
	if o1 == nil {
		return o2
	}

	if o2 == nil {
		return o1
	}

	if reflect.TypeOf(o1) != reflect.TypeOf(o2) {
		return nil
	}

	switch i := o1.(type) {
	default:
		return o1
	case []interface{}:
		if i2, ok := o2.([]interface{}); ok {
			i3 := append(i, i2...)
			return mergeElements(i3)
		}
		return mergeElements(i)
	case map[string]interface{}:
		if i2, ok := o2.(map[string]interface{}); ok {
			for k, v := range i2 {
				if v2, ok := i[k]; ok {
					i[k] = mergeObjects(v2, v)
				} else {
					i[k] = v
				}
			}
		}
		return i
	case map[interface{}]interface{}:
		if i2, ok := o2.(map[interface{}]interface{}); ok {
			for k, v := range i2 {
				if v2, ok := i[k]; ok {
					i[k] = mergeObjects(v2, v)
				} else {
					i[k] = v
				}
			}
		}
		return i
	}
}

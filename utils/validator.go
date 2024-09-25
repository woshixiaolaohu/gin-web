package utils

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Rules map[string][]string => key string {string}
type Rules map[string][]string

type RulesMap map[string]Rules

var CustomizeMap = make(map[string]Rules)

// RegisterRule
//
//	@function:		RegisterRule
//	@description:	注册自定义规则方案建议在路由初始化层即注册
//	@param:			key string, rule Rules
//	@return:		err error
func RegisterRule(key string, rule Rules) (err error) {
	if CustomizeMap[key] != nil {
		return errors.New(key + "已注册，无法重复注册")
	} else {
		CustomizeMap[key] = rule
		return nil
	}
}

// NotEmpty
//
//	@function:		NotEmpty
//	@description:	非空 不能为其对应类型的0值
//	@return:		string
func NotEmpty() string {
	return "notEmpty"
}

// RegexpMatch
//
//	@function:		RegexpMatch
//	@description:	正则校验 校验输入项是否满足正则表达式
//	@param:			rule string
//	@return:		string
func RegexpMatch(rule string) string {
	return "regexp=" + rule
}

// Lt
//
//	@function:		Lt
//	@description:	小于入参(<) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//	@param:			mark string
//	@return:		string
func Lt(mark string) string {
	return "lt=" + mark
}

// Le
//
//	@function:		Le
//	@description:	小于等于入参(<=) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//	@param:			mark string
//	@return:		string
func Le(mark string) string {
	return "le=" + mark
}

// Eq
//
//	@function:		Eq
//	@description:	等于入参(==) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//	@param:			mark string
//	@return:		string
func Eq(mark string) string {
	return "eq=" + mark
}

// Ne
//
//	@function:		Ne
//	@description:	不等于入参(!=)  如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//	@param:			mark string
//	@return:		string
func Ne(mark string) string {
	return "ne=" + mark
}

// Ge
//
//	@function:		Ge
//	@description:	大于等于入参(>=) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//	@param:			mark string
//	@return:		string
func Ge(mark string) string {
	return "ge=" + mark
}

// Gt
//
//	@function:		Gt
//	@description:	大于入参(>) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//	@param:			mark string
//	@return:		string
func Gt(mark string) string {
	return "gt=" + mark
}

// Verify
//
//	@function:		Verify
//	@description:	校验方法
//	@param:			st interface{}, roleMap Rules(入参实例，规则map)
//	@return:		err error
func Verify(st interface{}, roleMap Rules) (err error) {
	compareMap := map[string]bool{
		"lt": true,
		"le": true,
		"eq": true,
		"ne": true,
		"ge": true,
		"gt": true,
	}
	// 获取 reflect.Type 类型
	typ := reflect.TypeOf(st)
	// 获取 reflect.Value 类型
	val := reflect.ValueOf(st)
	// 获取 st 对应类别
	kd := val.Kind()
	if kd != reflect.Struct {
		return errors.New("expect struct")
	}
	// 返回 val 中的字段数
	num := val.NumField()
	// 遍历结构体所有字段
	for i := 0; i < num; i++ {
		// 返回 struct 的第 i 个字段
		tagVal := typ.Field(i)
		val := val.Field(i)
		if tagVal.Type.Kind() == reflect.Struct {
			if err = Verify(val.Interface(), roleMap); err != nil {
				return err
			}
		}
		if len(roleMap[tagVal.Name]) > 0 {
			for _, v := range roleMap[tagVal.Name] {
				switch {
				case v == "notEmpty":
					if isBlank(val) {
						return errors.New(tagVal.Name + "值不能为空")
					}
				case strings.Split(v, "=")[0] == "regexp":
					if !regexpMatch(strings.Split(v, "=")[1], val.String()) {
						return errors.New(tagVal.Name + "格式校验不通过")
					}
				case compareMap[strings.Split(v, "=")[0]]:
					if !compareVerify(val, v) {
						return errors.New(tagVal.Name + "长度或值不在合法范围内," + v)
					}
				}
			}
		}
	}
	return nil
}

// compareVerify
//
//	@function:		compareVerify
//	@description:	长度和数字的校验方法 根据类型自动校验
//	@param:			value reflect.Value, VerifyStr string
//	@return:		bool
func compareVerify(value reflect.Value, VerifyStr string) bool {
	switch value.Kind() {
	case reflect.String:
		// 类型别名
		// rune 类型表示中文符号
		// rune 4字节 32比特=> uint32=>Unicode      byte 1字节 8比特 => uint8=>ASCII
		return compare(len([]rune(value.String())), VerifyStr)
	case reflect.Slice, reflect.Array:
		return compare(value.Len(), VerifyStr)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), VerifyStr)
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), VerifyStr)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), VerifyStr)
	default:
		return false
	}
}

// isBlank
//
//	@function:		isBlank
//	@description:	非空校验
//	@param:			value reflect.Value
//	@return:		bool
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Int() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()

	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// compare
//
//	@function:		compare
//	@description:	比较函数
//	@param:			value interface{}, VerifyStr string
//	@return:		bool
func compare(value interface{}, VerifyStr string) bool {
	VerifyStrArr := strings.Split(VerifyStr, "=")
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		VInt, VErr := strconv.ParseInt(VerifyStrArr[1], 10, 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Int() < VInt
		case VerifyStrArr[0] == "le":
			return val.Int() <= VInt
		case VerifyStrArr[0] == "eq":
			return val.Int() == VInt
		case VerifyStrArr[0] == "ne":
			return val.Int() != VInt
		case VerifyStrArr[0] == "ge":
			return val.Int() >= VInt
		case VerifyStrArr[0] == "gt":
			return val.Int() > VInt
		default:
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		VInt, VErr := strconv.Atoi(VerifyStrArr[1])
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Uint() < uint64(VInt)
		case VerifyStrArr[0] == "le":
			return val.Uint() <= uint64(VInt)
		case VerifyStrArr[0] == "eq":
			return val.Uint() == uint64(VInt)
		case VerifyStrArr[0] == "ne":
			return val.Uint() != uint64(VInt)
		case VerifyStrArr[0] == "ge":
			return val.Uint() >= uint64(VInt)
		case VerifyStrArr[0] == "gt":
			return val.Uint() > uint64(VInt)
		default:
			return false
		}
	case reflect.Float32, reflect.Float64:
		VFloat, VErr := strconv.ParseFloat(VerifyStrArr[1], 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Float() < VFloat
		case VerifyStrArr[0] == "le":
			return val.Float() <= VFloat
		case VerifyStrArr[0] == "eq":
			return val.Float() == VFloat
		case VerifyStrArr[0] == "ne":
			return val.Float() != VFloat
		case VerifyStrArr[0] == "ge":
			return val.Float() >= VFloat
		case VerifyStrArr[0] == "gt":
			return val.Float() > VFloat
		default:
			return false
		}
	default:
		return false
	}
}

func regexpMatch(rule string, matchStr string) bool {
	return regexp.MustCompile(rule).MatchString(matchStr)
}

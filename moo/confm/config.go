package conf

import (
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
	"strings"
)

//类型定义
//配置
type config struct {
	//选项
	options
	//配置文件内容解码缓存
	contents
}

//配置内容
type contents = map[string]map[string]interface{}

//配置选项
type options struct {
	path string
	ext  string
	sep  string
}

//变量
//config 单例对象
var instance *config

//默认选项值
var defaultOptions = options{
	path: "./configs/",
	ext:  ".yml",
	sep:  ".",
}

//常量

//可导出包方法

// Instance 配置对象单例工厂
func Instance(options ...string) *config {
	if instance == nil {
		//创建实例
		instance = &config{
			contents: contents{},
			options: defaultOptions,
		}
		//配置选项
		for i, v := range options {
			switch i {
			case 0:
				instance.options.path = v
			case 1:
				instance.options.ext = v
			case 2:
				instance.options.sep = v
			}
		}
	}
	return instance
}

// Bool Getter|Setter
func (config) Bool(key string) bool {
	valueIf, err := instance.value(key)
	if err != nil {
		return false
	}
	value, ok := valueIf.(bool)
	if !ok {
		return false
	}
	return value
}

func (config) String(key string) string {
	valueIf, err := instance.value(key)
	if err != nil {
		return ""
	}
	return assertString(valueIf)
}

func (config) Int(key string) int {
	valueIf, err := instance.value(key)
	if err != nil {
		log.Println(err)
		return 0
	}
	value, ok := valueIf.(int)
	if !ok {
		return 0
	}
	return value
}

func (config) Float64(key string) float64 {
	valueIf, err := instance.value(key)
	if err != nil {
		return 0
	}
	return assertFloat64(valueIf)
}

func (config) Float32(key string) float32 {
	valueIf, err := instance.value(key)
	if err != nil {
		return 0
	}
	return assertFloat32(valueIf)
}

func (config) BoolSlice(key string) []bool {
	valueIf, err := instance.value(key)
	if err != nil {
		return []bool{}
	}
	//
	sli, ok := valueIf.([]interface{})
	if !ok {
		return []bool{}
	}
	value := make([]bool, len(sli))
	for i, vi := range sli {
		v, ok := vi.(bool)
		if ok {
			value[i] = v
		}
	}
	return value
}

func (config) IntSlice(key string) []int {
	valueIf, err := instance.value(key)
	if err != nil {
		return []int{}
	}
	//
	sli, ok := valueIf.([]interface{})
	if !ok {
		return []int{}
	}
	value := make([]int, len(sli))
	for i, vi := range sli {
		v, ok := vi.(int)
		if ok {
			value[i] = v
		}
	}
	return value
}

func (config) StringSlice(key string) []string {
	valueIf, err := instance.value(key)
	if err != nil {
		return []string{}
	}
	//
	sli, ok := valueIf.([]interface{})
	if !ok {
		return []string{}
	}
	value := make([]string, len(sli))
	for i, vi := range sli {
		value[i] = assertString(vi)
	}
	return value
}

func (config) Float64Slice(key string) []float64 {
	valueIf, err := instance.value(key)
	if err != nil {
		return []float64{}
	}
	//
	sli, ok := valueIf.([]interface{})
	if !ok {
		return []float64{}
	}
	value := make([]float64, len(sli))
	for i, vi := range sli {
		value[i] = assertFloat64(vi)
	}
	return value
}

func (config) Float32Slice(key string) []float32 {
	valueIf, err := instance.value(key)
	if err != nil {
		return []float32{}
	}
	//
	sli, ok := valueIf.([]interface{})
	if !ok {
		return []float32{}
	}

	value := make([]float32, len(sli))
	for i, vi := range sli {
		value[i] = assertFloat32(vi)
	}
	return value
}

func (config) BoolMap(key string) map[string]bool {
	valueIf, err := instance.value(key)
	if err != nil {
		return map[string]bool{}
	}
	//
	mp, ok := valueIf.(map[string]interface{})
	if !ok {
		return map[string]bool{}
	}
	value := map[string]bool{}
	for k, vk := range mp {
		value[k] = false
		v, ok := vk.(bool)
		if ok {
			value[k] = v
		}
	}
	return value
}

func (config) IntMap(key string) map[string]int {
	valueIf, err := instance.value(key)
	if err != nil {
		return map[string]int{}
	}
	//
	mp, ok := valueIf.(map[string]interface{})
	if !ok {
		return map[string]int{}
	}
	value := map[string]int{}
	for k, vk := range mp {
		value[k] = 0
		v, ok := vk.(int)
		if ok {
			value[k] = v
		}
	}
	return value
}

func (config) Float64Map(key string) map[string]float64 {
	valueIf, err := instance.value(key)
	if err != nil {
		return map[string]float64{}
	}
	//
	mp, ok := valueIf.(map[string]interface{})
	if !ok {
		return map[string]float64{}
	}

	value := map[string]float64{}
	for k, vk := range mp {
		value[k] = assertFloat64(vk)
	}
	return value
}

func (config) Float32Map(key string) map[string]float32 {
	valueIf, err := instance.value(key)
	if err != nil {
		return map[string]float32{}
	}
	//
	mp, ok := valueIf.(map[string]interface{})
	if !ok {
		return map[string]float32{}
	}

	value := map[string]float32{}
	for k, vk := range mp {
		value[k] = assertFloat32(vk)
	}
	return value
}

func (config) StringMap(key string) map[string]string {
	valueIf, err := instance.value(key)
	if err != nil {
		return map[string]string{}
	}
	//
	mp, ok := valueIf.(map[string]interface{})
	if !ok {
		return map[string]string{}
	}
	value := map[string]string{}
	for k, vk := range mp {
		value[k] = assertString(vk)
	}
	return value
}

//非导出方法

//利用 key 获取 值
func (config) value(key string) (interface{}, error) {
	filename, keys := instance.parseKey(key)
	if data, exists := instance.contents[filename]; !exists {
		var (
			content = []byte{}
			err     error
		)
		if content, err = instance.getContent(filename); err != nil {
			log.Println(err)
			return nil, err
		}
		if err = yaml.Unmarshal(content, &data); err != nil {
			log.Println(err)
			return nil, err
		}
		instance.contents[filename] = data
	}

	//解析
	for i, l, currLevel := 0, len(keys), instance.contents[filename]; i < l; i++ {
		//最后一级 key
		if i == l-1 {
			if value, exists := currLevel[keys[i]]; exists {
				return value, nil
			} else {
				return nil, errors.New("key is not exists")
			}
		}

		//不是最后一级，继续解析
		var exists bool
		currLevel, exists = currLevel[keys[i]].(map[string]interface{})
		if !exists {
			return nil, errors.New("key is not exists")
		}
	}
	return nil, nil
}

//解析 key
func (config) parseKey(key string) (string, []string) {
	strs := strings.Split(key, instance.options.sep)
	if len(strs) == 1 {
		strs = append([]string{"app"}, strs...)
	}
	return strs[0], strs[1:]
}

//读取配置文件
func (config) getContent(filename string) ([]byte, error) {
	return os.ReadFile(instance.options.path + filename + instance.options.ext)
}

//断言 Interface{} to string
func assertString(valueIf interface{}) (value string) {
	switch v := valueIf.(type) {
	case string:
		value = v
	case int:
		value = strconv.FormatInt(int64(v), 10) //string(v)
	case float64:
		value = strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		value = strconv.FormatFloat(float64(v), 'f', -1, 64)
	case bool:
		value = strconv.FormatBool(v)
	}
	return
}

//断言 Interface{} to float32
func assertFloat32(valueIf interface{}) (value float32) {
	switch v := valueIf.(type) {
	case float64:
		value = float32(v)
	case float32:
		value = v
	case int:
		value = float32(v)
	}
	return
}

//断言 Interface{} to float64
func assertFloat64(valueIf interface{}) (value float64) {
	switch v := valueIf.(type) {
	case float64:
		value = v
	case float32:
		value = float64(v)
	case int:
		value = float64(v)
	}
	return
}

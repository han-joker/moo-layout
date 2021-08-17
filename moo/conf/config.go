package conf

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
)

// 配置数据
type config struct {
	contents map[string]map[string]interface{}
}

// config 单例对象
var instance *config

var (
	path = "./configs/" // 配置文件目录
	ext  = ".json"
	sep  = "."
)

// 配置对象单例工厂
func Instance(options ...string) *config {
	if instance == nil {
		// 创建实例
		instance = &config{
			contents: map[string]map[string]interface{}{},
		}
		// 配置选项
		for i, v := range options {
			switch i {
			case 0:
				path = v
			case 1:
				ext = v
			case 2:
				sep = v
			}
		}
	}
	return instance
}

// Getter|Setter
func (c config) Bool(key string) bool {
	valueIf, err := c.value(key)
	if err != nil {
		return false
	}
	value, ok := valueIf.(bool)
	if !ok {
		return false
	}
	return value
}

func (c config) String(key string) string {
	valueIf, err := c.value(key)
	if err != nil {
		return ""
	}
	value, ok := valueIf.(string)
	if !ok {
		return ""
	}
	return value
}

func (c config) Int(key string) int {
	valueIf, err := c.value(key)
	if err != nil {
		return 0
	}
	value, ok := valueIf.(float64)
	if !ok {
		return 0
	}
	return int(value)
}

func (c config) Float64(key string) float64 {
	valueIf, err := c.value(key)
	if err != nil {
		return 0
	}
	value, ok := valueIf.(float64)
	if !ok {
		return 0
	}
	return value
}

func (c config) Float32(key string) float32 {
	valueIf, err := c.value(key)
	if err != nil {
		return 0
	}
	value, ok := valueIf.(float64)
	if !ok {
		return 0
	}
	return float32(value)
}

func (c config) BoolSlice(key string) []bool {
	valueIf, err := c.value(key)
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

func (c config) IntSlice(key string) []int {
	valueIf, err := c.value(key)
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
		v, ok := vi.(float64)
		if ok {
			value[i] = int(v)
		}
	}
	return value
}

func (c config) StringSlice(key string) []string {
	valueIf, err := c.value(key)
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
		v, ok := vi.(string)
		if ok {
			value[i] = v
		}
	}
	return value
}

func (c config) Float64Slice(key string) []float64 {
	valueIf, err := c.value(key)
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
		v, ok := vi.(float64)
		if ok {
			value[i] = v
		}
	}
	return value
}

func (c config) Float32Slice(key string) []float32 {
	valueIf, err := c.value(key)
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
		v, ok := vi.(float64)
		if ok {
			value[i] = float32(v)
		}
	}
	return value
}

func (c config) BoolMap(key string) map[string]bool {
	valueIf, err := c.value(key)
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
		v, ok := vk.(bool)
		if ok {
			value[k] = v
		}
	}
	return value
}

func (c config) IntMap(key string) map[string]int {
	valueIf, err := c.value(key)
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
		v, ok := vk.(float64)
		if ok {
			value[k] = int(v)
		}
	}
	return value
}

func (c config) Float64Map(key string) map[string]float64 {
	valueIf, err := c.value(key)
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
		v, ok := vk.(float64)
		if ok {
			value[k] = v
		}
	}
	return value
}

func (c config) Float32Map(key string) map[string]float32 {
	valueIf, err := c.value(key)
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
		v, ok := vk.(float64)
		if ok {
			value[k] = float32(v)
		}
	}
	return value
}

func (c config) StringMap(key string) map[string]string {
	valueIf, err := c.value(key)
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
		v, ok := vk.(string)
		if ok {
			value[k] = v
		}
	}
	return value
}

// 利用 key 获取 值
func (c *config) value(key string) (interface{}, error) {
	filename, keys := c.parseKey(key)
	if data, exists := instance.contents[filename]; !exists {
		var (
			content = []byte{}
			err     error
		)
		if content, err = c.getContent(filename); err != nil {
			log.Println(err)
			return nil, err
		}
		if err = json.Unmarshal(content, &data); err != nil {
			log.Println(err)
			return nil, err
		}
		c.contents[filename] = data
	}

	// 解析
	for i, l, currLevel := 0, len(keys), instance.contents[filename]; i < l; i++ {
		// 最后一级 key
		if i == l-1 {
			if value, exists := currLevel[keys[i]]; exists {
				return value, nil
			} else {
				return nil, errors.New("key is not exists")
			}
		}

		// 不是最后一级，继续解析
		var exists bool
		currLevel, exists = currLevel[keys[i]].(map[string]interface{})
		if !exists {
			return nil, errors.New("key is not exists")
		}
	}
	return nil, nil
}

// 解析 key
func (config) parseKey(key string) (string, []string) {
	strs := strings.Split(key, sep)
	if len(strs) == 1 {
		strs = append([]string{"app"}, strs...)
	}
	return strs[0], strs[1:]
}

// 读取配置文件
func (config) getContent(filename string) ([]byte, error) {
	return os.ReadFile(path + filename + ext)
}

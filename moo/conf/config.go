package conf

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
)

type config struct {
	contents map[string]map[string]interface{}
}

// config 单例对象
var instance *config

const (
	path = "./configs/" // 配置文件目录
	ext  = ".json"
	sep  = "."
)

func Instance() *config {
	if instance == nil {
		instance = &config{
			contents: map[string]map[string]interface{}{},
		}
	}
	return instance
}

// Getter|Setter
func (c *config) Bool(key string) bool {
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

func (c *config) String(key string) string {
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

func (c *config) Int(key string) int {
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

func (c *config) Float64(key string) float64 {
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
func (c *config) Float32(key string) float32 {
	valueIf, err := c.value(key)
	if err != nil {
		return 0
	}
	value, ok := valueIf.(float32)
	if !ok {
		return 0
	}
	return value
}

func (c *config) BoolSlice(key string) []bool {
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
	for i, l := 0, len(sli); i < l; i++ {
		v, ok := sli[i].(bool)
		if ok {
			value[i] = v
		}
	}
	return value
}

func (c *config) IntSlice(key string) []int {
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
	for i, l := 0, len(sli); i < l; i++ {
		v, ok := sli[i].(float64)
		if ok {
			value[i] = int(v)
		}
	}
	return value
}

func (c *config) StringSlice(key string) []string {
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
	for i, l := 0, len(sli); i < l; i++ {
		v, ok := sli[i].(string)
		if ok {
			value[i] = v
		}
	}
	return value
}

func (c *config) Float64Slice(key string) []float64 {
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
	for i, l := 0, len(sli); i < l; i++ {
		v, ok := sli[i].(float64)
		if ok {
			value[i] = v
		}
	}
	return value
}

func (c *config) Float32Slice(key string) []float32 {
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
	for i, l := 0, len(sli); i < l; i++ {
		v, ok := sli[i].(float32)
		if ok {
			value[i] = v
		}
	}
	return value
}

// 利用 key 获取 值
func (c *config) value(key string) (interface{}, error) {
	filename, keys := parseKey(key)
	if data, exists := instance.contents[filename]; !exists {
		var (
			content = []byte{}
			err     error
		)
		if content, err = os.ReadFile(path + filename + ext); err != nil {
			log.Println(err)
			return nil, err
		}
		if err = json.Unmarshal(content, &data); err != nil {
			log.Println(err)
			return nil, err
		}
		instance.contents[filename] = data
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

func parseKey(key string) (filename string, keys []string) {
	strs := strings.Split(key, sep)
	return strs[0], strs[1:]
}
func getContent(filename string) ([]byte, error) {
	return os.ReadFile(path + filename)
}

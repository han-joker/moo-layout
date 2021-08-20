package confm

import (
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
	"strings"
)

//变量
//实例池
var pool = map[string]*config{}

//默认选项值
var optionDefault = Option{
	Path: "./configs/",
	Ext:  ".yml",
	Sep:  ".",
}

//类型
//配置内容
type contents = map[string]map[string]interface{}
type config struct {
	option Option
	//配置文件内容解码缓存
	contents
}
type Option struct {
	Name string

	Path string
	Ext  string
	Sep  string
}

// New 创建对象
func New(option ...Option) *config {
	verifiedOption := optionVerify(option...)
	return create(verifiedOption)
}

// Get 存在直接返回，否则创建、存储再返回
func Get(option ...Option) *config {
	verifiedOption := optionVerify(option...)
	if !Has(verifiedOption.Name) {
		pool[verifiedOption.Name] = create(verifiedOption)
	}
	return pool[verifiedOption.Name]
}

// Has 存在返回 true，否则返回 false
func Has(name string) bool {
	_, has := pool[name]
	return has
}

func create(option Option) *config {
	return &config{
		option: option,
	}
}
func optionVerify(option ...Option) Option {
	opt := optionDefault
	if len(option) > 0 {
		opt.Name = option[0].Name
		if option[0].Path != "" {
			opt.Path = option[0].Path
		}
		if option[0].Ext != "" {
			opt.Ext = option[0].Ext
			if !strings.HasPrefix(opt.Ext, ".") {
				opt.Ext = "." + opt.Ext
			}
		}
		if option[0].Sep != "" {
			opt.Sep = option[0].Sep
		}
	}
	return opt
}

//可导出包方法

// Bool Getter|Setter
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
	return assertString(valueIf)
}

func (c config) Int(key string) int {
	valueIf, err := c.value(key)
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

func (c config) Float64(key string) float64 {
	valueIf, err := c.value(key)
	if err != nil {
		return 0
	}
	return assertFloat64(valueIf)
}

func (c config) Float32(key string) float32 {
	valueIf, err := c.value(key)
	if err != nil {
		return 0
	}
	return assertFloat32(valueIf)
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
		v, ok := vi.(int)
		if ok {
			value[i] = v
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
		value[i] = assertString(vi)
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
		value[i] = assertFloat64(vi)
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
		value[i] = assertFloat32(vi)
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
		value[k] = false
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
		value[k] = 0
		v, ok := vk.(int)
		if ok {
			value[k] = v
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
		value[k] = assertFloat64(vk)
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
		value[k] = assertFloat32(vk)
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
		value[k] = assertString(vk)
	}
	return value
}

//非导出方法

//利用 key 获取 值
func (c config) value(key string) (interface{}, error) {
	filename, keys := c.parseKey(key)
	if data, exists := c.contents[filename]; !exists {
		var (
			content = []byte{}
			err     error
		)
		if content, err = c.getContent(filename); err != nil {
			log.Println(err)
			return nil, err
		}
		if err = yaml.Unmarshal(content, &data); err != nil {
			log.Println(err)
			return nil, err
		}
		c.contents[filename] = data
	}

	//解析
	for i, l, currLevel := 0, len(keys), c.contents[filename]; i < l; i++ {
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
func (c config) parseKey(key string) (string, []string) {
	strs := strings.Split(key, c.option.Sep)
	if len(strs) == 1 {
		strs = append([]string{"app"}, strs...)
	}
	return strs[0], strs[1:]
}

//读取配置文件
func (c config) getContent(filename string) ([]byte, error) {
	return os.ReadFile(c.option.Path + filename + c.option.Ext)
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

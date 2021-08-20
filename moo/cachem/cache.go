package cachem

var pool = map[string]*cache{}

var optionDefault = Option{}

type cache struct {
	option Option
}

type Option struct {
	Name string
}

// New 创建对象
func New(option ...Option) *cache {
	verifiedOption := optionVerify(option...)
	return create(verifiedOption)
}

// Get 存在直接返回，否则创建、存储再返回
func Get(option ...Option) *cache {
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

func create(option Option) *cache {
	return &cache{
		option: option,
	}
}
func optionVerify(option ...Option) Option {
	opt := optionDefault
	if len(option) > 0 {
		opt = option[0]
	}
	return opt
}

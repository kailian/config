package config

type Options struct {
	FilePath string
	FileName string
}

func NewOptions() *Options {
	opt := &Options{
		".",
		"config",
	}
	return opt
}

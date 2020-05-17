package config

type Options struct {
	FilePath string
	FileName string
}

func (opt *Options) init() {
	if opt.FilePath == "" {
		opt.FilePath = "."
	}
	if opt.FileName == "" {
		opt.FileName = "config"
	}
}

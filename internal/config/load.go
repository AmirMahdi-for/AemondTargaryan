package config

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/structs"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"log"
	"strings"
)

const (
	delimeter      = "."
	seperator      = "__"
	envPrefix      = "AEMONDTARGARAYEN_"
	tagName        = "koanf"
	upTemplate     = "========== LoadConfiguration =========="
	bottomTemplate = "========================================"
)

func Load(print bool) *Config {
	k := koanf.New(delimeter)

	if err := k.Load(confmap.Provider(structs.Map(Default()), delimeter), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	if err := loadEnv(k); err != nil {
		log.Printf("error loading environment variables: %v", err)
	}

	config := Config{}
	var tag = koanf.UnmarshalConf{Tag: tagName}

	// بازگشایی تنظیمات به ساختار Config
	if err := k.UnmarshalWithConf("", &config, tag); err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

	// چاپ تنظیمات در صورت نیاز
	if print {
		log.Printf("%s\n%v\n%s", upTemplate, spew.Sdump(config), bottomTemplate)
	}

	return &config
}

func loadEnv(k *koanf.Koanf) error {
	// تابع تبدیل نام متغیرهای محیطی
	callback := func(source string) string {
		base := strings.ToLower(strings.TrimPrefix(source, envPrefix))
		return strings.ReplaceAll(base, delimeter, seperator)
	}

	// بارگذاری متغیرهای محیطی با استفاده از Provider
	if err := k.Load(env.Provider(envPrefix, delimeter, callback), nil); err != nil {
		return fmt.Errorf("error loading environment variables: %s", err)
	}

	return nil
}

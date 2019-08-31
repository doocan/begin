package setting

import (
	"github.com/go-ini/ini"
	"log"
)

type Database struct {
	Type string
	Path string
}

var DatabaseSetting = &Database{}

func SetUp(file string) {
	cfg, err := ini.Load(file)
	if err != nil {
		log.Fatalf("Fail to parse `conf/app.init`: %v.\n", err)
	}

	err = cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v.\n", err)
	}
}



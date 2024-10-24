package config

import (
	"github.com/smokecat/goweb-components/pkg/framework/go-zero/fw"
	"github.com/zeromicro/go-zero/core/logc"
	{{.authImport}}
)

type Config struct {
	Rest     rest.RestConf   `json:"rest"`
	Database fw.DatabaseConf `json:"database"`
	Logger   logc.LogConf    `json:"logger"`
	{{.auth}}
	{{.jwtTrans}}
}

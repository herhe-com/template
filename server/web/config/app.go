package config

import (
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Config()
	cfg.Add("app", map[string]any{
		"name":     cfg.Env("app.name", "UPER"),
		"version":  cfg.Env("app.version", "1.0.0"),
		"title":    cfg.Env("app.title", "UPER"),
		"domain":   cfg.Env("app.domain", "http://127.0.0.1:9600"),
		"debug":    cfg.Env("app.debug", false),
		"location": cfg.Env("app.location", "Asia/Shanghai"),
		"language": cfg.Env("app.language", "zh"),
		"node":     cfg.Env("app.node", 1),
	})
}

func Boot() {

}

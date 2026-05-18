package config

import (
	"github.com/herhe-com/framework/database/orm"
	"github.com/herhe-com/framework/database/redis"
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Config()
	cfg.Add("database", map[string]any{
		"orm": map[string]any{
			"default": cfg.Env("database.orm.default", "default"),
			"migration": map[string]any{
				"dir":   cfg.Env("database.orm.migration.dir", "/migration"),
				"table": cfg.Env("database.orm.migration.table", "sys_migration"),
			},
			"connections": map[string]any{
				"default": map[string]any{
					"driver":   cfg.Env("database.orm.connections.default.driver", orm.DriverMySQL),
					"username": cfg.Env("database.orm.connections.default.username", "root"),
					"password": cfg.Env("database.orm.connections.default.password", ""),
					"host":     cfg.Env("database.orm.connections.default.host", "127.0.0.1"),
					"port":     cfg.Env("database.orm.connections.default.port", "3306"),
					"db":       cfg.Env("database.orm.connections.default.db", "upper"),
					"charset":  cfg.Env("database.orm.connections.default.charset", "utf8mb4_unicode_ci"),
					"prefix":   cfg.Env("database.orm.connections.default.prefix", ""),
					"log_mode": cfg.Env("database.orm.connections.default.log_mode", "error"),
				},
			},
		},
		"redis": map[string]any{
			"default": cfg.Env("database.redis.default", "default"),
			"connections": map[string]any{
				"default": map[string]any{
					"driver":   cfg.Env("database.redis.connections.default.driver", redis.DriverRedis),
					"host":     cfg.Env("database.redis.connections.default.host", "127.0.0.1"),
					"port":     cfg.Env("database.redis.connections.default.port", "6379"),
					"username": cfg.Env("database.redis.connections.default.username", ""),
					"password": cfg.Env("database.redis.connections.default.password", ""),
					"db":       cfg.Env("database.redis.connections.default.db", 0),
				},
			},
		},
	})
}

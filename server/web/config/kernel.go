package config

import (
	"github.com/herhe-com/framework/console"
	"github.com/herhe-com/framework/console/consoles"
	cons "github.com/herhe-com/framework/contracts/console"
	"github.com/herhe-com/framework/contracts/service"
	"github.com/herhe-com/framework/database/orm"
	"github.com/herhe-com/framework/database/redis"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/filesystem"
	"github.com/herhe-com/framework/microservice/locker"
	"github.com/herhe-com/framework/microservice/snowflake"
	"github.com/herhe-com/framework/validation"
)

func init() {

	facades.Config().Add("kernel", map[string]any{
		"providers": []service.Provider{
			&orm.ServiceProvider{},
			&redis.ServiceProvider{},
			&filesystem.ServiceProvider{},
			&snowflake.ServiceProvider{},
			&locker.ServiceProvider{},
			&validation.ServiceProvider{},
			&console.ServiceProvider{},
		},
		"consoles": []cons.Provider{
			&consoles.ServerProvider{},
		},
	})
}

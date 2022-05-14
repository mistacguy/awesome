package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"zshf.private/configuration"
)

var (
	Config configuration.Config
	Db     *gorm.DB
	Rdb    *redis.Client
	Es     *elastic.Client
)

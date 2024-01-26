package localcache

import (
	"fmt"
	"github.com/coocood/freecache"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	cacheSize  = 100 * 1024 * 1024
	LocalCache = freecache.NewCache(cacheSize)
)

// 本次缓存中间件
func LocalCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			1. 如果key存在，则直接返回
			2. 如果key不存在，则下一步看看redis中有没有
		*/
		userid, _ := c.Get("userid")
		key := fmt.Sprintf("%s", userid)
		value, err := GetLocalCacheByUserId(key)
		if err != nil {
			log.Println("Miss Local Cache")
			c.Next()
		} else {
			log.Println("Hit Local Cache")
			c.JSON(http.StatusOK, gin.H{
				"message":     "success",
				"userprofile": value,
			})
			c.Abort()
		}
	}
}

func SetCache(key, value string, expire int) (err error) {
	err = LocalCache.Set([]byte(key), []byte(value), expire)
	if err != nil {
		return err
	}
	return nil
}

func GetCache(key string) (value []byte, err error) {
	value, err = LocalCache.Get([]byte(key))
	if err != nil {
		return nil, err
	}
	return value, nil
}

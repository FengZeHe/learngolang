// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/basicprojectv2/internal/repository"
	"github.com/basicprojectv2/internal/repository/cache"
	"github.com/basicprojectv2/internal/repository/dao"
	"github.com/basicprojectv2/internal/service"
	"github.com/basicprojectv2/internal/web"
	"github.com/basicprojectv2/internal/web/middleware"
	"github.com/basicprojectv2/ioc"
	"github.com/basicprojectv2/settings"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func InitializeApp() *gin.Engine {
	mysqlConfig := settings.InitMysqlConfig()
	db := ioc.InitDB(mysqlConfig)
	enforcer := ioc.InitMysqlCasbinEnforcer(db)
	userDAO := dao.NewUserDAO(db)
	redisConfig := settings.InitRedisConfig()
	cmdable := ioc.InitRedis(redisConfig)
	userCache := cache.NewUserCache(cmdable)
	userRepository := repository.NewCacheUserRepository(userDAO, userCache)
	casbinRoleCheck := middleware.NewCasbinRoleCheck(enforcer, userRepository)
	v := ioc.InitGinMiddlewares(casbinRoleCheck)
	userService := service.NewUserService(userRepository)
	codeCache := cache.NewCodeCache(cmdable)
	codeRepository := repository.NewCodeRepository(codeCache)
	smsService := ioc.InitSMSService()
	codeService := service.NewCodeService(codeRepository, smsService)
	userHandler := web.NewUserHandler(userService, codeService)
	sysDAO := dao.NewSysDAO(db)
	sysRepository := repository.NewSysRepository(sysDAO)
	sysService := service.NewSysService(sysRepository, enforcer)
	sysHandler := web.NewSysHandler(sysService)
	gormMenuDAO := dao.NewMenuDAO(db)
	menuRepository := repository.NewMenuRepository(gormMenuDAO)
	menuService := service.NewMenuService(menuRepository)
	menuHandler := web.NewMenuHandler(menuService)
	engine := ioc.InitWebServer(v, userHandler, sysHandler, menuHandler)
	return engine
}

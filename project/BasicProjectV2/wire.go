//go:build wireinject

package main

import (
	"github.com/IBM/sarama"
	"github.com/basicprojectv2/internal/events/article"
	"github.com/basicprojectv2/internal/repository"
	"github.com/basicprojectv2/internal/repository/cache"
	"github.com/basicprojectv2/internal/repository/dao"
	"github.com/basicprojectv2/internal/service"
	"github.com/basicprojectv2/internal/web"
	"github.com/basicprojectv2/internal/web/middleware"
	"github.com/basicprojectv2/ioc"
	"github.com/basicprojectv2/settings"
	"github.com/google/wire"
)

// ProvideSaramaConsumer提供sarama 消费者依赖
func ProvideSaramaConsumer(consumer sarama.Consumer, articleDAO dao.ArticleDAO) article.Consumer {
	return article.NewSaramaConsumer(consumer, articleDAO)
}

func ProvideSaramaConsumerClient() sarama.Consumer {
	// 初始化 Kafka 配置
	kafkaConfig := settings.InitSaramaConfig()
	// 初始化 Sarama 客户端
	client := ioc.InitSaramaClient(kafkaConfig)
	// 使用 InitConsumer 来创建 sarama.Consumer
	consumer := ioc.InitConsumer(client)
	return consumer
}

func InitializeApp() *App {
	wire.Build(
		// 读取配置
		settings.InitMysqlConfig, settings.InitRedisConfig, settings.InitSaramaConfig,

		// 第三方依赖部分
		ioc.InitDB, ioc.InitRedis, ioc.InitMysqlCasbinEnforcer, ioc.LoadI18nBundle,
		ioc.InitSaramaClient, ioc.InitSyncProducer,
		//ioc.InitConsumer,
		ProvideSaramaConsumerClient,
		ProvideSaramaConsumer,

		article.NewSaramaSyncProducer,
		//article.NewSaramaConsumer,

		// 测试Enforcer

		// Cache部分
		cache.NewCodeCache,
		cache.NewUserCache,

		// DAO部分
		dao.NewUserDAO,
		dao.NewSysDAO,
		dao.NewMenuDAO,
		dao.NewRoleDAO,
		dao.NewDraftDAO,
		dao.NewArticleDAO,

		// repository部分
		repository.NewCacheUserRepository,
		repository.NewCodeRepository,
		repository.NewSysRepository,
		repository.NewMenuRepository,
		repository.NewRoleRepository,
		repository.NewDraftRepository,
		repository.NewArticleRepository,

		// service部分
		ioc.InitSMSService,
		service.NewCodeService,
		service.NewUserService,
		service.NewSysService,
		service.NewMenuService,
		service.NewRoleService,
		service.NewDraftService,
		service.NewArticleService,

		//handler部分
		web.NewUserHandler,
		web.NewSysHandler,
		web.NewMenuHandler,
		web.NewRoleHandler,
		web.NewDraftHandler,
		web.NewArticleHandler,
		//wire.Bind(new(article.Producer), new(*article.SaramaSyncProducer)),

		// 中间件和路由
		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
		middleware.NewCasbinRoleCheck,
		wire.Struct(new(App), "*"),
	)

	return new(App)

}

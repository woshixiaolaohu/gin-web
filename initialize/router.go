package initialize

import (
	"gin-vue-admin/docs"
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

// Routers 初始化总路由
func Routers() *gin.Engine {
	// 创建不带中间件的路由
	Router := gin.New()
	// 恢复中间件
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	// 如果不想使用nginx代理 可以修改 web /.env.production下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面3行注释
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/assets", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.GVA_CONFIG.Local.StorePath)})
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	//global.GVA_LOG.Info("use middleware cors")
	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handle")

	//方便统一添加路由组前缀 多服务器上线使用

	publicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 健康检测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(publicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(publicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup, publicGroup)       //  注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)                    // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)                   // 注册用户相关路由
		systemRouter.InitMenuRouter(PrivateGroup)                   // 注册menu路由
		systemRouter.InitSystemRouter(PrivateGroup)                 // system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)                 // 权限相关路由
		systemRouter.InitAuthorityRouter(PrivateGroup)              // 注册角色路由
		systemRouter.InitSysDictionaryRouter(PrivateGroup)          // 字典管理
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)     // 操作记录
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)    // 字典详情管理
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)     // 权限按钮管理
		systemRouter.InitSysExportTemplateRouter(PrivateGroup)      // 导出模板
		systemRouter.InitAutoCodeRouter(PrivateGroup)               // 创建自动化代码
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)        // 自动化代码历史
		exampleRouter.InitCustomRouter(PrivateGroup)                // 客户路由
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由

	}
	// 插件路由安装
	InstallPlugin(PrivateGroup, publicGroup)
	global.GVA_LOG.Info("路由注册成功")
	return Router
}

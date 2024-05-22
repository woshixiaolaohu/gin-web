package cmd

import (
	"flag"
	"fmt"
	"gin-vue-admin/cmd/internal"
	"gin-vue-admin/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// Viper
// 优先级 命令行 > 环境变量 > 默认值

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		// 解析命令行函数，并将解析的结果存储到字符串变量
		flag.StringVar(&config, "c", "", "choose config file")
		flag.Parse()
		// 判断命令行参数是否为空
		if config == "" {
			// 判断internal.ConfigEnv是否为空
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("正在使用gin模式的%s环境名称，config路径为%s\n", gin.Mode(), config)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("正在使用gin模式的%s环境名称，config路径为%s\n", gin.Mode(), config)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("正在使用gin模式的%s环境名称，config路径为%s\n", gin.Mode(), config)
				}
			} else {
				//不为空的话，将configEnv赋值给config
				config = configEnv
				fmt.Printf("正在使用gin模式的%s环境名称，config路径为%s\n", internal.ConfigEnv, config)
			}
		} else {
			fmt.Printf("正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else {
		// 将可变参数的第一个值赋值给config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}
	// 新建viper
	v := viper.New()
	// 设置配置文件
	v.SetConfigFile(config)
	// 设置配置文件类型
	v.SetConfigType("yaml")
	// 加载配置文件
	err := v.ReadInConfig()
	if err != nil {
		// 程序停止运行
		panic(fmt.Errorf("Fatal Error Config File: %s\n", config))
	}
	// 监听配置文件修改
	v.WatchConfig()
	// 配置文件修改后执行方法
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config File Changed", in.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			panic(err)
		}
	})
	// root适配性 根据root位置去找对应的文件迁移位置，保证root路径有效
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}

package main

import (
	"fmt"
	"github.com/blog-service/global"
	"github.com/blog-service/internal/model"
	"github.com/blog-service/internal/routers"
	"github.com/blog-service/pkg/logger"
	"github.com/blog-service/pkg/setting"
	// _ "github.com/swaggo/gin-swagger/example/docs"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

// @title 博客系统
// @version 1.json.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 1048576
	}

	fmt.Println("1111:::==>", global.ServerSetting.HttpPort)
	fmt.Println("2222:::==>", global.DatabaseSetting.Password)

	// 打印日志
	global.Logger.Infof("%s:go-tour %s", "yida", "very-good")

	s.ListenAndServe()
}

func init() {
	// 配置文件
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v\n", err)
	}

	// 连接数据库
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v\n", err)
	}

	// 设置日志
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v\n", err)
	}
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,  // 文件最大大小是600M
		MaxAge:    10,   // 保存时间时10天
		LocalTime: true, // 时间格式是本地时间
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

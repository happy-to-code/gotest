package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/blog-service/global"
	"github.com/blog-service/internal/model"
	"github.com/blog-service/internal/routers"
	"github.com/blog-service/pkg/logger"
	"github.com/blog-service/pkg/setting"
	"github.com/blog-service/pkg/tracer"
	"os"
	"os/signal"
	"syscall"

	// _ "github.com/swaggo/gin-swagger/example/docs"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

// @title 博客系统
// @version bb.json.0
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
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err:%v", err)
		}
	}()
	quit := make(chan os.Signal)
	// 接收syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server ……")
	// 最大时间控制，用于通知该服务端它有5秒时间来处理原有的请求
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutDown:", err)
	}
	log.Println("Server exiting")

	fmt.Println("1111:::==>", global.ServerSetting.HttpPort)
	fmt.Println("2222:::==>", global.DatabaseSetting.Password)

	// 打印日志
	global.Logger.Infof("%s:go-tour %s", "yida", "very-good")

}

var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

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

	// 链路追踪
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v\n", err)
	}

	err = setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
}
func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}
func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
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

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

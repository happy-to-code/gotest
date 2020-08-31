package main

import (
	"fmt"
	"github.com/blok-service/global"
	"github.com/blok-service/internal/routers"
	"github.com/blok-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

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
	s.ListenAndServe()
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v\n", err)
	}
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

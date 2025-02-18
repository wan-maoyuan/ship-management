package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"ship-management/pkg/config"
	"ship-management/pkg/database"
	"ship-management/pkg/global"
	"ship-management/pkg/web"
	"syscall"

	"github.com/go-kratos/kratos/v2"
	"github.com/sirupsen/logrus"
)

func init() {
	c, err := config.New()
	if err != nil {
		logrus.Fatalf("配置信息加载失败: %v", err)
	}

	c.Show()
}

func main() {
	global.ShowProgramInfo()

	srv := web.NewService()
	app := newApp(srv)

	go func() {
		if err := app.Run(); err != nil {
			logrus.Fatalf("ship_management 程序启动失败: %v", err)
		}
	}()

	stopSingal := make(chan os.Signal, 1)
	signal.Notify(stopSingal, syscall.SIGINT, syscall.SIGTERM)
	<-stopSingal
	logrus.Info("ship_management 程序正在停止 ......")

	if err := app.Stop(); err != nil {
		logrus.Errorf("ship_management 程序停止失败: %v", err)
	}
}

func newApp(gs *web.Service) *kratos.App {
	return kratos.New(
		kratos.Name("ship_management"),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs),
		kratos.AfterStart(AfterStart),
		kratos.AfterStop(AfterStop),
	)
}

func AfterStart(ctx context.Context) error {
	if err := database.InitDatabase(config.Get().Server.SqlserverDsn); err != nil {
		return fmt.Errorf("数据库初始化失败: %v", err)
	}

	return nil
}

func AfterStop(ctx context.Context) error {
	c := config.Get()
	c.Log.Close()

	return nil
}

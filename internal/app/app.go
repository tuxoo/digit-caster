package app

import (
	"context"
	"digit-caster/internal/config"
	"digit-caster/internal/server"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	httpServer := server.NewHTTPServer(cfg, httpHandlers.Init(cfg.HTTP))

	go func() {
		if err := httpServer.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logrus.Print("Digit Caster application has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Digit Caster application shutting down")

	if err := httpServer.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on http server shutting down: %s", err.Error())
	}
}

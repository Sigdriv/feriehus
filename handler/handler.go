package handler

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Handler struct {
	Config    Config `yaml:",inline"`
	Validator *validator.Validate
}

type Config struct {
	Port string `yaml:"port" validate:"required"`
}

func CreateHandler() (srv Handler) {
	data, err := os.ReadFile("./cfg/cfg.yml")
	if err != nil {
		logrus.Fatalf("Failed reading config file >> %v", err)
	}

	err = yaml.Unmarshal(data, &srv)
	if err != nil {
		logrus.Fatalf("Failed parsing config file >> %v", err)
	}

	srv.resolveFileReferences(&srv)

	srv.Validator = validator.New()
	err = srv.Validator.Struct(srv.Config)
	if err != nil {
		logrus.Fatalf("Config validation failed >> %v", err)
	}

	return srv
}

func (srv *Handler) CreateGinGroup() {
	router := gin.New()

	apiRouter := router.Group("/api")
	apiRouter.GET("/apartments/:id", srv.HandleGetApartment)

	apiRouter.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()

		logrus.WithFields(logrus.Fields{
			"status":  c.Writer.Status(),
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
			"latency": time.Since(start),
		}).Info("Request handled")
	})

	router.Static("/template", "./template")
	router.Static("/images", "./data/images")

	runner := fmt.Sprintf("localhost:%s", srv.Config.Port)
	router.Run(runner)
}

func (*Handler) getLog(c *gin.Context) *logrus.Logger {
	url := c.Request.URL
	logger := logrus.New()

	logger.WithFields(logrus.Fields{
		"method": c.Request.Method,
		"url":    url.String(),
	})

	return logger
}

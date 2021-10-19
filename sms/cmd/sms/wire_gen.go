// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/ZQCard/kratos-service-base/sms/internal/biz"
	"github.com/ZQCard/kratos-service-base/sms/internal/conf"
	"github.com/ZQCard/kratos-service-base/sms/internal/data"
	"github.com/ZQCard/kratos-service-base/sms/internal/server"
	"github.com/ZQCard/kratos-service-base/sms/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	smsRepo := data.NewSmsRepo(dataData, logger)
	smsUseCase := biz.NewSmsUseCase(smsRepo, logger)
	smsService := service.NewSmsService(smsUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, smsService, logger)
	grpcServer := server.NewGRPCServer(confServer, smsService, tracerProvider, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/ZQCard/kratos-service-base/coupon/internal/biz"
	"github.com/ZQCard/kratos-service-base/coupon/internal/conf"
	"github.com/ZQCard/kratos-service-base/coupon/internal/data"
	"github.com/ZQCard/kratos-service-base/coupon/internal/server"
	"github.com/ZQCard/kratos-service-base/coupon/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(discovery)
	dataData, cleanup, err := data.NewData(confData, logger, userClient)
	if err != nil {
		return nil, nil, err
	}
	couponRepo := data.NewCouponRepo(dataData, logger)
	couponUseCase := biz.NewCouponUseCase(couponRepo, logger)
	couponService := service.NewCouponService(couponUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, couponService, logger)
	grpcServer := server.NewGRPCServer(confServer, couponService, tracerProvider, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
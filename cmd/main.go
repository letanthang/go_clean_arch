package main

import (
	"app/internal/config"
	"app/internal/domain/entity"
	"app/internal/domain/usecase"
	"app/internal/framework/route"
	"app/internal/handler"
	"app/internal/infrastructure/db/mysql"
	"app/internal/infrastructure/db/query_builder"
	gormmysql "app/pkg/gorm_mysql"
	"context"
	"fmt"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	db, err := gormmysql.InitDatabase(gormmysql.Config{
		User:         cfg.DBUser,
		Password:     cfg.DBPass,
		Host:         cfg.DBHost,
		Port:         cfg.DBPort,
		DatabaseName: cfg.DBName,
	})
	if err != nil {
		panic(err)
	}

	productRepo := mysql.NewProductRepository(db)
	getUsecase := usecase.NewProductGetUsecase(productRepo)
	listUsecase := usecase.NewProductListUsecase(productRepo)
	productHandler := handler.NewProductHandler(getUsecase, listUsecase)

	r := route.NewRouter(productHandler)
	if err := r.Run(":" + cfg.PORT); err != nil {
		panic(fmt.Errorf("failed to run server: %w", err))
	}
}

func mainDebug() {

	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	db, err := gormmysql.InitDatabase(gormmysql.Config{
		User:         cfg.DBUser,
		Password:     cfg.DBPass,
		Host:         cfg.DBHost,
		Port:         cfg.DBPort,
		DatabaseName: cfg.DBName,
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	productRepo := mysql.NewProductRepository(db)
	if err := productRepo.Create(ctx, &entity.Product{Name: "Product 1", SKU: "Product1"}); err != nil {
		panic(err)
	}

	product1, err := productRepo.Get(ctx, query_builder.NewProductByIDQuery(1))
	if err != nil {
		panic(err)
	}

	fmt.Println(product1)
}

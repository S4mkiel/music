package db

import (
	"context"

	"github.com/S4mkiel/music/domain/entity"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresModule = fx.Module("postgres", fx.Provide(NewClient), fx.Invoke(HookDatabase))

func NewClient(postgresConfig Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(postgresConfig.ConnectionString()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func HookDatabase(lc fx.Lifecycle, db *gorm.DB, logger *zap.SugaredLogger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			dbDriver, err := db.DB()
			if err != nil {
				logger.Fatal("Failed to get DB driver", zap.Error(err))
				return err
			}

			err = dbDriver.Ping()
			if err != nil {
				logger.Fatal("failed to ping database", zap.Error(err))
				return err
			}

			err = enableUUIDExtension(db)
			if err != nil {
				logger.Fatal("failed to enable uuid extension", zap.Error(err))
				return err
			}


			err = migrate(db)
			if err != nil {
				logger.Fatal("failed to migrate database", zap.Error(err))
				return err
			}

			logger.Info("Database OK!")
			return nil
		},
		OnStop: func(context.Context) error {
			dbDriver, err := db.DB()
			if err != nil {
				logger.Fatal("Failed to get DB driver", zap.Error(err))

			}

			err = dbDriver.Close()
			if err != nil {
				logger.Fatal("failed to close database connection", zap.Error(err))

			}
			return nil
		},
	})
}

func enableUUIDExtension(db *gorm.DB) error {
	_, err := db.Raw(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Rows()
	return err
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.PdfFile{},
	)
	return err
}

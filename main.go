package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/S4mkiel/music/domain/module"
	"github.com/S4mkiel/music/infra/config"
	"github.com/S4mkiel/music/infra/db"
	"github.com/S4mkiel/music/infra/http"
	"github.com/S4mkiel/music/infra/logger"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

// @title Music Example API
// @version 0.1
// @description This is a endpoint's example
// @termsOfService italo.me
// @contact.name Italo Esdrass
// @contact.email suport@italo.me
// @license.name Italo 1.0
// @license.url italoesdrass.me
// @host localhost:8000/api/v1
// @BasePath /
func main() {
	if os.Getenv("ENV") != "production" {
		LoadConfig()
	}
	fx.New(
		module.Service,
		http.Module,
		logger.Module,
		config.Module,
		db.Module,
	).Run()
}

func LoadConfig() {
	_, b, _, _ := runtime.Caller(0)

	basepath := filepath.Dir(b)

	err := godotenv.Load(fmt.Sprintf("%v/.env", basepath))
	if err != nil {
		panic(err)
	}
}

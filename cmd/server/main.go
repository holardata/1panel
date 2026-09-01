package main

import (
	"fmt"
	"os"

	_ "net/http/pprof"

	"github.com/1Panel-dev/1Panel/cmd/server/cmd"
	_ "github.com/1Panel-dev/1Panel/cmd/server/docs"
)

// @title 1Panel
// @version 1.0
// @description Top-Rated Web-based Linux Server Management Tool
// @termsOfService http://swagger.io/terms/
// @license.name GPL-3.0
// @license.url https://www.gnu.org/licenses/gpl-3.0.html
// @BasePath /api/v1
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @description JWT Authorization header. Format: `Authorization: Bearer <token>`
// @type apiKey
// @in header
// @name Authorization

//go:generate env GOFLAGS=-mod=mod go run github.com/swaggo/swag/cmd/swag init -o ./docs -g main.go -d ../../backend -g ../cmd/server/main.go
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"

	"github.com/axosec/auth/internal/config"
	"github.com/axosec/auth/internal/data/db"
	"github.com/axosec/core/crypto/token"
)

var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"
)

func main() {
	fmt.Printf(`Starting Axosec Auth
Version: %s
Commit: %s
Built:  %s
`, Version, GitCommit, BuildDate)
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("failed to load config: %s\n", err)
		return
	}

	privateKey, publicKey, err := token.LoadKeysFromFiles(cfg.JWT.PrivateKeyPath, cfg.JWT.PublicKeyPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = token.NewJWTManager(privateKey, publicKey, cfg.JWT.Issuer)

	_, err = db.NewConnection(cfg.Database)
	if err != nil {
		fmt.Printf("failed to connect to database: %s", err)
		return
	}
}

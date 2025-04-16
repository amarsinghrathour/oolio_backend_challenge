package migratecmd

import (
	"fmt"
	"github.com/amarsinghrathour/oolio_backend_challenge/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func RunMigration() {

	cfg := config.LoadConfig("../.env")
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	m, err := migrate.New("file://../migrations", dbUrl)
	if err != nil {
		log.Fatalf("❌ Migration init error: %v", err)
	}

	log.Println("🚀 Running DB migrations...")

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("❌ Migration error: %v", err)
	}

	log.Println("✅ Database migrated successfully.")
}

package conf

import (
	"database/sql"
	"go-clean-arch/internal/platform/db"
	"go-clean-arch/internal/platform/environment"
	"go-clean-arch/internal/platform/log"

	"github.com/DATA-DOG/go-sqlmock"
)

var dep *dependencies

type dependencies struct {
	DB *sql.DB
}

// Here we can instantiate dependencies.
func init() {
	deployEnv := environment.GetCurrentEnvironment() != environment.Development

	db, err := buildDatabase(deployEnv)

	if !deployEnv && err != nil {
		log.Info("Mocking the database...")
		db, _, err = sqlmock.New()
	}

	if err != nil {
		log.Errorf("Error trying to setup database connection.")
	}

	dep = &dependencies{
		db,
	}

}

func buildDatabase(deployEnv bool) (*sql.DB, error) {
	user := environment.GetEnvOrConfig("DATABASE_USER")
	pass := environment.GetEnvOrConfig("DATABASE_PASS")
	host := environment.GetEnvOrConfig("DATABASE_HOST")
	port := environment.GetEnvOrConfig("DATABASE_PORT")
	name := environment.GetEnvOrConfig("DATABASE_NAME")

	db, err := db.New(user, pass, host, port, name)

	if err != nil {
		log.Errorf("Error trying to setup database connection.")
	}

	if !deployEnv && err != nil {
		log.Info("Mocking the database...")
		db, _, err = sqlmock.New()
	}

	return db, err
}

func GetDependencies() *dependencies {
	return dep
}

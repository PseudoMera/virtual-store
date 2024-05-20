package shared

import (
	"context"
	"os"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// SetupPostgresClient creates a postgres container and reads all sql files from the schema directory to initialize a postgres database.
func SetupPostgresClient(ctx context.Context, projectRootPath string) (*PostgresDB, *postgres.PostgresContainer, error) {
	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:16.1-alpine"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, nil, err
	}

	connectionString, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, nil, err
	}

	db, err := NewPostgresDatabase(context.Background(), connectionString)
	if err != nil {
		return nil, nil, err
	}

	schemas := []string{"schema/schema.sql", "schema/triggers.sql"}

	for _, schema := range schemas {
		schemaSQLStatements, err := loadSQLFile(projectRootPath + schema)
		if err != nil {
			return nil, nil, err
		}

		_, err = db.db.Exec(context.Background(), schemaSQLStatements)
		if err != nil {
			return nil, nil, err
		}
	}

	return db, container, nil
}

// loadEntireSQLFile loads the entire file to execute as one big query
// If not done this way we would have to implement a custom splitter
func loadSQLFile(filePath string) (string, error) {
	// Read the entire SQL file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), err
}

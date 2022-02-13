package db

import (
	"codegram/ent"
	"context"
	"database/sql"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func Open(db_url string) *ent.Client {
	db, err := sql.Open("pgx", db_url)

	if err != nil {
		log.Fatal(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func SetupDb() *ent.Client {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading env file, maybe missing ?")
	}

	db_url := os.Getenv("DB_URL")

	client := Open(db_url)
	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

package utility

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func DbUp() {
	var ctx = gctx.New()
	value, err := g.Cfg().GetWithEnv(ctx, "database.default")
	if err != nil {
		panic(err)
	}
	dbConfig := value.Map()

	name := dbConfig["name"].(string)
	// mysql connection string format: user:password@tcp(host:port)/dbname
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig["user"], dbConfig["pass"], dbConfig["host"], dbConfig["port"], dbConfig["name"])

	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations/mysql",
		name, driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			g.Log().Info(ctx, "Migration no change")
		} else {
			g.Log().Error(ctx, "Error while running migration: %v", err)
			panic(err)
		}
	} // or m.Steps(2) if you want to explicitly set the number of migrations to run
	g.Log().Info(ctx, "Migration completed")
}

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hunght/gqlgen-sqlc-example/global"

	"github.com/hunght/gqlgen-sqlc-example/gqlgen"
	"github.com/hunght/gqlgen-sqlc-example/pg"
)

func main() {
	// initialize the config
	global.InitConfig()
	// create postgresql information for connect
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", global.Conf.Postgresql.Host, global.Conf.Postgresql.Port, global.Conf.Postgresql.User, global.Conf.Postgresql.Password, global.Conf.Postgresql.DbName, global.Conf.Postgresql.SSLMode)
	// initialize the db
	db, err := pg.Open(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize the repository
	repo := pg.NewRepository(db)

	// configure the server
	mux := http.NewServeMux()
	mux.Handle("/", gqlgen.NewPlaygroundHandler("/query"))
	mux.Handle("/query", gqlgen.NewHandler(repo))

	// run the server
	port := ":8080"
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
	fmt.Fprintln(os.Stderr, http.ListenAndServe(port, mux))
}

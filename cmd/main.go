package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/bootcamp-go/desafio-cierre-db.git/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	db, err := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/fantasy_products")
	if err != nil {
		panic(err)
	}

	router.NewRouter(r, db).MapRoutes()

	r.Run()

}

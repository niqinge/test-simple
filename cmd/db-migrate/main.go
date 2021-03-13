package main

import (
	"flag"
	"fmt"
	"github.com/niqinge/test-simple/pkg/model"
	"github.com/niqinge/utils/mysql"
)

var do string
var db string

func main() {
	flag.StringVar(&do, "do", "", "do")
	flag.StringVar(&db, "db", "", "dbName")
	flag.Parse()
	
	// step1 creat db connection
	migrateOrm := mysql.NewMigrateOrm(mysql.NewLocalDBConf(db))
	// step2
	migrateOrm.SetModels(&model.User{})
	
	switch do {
	case "create":
		if err := migrateOrm.CreateDB(); err != nil {
			panic(err)
		}
	case "migrate":
		if err := migrateOrm.MigrateDB(); err != nil {
			panic(err)
		}
	default:
		panic(fmt.Sprintf("no support do:%s", do))
	}
}

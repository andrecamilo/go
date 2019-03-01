package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/nakagami/firebirdsql"
)

type usuario struct {
	Id   int
	Nome string
}

type linhatabela struct {
	Codigo             int    `db:"CODIGO"`
	Identificadorlinha string `db:"IDENTIFICADOR_LINHA"`
}

func main() {

	//TestMySql()
	TestFirebird()

	fmt.Scanln()
}

func TestFirebird() {
	//db, err := sql.Open("firebirdsql",
	//"SYSDBA:masterkey@192.168.231.208/home/rangelsantos/discod/desenvolvimento/dbsfirebird/forquilinha/avls_20190128.fdb")

	db, err := sqlx.Connect("firebirdsql",
		"SYSDBA:masterkey@192.168.231.208/home/rangelsantos/discod/desenvolvimento/dbsfirebird/forquilinha/avls_20190128.fdb")

	if err != nil {
		log.Fatalln(err)
	}

	//fcvsSql(db.DB)
	fcvsSqlx(db)

	defer db.Close()
}

func fcvsSqlx(db *sqlx.DB) {
	people := []linhatabela{}
	db.Select(&people, `
	SELECT LT.CODIGO, LT.IDENTIFICADOR_LINHA
	  FROM LINHAS_TABELAS LT 
	 WHERE LT.codigo = ?`, 873150)

	fmt.Println(people)
}

func fcvsSql(conn *sql.DB) {
	rows, _ := conn.Query(`
		SELECT LT.CODIGO, LT.IDENTIFICADOR_LINHA
		  FROM LINHAS_TABELAS LT 
		 WHERE LT.codigo = ?`,
		873150)

	defer rows.Close()

	for rows.Next() {
		var u linhatabela
		rows.Scan(&u.Codigo)
		//rows.Scan(&u.Codigo, &u.Identificadorlinha)
		fmt.Println(u)
	}
}

func TestMySql() {
	//db, err := sql.Open("mysql", "root:usuubu@tcp(192.168.231.208)/cursogo")
	db, err := sqlx.Connect("mysql", "root:usuubu@tcp(192.168.231.208)/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//getUsuariosSql(db.DB)
	getUsuariosSqlx(db)
}

func getUsuariosSqlx(db *sqlx.DB) {
	people := []usuario{}
	db.Select(&people, "select id, nome from usuarios where id = ?", 1)

	fmt.Println(people)
}

func getUsuariosSql(db *sql.DB) {
	rows, _ := db.Query("select id, nome from usuarios where id = ?", 1)

	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.Id, &u.Nome)
		fmt.Println(u)
	}
}
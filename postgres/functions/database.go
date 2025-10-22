package functions

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Categoria struct {
	ID   int64
	Nome string
}

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=0000 dbname=godatabase sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}

func Insert(c Categoria) (id int64, err error) {
	conn, err := OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO categorias (nome) VALUES ($1) RETURNING id`

	err = conn.QueryRow(sql, c.Nome).Scan(&id)

	return id, err
}

func Get(id int64) (c Categoria, err error) {
	conn, err := OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM categorias WHERE id=$1`, id)

	err = row.Scan(&c.ID, &c.Nome)

	return
}

package user

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type ZiDB struct {
	*sql.DB
}

func OpenDB() *ZiDB {
	db, err := sql.Open("mysql",
		"cazzi:password@tcp(127.0.0.1:3306)/cazzi_users")
	if err != nil {
		log.Fatal(err)
	}
	return &ZiDB{
		db,
	}
}

func (zdb *ZiDB) Close() {
	zdb.Close()
}

func (zdb *ZiDB) New(u User) error {
	_, err := zdb.Exec(
		"INSERT INTO users(id, username) VALUES(?, ?)",
		u.ID.String(),
		u.Username,
	)
	return err
}

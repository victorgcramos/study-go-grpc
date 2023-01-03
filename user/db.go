package user

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type ZiDB struct {
	db *sql.DB
}

func NewDB() ZiDB {
	db, err := sql.Open("mysql",
		"cazzi:password@tcp(127.0.0.1:3306)/cazzi_users")
	if err != nil {
		log.Fatal(err)
	}
	return ZiDB{
		db: db,
	}
}

func (zdb *ZiDB) Close() {
	zdb.Close()
}

func (zdb *ZiDB) New(u User) error {
	_, err := zdb.db.Exec(
		"INSERT INTO users(id, username) VALUES(?, ?)",
		u.ID.String(),
		u.Username,
	)
	return err
}

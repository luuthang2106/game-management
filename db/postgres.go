package db

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/storage/postgres/v3"
)

var pg *postgres.Storage

func init() {
	pg = postgres.New(postgres.Config{
		ConnectionURI: "postgres://vcenzxal:1kX_EN9ewICl7A1UNQV1W2iLnW_3_ugP@floppy.db.elephantsql.com/vcenzxal",
		Database:      "test_db",
	})
	migrate()
}
func migrate() {
	// Đọc nội dung của file SQL
	sqlFile, err := os.ReadFile("db/migration.sql")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Không thể đọc file SQL: %v\n", err)
		os.Exit(1)
	}

	// Chuyển nội dung file SQL thành một chuỗi
	sqlQuery := string(sqlFile)
	// Thực thi truy vấn SQL
	_, err = pg.Conn().Exec(context.Background(), sqlQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Lỗi khi thực thi truy vấn SQL: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("File SQL đã được thực thi thành công.")
}
func GetStorage() *postgres.Storage {
	return pg
}

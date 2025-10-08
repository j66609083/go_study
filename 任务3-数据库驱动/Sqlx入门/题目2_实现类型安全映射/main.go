package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Id     int
	Title  *string
	Author sql.NullString
	Price  sql.NullFloat64
}

func (b Book) String() string {
	return fmt.Sprintf("Book{Id: %d, Title: %v, Author: %v, Price: %.2f}",
		b.Id, *(b.Title), b.Author, b.Price)
}

/*
编写Go代码，使用Sqlx执行一个复杂的查询，
例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
func getBookByPriceGreaterThan(db *sqlx.DB, price float64) ([]Book, error) {
	var books []Book
	err := db.Select(&books, "SELECT id, title, author, price FROM books WHERE price > ?", price)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func main() {
	db := sqlx.MustConnect("mysql", "root:123456@tcp(localhost:3306)/go_study")
	defer db.Close()

	books, err := getBookByPriceGreaterThan(db, 50)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range books {
		log.Printf("%+v\n", e)
	}
}

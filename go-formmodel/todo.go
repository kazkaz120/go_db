package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
}

func CreateTasks(c echo.Context) error {

	name := c.FormValue("name")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: name})

	var pro2 []string
	product := []Product{}
	db.Find(&product)
	for _, pro := range product {
		pro2 = append(pro2, pro.Code)
		fmt.Println(pro.Code)
	}
	//	fmt.Println(a)
	//	return nil
	return c.JSON(http.StatusOK, pro2)
}

func main() {

	e := echo.New()

	e.File("/", "public/index.html")
	e.POST("/writeout", CreateTasks)
	e.Start(":8080")

	// Read
	//	var product Product
	//	db.First(&product, 1) // find product with integer primary key
	//	db.First(&product, "code = ?", "D42")
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Cars struct {
	Id        int
	Name      string
	Model     string
	Year      int
	DeleteAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func dbConn() (DB *gorm.DB) {

	dsn := "root:root@tcp(127.0.0.1:8889)/crud_cars?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
	})

	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Cars{})
	return db
}

func getRecords(c *gin.Context) {

	db := dbConn()

	var cars []Cars
	db.Find(&cars)

	c.JSON(200, &cars)

}

func insertRecods(c *gin.Context) {

	log.Println("test")
	data, err := ioutil.ReadAll(c.Request.Body)

	//check why i need doing it
	//I guess that it use to convert to string this endpoint returns Json format
	records := Cars{}
	json.Unmarshal([]byte(data), &records)

	var name = records.Name
	var model = records.Model
	var year = records.Year

	if err != nil {
		log.Panicf("error: %s", err)
		c.JSON(500, "not ok records")
	}

	db := dbConn()

	car := Cars{Name: name, Model: model, Year: year, CreatedAt: time.Now()}
	db.Select("name", "model", "year", "created_at").Create(&car)

	c.JSON(200, "Created records")

}

func updateRecods(c *gin.Context) {

	db := dbConn()

	db.Model(&Cars{}).Where("id = 1").Updates(Cars{Name: "helloTEST", Model: "modelTest", Year: 555555})

	c.JSON(200, "Updated records")

}

func deleteRecods(c *gin.Context) {

	db := dbConn()

	db.Delete(&Cars{}, 2)

	c.JSON(200, "Deleted records")

}

func main() {
	log.Println("Server started on: http://localhost:8080")

	server := gin.Default()

	server.Use(cors.Default())

	server.GET("/", getRecords)
	server.POST("/insert", insertRecods)
	server.PUT("/update", updateRecods)
	server.DELETE("/delete", deleteRecods)

	server.Run(":3000")

}

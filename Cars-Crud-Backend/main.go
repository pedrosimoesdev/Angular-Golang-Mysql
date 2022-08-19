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
	// to do there are issue with id and year int to
	Id        string
	Name      string
	Model     string
	Year      string
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

	data, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Panicf("error: %s", err)
		c.JSON(500, "not ok records")
	}

	//check why i need doing it
	//I guess that it use to convert to string this endpoint returns Json format
	records := Cars{}
	json.Unmarshal([]byte(data), &records)

	var name = records.Name
	var model = records.Model
	var year = records.Year

	db := dbConn()

	car := Cars{Name: name, Model: model, Year: year, CreatedAt: time.Now()}
	db.Select("name", "model", "year", "created_at").Create(&car)

	c.JSON(200, "Created records")

}

func updateRecods(c *gin.Context) {

	log.Println("test")

	data, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Panicf("error: %s", err)
		c.JSON(500, "not ok records")
	}

	//check why i need doing it
	//I guess that it use to convert to string this endpoint returns Json format
	records := Cars{}
	json.Unmarshal([]byte(data), &records)

	var id = records.Id
	var name = records.Name
	var model = records.Model
	var year = records.Year

	log.Println(records)
	log.Println(id)
	log.Println(name)
	log.Println(year)
	log.Println(model)

	db := dbConn()

	db.Model(&Cars{}).Where("id = ?", id).Updates(Cars{Name: name, Model: model, Year: year})

	c.JSON(200, "Updated records")

}

func deleteRecods(c *gin.Context) {

	data, err := ioutil.ReadAll(c.Request.Body)

	var id = string(data)
	log.Println("ID:", id)

	if err != nil {
		log.Panicf("error: %s", err)
		c.JSON(500, "not ok records")
	}

	db := dbConn()

	db.Delete(&Cars{}, id)

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

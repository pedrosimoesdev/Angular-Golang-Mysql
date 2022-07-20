package main

import (
	"log"
	"time"

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

	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/crud_cars")
	/*
		db.SetConnMaxLifetime(time.Minute * 4) // <-- this
		if err != nil {
			log.Println("error")

			panic(err.Error())
		}
	*/

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
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

	/*
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM cars")

		if err != nil {
			log.Println("error 2 ")
			panic(err.Error())
		}

		c.Request.Context()
		car := Cars{}
		var response []Cars

		for selDB.Next() {
			var id int
			var name string
			var model string
			var year int
			err = selDB.Scan(&id, &name, &model, &year)
			if err != nil {
				log.Println("error 3")
				panic(err.Error())
			}
			car.id = id
			car.name = name
			car.model = model
			car.year = year

			//used to fill array
			response = append(response, car)
		}

		//usersBytes, _ := json.Marshal(&response)

		log.Println(response)

		c.JSON(200, gin.H{
			"response": response,
		})
	*/

}

func insertRecods(c *gin.Context) {

	db := dbConn()

	car := Cars{Name: "car orm", Model: "model ORM ", Year: 255555, CreatedAt: time.Now()}
	db.Select("name", "model", "year", "created_at").Create(&car)

	c.JSON(200, "Created records")

	/*
		sql := "INSERT INTO cars(name, model,year) VALUES (?,?,?)"
		res, err := db.Exec(sql, "Opel", "test", 1997)

		if err != nil {
			panic(err.Error())
		}

		lastId, err := res.LastInsertId()

		if err != nil {
			log.Fatal(err)
			log.Fatal(lastId)
		}

		fmt.Fprintln(w, "insert Records")
		log.Println("Records Saved")
	*/

}

func updateRecods(c *gin.Context) {

	db := dbConn()

	db.Model(&Cars{}).Where("id = 1").Updates(Cars{Name: "helloTEST", Model: "modelTest", Year: 555555})

	c.JSON(200, "Updated records")

	/*

		sql := `UPDATE cars
		       SET name = ? , model = ?, year = ?
		       WHERE id = ?;`
		_, err := db.Exec(sql, "opelUpdate", "modelUpdated", 2022, 6)

		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintln(w, "Updated Record")
		log.Println("Record Updated")
	*/
}

func deleteRecods(c *gin.Context) {

	db := dbConn()

	db.Delete(&Cars{}, 2)

	c.JSON(200, "Deleted records")

	//db := dbConn()
	/*
		defer db.Close()

		sql := `DELETE FROM cars WHERE id = ?;`
		_, err := db.Exec(sql, 6)

		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintln(w, "Deleted Record")
		log.Println("Deleted Record")
	*/
}

func main() {
	log.Println("Server started on: http://localhost:8080")

	//http.HandleFunc("/update", updateRecods)
	//http.HandleFunc("/delete", deleteRecods)

	server := gin.Default()

	server.GET("/", getRecords)
	server.POST("/insert", insertRecods)
	server.PUT("/update", updateRecods)
	server.DELETE("/delete", deleteRecods)

	server.Run(":8080")

}

package main

import (
	// "github.com/YosukeCSato/go_sample/routes"

	"fmt"
	"m/routes"
	"m/sessions"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	ID          int    `gorm:"primary_key;not null"`
	ProductName string `gorm:"type:varchar(200);not null"`
	Memo        string `gorm:"type:varchar(400)"`
	Status      string `gorm:"type:char(2);not null"`
}

func getGormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "Shopping"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		println("はんんあんっはんっはｎ")
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	db.LogMode(true)

	db.SingularTable(true)

	db.AutoMigrate(&Product{})

	fmt.Println("db connected: ", &db)
	return db
}

func insertProduct(registerProduct *Product) {
	db := getGormConnect()

	db.Create(&registerProduct)
	defer db.Close()
}

func main() {

	var product = Product{
		ProductName: "Test",
		Memo:        "memoです",
		Status:      "01",
	}

	insertProduct(&product)

	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	store := sessions.NewDummyStore()
	router.Use(sessions.StartDefaultSession(store))

	user := router.Group("/user")
	{
		user.POST("/test", routes.Test)
		user.POST("/signup", routes.UserSignUp)
		user.POST("/login", routes.UserLogIn)
	}

	router.GET("/", routes.Home)
	router.GET("/login", routes.LogIn)
	router.GET("/signup", routes.SignUp)
	router.GET("/loggedin", routes.LoggedIn)
	router.NoRoute(routes.NoRoute)

	router.Run(":8080")

}

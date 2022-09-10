package main
import (
	"database/sql"
	"time"
    "log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
)


func udr_db () {
    db, err := sql.Open("mysql", "root:mysql123@tcp(localhost:3306)/")

    if err != nil {
	   panic(err)
    }

    // See "Important settings" section.
    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)
}

func initialization(e *echo.Echo) {
    udr_db()
    db_create("subscriber_data")
	log.Print("Executing Initalization tasks")
	e.POST("/nudr-dr/v1/subscriber_data", create_subscriber_data)
	//e.GET("/users/:ueId", getUser)
	//e.PUT("/users/:ueId", updateUser)
	//e.DELETE("/users/:ueId", deleteUser)
}

// e.POST("/registerUser", registerUser)
func create_subscriber_data(c echo.Context) error {
	// Get name and email
	ueId := c.FormValue("ueId")
	return c.String(http.StatusOK, "ueId:"+ueId)
}

func db_create(name string) {

    log.Print("create database")
    db, err := sql.Open("mysql", "root:mysql123@tcp(localhost:3306)/")
    if err != nil {
        panic(err)
    }
    defer db.Close()
 
    _,err = db.Exec("CREATE DATABASE "+name)
    if err != nil {
        panic(err)
    }
 
    _,err = db.Exec("USE "+name)
    if err != nil {
        panic(err)
    }
 
    _,err = db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
    if err != nil {
        panic(err)
    }
 }

func main() {

    // Create a new Echo web server instance
	e := echo.New()
	// Initialize required handlers
	log.Print("Initialize..")
	initialization(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "5G UDR Service")
	})

	// Start web server at port 1000
	e.Logger.Fatal(e.Start(":1000"))
    
}

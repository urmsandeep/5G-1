package main
import (
    "log"
	"net/http"

	"github.com/labstack/echo/v4"
)


func initialization(e *echo.Echo) {
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

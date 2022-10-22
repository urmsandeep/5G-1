package main
import (
    "log"
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
)

type policy struct {
	Name string `json:max-data,ommitempty"`
}

type subscriberData struct {
	Name    string `json:"name,ommitempty"`
	IMSI    string `json:"imsi,ommitempty"`
	Policy  policy
}

func initialization(e *echo.Echo) {
	log.Print("Executing Initalization tasks")
	e.POST("/nudr-dr/v1/subscriber_data", create_subscriber_data)
	//e.GET("/users/:ueId", getUser)
	//e.PUT("/users/:ueId", updateUser)
	//e.DELETE("/users/:ueId", deleteUser)
}

// e.POST("/createSubscriber", createUser)
func create_subscriber_data(c echo.Context) error {
	m := echo.Map{}
    if err := c.Bind(&m); err != nil {
        return err
    }
    new_data := subscriberData{
        Name:    m["name"].(string),
        IMSI:    m["imsi"].(string),
		Policy:  m["policy"].(policy),
    }
    js, _ := json.Marshal(new_data)
	
    return c.JSON(http.StatusOK, string(js))
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

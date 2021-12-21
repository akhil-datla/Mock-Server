package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	PortPtr := flag.Int("port", 8080, "a port number to start server")
	FilePtr := flag.String("path", filepath.Join("test.txt"), "file path to file")
	flag.Parse()
	fmt.Print(Banner())
	fmt.Printf("Starting server on port %d", *PortPtr)

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())

	data, err := ioutil.ReadFile(*FilePtr)

	if err != nil {
		fmt.Println(err)
		return
	}

	e.GET("/file", func(c echo.Context) error {
		return c.String(http.StatusOK, string(data))
	})

	e.GET("/body", func(c echo.Context) error {
		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			fmt.Print(err)
		}
		return c.JSON(http.StatusOK, m)
	})

	e.GET("/header", func(c echo.Context) error {
		hMap := c.Request().Header
		return c.JSON(http.StatusOK, hMap)
	})

	portNum := fmt.Sprintf("%s%d", ":", *PortPtr)
	e.Start(portNum)
}

//Banner creates a banner
func Banner() string {
	banner := `
__  __  ___   ____ _  ______  _____ ______     _______ ____  
|  \/  |/ _ \ / ___| |/ / ___|| ____|  _ \ \   / / ____|  _ \ 
| |\/| | | | | |   | ' /\___ \|  _| | |_) \ \ / /|  _| | |_) |
| |  | | |_| | |___| . \ ___) | |___|  _ < \ V / | |___|  _ < 
|_|  |_|\___/ \____|_|\_\____/|_____|_| \_\ \_/  |_____|_| \_\
								      
Mock Server - A simple mock server
(c)2021 - Created by Akhil Datla
`
	return banner

}

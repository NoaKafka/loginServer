package main

import (
	"loginServer/handler"

	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
)

func serveHome(c echo.Context) error {

	//http.ServeFile(w, r, "home.html")
	return c.File("home.html")
}

func main() {

	// godotenv는 로컬 개발환경에서 .env를 통해 환경변수를 읽어올 때 쓰는 모듈이다.
	// 프로덕션 환경에서는 필요하지 않음.
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	e := echo.New()
	// 회원가입 API
	e.POST("/api/signup", handler.SignUp)

	e.POST("/api/signin", handler.SignIn)

	e.Logger.Fatal(e.Start(":1323"))
}

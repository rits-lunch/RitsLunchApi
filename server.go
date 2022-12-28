package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", getHello)

	e.GET("/shop", getAllShops)

	e.GET("/shop/:shopId", getShopDetail)

	e.Logger.Fatal(e.Start(":3000"))
}
func getHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello RitsLunch")
}

func getAllShops(c echo.Context) error {
	return c.String(http.StatusOK, "店舗一覧")
}

func getShopDetail(c echo.Context) error {
	id := c.Param("shopId")
	return c.String(http.StatusOK, id)
}

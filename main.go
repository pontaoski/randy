package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	rand.Seed(time.Now().Unix())

	e := echo.New()
	e.GET("/*", func(c echo.Context) error {
		urls := strings.TrimPrefix(c.Request().URL.Path, "/")
		var data struct {
			URLs []string `json:"urls"`
		}
		str, err := url.QueryUnescape(urls)
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(str), &data)
		if err != nil {
			return err
		}
		if len(data.URLs) == 0 {
			return c.String(http.StatusOK, "you cant have empty urls")
		}
		idx := rand.Intn(len(data.URLs))
		return c.Redirect(http.StatusTemporaryRedirect, data.URLs[idx])
	})
	e.Logger.Fatal(e.Start(":1515"))
}

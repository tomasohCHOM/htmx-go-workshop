package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

type Count struct {
	Count int
}

type Response struct {
	PokemonData []PokemonData `json:"results"`
}

type PokemonData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokemon struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()

	count := Count{Count: 69}

	// GET endpoint to get a list of pokemon with limit = 20
	e.GET("/", func(c echo.Context) error {
		res, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=20")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		resData, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var resObject Response
		json.Unmarshal(resData, &resObject)
		fmt.Println(resObject)

		return c.Render(http.StatusOK, "index.html", count)
	})

	// GET endpoint to get a specific pokemon based on their id number
	e.GET("/pokemon/:id", func(c echo.Context) error {
		pokemonNumber := c.Param("id")
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonNumber)
		res, err := http.Get(url)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		resData, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var resObject Pokemon
		json.Unmarshal(resData, &resObject)
		fmt.Println(resObject)

		return c.Render(http.StatusOK, "index.html", count)
	})

	e.Logger.Fatal(e.Start(":3000"))
}

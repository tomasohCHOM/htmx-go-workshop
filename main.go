package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"htmx-go-workshop/internal/pokedex"
)

type Templates struct {
	templates *template.Template
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

	// GET endpoint to get a list of Pokemon.
	e.GET("/pokemon", func(c echo.Context) error {

		return c.Render(http.StatusOK, "index.html", pokedex.GetPokedex())
	})

	// GET endpoint to get a specific pokemon based on their id number
	e.GET("/pokemon/:id", func(c echo.Context) error {
		pokedexNumberString := c.Param("id")
		pokedexNumber, err := strconv.Atoi(pokedexNumberString)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		entry, err := pokedex.GetPokemonEntry(pokedexNumber)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		return c.Render(http.StatusOK, "index.html", entry)
	})

	e.Logger.Fatal(e.Start(":3000"))
}

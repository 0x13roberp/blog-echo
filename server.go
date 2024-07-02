package main

import (
	"blog/handlers"
	"blog/models"
	"html/template"
	"io"

	"net/http"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// para renderizar los html
type TemplateRenderer struct {
	templates *template.Template
}

// render renderiza una plantilla utilizando datos
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatalf("environment variable not connected")
	}

	// conexion con la db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting with database %v", err)
	}

	// verificar conexion de db
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database %v", err)
	}
	defer sqlDB.Close() // cerrar la conexion con la db

	// migrar el esquema
	db.AutoMigrate(&models.Post{})

	e := echo.New()

	// middleware
	e.Use(middleware.Logger())  // registra los logs de las solicitudes y respuestas http
	e.Use(middleware.Recover()) // se encarga de recuperar el server si sucede un panic

	// motor de plantilla
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// handler con la base de datos
	h := &handlers.Handler{DB: db}

	// rutas
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to my blog")
	})

	e.GET("/posts", h.GetPost)

	e.GET("/posts/html", h.GetPostHtml)

	e.Logger.Fatal(e.Start(":8080"))
}

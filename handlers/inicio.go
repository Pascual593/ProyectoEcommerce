package handlers

import (
	"ProyectoEcommerce/models"
	"html/template"
	"log"
	"net/http"
)

type InicioData struct {
	Title     string
	Productos []models.Producto
}

func InicioHandler(w http.ResponseWriter, r *http.Request) {
	productos, err := models.ListarProductos()
	if err != nil {
		http.Error(w, "Error al cargar productos", http.StatusInternalServerError)
		log.Println("❌ Error al listar productos:", err)
		return
	}

	data := InicioData{
		Title:     "Bienvenido a la tienda",
		Productos: productos,
	}

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/inicio.html"))
	tmpl.ExecuteTemplate(w, "base.html", data)
}

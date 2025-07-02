package handlers

import (
	"ProyectoEcommerce/models"
	"html/template"
	"net/http"
	"strconv"
)

// DashboardData estructura los datos que se enviarán a la plantilla HTML.
// Contiene el título de la página y el nombre del usuario.
type DashboardData struct {
	Title  string
	Nombre string
	Rol    string
}

// DashboardHandler controla el acceso al panel principal.
// Solo permite el ingreso si existe una cookie de sesión válida.
// Si la sesión es válida, renderiza el dashboard mostrando el nombre del usuario.
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener la cookie de sesión
	cookie, err := r.Cookie("usuario_id")
	if err != nil {
		// Si no hay cookie, redirige al login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Convertir el valor de la cookie a entero (ID del usuario)
	id, err := strconv.Atoi(cookie.Value)
	if err != nil || id == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Obtener el nombre del usuario desde la base de datos
	nombre, rol, err := models.ObtenerNombreYRolPorID(id)
	if err != nil {
		http.Error(w, "No se pudo obtener el nombre del usuario", http.StatusInternalServerError)
		return
	}

	// Datos que se pasan a la plantilla HTML
	data := DashboardData{
		Title:  "Dashboard",
		Nombre: nombre,
		Rol:    rol,
	}

	// Cargar la plantilla base y la vista del dashboard
	tmpl, err := template.ParseFiles("templates/base.html", "templates/dashboard.html")
	if err != nil {
		http.Error(w, "Error al cargar plantilla", http.StatusInternalServerError)
		return
	}

	// Renderizar la plantilla usando "base.html" como entrada principal
	tmpl.ExecuteTemplate(w, "base.html", data)
}

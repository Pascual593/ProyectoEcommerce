package handlers

import (
	"ProyectoEcommerce/models"
	"fmt"
	"html/template"
	"net/http"
)

// LoginData estructura los datos enviados a la plantilla del login.
// En este caso, solo contiene el título de la página.
type LoginData struct {
	Title string
	Error string
}

// LoginHandler maneja la visualización y procesamiento del formulario de inicio de sesión.
// En método GET muestra el formulario, en POST valida credenciales y establece la sesión.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Si se envía el formulario (POST)
	if r.Method == http.MethodPost {
		// Obtener email y contraseña ingresados por el usuario
		email := r.FormValue("email")
		contrasena := r.FormValue("password")

		// Verificar credenciales con la función del modelo
		userID, err := models.VerificarCredenciales(email, contrasena)
		if err != nil {
			// Si las credenciales son inválidas, devolver error 401
			data := LoginData{
				Title: "Iniciar sesión",
				Error: "❌ Usuario o contraseña incorrectos. Intenta nuevamente.",
			}

			tmpl, err := template.ParseFiles("templates/base.html", "templates/login.html")
			if err != nil {
				http.Error(w, "Error al cargar plantilla", http.StatusInternalServerError)
				return
			}
			tmpl.ExecuteTemplate(w, "base.html", data)
			return
		}

		// Crear una cookie de sesión con el ID del usuario
		cookie := &http.Cookie{
			Name:  "usuario_id",              // Nombre de la cookie
			Value: fmt.Sprintf("%d", userID), // Valor: el ID del usuario convertido a texto
			Path:  "/",                       // Disponible en toda la app
		}
		// Asignar cookie al navegador
		http.SetCookie(w, cookie)

		// Redirigir al dashboard tras el login exitoso
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	// Si es GET (visita la página sin enviar formulario)
	// Preparar los datos que se pasan a la plantilla (el título en este caso)
	data := LoginData{
		Title: "Iniciar sesión",
	}

	// Cargar las plantillas base y específica del login
	tmpl, err := template.ParseFiles("templates/base.html", "templates/login.html")
	if err != nil {
		// Si hay error al cargar la plantilla, mostrar mensaje interno
		http.Error(w, "Error al cargar plantilla", http.StatusInternalServerError)
		return
	}

	// Renderizar la plantilla "base.html" con los datos del título
	tmpl.ExecuteTemplate(w, "base.html", data)
}

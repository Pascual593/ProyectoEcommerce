package handlers

import (
	"net/http"
)

// LogoutHandler elimina la cookie de sesión del usuario para cerrar su sesión.
// Luego lo redirige nuevamente al login.

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Borrar la cookie "usuario_id" seteándola con valor vacío y expiración inmediata
	cookie := &http.Cookie{
		Name:   "usuario_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Elimina la cookie
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

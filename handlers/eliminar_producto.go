package handlers

import (
	"log"
	"net/http"
	"strconv"

	"ProyectoEcommerce/models"

	"github.com/gorilla/mux"
)

// EliminarProductoHandler borra un producto según su ID
func EliminarProductoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = models.DeleteProducto(id)
	if err != nil {
		log.Println("❌ Error al eliminar producto:", err)
		http.Error(w, "No se pudo eliminar el producto", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/productos/listar", http.StatusSeeOther)
}

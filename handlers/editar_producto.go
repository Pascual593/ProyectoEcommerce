package handlers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"ProyectoEcommerce/models"

	"github.com/gorilla/mux"
)

// EditarProductoHandler maneja la visualización y actualización de un producto
func EditarProductoHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer el ID desde la URL
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		// Obtener datos del formulario
		nombre := r.FormValue("nombre")
		descripcion := r.FormValue("descripcion")
		precioStr := r.FormValue("precio")
		categoriaStr := r.FormValue("categoria_id")
		stockStr := r.FormValue("stock")
		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			http.Error(w, "Stock inválido", http.StatusBadRequest)
			return
		}

		precio, err := strconv.ParseFloat(precioStr, 64)
		if err != nil {
			http.Error(w, "Precio inválido", http.StatusBadRequest)
			return
		}

		categoriaID, err := strconv.Atoi(categoriaStr)
		if err != nil {
			http.Error(w, "Categoría inválida", http.StatusBadRequest)
			return
		}

		// Procesar imagen
		file, handler, err := r.FormFile("imagen")
		imagenNombre := r.FormValue("imagen_actual") // ← del input oculto

		if err == nil {
			defer file.Close()

			os.MkdirAll("uploads", os.ModePerm)

			imagenNombre = fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
			ruta := filepath.Join("uploads", imagenNombre)

			destino, err := os.Create(ruta)
			if err != nil {
				http.Error(w, "Error al guardar la imagen", http.StatusInternalServerError)
				return
			}
			defer destino.Close()

			_, err = io.Copy(destino, file)
			if err != nil {
				http.Error(w, "Error al guardar la imagen", http.StatusInternalServerError)
				return
			}
		}

		// Crear struct actualizado
		p := models.Producto{
			ID:          id,
			Nombre:      nombre,
			Descripcion: descripcion,
			Precio:      precio,
			CategoriaID: categoriaID,
			Imagen:      imagenNombre,
			Stock:       stock,
		}

		err = models.UpdateProducto(p)
		if err != nil {
			log.Println("❌ Error al actualizar producto:", err)
			http.Error(w, "No se pudo actualizar el producto", http.StatusInternalServerError)
			return
		}

		// Redirigir al listado tras actualizar
		http.Redirect(w, r, "/productos/listar", http.StatusSeeOther)
		return
	}

	// GET: Obtener producto por ID y cargar formulario
	producto, err := models.GetProductoByID(id)
	if err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	categorias, err := models.GetAllCategorias()
	if err != nil {
		http.Error(w, "No se pudieron cargar las categorías", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title      string
		Producto   models.Producto
		Categorias []models.Categoria
	}{
		Title:      "Editar producto",
		Producto:   producto,
		Categorias: categorias,
	}

	tmpl, err := template.ParseFiles("templates/base.html", "templates/editar_producto.html")
	if err != nil {
		http.Error(w, "Error al cargar plantilla", http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base.html", data)
}

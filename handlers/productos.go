package handlers

import (
	// ...
	"ProyectoEcommerce/models"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// ProductoFormData contiene los datos enviados a la plantilla de productos.
// Incluye título de la página y lista de categorías para el select.
type ProductoFormData struct {
	Title      string
	Categorias []models.Categoria
}

// ProductoHandler maneja el registro de nuevos productos.
// Si es GET, muestra el formulario. Si es POST, guarda el producto.
func ProductoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// 📥 Obtener datos del formulario
		nombre := r.FormValue("nombre")
		descripcion := r.FormValue("descripcion")
		precioStr := r.FormValue("precio")
		categoriaIDStr := r.FormValue("categoria_id")
		stockStr := r.FormValue("stock")
		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			http.Error(w, "Stock inválido", http.StatusBadRequest)
			return
		}

		// 🔄 Validar y convertir datos
		precio, err := strconv.ParseFloat(precioStr, 64)
		if err != nil {
			http.Error(w, "Precio inválido", http.StatusBadRequest)
			return
		}

		categoriaID, err := strconv.Atoi(categoriaIDStr)
		if err != nil {
			http.Error(w, "Categoría inválida", http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("imagen")
		var nombreArchivo string

		if err == nil && handler != nil {
			defer file.Close()

			// Crear carpeta si no existe
			os.MkdirAll("uploads", os.ModePerm)

			// Generar nombre único para la imagen
			nombreArchivo = fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
			ruta := filepath.Join("uploads", nombreArchivo)

			destino, err := os.Create(ruta)
			if err != nil {
				http.Error(w, "Error al guardar imagen", http.StatusInternalServerError)
				return
			}
			defer destino.Close()

			_, err = io.Copy(destino, file)
			if err != nil {
				http.Error(w, "Error al guardar imagen", http.StatusInternalServerError)
				return
			}
		}

		// 📦 Construir struct Producto con los datos recibidos
		producto := models.Producto{
			Nombre:      nombre,
			Descripcion: descripcion,
			Precio:      precio,
			CategoriaID: categoriaID,
			Imagen:      nombreArchivo,
			Stock:       stock,
		}

		// 💾 Guardar producto en base de datos
		err = models.InsertProducto(producto)
		if err != nil {
			log.Println("❌ Error al guardar producto:", err)
			http.Error(w, "No se pudo guardar el producto", http.StatusInternalServerError)
			return
		}

		// 🔁 Redirigir al mismo formulario tras guardar
		http.Redirect(w, r, "/productos", http.StatusSeeOther)
		return
	}

	// 🧾 Si es GET, cargar las categorías para mostrar en el select
	categorias, err := models.GetAllCategorias()
	if err != nil {
		http.Error(w, "No se pudieron cargar las categorías", http.StatusInternalServerError)
		return
	}

	// 📄 Datos que se pasarán a la plantilla HTML
	data := ProductoFormData{
		Title:      "Registrar producto",
		Categorias: categorias,
	}

	// 🧠 Renderizar la plantilla usando base.html + productos.html
	tmpl, err := template.ParseFiles("templates/base.html", "templates/productos.html")
	if err != nil {
		http.Error(w, "Error al cargar plantilla", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base.html", data)
}
func VerProductoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	producto, err := models.GetProductoByID(id)
	if err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}
	log.Println("🧪 Producto cargado:", producto.Nombre, "Stock:", producto.Stock)

	data := struct {
		Title    string
		Producto models.Producto
	}{
		Title:    "Detalle de producto",
		Producto: producto,
	}

	tmpl, _ := template.ParseFiles("templates/base.html", "templates/ver_producto.html")
	tmpl.ExecuteTemplate(w, "base.html", data)
}

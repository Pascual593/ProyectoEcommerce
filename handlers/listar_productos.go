package handlers

import (
	"ProyectoEcommerce/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Función para resaltar coincidencias
func resaltarCoincidencia(texto, query string) template.HTML {
	if query == "" {
		return template.HTML(template.HTMLEscapeString(texto))
	}

	queryLower := strings.ToLower(query)
	textoLower := strings.ToLower(texto)

	// Encontrar posición
	index := strings.Index(textoLower, queryLower)
	if index == -1 {
		return template.HTML(template.HTMLEscapeString(texto))
	}

	// Dividir texto en partes
	inicio := texto[:index]
	coincidencia := texto[index : index+len(query)]
	final := texto[index+len(query):]

	// Escapar partes por separado para evitar XSS
	highlighted := template.HTMLEscapeString(inicio) +
		"<mark>" + template.HTMLEscapeString(coincidencia) + "</mark>" +
		template.HTMLEscapeString(final)

	return template.HTML(highlighted)
}

// ListarProductosHandler muestra todos los productos en tabla
func ListarProductosHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "carrito")
	usuarioID, ok := session.Values["usuario_id"].(int)

	var rol string
	if ok {
		nombre, rolObtenido, err := models.ObtenerNombreYRolPorID(usuarioID)
		if err != nil {
			log.Println("❌ Error al obtener el rol:", err)
		} else {
			log.Printf("🧪 Nombre: %s | Rol detectado: %s", nombre, rolObtenido)
			rol = rolObtenido
		}
	}
	log.Println("🧪 Rol detectado:", rol)
	query := r.URL.Query().Get("q")
	categoriaStr := r.URL.Query().Get("categoria")

	categoriaID, _ := strconv.Atoi(categoriaStr) // ← Si no hay, queda 0

	// Obtener productos según los filtros
	var productos []models.Producto
	var err error
	if query != "" || categoriaStr != "" {
		productos, err = models.BuscarProductosFiltrado(query, categoriaID)
	} else {
		productos, err = models.GetAllProductosConCategoria()
	}
	if err != nil {
		http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
		return
	}

	// Obtener todas las categorías para el <select>
	categorias, err := models.GetAllCategorias()
	if err != nil {
		http.Error(w, "Error al cargar categorías", http.StatusInternalServerError)
		return
	}

	// Enviar todo al template
	data := struct {
		Title                 string
		Query                 string
		Categorias            []models.Categoria
		CategoriaSeleccionada int
		Productos             []models.Producto
		Rol                   string
	}{
		Title:                 "Lista de productos",
		Query:                 query,
		Categorias:            categorias,
		CategoriaSeleccionada: categoriaID,
		Productos:             productos,
		Rol:                   rol,
	}

	funcMap := template.FuncMap{
		"resaltar": resaltarCoincidencia,
		"add":      func(a, b int) int { return a + b },
		"sub":      func(a, b int) int { return a - b },
	}

	tmpl, err := template.New("base.html").Funcs(funcMap).
		ParseFiles("templates/base.html", "templates/listar_productos.html")
	if err != nil {
		http.Error(w, "Error al cargar plantilla", http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base.html", data)
}

package handlers

import (
	"ProyectoEcommerce/models"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("clave-secreta-segura"))

func init() {
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 24, // 1 día en segundos
		HttpOnly: true,
	}
}

func AgregarAlCarritoHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "carrito")

	vars := mux.Vars(r)
	idStr := vars["id"]
	_, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	cantidadStr := r.FormValue("cantidad")
	cantidad, _ := strconv.Atoi(cantidadStr)
	if cantidad < 1 {
		cantidad = 1
	}

	// 🔄 Reconstrucción del carrito si ya existe
	carrito := make(map[string]int)
	if val, ok := session.Values["carrito"].(map[interface{}]interface{}); ok {
		for k, v := range val {
			ks, kOk := k.(string)
			vi, vOk := v.(int)
			if kOk && vOk {
				carrito[ks] = vi
			}
		}
	}

	// 🛒 Agregamos el nuevo producto
	carrito[idStr] += cantidad

	// 💾 Convertimos a formato compatible con la sesión
	storeMap := make(map[string]interface{})
	for k, v := range carrito {
		storeMap[k] = v
	}
	session.Values["carrito"] = storeMap
	if err := session.Save(r, w); err != nil {
		log.Println("❌ Error al guardar la sesión:", err)
	}

	log.Println("🛒 Carrito en sesión después de agregar:", session.Values["carrito"])
	log.Println("🔐 Todas las variables en sesión:", session.Values)

	http.Redirect(w, r, "/carrito/ver", http.StatusSeeOther)
}

func VerCarritoHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "carrito")
	raw := session.Values["carrito"]
	carrito := make(map[string]int)

	if rawMap, ok := raw.(map[string]interface{}); ok {
		log.Println("📦 Convertido desde map[string]interface{}")
		for k, v := range rawMap {
			if vi, ok := v.(int); ok {
				carrito[k] = vi
			}
		}
	} else if casted, ok := raw.(map[string]int); ok {
		log.Println("📦 Convertido desde map[string]int")
		carrito = casted
	} else {
		log.Println("⚠️ No se pudo convertir el carrito:", raw)
	}

	var productos []models.Producto
	var total float64

	for idStr, cantidad := range carrito {
		id, _ := strconv.Atoi(idStr)
		prod, err := models.GetProductoByID(id)
		if err != nil {
			log.Println("❌ No se encontró el producto con ID:", id, ":", err)
			continue // ignorar productos que no existan
		}
		log.Println("✔️ Producto recuperado del modelo:", prod.Nombre, "Cantidad:", cantidad)
		prod.Stock = cantidad // usamos .Stock como "cantidad en carrito" aquí
		productos = append(productos, prod)
		total += prod.Precio * float64(cantidad)
	}

	data := struct {
		Title     string
		Productos []models.Producto
		Total     float64
	}{
		Title:     "Carrito de Compras",
		Productos: productos,
		Total:     total,
	}
	funcMap := template.FuncMap{
		"mulFloat": func(a float64, b int) float64 {
			return a * float64(b)
		},
	}
	tmpl, err := template.New("base.html").
		Funcs(funcMap).
		ParseFiles("templates/base.html", "templates/ver_carrito.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla del carrito", http.StatusInternalServerError)
		return
	}
	log.Println("✅ Llegamos al render del carrito")
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("❌ Error al ejecutar plantilla:", err)
	}
}
func EliminarDelCarritoHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "carrito")
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Reconstruir carrito
	carrito := map[string]interface{}{}
	if session.Values["carrito"] != nil {
		carrito = session.Values["carrito"].(map[string]interface{})
	}

	// Borrar el producto
	delete(carrito, idStr)
	session.Values["carrito"] = carrito
	session.Save(r, w)

	http.Redirect(w, r, "/carrito/ver", http.StatusSeeOther)
}
func VaciarCarritoHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "carrito")
	session.Values["carrito"] = make(map[string]interface{})
	session.Save(r, w)
	http.Redirect(w, r, "/carrito/ver", http.StatusSeeOther)
}

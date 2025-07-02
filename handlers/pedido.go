package handlers

import (
	"ProyectoEcommerce/database"
	"ProyectoEcommerce/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Asegurate de tener esta variable en tu archivo handlers/sesion.go o similar

func ConfirmarPedidoHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "carrito")
	raw := session.Values["carrito"]

	// Reconstruir el carrito
	carrito := make(map[string]int)
	if rawMap, ok := raw.(map[string]interface{}); ok {
		for k, v := range rawMap {
			if vi, ok := v.(int); ok {
				carrito[k] = vi
			}
		}
	}

	if len(carrito) == 0 {
		http.Redirect(w, r, "/carrito/ver", http.StatusSeeOther)
		return
	}

	// Conexión temporal a la base de datos
	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Error de conexión a la base de datos", http.StatusInternalServerError)
		log.Println("❌ Conexión fallida:", err)
		return
	}
	defer db.Close()

	// Calcular totales y preparar detalles
	var total float64
	var detalles []models.DetallePedido

	for idStr, cantidad := range carrito {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("❌ ID inválido en carrito:", idStr)
			continue
		}

		producto, err := models.GetProductoByID(id)
		if err != nil {
			log.Println("❌ Error al obtener producto:", err)
			continue
		}

		subtotal := float64(cantidad) * producto.Precio
		total += subtotal

		detalles = append(detalles, models.DetallePedido{
			ProductoID: producto.ID,
			Cantidad:   cantidad,
			Precio:     producto.Precio, // ⚠️ Usás Precio en lugar de PrecioUnitario
		})
	}

	// Registrar el pedido
	res, err := db.Exec(`INSERT INTO pedidos (fecha, total, estado) VALUES (NOW(), ?, 'Pendiente')`, total)
	if err != nil {
		http.Error(w, "Error al registrar el pedido", http.StatusInternalServerError)
		log.Println("❌ Insert pedido:", err)
		return
	}

	pedidoID, _ := res.LastInsertId()

	// Insertar detalles del pedido
	for _, detalle := range detalles {
		_, err := db.Exec(`
			INSERT INTO detalles_pedido (orden_id, producto_id, cantidad, precio_unitario)
			VALUES (?, ?, ?, ?)`,
			pedidoID, detalle.ProductoID, detalle.Cantidad, detalle.Precio)
		if err != nil {
			http.Error(w, "Error al registrar detalle", http.StatusInternalServerError)
			log.Println("❌ Insert detalle:", err)
			return
		}
	}

	// Limpiar carrito
	delete(session.Values, "carrito")
	session.Save(r, w)

	http.Redirect(w, r, "/pedido/confirmado", http.StatusSeeOther)
}

func PedidoConfirmadoHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/pedido_confirmado.html"))
	tmpl.Execute(w, nil)
}

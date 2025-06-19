package models

//autor: Pascual Campos
//fecha: 15/06/2025
//tema Proyecto Ecommerce
//Avance de proyecto definicion de struct, funciones , conexion de modelo con la bse de datos
import (
	"ProyectoEcommerce/database"
	"database/sql"
	"log"
)

type DetallePedido struct {
	ID         int
	PedidoID   int // Relación con pedidos
	ProductoID int // Relación con productos
	Cantidad   int
	Precio     float64
}

// Funcion para ingresar detalle del pedido
func InsertDetallePedido(pedidoID, productoID, cantidad int, precio float64) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	// Validar que el pedido y el producto existen
	var existePedido, existeProducto bool
	DB.QueryRow("SELECT EXISTS(SELECT 1 FROM pedidos WHERE id = ?)", pedidoID).Scan(&existePedido)
	DB.QueryRow("SELECT EXISTS(SELECT 1 FROM productos WHERE id = ?)", productoID).Scan(&existeProducto)

	if !existePedido || !existeProducto {
		log.Println(" Pedido o producto no existen:", pedidoID, productoID)
		return nil
	}

	// Insertar el detalle del pedido
	query := "INSERT INTO detalle_pedido (pedido_id, producto_id, cantidad, precio) VALUES (?, ?, ?, ?)"
	_, err = DB.Exec(query, pedidoID, productoID, cantidad, precio)
	if err != nil {
		log.Println(" Error al insertar detalle del pedido:", err)
		return err
	}

	log.Println(" Producto agregado al pedido:", pedidoID)
	return nil
}

// funcion para busqueda de registro por id
func GetDetallePedidoById(id int) (DetallePedido, error) {
	var detalle DetallePedido
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return detalle, err
	}
	defer DB.Close()

	// Consultar detalle de pedido por ID
	query := "SELECT id, pedido_id, producto_id, cantidad, precio FROM detalle_pedido WHERE id = ?"
	row := DB.QueryRow(query, id)
	err = row.Scan(&detalle.ID, &detalle.PedidoID, &detalle.ProductoID, &detalle.Cantidad, &detalle.Precio)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(" No se encontró ningún detalle de pedido con ID:", id)
			return detalle, nil
		}
		log.Println(" Error al obtener detalle de pedido:", err)
		return detalle, err
	}

	log.Println(" Detalle de pedido obtenido correctamente, ID:", detalle.ID)
	return detalle, nil
}

// funcion para mostrar todos los detalles de pedido
func GetAllDetallesByPedido(pedidoID int) ([]DetallePedido, error) {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return nil, err
	}
	defer DB.Close()

	// Consultar todos los detalles de un pedido
	query := "SELECT id, pedido_id, producto_id, cantidad, precio FROM detalle_pedido WHERE pedido_id = ?"
	rows, err := DB.Query(query, pedidoID)
	if err != nil {
		log.Println(" Error al obtener detalles de pedido:", err)
		return nil, err
	}
	defer rows.Close()

	var detalles []DetallePedido
	for rows.Next() {
		var detalle DetallePedido
		err = rows.Scan(&detalle.ID, &detalle.PedidoID, &detalle.ProductoID, &detalle.Cantidad, &detalle.Precio)
		if err != nil {
			log.Println(" Error al leer detalle de pedido:", err)
			return nil, err
		}
		detalles = append(detalles, detalle)
	}

	if len(detalles) == 0 {
		log.Println(" No se encontraron productos en el pedido:", pedidoID)
	}

	log.Println(" Detalles de pedido obtenidos correctamente:", len(detalles))
	return detalles, nil
}

// funcion para editar detalle del pedido
func UpdateDetallePedido(id int, nuevaCantidad int, nuevoPrecio float64) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println("❌ Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	// Verificar que la cantidad sea válida
	if nuevaCantidad <= 0 {
		log.Println("⚠️ La cantidad debe ser mayor a 0")
		return nil
	}

	// Actualizar la cantidad y el precio del producto en el pedido
	query := "UPDATE detalle_pedido SET cantidad = ?, precio = ? WHERE id = ?"
	result, err := DB.Exec(query, nuevaCantidad, nuevoPrecio, id)
	if err != nil {
		log.Println(" Error al actualizar detalle de pedido:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(" Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println(" No se encontró ningún detalle de pedido con ID:", id)
		return nil
	}

	log.Println(" Detalle de pedido actualizado correctamente:", id, "→ Cantidad:", nuevaCantidad, "→ Precio:", nuevoPrecio)
	return nil
}

// eliminar el detalle de pedido de un producto
func DeleteDetallePedido(id int) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	// Preparar la consulta SQL para eliminar el detalle del pedido
	query := "DELETE FROM detalle_pedido WHERE id = ?"
	result, err := DB.Exec(query, id)
	if err != nil {
		log.Println(" Error al eliminar detalle de pedido:", err)
		return err
	}

	// Verificar que realmente se eliminó un detalle
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(" Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println(" No se encontró ningún detalle de pedido con ID:", id)
		return nil
	}

	log.Println(" Detalle de pedido eliminado correctamente:", id)
	return nil
}

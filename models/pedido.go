package models

//autor: Pascual Campos
//fecha: 15/06/2025
//tema Proyecto Ecommerce
//Avance de proyecto definicion de struct, funciones , conexion de modelo con la bse de datos
import (
	"ProyectoEcommerce/database"
	"database/sql"
	"log"
	"time"
)

type Pedido struct {
	ID           int
	UsuarioID    int
	Fecha        time.Time
	Total        float64
	Estado       string // Estado del pedido (pendiente, procesando, enviado, entregado)
	MetodoPagoID int    // Relación con métodos de pago
}

// funcion para el ingreso de un nuevo pedido
func InsertPedido(usuarioID int, total float64, metodoPagoID int) (int, error) {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return 0, err
	}
	defer DB.Close()

	fecha := time.Now()
	estado := "pendiente"

	query := "INSERT INTO pedidos (usuario_id, fecha, total, estado, metodo_pago_id) VALUES (?, ?, ?, ?, ?)"
	result, err := DB.Exec(query, usuarioID, fecha, total, estado, metodoPagoID)
	if err != nil {
		log.Println(" Error al registrar pedido:", err)
		return 0, err
	}

	// Obtener el ID del pedido recién insertado
	pedidoID, err := result.LastInsertId()
	if err != nil {
		log.Println(" Error al obtener ID del pedido:", err)
		return 0, err
	}

	log.Println(" Pedido registrado correctamente, ID:", pedidoID)
	return int(pedidoID), nil
}

// funcion para obtener el registro de un pedido por id
func GetPedidoById(id int) (Pedido, error) {
	var pedido Pedido
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return pedido, err
	}
	defer DB.Close()

	// Consultar pedido por ID
	query := "SELECT id, usuario_id, fecha, total, estado, metodo_pago_id FROM pedidos WHERE id = ?"
	row := DB.QueryRow(query, id)
	err = row.Scan(&pedido.ID, &pedido.UsuarioID, &pedido.Fecha, &pedido.Total, &pedido.Estado, &pedido.MetodoPagoID)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(" No se encontró ningún pedido con ID:", id)
			return pedido, nil
		}
		log.Println(" Error al obtener pedido:", err)
		return pedido, err
	}

	log.Println(" Pedido obtenido correctamente, ID:", pedido.ID)
	return pedido, nil
}

// funcion para listar todos los registros de los pedidos
func GetAllPedidos(usuarioID int) ([]Pedido, error) {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return nil, err
	}
	defer DB.Close()

	// Consultar todos los pedidos de un usuario
	query := "SELECT id, usuario_id, fecha, total, estado, metodo_pago_id FROM pedidos WHERE usuario_id = ?"
	rows, err := DB.Query(query, usuarioID)
	if err != nil {
		log.Println(" Error al obtener pedidos:", err)
		return nil, err
	}
	defer rows.Close()

	var pedidos []Pedido
	for rows.Next() {
		var pedido Pedido
		err = rows.Scan(&pedido.ID, &pedido.UsuarioID, &pedido.Fecha, &pedido.Total, &pedido.Estado, &pedido.MetodoPagoID)
		if err != nil {
			log.Println(" Error al leer pedido:", err)
			return nil, err
		}
		pedidos = append(pedidos, pedido)
	}

	if len(pedidos) == 0 {
		log.Println(" No hay pedidos registrados para el usuario:", usuarioID)
	}

	log.Println(" Pedidos obtenidos correctamente:", len(pedidos))
	return pedidos, nil
}

// funcion para actualizar el estado de un pedido
func UpdateEstadoPedido(id int, nuevoEstado string) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	// Validar que el estado sea uno de los permitidos
	estadosPermitidos := map[string]bool{
		"pendiente":  true,
		"procesando": true,
		"enviado":    true,
		"entregado":  true,
		"cancelado":  true,
	}

	if !estadosPermitidos[nuevoEstado] {
		log.Println(" Estado inválido:", nuevoEstado)
		return nil
	}

	// Actualizar el estado del pedido
	query := "UPDATE pedidos SET estado = ? WHERE id = ?"
	result, err := DB.Exec(query, nuevoEstado, id)
	if err != nil {
		log.Println(" Error al actualizar estado del pedido:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(" Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println(" No se encontró ningún pedido con ID:", id)
		return nil
	}

	log.Println(" Estado del pedido actualizado correctamente:", id, "→", nuevoEstado)
	return nil
}

// funcion para eliminar el registro de un pedido
func DeletePedido(id int) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	// Preparar la consulta SQL para eliminar el pedido
	query := "DELETE FROM pedidos WHERE id = ?"
	result, err := DB.Exec(query, id)
	if err != nil {
		log.Println(" Error al eliminar pedido:", err)
		return err
	}

	// Verificar que realmente se eliminó un pedido
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(" Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println(" No se encontró ningún pedido con ID:", id)
		return nil
	}

	log.Println(" Pedido eliminado correctamente:", id)
	return nil
}

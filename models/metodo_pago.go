package models

//autor: Pascual Campos
//fecha: 15/06/2025
//tema Proyecto Ecommerce
//Avance de proyecto definicion de struct, funciones , conexion de modelo con la bse de datos
import (
	"ProyectoEcommerce/database"
	"log"
)

type MetodoPago struct {
	ID     int
	Nombre string // Ej: "Tarjeta de crédito", "PayPal", "Transferencia bancaria"
	Activo bool   // Indica si el método de pago está disponible
}

// funcion para ingresar un nuevo registro de metodo de pago
func InsertMetodoPago(nombre string, activo bool) (int, error) {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return 0, err
	}
	defer DB.Close()

	query := "INSERT INTO metodos_pago (nombre, activo) VALUES (?, ?)"
	result, err := DB.Exec(query, nombre, activo)
	if err != nil {
		log.Println(" Error al registrar método de pago:", err)
		return 0, err
	}

	// Obtener el ID del método de pago recién insertado
	metodoPagoID, err := result.LastInsertId()
	if err != nil {
		log.Println(" Error al obtener ID del método de pago:", err)
		return 0, err
	}

	log.Println(" Método de pago registrado correctamente, ID:", metodoPagoID)
	return int(metodoPagoID), nil
}

// Funcion para modificar metodo de pago
func UpdateMetodoPago(id int, nuevoNombre string, activo bool) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	// Actualizar el nombre y estado del método de pago
	query := "UPDATE metodos_pago SET nombre = ?, activo = ? WHERE id = ?"
	result, err := DB.Exec(query, nuevoNombre, activo, id)
	if err != nil {
		log.Println(" Error al actualizar método de pago:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(" Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println(" No se encontró ningún método de pago con ID:", id)
		return nil
	}

	log.Println(" Método de pago actualizado correctamente:", id, "→ Nombre:", nuevoNombre, "→ Activo:", activo)
	return nil
}

//Funcion para eliminar un metodo de pago mediante el id
//Confirmar si realmente se eliminó algo.
//Registrar cualquier error durante el proceso.

func DeleteMetodoPago(id int) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	query := "DELETE FROM metodos_pago WHERE id = ?"
	result, err := DB.Exec(query, id)
	if err != nil {
		log.Println(" Error al eliminar método de pago:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(" Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println(" No se encontró ningún método de pago con ID:", id)
		return nil
	}

	log.Println(" Método de pago eliminado correctamente:", id)
	return nil
}

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

type Producto struct {
	ID          int
	Nombre      string
	Descripcion string
	Precio      float64
	Stock       int
	CategoriaID int
}

// Funcion para ingresar un nuevo registro de un producto
func InsertProducto(nombre, descripcion string, precio float64, stock int, categoriaID int) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	query := "INSERT INTO productos (nombre, descripcion, precio, stock, categoria_id) VALUES (?, ?, ?, ?, ?)"
	_, err = DB.Exec(query, nombre, descripcion, precio, stock, categoriaID)
	if err != nil {
		log.Println("Error al insertar producto:", err)
		return err
	}

	log.Println(" Producto registrado correctamente:", nombre)
	return nil
}

// funcion para obtener el registro de un producto por id
func GetProductoById(id int) (Producto, error) {
	var producto Producto
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return producto, err
	}
	defer DB.Close()

	// Consultar producto por ID
	query := "SELECT id, nombre, descripcion, precio, stock, categoria_id FROM productos WHERE id = ?"
	row := DB.QueryRow(query, id)
	err = row.Scan(&producto.ID, &producto.Nombre, &producto.Descripcion, &producto.Precio, &producto.Stock, &producto.CategoriaID)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(" No se encontró ningún producto con ID:", id)
			return producto, nil
		}
		log.Println(" Error al obtener producto:", err)
		return producto, err
	}

	log.Println(" Producto obtenido correctamente:", producto.Nombre)
	return producto, nil
}

// funcion para para editar los datos de un producto
func UpdateProducto(id int, nombre, descripcion string, precio float64, stock int, categoriaID int) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	// Preparar la consulta SQL para actualizar el producto
	query := "UPDATE productos SET nombre = ?, descripcion = ?, precio = ?, stock = ?, categoria_id = ? WHERE id = ?"
	_, err = DB.Exec(query, nombre, descripcion, precio, stock, categoriaID, id)
	if err != nil {
		log.Println(" Error al actualizar producto:", err)
		return err
	}

	log.Println(" Producto actualizado correctamente:", id)
	return nil
}

// funcion para eliminar el registro de un producto
func DeleteProducto(id int) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	// Preparar la consulta SQL para eliminar el producto
	query := "DELETE FROM productos WHERE id = ?"
	result, err := DB.Exec(query, id)
	if err != nil {
		log.Println(" Error al eliminar producto:", err)
		return err
	}

	// Verificar que realmente se eliminó un producto
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(" Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println(" No se encontró ningún producto con ID:", id)
		return nil
	}

	log.Println(" Producto eliminado correctamente:", id)
	return nil
}

// funcion para listar el registro de todos los productos
func GetAllProductos() ([]Producto, error) {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return nil, err
	}
	defer DB.Close()

	// Consultar todos los productos
	query := "SELECT id, nombre, descripcion, precio, stock, categoria_id FROM productos"
	rows, err := DB.Query(query)
	if err != nil {
		log.Println(" Error al obtener productos:", err)
		return nil, err
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var producto Producto
		err = rows.Scan(&producto.ID, &producto.Nombre, &producto.Descripcion, &producto.Precio, &producto.Stock, &producto.CategoriaID)
		if err != nil {
			log.Println(" Error al leer producto:", err)
			return nil, err
		}
		productos = append(productos, producto)
	}

	if len(productos) == 0 {
		log.Println(" No hay productos registrados en la base de datos")
	}

	log.Println(" Productos obtenidos correctamente:", len(productos))
	return productos, nil
}

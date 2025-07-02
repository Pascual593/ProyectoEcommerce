package models

import (
	"ProyectoEcommerce/database"
	"database/sql"
	"log"
)

// Producto representa un producto de la tienda.
// Incluye campos compatibles con la tabla `productos`.
type Producto struct {
	ID              int
	Nombre          string
	Descripcion     string
	Precio          float64
	CategoriaID     int
	CategoriaNombre string // ← nuevo campo para mostrar en la tabla
	Imagen          string
	Stock           int
}

// InsertProducto guarda un nuevo producto en la base de datos.
// Recibe una instancia de Producto con los datos completos.
func InsertProducto(p Producto) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	// Sentencia SQL para insertar un nuevo producto
	query := `INSERT INTO productos (nombre, descripcion, precio, categoria_id, imagen, stock)
          VALUES (?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(query, p.Nombre, p.Descripcion, p.Precio, p.CategoriaID, p.Imagen, p.Stock)
	if err != nil {
		log.Println("❌ Error al insertar producto:", err)
	}
	return err
}

// GetAllProductos devuelve todos los productos registrados en la base de datos
func GetAllProductos() ([]Producto, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
	SELECT p.id, p.nombre, p.descripcion, p.precio, p.categoria_id, p.stock
	FROM productos p
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []Producto

	for rows.Next() {
		var p Producto
		err := rows.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.CategoriaID, &p.Stock)
		if err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}

	return productos, nil
}

// GetAllProductosConCategoria trae productos junto a su categoría (JOIN)
func GetAllProductosConCategoria() ([]Producto, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
	SELECT p.id, p.nombre, p.descripcion, p.precio, p.imagen, c.id, c.nombre, p.stock
	FROM productos p
	JOIN categorias c ON p.categoria_id = c.id
	ORDER BY p.id DESC
`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []Producto

	for rows.Next() {
		var p Producto
		var imagen sql.NullString
		err := rows.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &imagen, &p.CategoriaID, &p.CategoriaNombre, &p.Stock)
		if err != nil {
			log.Println("Error al escanear producto:", err)
			return nil, err
		}
		if imagen.Valid {
			p.Imagen = imagen.String
		} else {
			p.Imagen = "" // sin imagen
		}

		productos = append(productos, p)
	}

	return productos, nil
}

// GetProductoByID devuelve los datos de un producto según su ID
func GetProductoByID(id int) (Producto, error) {
	db, err := database.Connect()
	if err != nil {
		return Producto{}, err
	}
	defer db.Close()

	query := `
		SELECT p.id, p.nombre, p.descripcion, p.precio, p.stock, p.categoria_id, c.nombre, p.imagen
		FROM productos p
		JOIN categorias c ON p.categoria_id = c.id
		WHERE p.id = ?
	`

	row := db.QueryRow(query, id)

	var p Producto
	var imagen sql.NullString

	err = row.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.Stock, &p.CategoriaID, &p.CategoriaNombre, &imagen)
	if err != nil {
		return Producto{}, err
	}

	if imagen.Valid {
		p.Imagen = imagen.String
	} else {
		p.Imagen = ""
	}

	return p, nil
}

// UpdateProducto actualiza los datos de un producto existente
func UpdateProducto(p Producto) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
	UPDATE productos
	SET nombre = ?, descripcion = ?, precio = ?, categoria_id = ?, stock = ?
	WHERE id = ?
`
	_, err = db.Exec(query, p.Nombre, p.Descripcion, p.Precio, p.CategoriaID, p.Stock, p.ID)
	return err
}

// DeleteProducto elimina un producto de la base de datos por ID
func DeleteProducto(id int) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM productos WHERE id = ?"
	_, err = db.Exec(query, id)
	return err
}

func BuscarProductos(query string) ([]Producto, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	querySQL := `
		SELECT p.id, p.nombre, p.descripcion, p.precio, c.id, c.nombre, p.stock
		FROM productos p
		JOIN categorias c ON p.categoria_id = c.id
		WHERE LOWER(p.nombre) LIKE LOWER(?) 
  		OR LOWER(p.descripcion) LIKE LOWER(?)
   		OR LOWER(c.nombre) LIKE LOWER(?)
		ORDER BY p.id DESC
	`

	like := "%" + query + "%"
	rows, err := db.Query(querySQL, like, like, like)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var p Producto
		err := rows.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.CategoriaID, &p.CategoriaNombre, &p.Stock)
		if err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}

	return productos, nil
}
func BuscarProductosFiltrado(query string, categoriaID int) ([]Producto, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	querySQL := `
		SELECT p.id, p.nombre, p.descripcion, p.precio, c.id, c.nombre, p.stock
		FROM productos p
		JOIN categorias c ON p.categoria_id = c.id
		WHERE (
			LOWER(p.nombre) LIKE LOWER(?)
			OR LOWER(p.descripcion) LIKE LOWER(?)
			OR LOWER(c.nombre) LIKE LOWER(?)
		)
	`

	params := []interface{}{"%" + query + "%", "%" + query + "%", "%" + query + "%"}

	if categoriaID > 0 {
		querySQL += " AND c.id = ?"
		params = append(params, categoriaID)
	}

	querySQL += " ORDER BY p.id DESC"

	rows, err := db.Query(querySQL, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var p Producto
		err := rows.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.CategoriaID, &p.CategoriaNombre, &p.Stock)
		if err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}

	return productos, nil
}
func ListarProductos() ([]Producto, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
		SELECT p.id, p.nombre, p.descripcion, p.precio, c.id, c.nombre, p.imagen, p.stock
		FROM productos p
		JOIN categorias c ON p.categoria_id = c.id
		ORDER BY p.id DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []Producto

	for rows.Next() {
		var p Producto
		var imagen sql.NullString

		err := rows.Scan(
			&p.ID, &p.Nombre, &p.Descripcion, &p.Precio,
			&p.CategoriaID, &p.CategoriaNombre, &imagen, &p.Stock,
		)
		if err != nil {
			log.Println("❌ Error al escanear producto:", err)
			continue
		}

		p.Imagen = ""
		if imagen.Valid {
			p.Imagen = imagen.String
		}

		productos = append(productos, p)
	}

	return productos, nil
}

package models

import (
	"ProyectoEcommerce/database"
)

// Categoria representa una categoría de productos
type Categoria struct {
	ID     int
	Nombre string
}

// GetAllCategorias obtiene todas las categorías desde la base de datos
func GetAllCategorias() ([]Categoria, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, nombre FROM categorias ORDER BY nombre ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []Categoria
	for rows.Next() {
		var c Categoria
		if err := rows.Scan(&c.ID, &c.Nombre); err != nil {
			return nil, err
		}
		categorias = append(categorias, c)
	}

	return categorias, nil
}

package models

//autor: Pascual Campos
//fecha: 15/06/2025
//tema Proyecto Ecommerce
//Avance de proyecto definicion de struct, funciones , conexion de modelo con la bse de datos
import (
	"ProyectoEcommerce/database"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Usuarios struct {
	ID         int
	Nombre     string
	Email      string
	Contraseña string
	Direccion  string
	Telefono   string
}

// funcion para obtener usuario por id
func GetUsuariosById(id int) (Usuarios, error) {
	var usuarios Usuarios
	DB, err := database.Connect()
	if err != nil {
		log.Println("Error al conectar a la base de datos: ", err)
		return usuarios, err
	}
	defer DB.Close()

	//consultar a la bd con respecto al id
	stmt, err := DB.Prepare("SELECT id, nombre, email, contraseña, direccion, telefono FROM usuarios WHERE id = ?")
	if err != nil {
		log.Println("Error al preparar la consulta:", err)
		return usuarios, err
	}
	defer stmt.Close()
	//obtenida la consulta la enviamos hacia el template
	row := stmt.QueryRow(id)
	err = row.Scan(&usuarios.ID, &usuarios.Nombre, &usuarios.Email, &usuarios.Contraseña, &usuarios.Direccion, &usuarios.Telefono)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Nose encontro el usuario con el id:", id)
			return usuarios, nil
		}
		log.Println("Error al obtener el usuario:", err)
		return usuarios, nil

	}
	log.Println("usuario obtenido: ", usuarios)
	return usuarios, nil
}

// crear metodo para obtener todos los usuarios
func GetAllUsuarios() ([]Usuarios, error) {
	//declaramos la variable
	var usuario []Usuarios
	//obtenemos la conexion
	DB, err := database.Connect()
	if err != nil {
		log.Println("Error al conectar a la base de datos:", err)
		return usuario, err
	}
	//cerramos la conexion
	defer DB.Close()
	//obtenemos todos los usuarios
	filas, err := DB.Query("SELECT id, nombre, email, contraseña, direccion, telefono FROM usuarios")
	if err != nil {
		log.Println("Error al preparar la consulta:", err)
		return usuario, err
	}
	//cerramos la conexion
	defer filas.Close()
	//obtenemos todos los usuarios
	for filas.Next() {
		var usuarios Usuarios
		err = filas.Scan(&usuarios.ID, &usuarios.Nombre, &usuarios.Email, &usuarios.Contraseña, &usuarios.Direccion, &usuarios.Telefono)
		if err != nil {
			log.Println("Error al obtener el usuario:", err)
			return usuario, err

		}
		usuario = append(usuario, usuarios)
	}
	//verificar si hay algun error en el for
	if err = filas.Err(); err != nil {
		log.Println("Error al obtener los usuarios:", err)
		return usuario, err
	}
	log.Println("usuarios obtenidas:", usuario)
	return usuario, nil

}

/*
// funcion para insertar un usuario
func InsertUsuario(nombre, email, contraseña, direccion, telefono string) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println("Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	// Asignamos valores predeterminados
	rol := "cliente" // Por defecto, los usuarios son clientes
	query := "INSERT INTO usuarios (nombre, email, contraseña, direccion, telefono, rol, creado_en) VALUES (?, ?, ?, ?, ?, ?, NOW())"

	_, err = DB.Exec(query, nombre, email, contraseña, direccion, telefono, rol)
	if err != nil {
		log.Println("Error al insertar usuario:", err)
		return err
	}

	log.Println("Usuario registrado correctamente:", email)
	return nil
}
*/

// funcion para actualizar los datos del usuario
func UpdateUsuario(id int, nombre, email, contraseña, direccion, telefono, rol string) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println("Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	// Preparar la consulta SQL para actualizar el usuario
	query := "UPDATE usuarios SET nombre = ?, email = ?, contraseña = ?, direccion = ?, telefono = ?, rol = ? WHERE id = ?"
	_, err = DB.Exec(query, nombre, email, contraseña, direccion, telefono, rol, id)
	if err != nil {
		log.Println("Error al actualizar usuario:", err)
		return err
	}

	log.Println(" Usuario actualizado correctamente:", id)
	return nil
}

// funcion para eliminar el registro de un usuario
func DeleteUsuario(id int) error {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return err
	}
	defer DB.Close()

	// Preparar la consulta SQL para eliminar el usuario
	query := "DELETE FROM usuarios WHERE id = ?"
	result, err := DB.Exec(query, id)
	if err != nil {
		log.Println(" Error al eliminar usuario:", err)
		return err
	}

	// Verificar que realmente se eliminó un usuario
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(" Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println(" No se encontró ningún usuario con el ID:", id)
		return nil
	}

	log.Println(" Usuario eliminado correctamente:", id)
	return nil
}

// funcion para loguearse mediante el email verificando contraseña
func LoginUsuario(email, contraseña string) (Usuarios, error) {
	var usuario Usuarios
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return usuario, err
	}
	defer DB.Close()

	// Buscar usuario en la base de datos
	query := "SELECT id, nombre, email, contraseña, direccion, telefono FROM usuarios WHERE email = ?"
	row := DB.QueryRow(query, email)
	err = row.Scan(&usuario.ID, &usuario.Nombre, &usuario.Email, &usuario.Contraseña, &usuario.Direccion, &usuario.Telefono)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(" Usuario no encontrado con email:", email)
			return usuario, nil
		}
		log.Println(" Error al obtener usuario:", err)
		return usuario, err
	}

	// Validar contraseña
	if usuario.Contraseña != contraseña {
		log.Println(" Contraseña incorrecta para el usuario:", email)
		return Usuarios{}, nil
	}

	log.Println(" Usuario autenticado correctamente:", email)
	return usuario, nil
}

// InsertUsuario registra un nuevo usuario con su contraseña cifrada
func InsertUsuario(nombre, email, contrasena, rol string) (int64, error) {
	DB, err := database.Connect()
	if err != nil {
		log.Println(" Error al conectar a la base de datos:", err)
		return 0, err
	}
	defer DB.Close()

	// Generar hash de la contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(contrasena), bcrypt.DefaultCost)
	if err != nil {
		log.Println(" Error al encriptar contraseña:", err)
		return 0, err
	}

	query := `
		INSERT INTO usuarios (nombre, email, contraseña, rol)
		VALUES (?, ?, ?, ?)
	`
	result, err := DB.Exec(query, nombre, email, string(hash), rol)
	if err != nil {
		log.Println(" Error al registrar usuario:", err)
		return 0, err
	}

	id, _ := result.LastInsertId()
	log.Println(" Usuario registrado con ID:", id)
	return id, nil
}

// verificar credenciales de correo y contraseña para el login
func VerificarCredenciales(email, contrasena string) (int, error) {
	DB, err := database.Connect()
	if err != nil {
		log.Println("❌ Error al conectar a la base de datos:", err)
		return 0, err
	}
	defer DB.Close()

	var id int
	var hash string
	err = DB.QueryRow("SELECT id, contraseña FROM usuarios WHERE email = ?", email).Scan(&id, &hash)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("⚠️ Usuario no encontrado")
		} else {
			log.Println("❌ Error al obtener usuario:", err)
		}
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(contrasena))
	if err != nil {
		log.Println("⛔ Contraseña incorrecta")
		return 0, err
	}

	log.Println("✅ Usuario autenticado con ID:", id)
	return id, nil
}

// GetNombreUsuarioByID devuelve el nombre del usuario correspondiente al ID proporcionado.
// Se usa para personalizar vistas como el dashboard.
func GetNombreUsuarioByID(id int) (string, error) {
	DB, err := database.Connect()
	if err != nil {
		return "", err
	}
	defer DB.Close()

	var nombre string
	err = DB.QueryRow("SELECT nombre FROM usuarios WHERE id = ?", id).Scan(&nombre)
	if err != nil {
		return "", err
	}
	return nombre, nil
}

// ObtenerNombreYRolPorID consulta el nombre y el rol de un usuario según su ID.
// ObtenerNombreYRolPorID devuelve el nombre y rol de un usuario según su ID.
func ObtenerNombreYRolPorID(id int) (string, string, error) {
	DB, err := database.Connect()
	if err != nil {
		log.Println("❌ Error al conectar a la base de datos:", err)
		return "", "", err
	}
	defer DB.Close()

	var nombre, rol string
	query := "SELECT nombre, rol FROM usuarios WHERE id = ?"
	err = DB.QueryRow(query, id).Scan(&nombre, &rol)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("⚠️ Usuario no encontrado con ID:", id)
			return "", "", nil
		}
		log.Println("❌ Error al obtener nombre y rol:", err)
		return "", "", err
	}

	log.Printf("✅ Usuario ID %d - Nombre: %s | Rol: %s\n", id, nombre, rol)
	log.Printf("🧪 [func] Retornando nombre='%s', rol='%s'", nombre, rol)
	return nombre, rol, nil
}
func GetRolUsuarioByID(id int) (string, error) {
	DB, err := database.Connect()
	if err != nil {
		log.Println("❌ Error al conectar a la base de datos:", err)
		return "", err
	}
	defer DB.Close()

	var rol string
	err = DB.QueryRow("SELECT rol FROM usuarios WHERE id = ?", id).Scan(&rol)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("⚠️ Usuario no encontrado con ID:", id)
			return "", nil
		}
		log.Println("❌ Error al obtener el rol:", err)
		return "", err
	}

	log.Printf("🔐 Rol del usuario con ID %d: %s", id, rol)
	return rol, nil
}

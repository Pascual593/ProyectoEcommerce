package main

//autor: Pascual Campos
//fecha: 15/06/2025
//tema Proyecto Ecommerce
//Avance de proyecto definicion de struct, funciones , conexion de modelo con la bse de datos

import (
	"encoding/gob"

	"ProyectoEcommerce/database"
	"ProyectoEcommerce/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	gob.Register(map[string]interface{}{})
	//verifica si se ejecuta la conexion con la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	//si la conexion es exitosa, se imprime el mensaje
	//defer espera la ejecucion de todos los procesos para luego cerrar la conexion
	//defer tiempo de espera
	defer db.Close()
	log.Println("Conexion a la base de datos establecida correctamente")
	//--------------------------------------------------------------
	//Prueba de inserción de usuario antes de inicializar el servidor
	//--------------------------------------------------------------
	/*
		err = models.InsertUsuario("Roberto Mera", "roberto@email.com", "claveSegura555", "calderon", "0998337444")
		if err != nil {
			log.Println(" Error al registrar usuario:", err)
		} else {
			log.Println(" Usuario registrado exitosamente")
		}
	*/
	//---------------------------------------------------------------------
	//-----Prueba para Moficar Usuario antes de inicializar el servidor
	/*
		err = models.UpdateUsuario(1, "Juan Pérez", "juan@email.com", "nuevaClave123", "Av. Principal", "0999999999", "cliente")
		if err != nil {
			log.Println(" Fallo al actualizar usuario")
		} else {
			log.Println(" Usuario actualizado exitosamente")
		}
	*/
	//-------
	// ------------------------------------------------------------------------------------------
	/*
		// Intentar eliminar un usuario con ID 1
		err = models.DeleteUsuario(4)
		if err != nil {
			log.Println(" Fallo al eliminar usuario")
		} else {
			log.Println(" Usuario eliminado exitosamente")
		}
	*/
	//---------------------------------------------------------------------
	/*
		// Intentar iniciar sesión con un usuario
		usuario, err := models.LoginUsuario("carlos@email.com", "claveSegura123")
		if err != nil {
			log.Println(" Fallo en la autenticación")
		} else if usuario.ID == 0 {
			log.Println(" Usuario o contraseña incorrectos")
		} else {
			log.Println(" Sesión iniciada con éxito:", usuario.Email)
		}
	*/
	//---------------------------------------------------------------------------
	/*
		// Insertar un producto con categoría existente
		err = models.InsertProducto("camiseta cuello redondo", "color rojo S", 11.99, 20, 1) // Categoría ID = 1
		if err != nil {
			log.Println(" Fallo al registrar producto")
		} else {
			log.Println(" Producto registrado exitosamente")
		}
	*/
	//-----------------------------------------------------------------------------------------
	/*
		// Intentar obtener un producto con ID 1
		producto, err := models.GetProductoById(1)
		if err != nil {
			log.Println(" Fallo al obtener producto")
		} else if producto.ID == 0 {
			log.Println(" Producto no encontrado")
		} else {
			log.Println(" Producto obtenido con éxito:", producto.Nombre)
		}
	*/
	//----------------------------------------------------------------------
	/*
		// Intentar actualizar un producto con ID 1
		err = models.UpdateProducto(1, "camiseta polo", "color rojo talla L", 15.99, 10, 1)
		if err != nil {
			log.Println(" Fallo al actualizar producto")
		} else {
			log.Println(" Producto actualizado exitosamente")
		}
	*/
	//----------------------------------------------------------------------------------------
	/*
		// Intentar eliminar un producto con ID 1
		err := models.DeleteProducto(1)
		if err != nil {
			log.Println(" Fallo al eliminar producto")
		} else {
			log.Println(" Producto eliminado exitosamente")
		}

	*/
	//------------------------------------------------------------------------------
	/*
		// Obtener todos los productos
		productos, err := models.GetAllProductos()
		if err != nil {
			log.Println(" Fallo al obtener productos")
		} else if len(productos) == 0 {
			log.Println(" No hay productos disponibles")
		} else {
			log.Println(" Productos obtenidos con éxito:")
			for _, producto := range productos {
				log.Println("- ", producto.Nombre, "| Precio:", producto.Precio)
			}
		}
	*/
	//----------------------------------------------------------------------
	/*
		// Intentar registrar un pedido
		pedidoID, err := models.InsertPedido(8, 15.99, 1) // Usuario ID = 1, Método de pago ID = 1
		if err != nil {
			log.Println(" Fallo al registrar pedido")
		} else {
			log.Println(" Pedido registrado exitosamente con ID:", pedidoID)
		}
	*/
	//----------------------------------------------------------------------
	/*
		// Intentar obtener un pedido con ID 1
		pedido, err := models.GetPedidoById(1)
		if err != nil {
			log.Println(" Fallo al obtener pedido")
		} else if pedido.ID == 0 {
			log.Println(" Pedido no encontrado")
		} else {
			log.Println(" Pedido obtenido con éxito, Total:", pedido.Total, "Estado:", pedido.Estado)
		}
	*/
	//--------------------------------------------------------------------
	/*
		// Obtener todos los pedidos de un usuario (ID = 1)
		pedidos, err := models.GetAllPedidos(1)
		if err != nil {
			log.Println(" Fallo al obtener pedidos")
		} else if len(pedidos) == 0 {
			log.Println(" No hay pedidos para este usuario")
		} else {
			log.Println(" Pedidos encontrados:")
			for _, pedido := range pedidos {
				log.Println("- ID:", pedido.ID, "| Total:", pedido.Total, "| Estado:", pedido.Estado)
			}
		}

	*/
	//------------------------------------------------------------------------
	/*
		// Intentar cambiar el estado de un pedido con ID 1
		err := models.UpdateEstadoPedido(1, "enviado")
		if err != nil {
			log.Println(" Fallo al actualizar estado del pedido")
		} else {
			log.Println(" Estado del pedido actualizado exitosamente")
		}
	*/
	//--------------------------------------------------------------------
	/*
		// Intentar eliminar un pedido con ID 1
		err := models.DeletePedido(1)
		if err != nil {
			log.Println(" Fallo al eliminar pedido")
		} else {
			log.Println(" Pedido eliminado exitosamente")
		}
	*/
	//-----------------------------------------------------------------------
	/*
		// Agregar un producto a un pedido existente
		err := models.InsertDetallePedido(1, 3, 2, 500.00) // Pedido ID = 1, Producto ID = 3, Cantidad = 2, Precio = 500.00
		if err != nil {
			log.Println(" Fallo al agregar producto al pedido")
		} else {
			log.Println(" Producto agregado al pedido exitosamente")
		}
	*/
	//---------------------------------------------------------------------

	//---------------------------------------------------------------------
	/*
		// Intentar obtener un detalle de pedido con ID 1
		detalle, err := models.GetDetallePedidoById(1)
		if err != nil {
			log.Println(" Fallo al obtener detalle de pedido")
		} else if detalle.ID == 0 {
			log.Println(" Detalle de pedido no encontrado")
		} else {
			log.Println(" Detalle de pedido obtenido con éxito, Producto ID:", detalle.ProductoID, "Cantidad:", detalle.Cantidad)
		}
	*/
	/*
		id, err := models.InsertUsuario("Cesar", "cesar@mail.com", "1234secure", "admin")
		if err != nil {
			log.Println(" Registro fallido:", err)
		} else {
			log.Println(" Usuario insertado con ID:", id)
		}
	*/
	/*
		id, err := models.VerificarCredenciales("samantha@mail.com", "1234secure")
		if err != nil {
			log.Println("❌ Login fallido:", err)
		} else {
			log.Println("🔐 Login exitoso. Usuario ID:", id)
		}
	*/

	//---------------------------------------------------------------------
	//agrgar las consultas a la base de datos para retornarlas como api al sistema principal
	//--------------------------------------------------------------
	r := mux.NewRouter()
	// Servir archivos estáticos desde /uploads/
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads/"))))
	//definimos la ruta de navegacion para el sistema pero
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/dashboard", handlers.DashboardHandler).Methods("GET")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")
	r.HandleFunc("/productos", handlers.ProductoHandler).Methods("GET", "POST")
	r.HandleFunc("/productos/listar", handlers.ListarProductosHandler).Methods("GET")
	r.HandleFunc("/productos/editar/{id}", handlers.EditarProductoHandler).Methods("GET", "POST")
	r.HandleFunc("/productos/eliminar/{id}", handlers.EliminarProductoHandler).Methods("GET")
	r.HandleFunc("/productos/ver/{id}", handlers.VerProductoHandler)
	r.HandleFunc("/carrito/agregar/{id}", handlers.AgregarAlCarritoHandler).Methods("POST")
	r.HandleFunc("/carrito/ver", handlers.VerCarritoHandler).Methods("GET")
	r.HandleFunc("/pedido/confirmar", handlers.ConfirmarPedidoHandler).Methods("POST")
	r.HandleFunc("/pedido/confirmado", handlers.PedidoConfirmadoHandler).Methods("GET")
	r.HandleFunc("/inicio", handlers.InicioHandler).Methods("GET")
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))
	r.HandleFunc("/productos/nuevo", handlers.ProductoHandler)
	// inicializar el servidor web con las rutas y controladores
	log.Println("servidor de la bd inicializado")
	//http/listenandserve inicia un servidor HTTP en el puerto 8081
	//nil es el manejador por defecto que manejalas peticiones
	//las rutas con gorilla/mux
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)

	}

}

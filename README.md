# 🛒 ProyectoEcommerce — Aplicación Web de Comercio Electrónico

## 👥 Integrantes del grupo

- **Nombre:** Pascual Chura  
- **Carrera:** [tu carrera aquí]  
- **Materia:** Programación  
- **Universidad:** Universidad Internacional del Ecuador (UIDE)  
- **Semestre:** [coloca el semestre o la fecha de entrega]  
- **Docente:** [nombre del docente, si querés incluirlo]

---

## 🎯 Objetivo del proyecto

Desarrollar una aplicación web funcional que permita a usuarios explorar un catálogo de productos, autenticarse en el sistema, gestionar su carrito de compras y generar pedidos, integrando servicios web y persistencia de datos.

El sistema busca simular el flujo completo de un ecommerce básico, donde se aplican los conceptos aprendidos en las 4 unidades de la materia.

---

## 🚀 Tecnologías utilizadas

- **Lenguaje principal:** Go (Golang)
- **Framework web:** `net/http` + `gorilla/mux`
- **Base de datos:** SQLite
- **Motor de plantillas:** `html/template`
- **Serialización de datos:** JSON (servicios web RESTful)
- **Gestión de sesiones:** Cookies HTTP
- **Control de versiones:** Git y GitHub

---

## 🧩 Funcionalidades implementadas

### Funciones principales para usuarios:
- 🔍 Ver catálogo de productos
- 🔐 Iniciar y cerrar sesión
- 🧺 Agregar productos al carrito
- 📦 Confirmar pedidos realizados

### Funciones para administradores:
- ➕ Registrar nuevos productos
- ✏️ Editar y eliminar productos existentes
- 👥 Listar usuarios
- 📈 Ver panel de administración

---

## 🌐 Servicios Web implementados (JSON)

El sistema incluye 8 servicios web que responden en formato JSON:

| Ruta                   | Método | Descripción                     |
|------------------------|--------|---------------------------------|
| `/api/productos`       | GET    | Listar todos los productos      |
| `/api/producto/{id}`   | GET    | Obtener un producto por ID      |
| `/api/producto`        | POST   | Crear un nuevo producto         |
| `/api/producto/{id}`   | PUT    | Actualizar un producto          |
| `/api/producto/{id}`   | DELETE | Eliminar producto               |
| `/api/usuarios`        | GET    | Listar todos los usuarios       |
| `/api/login`           | POST   | Iniciar sesión                  |
| `/api/pedido`          | POST   | Confirmar pedido desde carrito  |

---

## 🔮 Visualización del futuro

Imaginamos este sistema evolucionando hacia una plataforma colaborativa para pequeños comercios, integrando microservicios, autenticación basada en tokens JWT, análisis de tendencias con inteligencia artificial y soporte multiplataforma móvil. Sería una solución SaaS lista para escalar en la nube ☁️🚀

---

## 📽️ Presentación

- 🎞️ Video demostrativo: https://www.dropbox.com/scl/fi/rocmtu8ej3zsx0t13xnac/Grabar_2025_07_01_23_06_26_444.mp4?rlkey=gdmma5o126d5hw1t0fc5y2zfb&st=5c4axgxh&dl=0
- 🖼️ Presentación en PowerPoint/Canva: 

---

## 🧠 Reflexión final

Este proyecto integró todos los aprendizajes de la materia: desde rutas y controladores en Go, plantillas HTML, manejo de sesiones, hasta la creación de servicios web RESTful y diseño visual.

A lo largo del desarrollo enfrentamos retos como la gestión de sesiones seguras, validación de formularios y diseño responsive. Gracias a estos desafíos, se reforzaron habilidades prácticas en programación real con un enfoque web moderno y funcional.

---

## 📁 Cómo ejecutar el proyecto

1. Cloná el repositorio:

```bash
git clone https://github.com/Pascual593/ProyectoEcommerce.git
cd ProyectoEcommerce

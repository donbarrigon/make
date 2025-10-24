# Make CLI

Asistente de línea de comandos inspirado en Artisan de Laravel para proyectos Go.

## 📋 Descripción

Make CLI es una herramienta que facilita la creación y gestión de proyectos Go siguiendo patrones similares a Laravel Artisan. Permite generar automáticamente estructuras de código, migraciones, controladores y más.

## 🚀 Instalación

```bash
go install github.com/donmarrigon/make@latest
```

## 📖 Uso Básico

```bash
make <comando> [<otro_comando>] [<dominio>.]<nombre>
```

- `<comando>` y `<nombre>` son **obligatorios**
- `<otro_comando>` y `<dominio>` son **opcionales**
- El nombre debe estar en **snake_case**

## 🛠️ Comandos Principales

### Ayuda

```bash
make help
```
Muestra todos los comandos disponibles.

---

## 🏗️ Gestión de Proyectos

### Crear Nuevo Proyecto

```bash
make project
```
Inicia un asistente interactivo para crear un nuevo proyecto desde cero.

### Ejecutar el Proyecto

```bash
make run            # Ejecuta el proyecto en modo desarrollo
make build          # Compila todo el proyecto en /dist
make merge:upstream # Hace merge con upstream (solo para forks)
```

---

## 💻 Comandos de Desarrollo

### Base de Datos

```bash
make model nombre_modelo              # Crea un template de modelo
make migration crear_tabla_usuarios   # Crea una migración
make seed usuarios_seed               # Crea un seeder
make repository user_repository       # Crea un repositorio
make resource user_resource           # Crea un recurso
```

### Vistas y Frontend

```bash
make view nombre_vista        # Crea templates de vista
make page home_page           # Crea una página
make component card           # Crea un componente
```

### Lógica de Negocio

```bash
make controller user_controller   # Crea un controlador
make middleware auth_middleware   # Crea un middleware
make policy user_policy           # Crea una policy
make route api_route              # Crea rutas
make service email_service        # Crea un service
make validator user_validator     # Crea un validator
```

---

## 🎯 Comandos Combinados

Estos comandos crean múltiples archivos relacionados de una sola vez:

### DB - Estructura Completa de Base de Datos
```bash
make db user
```
Crea: **migrations + seeds + model**

### Handler - Capa de Manejo HTTP Completa
```bash
make handler user
```
Crea: **controller + policy + route + validator**

### UI - Interfaz de Usuario Completa
```bash
make ui dashboard
```
Crea: **view + js + css**

### MVC - Aplicación Web Completa
```bash
make mvc producto
```
Crea: **db + ui + handler** (todos los archivos necesarios para una entidad completa)

### API - Backend API Completo
```bash
make api producto
```
Crea: **db + handler** (ideal para APIs REST)

---

## 📁 Estructura de Proyecto

```
.
├── main.go
├── go.mod
└── internal/
    ├── config/
    ├── database/
    │   ├── migrations/
    │   └── seeds/
    ├── models/
    ├── controllers/
    ├── middlewares/
    ├── policies/
    ├── routes/
    ├── services/
    ├── validators/
    ├── repositories/
    ├── resources/
    └── views/
```

---

## 💡 Ejemplos de Uso

### Crear un CRUD completo de usuarios

```bash
# Opción 1: Comando único
make mvc user

# Opción 2: Paso a paso
make model user
make migration create_users_table
make seed users_seed
make controller user_controller
make view users
```

### Crear una API REST

```bash
make api product
```

### Agregar autenticación

```bash
make middleware auth
make policy auth_policy
```

---

## 🤝 Contribuir

Si encuentras algún error o quieres contribuir, siéntete libre de abrir un issue o pull request.

---

## 📄 Licencia

MIT License

---

## ✨ Características

- ✅ Generación automática de código
- ✅ Patrones consistentes
- ✅ Inspirado en Laravel Artisan
- ✅ Soporte para arquitectura MVC
- ✅ Comandos combinados para productividad
- ✅ Estructura de proyecto organizada
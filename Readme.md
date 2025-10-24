# Make CLI

Asistente de línea de comandos inspirado en Artisan de Laravel para proyectos Go.

## 📋 Descripción

Make CLI es una herramienta que facilita la creación y gestión de proyectos Go siguiendo patrones similares a Laravel Artisan. Permite generar automáticamente estructuras de código, migraciones, controladores y más.

## 🚀 Instalación

```bash
go install github.com/donmarrigon/make@latest
```

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

## 📖 Uso Básico

```bash
make <comando> [<otro_comando>] [<dominio>.]<nombre>
```

- `<comando>` y `<nombre>` son **obligatorios**
- `<otro_comando>` y `<dominio>` son **opcionales**
- El `<nombre>` debe estar en **snake_case** y en **singular**

### Base de Datos

```bash
make model product             # Crea un template de modelo
make migration user            # Crea una migración
make seed role                 # Crea un seeder
make repository profile        # Crea un repositorio
make resource user_profile     # Crea un recurso
```

### Vistas y Frontend

```bash
make view category             # Crea templates de vista
make page home                 # Crea una página
make component card            # Crea un componente
```

### Lógica de Negocio

```bash
make controller user           # Crea un controlador
make middleware auth           # Crea un middleware
make policy user               # Crea una policy
make route inventory           # Crea rutas
make service emial             # Crea un service
make validator user_store      # Crea un validator
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
Crea: **pages_crud + js + css**

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
├── go.sum
├── biome.json
├── bun.lock
├── package.json
├── tsconfig.json
├── .env
├── .gitignore
├── internal/
│   ├── app/
│   │   ├── data/
│   │   │   ├── model/
│   │   │   ├── repository/       
│   │   │   └── resource/
│   │   ├── handler/          
│   │   │   ├── controller/
│   │   │   ├── middleware/
│   │   │   ├── policy/
│   │   │   ├── route/
│   │   │   ├── service/
│   │   │   └── validator/
│   │   └── ui/
│   │       ├── components/
│   │       ├── css/
│   │       ├── js/
│   │       ├── layout/
│   │       ├── pages/
│   │       └── view/        # archivos de pages compilados en formato go
│   └── database/
│       ├── data
│       │   ├── migrations/  # archivos de migracion
│       │   └── seeds/       # carpeta con los seeds
│       └── handler          # esta capa solo son endpoints para correr las migracion no requiere intervencion
│           ├── controller/
│           ├── middleware/
│           └── route/
├── node_modules/            # carpeta con cientos de gigas de librerias
├── public/                  # contiene el css, js compilado para el modo desarrollo (no usar)
├── tmp/                     # archivos temporales logs, sessiones, trackers de la migracion y seeds
```

---

## 💡 Ejemplos de Uso

### Crear un CRUD completo de usuarios

```bash
# Opción 1: Comando único
make mvc user

# Opción 2: Paso a paso
make model migration controller view user
```

### Crear una API-REST o API-MSGPACK

```bash
make api product
```

### Agregar autenticación

```bash
make middleware auth
make policy auth
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
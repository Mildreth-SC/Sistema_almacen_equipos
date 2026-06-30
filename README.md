# Sistema de Almacén de Equipos Tecnológicos

API en Go para Portotech, una tienda de cómputo en Portoviejo. La idea es dejar de anotar todo a mano y tener en un solo lugar el inventario de piezas, las devoluciones y los mantenimientos de los equipos.

## Repositorio y anexos

**Código fuente (GitHub):**  
https://github.com/Mildreth-SC/Sistema_almacen_equipos

**Documentos y entregables:**

- **Documento técnico** (Google Docs): [Sistema de Gestión de Soporte Técnico Portotech](https://docs.google.com/document/d/1U5uDGMQWNVDAJj211U_RJoV_4hpx0F0u5VdUIjA0Eu8/edit?usp=sharing)
- **Video demo Postman — inventario de piezas** (Mildreth Guanoluisa): https://canva.link/bsz189339rbh0wa
- **Video demo Postman — devoluciones** (Ivanna Zamora): _pendiente_
- **Video demo Postman — mantenimientos** (José Mieles): https://1drv.ms/f/c/a17811b6a61c315d/IgCumwojNjnXRYwofe-WlcTQAVoslMW9OtYk1AMPZNEN41Y 

## El problema

Hoy pierden historial, cuesta saber en qué estado está cada reparación y el cliente muchas veces no recibe información clara.

## El equipo

Somos tres integrantes y cada uno tiene su módulo:

- **Mildreth Guanoluisa** — inventario de piezas (`/api/v1/inventario-piezas`)
- **Ivanna Zamora** — devoluciones y garantías (`/api/v1/devoluciones`)
- **José Mieles** — mantenimiento de equipos (`/api/v1/mantenimientos`)

Al inicio mi módulo iba a ser seguimiento técnico, pero lo cambiamos a inventario de piezas porque sin control de repuestos el taller no funciona bien. Los tres módulos ya tienen CRUD con Chi, validación y SQLite.

## Tecnologías

Usamos Go, Chi como router, GORM con SQLite. La base se guarda en `cmd/api/almacen.db` y se crea sola al correr el servidor. También hay middleware de CORS para probar desde el navegador.

## Cómo correrlo

```bash
go mod tidy
go run ./cmd/api
```

Queda en `http://localhost:8080`. Si abres solo la raíz sale 404, es normal — la API está en `/api/v1/...`.

La primera vez crea la base y mete datos de ejemplo si las tablas están vacías.

## Endpoints

Inventario de piezas (Mildreth):

```
GET    /api/v1/inventario-piezas
GET    /api/v1/inventario-piezas/{id}
POST   /api/v1/inventario-piezas
PUT    /api/v1/inventario-piezas/{id}
DELETE /api/v1/inventario-piezas/{id}
PATCH  /api/v1/inventario-piezas/{id}/stock   →  body: {"delta": 5}
```

Devoluciones (Ivanna):

```
GET    /api/v1/devoluciones
GET    /api/v1/devoluciones/{id}
POST   /api/v1/devoluciones
PUT    /api/v1/devoluciones/{id}
DELETE /api/v1/devoluciones/{id}
```

Mantenimientos (José):

```
GET    /api/v1/mantenimientos
GET    /api/v1/mantenimientos/{id}
POST   /api/v1/mantenimientos
PUT    /api/v1/mantenimientos/{id}
DELETE /api/v1/mantenimientos/{id}
```

## Cómo está organizado el código

El punto de entrada es `cmd/api/main.go`. Los modelos van en `internal/models/`, los handlers en `internal/handlers/`, y la capa de datos en `internal/storage/` con una interfaz `Almacen` y la implementación en SQLite. El CORS está en `internal/middleware/`.

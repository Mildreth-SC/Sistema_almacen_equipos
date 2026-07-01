# Sistema de Almacén de Equipos Tecnológicos

API REST en Go para **Portotech**, tienda de cómputo en Portoviejo. Centraliza inventario de piezas, devoluciones y mantenimientos de equipos.

## Repositorio y anexos

**Código fuente:** https://github.com/Mildreth-SC/Sistema_almacen_equipos

**Documentos y entregables:**

- **Documento técnico** (Google Docs): [Sistema de Gestión de Soporte Técnico Portotech](https://docs.google.com/document/d/1U5uDGMQWNVDAJj211U_RJoV_4hpx0F0u5VdUIjA0Eu8/edit?usp=sharing)
- **Video demo Postman — inventario de piezas** (Mildreth Guanoluisa): https://canva.link/bsz189339rbh0wa
- **Video demo Postman — devoluciones** (Ivanna Zamora): _pendiente_
- **Video demo Postman — mantenimientos** (José Mieles): https://1drv.ms/f/c/a17811b6a61c315d/IgCumwojNjnXRYwofe-WlcTQAVoslMW9OtYk1AMPZNEN41Y

## El problema

Hoy pierden historial, cuesta saber en qué estado está cada reparación y el cliente muchas veces no recibe información clara.

## El equipo

| Integrante | Módulo | Ruta |
|------------|--------|------|
| Mildreth Guanoluisa | Inventario de piezas | `/api/v1/inventario-piezas` |
| Ivanna Zamora | Devoluciones y garantías | `/api/v1/devoluciones` |
| José Mieles | Mantenimiento de equipos | `/api/v1/mantenimientos` |

## Tecnologías

- **Go** + **Chi** (router)
- **GORM** + **SQLite** (`cmd/api/almacen.db`)
- **JWT** + **bcrypt** (autenticación)
- Arquitectura en capas: `handlers` → `service` → `storage`

## Arquitectura

```
cmd/api/main.go          → wiring, migrate, router
internal/models/         → structs de dominio
internal/service/        → reglas de negocio y validaciones
internal/storage/        → interfaces + GORM + memoria (tests)
internal/handlers/       → HTTP/JSON
internal/middleware/     → CORS + Auth JWT
```

Los tres módulos comparten el catálogo de **Pieza** (`pieza_id`) y **Cliente** (`cliente_id`) en devoluciones y mantenimientos.

**Usuario** = empleado de Portotech (login JWT). **Cliente** = persona que compra, devuelve o deja equipos en el taller.

## Cómo correrlo

```bash
go mod tidy
go run ./cmd/api
```

Servidor en `http://localhost:8080`.

Variable opcional para producción:

```bash
set JWT_SECRET=tu-secreto-seguro
go run ./cmd/api
```

**Nota:** Si cambias el modelo de datos, borra `cmd/api/almacen.db` y reinicia para regenerar el esquema y los datos de ejemplo.

```bash
del cmd\api\almacen.db
go run ./cmd/api
```

## Autenticación

Todas las rutas CRUD requieren token JWT.

### 1. Registrar usuario (primera vez)

```http
POST /api/v1/auth/registrar
Content-Type: application/json

{"email":"admin@portotech.com","password":"secret123"}
```

### 2. Login

```http
POST /api/v1/auth/login
Content-Type: application/json

{"email":"admin@portotech.com","password":"secret123"}
```

Respuesta: `{"token":"eyJ..."}`

### 3. Usar token en peticiones

```
Authorization: Bearer <token>
```

## Endpoints

### Auth (públicos)

```
POST /api/v1/auth/registrar
POST /api/v1/auth/login
```

### Inventario piezas (protegidos)

```
GET    /api/v1/inventario-piezas
GET    /api/v1/inventario-piezas/{id}
POST   /api/v1/inventario-piezas
PUT    /api/v1/inventario-piezas/{id}
DELETE /api/v1/inventario-piezas/{id}
PATCH  /api/v1/inventario-piezas/{id}/stock   →  {"delta": 5}
```

### Clientes (protegidos)

Catálogo de clientes de la tienda. Registrar un cliente antes de crear devoluciones o mantenimientos.

```
GET    /api/v1/clientes
GET    /api/v1/clientes/{id}
POST   /api/v1/clientes
PUT    /api/v1/clientes/{id}
DELETE /api/v1/clientes/{id}
```

### Devoluciones — Ivanna Zamora (protegidos)

```
GET    /api/v1/devoluciones                    ?estado=PENDIENTE
GET    /api/v1/devoluciones/{id}
POST   /api/v1/devoluciones
PUT    /api/v1/devoluciones/{id}
PATCH  /api/v1/devoluciones/{id}/estado
DELETE /api/v1/devoluciones/{id}
```

Estados: `PENDIENTE`, `APROBADA`, `RECHAZADA`

Resolver devolución (PATCH):

```json
{
  "estado": "APROBADA",
  "resolucion": "cambio",
  "atendido_por": "Ivanna Zamora"
}
```

### Mantenimientos — José Mieles (protegidos)

```
GET    /api/v1/mantenimientos                  ?estado=PENDIENTE
GET    /api/v1/mantenimientos/{id}
POST   /api/v1/mantenimientos
PUT    /api/v1/mantenimientos/{id}
PATCH  /api/v1/mantenimientos/{id}/estado
DELETE /api/v1/mantenimientos/{id}
```

Flujo de estados: `PENDIENTE` → `EN_PROCESO` → `LISTO` → `ENTREGADO`

Avanzar estado (PATCH):

```json
{"estado": "EN_PROCESO"}
```

## Ejemplos de body JSON

### Crear cliente

```json
{
  "nombre": "María López",
  "cedula": "0923456789",
  "telefono": "0991234567",
  "email": "maria.lopez@email.com",
  "direccion": "Portoviejo, Cdla. Kennedy"
}
```

### Crear pieza

```json
{
  "numero_serial": "SN-KING-RAM-8G",
  "codigo_barras": "BAR-10042",
  "nombre": "RAM DDR4 8GB",
  "categoria": "RAM",
  "marca": "Kingston",
  "modelo": "DDR4-2666",
  "garantia_meses": 12,
  "stock": 10,
  "stock_minimo": 2,
  "precio_compra": 28.00,
  "precio_venta": 45.00,
  "proveedor": "TechParts SA",
  "ubicacion": "Estante A3",
  "estado": "DISPONIBLE"
}
```

### Crear devolución

```json
{
  "pieza_id": "<uuid-de-pieza>",
  "cliente_id": "<uuid-de-cliente>",
  "numero_factura": "FAC-2024-089",
  "motivo": "DEFECTUOSO",
  "descripcion": "RAM no reconocida por la BIOS"
}
```

Motivos: `DEFECTUOSO`, `EQUIVOCADO`, `GARANTIA`

### Crear mantenimiento

```json
{
  "cliente_id": "<uuid-de-cliente>",
  "equipo_descripcion": "Laptop HP 15, negro",
  "numero_serial": "HP-CLIENTE-9988",
  "falla_reportada": "No enciende",
  "tipo": "CORRECTIVO",
  "tecnico": "Juan Pérez",
  "costo": 45.00,
  "anticipo": 20.00,
  "pieza_id": "<uuid-opcional>"
}
```

## Pruebas (Actividad C1 — Testing)

Cada integrante tiene **3 tipos de test** en su módulo:

| Integrante | Service (mock) | Handler (httptest) | Repositorio (GORM `:memory:`) |
|------------|----------------|--------------------|---------------------------------|
| Mildreth | `service/pieza_mock_test.go` | `handlers/inventario_piezas_handler_test.go` | `storage/pieza_sqlite_test.go` |
| Ivanna | `service/devolucion_mock_test.go` | `handlers/devolucion_handler_test.go` | `storage/devolucion_sqlite_test.go` |
| José | `service/mantenimiento_mock_test.go` | `handlers/mantenimiento_handler_test.go` | `storage/mantenimiento_sqlite_test.go` |

```bash
go test ./... -cover
```

- **Mock:** no guarda; verifica que dato inválido **no llega** al repositorio.
- **Fake (`AlmacenMemoria`):** guarda en mapa para httptest sin SQLite real.
- **GORM `:memory:`:** SQLite en RAM; crear → buscar/listar lo refleja.
- **401:** rutas protegidas sin `Authorization: Bearer` responden 401.

## Códigos HTTP

| Código | Cuándo |
|--------|--------|
| 200 | OK |
| 201 | Recurso creado |
| 400 | Validación fallida |
| 401 | Sin token o token inválido |
| 404 | Recurso no encontrado |
| 409 | Duplicado, devolución ya resuelta o transición de estado inválida |

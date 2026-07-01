# Base de datos — Portotech

SQLite de producción/desarrollo:

```
db/almacen.db
```

Se crea automáticamente al ejecutar `go run ./cmd/api`.

Si cambias el modelo (`internal/models/`), borra `almacen.db` y reinicia el servidor para regenerar esquema y datos de ejemplo.

```powershell
del db\almacen.db
go run ./cmd/api
```

Este archivo `.db` está en `.gitignore` (no se sube a Git).

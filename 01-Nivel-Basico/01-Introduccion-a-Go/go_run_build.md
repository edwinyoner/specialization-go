# Compilación y Ejecución

Guía completa de los comandos básicos para ejecutar y compilar programas en Go.

## ▶️ Comandos Principales

### 1. `go run` - Ejecutar sin compilar
- **Propósito**: Compila y ejecuta temporalmente
- **Uso ideal**: Desarrollo rápido
- **No genera** archivo ejecutable

```bash
   # Ejecutar archivo directamente
   go run hola_mundo.go

   # Ejecutar múltiples archivos
   go run main.go utils.go
```

### 2. `go build` - Compilar a binario
- **Propósito**: Genera ejecutable permanente
- **Opciones clave**:
    - `-o`: Nombre del output
    - `-v`: Modo verbose
    - `-race`: Detección de race conditions

```bash
   # Compilación básica (nombre automático)
   go build main.go  # Genera 'main' (Linux) o 'main.exe' (Windows)

   # Especificar nombre de salida
   go build -o mi-app main.go

   # Compilar para otras plataformas
   GOOS=windows GOARCH=amd64 go build -o app.exe  # Windows 64-bit
   GOOS=darwin GOARCH=arm64 go build -o app-m1     # macOS Apple Silicon
```

### 3. `go install` - Instalar como comando
- **Propósito**: Compila y mueve el binario a `$GOBIN`
- **Path por defecto**:
    - Linux/macOS: `~/go/bin`
    - Windows: `%USERPROFILE%\go\bin`

```bash
   # Instalar y hacer disponible globalmente
   go install .

   # Verificar instalación
   ls $GOBIN/mi-app
```

## 🛠 Flags Útiles

| Flag           | Descripción                          | Ejemplo                   |
|----------------|--------------------------------------|---------------------------|
| `-ldflags`     | Inyectar valores al compilar         | `-ldflags "-X main.Version=1.2"` |
| `-tags`        | Habilitar build tags                 | `-tags dev,debug`         |
| `-a`           | Recompilar paquetes dependientes     | `go build -a`             |
| `-work`        | Mostrar temp directory               | `go build -work`          |

## 📊 Comparación de Comandos

| Característica       | `go run`      | `go build`    | `go install`  |
|----------------------|---------------|---------------|---------------|
| Genera ejecutable    | No            | Sí            | Sí            |
| Ubicación output     | Temp          | Current dir   | $GOBIN        |
| Ideal para           | Pruebas rápidas | Distribución | Herramientas CLI |
| Requiere main package| Sí            | Sí            | Sí            |

## 🚨 Errores Comunes y Soluciones

### 1. **"No required module provides package"**
```bash
   # Ejecutar desde directorio raíz del módulo
   go mod init mi-proyecto
```

### 2. **"Cannot find package"**
```bash
   # Descargar dependencias faltantes
   go mod tidy
```

### 3. **Ejecutable no funciona en otra máquina**
```bash
   # Compilar estáticamente (sin dependencias C)
   CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"'
```

## 🏭 Ejemplo de Workflow Completo

1. Crear programa:
```go
// main.go
package main

import "fmt"

var Version = "desarrollo"

func main() {
    fmt.Printf("App v%s\n", Version)
}
```

2. Compilar con versión específica:
```bash
   go build -ldflags "-X main.Version=1.0" -o app-v1
```

3. Verificar:
```bash
   ./app-v1
   # Salida: App v1.0
```

## 📌 Mejores Prácticas

1. **Compilación cruzada**:
   ```bash
      # Lista de sistemas soportados
      go tool dist list
   ```

2. **Compresión de binarios**:
   ```bash
      # Reducir tamaño con upx
      upx --best mi-app
   ```

3. **Versionado automático**:
   ```bash
      # Inyectar versión desde Git
      VERSION=$(git describe --tags)
      go build -ldflags "-X main.Version=$VERSION"
   ```
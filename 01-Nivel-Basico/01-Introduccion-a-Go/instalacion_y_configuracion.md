```
# Instalación y Configuración de Go

Guía paso a paso para instalar Go y configurar tu entorno de desarrollo.

## 📥 Instalación de Go

### Requisitos Previos
- 500 MB de espacio en disco
- Permisos de administrador (para instalación global)
- Conexión a internet

### Pasos de Instalación por Sistema Operativo

#### 💻 Windows
1. **Descargar instalador**:
   - Visita [golang.org/dl](https://golang.org/dl/)
   - Descarga el archivo `.msi` para Windows
2. **Ejecutar instalador**:
   ```powershell
 # Ejecutar como administrador
    go-windows-amd64.m
    ```
3. **Configurar PATH** (automático en instalación estándar)

#### 💻 Linux (Debian/Ubuntu)
```bash

# Opción 1: Usar repositorio
sudo apt update
sudo apt install golang-go

# Opción 2: Instalar versión específica (ejemplo 1.21.0)
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

#### 💻 MacOS

```bash

# Opción 1: Homebrew
brew install go

# Opción 2: Instalador manual
# Descargar .pkg desde https://go.dev/dl/
# Ejecutar el instalador gráfico
```

### Verificar Instalación

```bash

go version
# Salida esperada: go version go1.21.0 linux/amd64
```

## ⚙ Configuración del Entorno

### Variables de Entorno Clave

| Variable   | Descripción                           | Valor Recomendado            |
| ---------- | -------------------------------------- | ---------------------------- |
| `GOPATH` | Workspace de Go (obsoleto en módulos) | `$HOME/go`                 |
| `GOROOT` | Directorio de instalación             | `/usr/local/go` (auto-set) |
| `GOBIN`  | Ejecutables instalados                 | `$GOPATH/bin`              |

### Configurar Workspace Moderno (Módulos)

1. Crear directorio del proyecto:
   ```bash

   ```

mkdir specialization-go && cd specialization-go

```
2. Inicializar módulo:
   ```bash
go mod init github.com/edwinyoner/specialization-go
```

### IDE/Editores Recomendados

1. **Visual Studio Code** (Configuración mínima):
   - Instalar extensiones:
     - [Go](https://marketplace.visualstudio.com/items?itemName=golang.go)
     - [Code Runner](https://marketplace.visualstudio.com/items?itemName=formulahendry.code-runner)
2. **GoLand** (JetBrains) - IDE profesional

## 🚨 Solución de Problemas Comunes

### Error: `go: command not found`

- **Causa**: PATH no configurado
- **Solución**:
  ```bash

  ```

# Linux/macOS

  echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc  # o ~/.bashrc
  source ~/.zshrc

```

### Error de permisos (Linux/macOS)
```bash

# Dar permisos al directorio go
sudo chown -R $USER:$USER /usr/local/go
```

### Versión incorrecta

```bash

# Eliminar instalaciones previas
sudo rm -rf /usr/local/go
# Reinstalar versión correcta
```

## 🔄 Post-Instalación Recomendada

1. Actualizar herramientas:
   ```bash

   ```

go install golang.org/x/tools/cmd/goimports@latest

```
2. Configurar formato automático en tu IDE
3. Verificar entorno completo:
   ```bash
go env
```

## 🧪 Primer Ejecución (Prueba)

```bash

# Crear archivo hello.go
echo 'package main\n\nimport "fmt"\n\nfunc main() {\n\tfmt.Println("¡Go instalado correctamente!")\n}' > hello.go

# Ejecutar
go run hello.go
```

## 📌 Notas Importantes sobre Módulos Go (1.13+)

- Desde Go 1.13+ se recomienda usar **Módulos**

### 1. **Módulos Go vs GOPATH**

- **GOPATH (Obsoleto)**:

  - Sistema antiguo que requería estructura de directorios específica (`src`, `bin`, `pkg`)
  - Dificultaba el manejo de dependencias y versiones

  ```bash

  ```

# Estructura antigua

   GOPATH/
     ├── src/
     │    └── github.com/edwinyoner/specialization-go
     ├── bin/
     └── pkg/

```

- **Módulos Go (Moderno)**:
  - Sistema actual desde Go 1.13 (activado por defecto en 1.16+)
  - Permite trabajar en cualquier directorio
  - Maneja dependencias con archivo `go.mod`
   ```bash
 # Inicializar módulo en cualquier directorio
   go mod init github.com/edwinyoner/specialization-go
```

### 2. **Características Clave de los Módulos**

- **Dependencias Versionadas**:
  - Versiones semánticas (`v1.2.3`)
  - Archivo `go.sum` para hashes de seguridad
- **Reproducibilidad**:
  ```bash

  ```

 go mod tidy       # Sincroniza dependencias
   go mod vendor     # Crea vendor/ para builds reproducibles

```
- **Compatibilidad**:
  - Soporte para _Semantic Import Versioning_
  - Mecanismo de actualización segura:
  ```bash
 go get -u         # Actualiza a versiones menores/parches
   go get foo@v1.2.3 # Versión específica
```

### 3. **Flujo de Trabajo Recomendado**

1. Crear nuevo proyecto fuera de `GOPATH`
2. Inicializar módulo:
   ```bash

   ```

go mod init `<module-path>`

```
3. Desarrollar normalmente, las dependencias se gestionan automáticamente
4. Comandos esenciales:
   ```bash
go mod download  # Descarga dependencias
go list -m all   # Lista todas las dependencias
go mod verify    # Verifica integridad de dependencias
```

### 4. **Ejemplo de archivo go.mod**

```mod
module github.com/edwinyoner/specialization-go

go 1.21

require (
    github.com/gorilla/mux v1.8.0
    golang.org/x/text v0.3.8
)

replace github.com/old/lib => ../local/lib  // Sustitución local
```

### 5. **Consejos de Migración**

- Para proyectos existentes en GOPATH:

```bash
   cd $GOPATH/src/mi-proyecto
   go mod init     # Convierte a módulo
   go mod tidy     # Actualiza dependencias
```

- Si necesitas desactivar módulos temporalmente:

```bash
  export GO111MODULE=off
```

**Importante**: Los módulos son ahora el estándar oficial para gestión de dependencias.

---

▶ **Siguiente paso**: [Estructura Básica de un Programa Go](./estructura_basica.md)

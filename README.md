# GitHub User Activity CLI

![Go Version](https://img.shields.io/badge/Go-1.25.5-blue)
![License](https://img.shields.io/github/license/maikreyes/GitHub-user-activity)

Una herramienta de línea de comandos (CLI) escrita en **Go** que obtiene y muestra la actividad reciente de usuarios de GitHub.

Este proyecto es una solución para el reto [GitHub User Activity](https://roadmap.sh/projects/github-user-activity) de **roadmap.sh**.

## 🚀 Características

- **Consulta en tiempo real:** Utiliza la API pública de GitHub para traer los últimos eventos.
- **Soporte de Eventos:** Identifica y formatea acciones como:
  - `PushEvent` (Commits en repositorios)
  - `IssuesEvent` (Apertura de nuevas incidencias)
  - `WatchEvent` (Estrellas dadas a repositorios)
  - `CreateEvent`, `PullRequestEvent`, entre otros.
- **Ligero y Rápido:** Compilado como un binario nativo gracias a Go.
- **Manejo de Errores:** Gestión clara de usuarios inexistentes o errores de conexión.

## 📋 Requisitos Previos

Para compilar y ejecutar este proyecto necesitarás:

- **Go 1.25.5** o superior instalado en tu sistema.
- Conexión a Internet (para consultar la API de GitHub).

## 🛠️ Instalación y Compilación

1. Clona el repositorio:

```Bash
git clone [https://github.com/maikreyes/GitHub-user-activity.git](https://github.com/maikreyes/GitHub-user-activity.git)
cd GitHub-user-activity
```

2. Descarga las dependencias (si las hay) y verifica el módulo:

```Bash
go mod tidy
Compila el binario:
```

```Bash
go build -o github-activity
Esto generará un ejecutable llamado github-activity (o github-activity.exe en Windows) en la raíz del proyecto.
```

💻 Uso
Ejecuta la herramienta desde tu terminal pasando el nombre de usuario de GitHub como argumento.

Sintaxis

```Bash
./github-activity <usuario>
```

Ejemplos

Ver tu propia actividad:

```Bash
./github-activity maikreyes
```

Ver la actividad del creador de Linux:

```Bash
./github-activity torvalds
```

Salida Esperada:

```Plaintext
Output:
- Fetching activity for GitHub user: maikreyes
- Push Events: 9
- Create Events: 3
- Repositories: [maikreyes/expense-tracker maikreyes/GitHub-user-activity maikreyes/OpenGl-GraphicMotorTest maikreyes/go-vue-journey]
```

🤝 Contribuciones
Las contribuciones son bienvenidas. Si deseas mejorar el código o añadir soporte para más eventos:

Haz un Fork del proyecto.

Crea una rama (git checkout -b feature/NuevaFuncionalidad).

Haz Commit de tus cambios (git commit -m 'Agrega nueva funcionalidad').

Haz Push a la rama (git push origin feature/NuevaFuncionalidad).

Abre un Pull Request.

📄 Licencia
Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.

Hecho por Maik Reyes

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
   ```bash
   git clone [https://github.com/maikreyes/GitHub-user-activity.git](https://github.com/maikreyes/GitHub-user-activity.git)
   cd GitHub-user-activity

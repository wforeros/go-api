# Tasks API

Proyecto básico para manejo de tareas de usuarios haciendo uso de GO y PostgreSQL.

El objetivo de este es realizar pruebas de integración con Go y PostgreSQL por lo que es un proyecto pequeño y no diseñado para producción sino para ser una guía básica.

## Requisitos

- Go 1.16
- PostgreSQL 13

## Como correr

1. Clonar el repositorio
2. Crear una base de datos en PostgreSQL
3. Setear los valores de conexión a la base de datos en el archivo `db/connection.go`
4. Correr el comando `go run main.go` o usar [air](https://github.com/air-verse/air) para facilitar el desarrollo

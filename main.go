package main

import (
	"bufio"
	"fmt"
	"gestion-libros/db"
	"gestion-libros/models"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Inicializar la conexión con la base de datos
	db.InitDB()

	// Crear la tabla de libros si no existe
	err := models.CrearTablaLibros()
	if err != nil {
		panic(err)
	}
	// BORRAR todos los libros y reiniciar el contador AUTO_INCREMENT (para pruebas)
	db.DB.Exec("DELETE FROM libros;")
	db.DB.Exec("ALTER TABLE libros AUTO_INCREMENT = 1;")

	// Lista de libros de ejemplo (se insertan solo si no existen por título)
	librosEjemplo := []models.Libro{
		{Titulo: "Go Básico", Autor: "Juan Pérez", Descripcion: "Aprende Go desde cero.", Disponible: true},
		{Titulo: "Go Intermedio", Autor: "Henry Campaña", Descripcion: "Temas intermedios de Go.", Disponible: true},
		{Titulo: "Go Avanzado", Autor: "David Vela", Descripcion: "Temas avanzados de Go.", Disponible: true},
		{Titulo: "Go Experto", Autor: "Richard Alejandro", Descripcion: "Temas completos de Go.", Disponible: true},
		{Titulo: "Go Profesional", Autor: "Laura Cedeño", Descripcion: "Nivel profesional de Go.", Disponible: true},
		{Titulo: "Go Testing", Autor: "Carlos Ruiz", Descripcion: "Pruebas unitarias y TDD con Go.", Disponible: true},
	}

	// Insertar libros si no están duplicados (basado en el título)
	for _, libro := range librosEjemplo {
		existe, err := models.ExisteLibroPorTitulo(libro.Titulo)
		if err != nil {
			fmt.Println("Error al verificar libro:", err)
			continue
		}
		if !existe {
			err = libro.Insertar()
			if err != nil {
				fmt.Println("Error al insertar libro:", err)
			}
		}
	}

	// Menú interactivo
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n ---------------MENÚ DEL SISTEMA----------------")
		fmt.Println("1. Listar libros")
		fmt.Println("2. Prestar libro")
		fmt.Println("3. Devolver libro")
		fmt.Println("4. Salir")
		fmt.Print("Seleccione una opción: ")

		opcionStr, _ := reader.ReadString('\n')
		opcion := strings.TrimSpace(opcionStr)

		switch opcion {
		case "1":
			libros, err := models.ListarLibros()
			if err != nil {
				fmt.Println("Error al listar libros:", err)
				continue
			}
			fmt.Println("\n -------------Lista de libros:--------------------")
			for _, l := range libros {
				estado := "Disponible"
				if !l.Disponible {
					estado = "Prestado"
				}
				fmt.Printf("ID: %d | Título: %s | Autor: %s | Estado: %s\n", l.ID, l.Titulo, l.Autor, estado)
			}

		case "2":
			fmt.Print("Ingrese el ID del libro a prestar: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("ID inválido.")
				continue
			}
			err = models.PrestarLibro(id)
			if err != nil {
				fmt.Println("Error al prestar libro:", err)
			} else {
				fmt.Println("Libro prestado con éxito.")
			}

		case "3":
			fmt.Print("Ingrese el ID del libro a devolver: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("ID inválido.")
				continue
			}
			err = models.DevolverLibro(id)
			if err != nil {
				fmt.Println("Error al devolver libro:", err)
			} else {
				fmt.Println("Libro devuelto con éxito.")
			}

		case "4":
			fmt.Println("Saliendo del sistema. ¡Hasta pronto!")
			return

		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

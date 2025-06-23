package handlers

import (
	"gestion-libros/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Home muestra la página de bienvenida
func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
	}
}

// MostrarLibros muestra el listado de libros
func MostrarLibros(w http.ResponseWriter, r *http.Request) {
	libros, err := models.ListarLibros()
	if err != nil {
		http.Error(w, "Error al obtener los libros", http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.Execute(w, libros)
	if err != nil {
		http.Error(w, "Error al renderizar el listado", http.StatusInternalServerError)
	}
}

// FormularioNuevoLibro muestra el formulario para agregar libros
func FormularioNuevoLibro(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/nuevo.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error al renderizar el formulario", http.StatusInternalServerError)
	}
}

// GuardarLibro guarda el nuevo libro enviado desde el formulario
func GuardarLibro(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		libro := models.Libro{
			Titulo:      r.FormValue("titulo"),
			Autor:       r.FormValue("autor"),
			Descripcion: r.FormValue("descripcion"),
			Disponible:  true,
		}
		err := libro.Insertar()
		if err != nil {
			log.Println("Error al insertar libro:", err)
			http.Error(w, "Error al guardar el libro", http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/libros", http.StatusSeeOther)
}

// PrestarLibroHandler cambia el estado del libro a prestado
func PrestarLibroHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err == nil {
		models.PrestarLibro(id)
	}
	http.Redirect(w, r, "/libros", http.StatusSeeOther)
}

// DevolverLibroHandler cambia el estado del libro a disponible
func DevolverLibroHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err == nil {
		models.DevolverLibro(id)
	}
	http.Redirect(w, r, "/libros", http.StatusSeeOther)
}

// EliminarLibroHandler elimina un libro por ID
func EliminarLibroHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err == nil {
		models.EliminarLibro(id)
	}
	http.Redirect(w, r, "/libros", http.StatusSeeOther)
}

package models

import (
	"database/sql"
	"errors"
	"gestion-libros/db"
)

// Libro representa la estructura del libro en la base de datos
type Libro struct {
	ID          int
	Titulo      string
	Autor       string
	Descripcion string
	Disponible  bool
}

// CrearTablaLibros crea la tabla si no existe
func CrearTablaLibros() error {
	query := `
	CREATE TABLE IF NOT EXISTS libros (
		id INT AUTO_INCREMENT PRIMARY KEY,
		titulo VARCHAR(100) NOT NULL,
		autor VARCHAR(100) NOT NULL,
		descripcion VARCHAR(255),
		disponible BOOLEAN DEFAULT TRUE
	);`
	_, err := db.DB.Exec(query)
	return err
}

// Insertar agrega un nuevo libro a la base de datos
func (l *Libro) Insertar() error {
	query := "INSERT INTO libros (titulo, autor, descripcion, disponible) VALUES (?, ?, ?, ?)"
	result, err := db.DB.Exec(query, l.Titulo, l.Autor, l.Descripcion, l.Disponible)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	l.ID = int(id)
	return nil
}

// ExisteLibroPorTitulo verifica si ya existe un libro con el mismo título
func ExisteLibroPorTitulo(titulo string) (bool, error) {
	query := "SELECT COUNT(*) FROM libros WHERE titulo = ?"
	var count int
	err := db.DB.QueryRow(query, titulo).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ListarLibros devuelve todos los libros de la base
func ListarLibros() ([]Libro, error) {
	query := "SELECT id, titulo, autor, descripcion, disponible FROM libros"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libros []Libro
	for rows.Next() {
		var l Libro
		err := rows.Scan(&l.ID, &l.Titulo, &l.Autor, &l.Descripcion, &l.Disponible)
		if err != nil {
			return nil, err
		}
		libros = append(libros, l)
	}
	return libros, nil
}

// BuscarLibroPorID busca un libro según su ID
func BuscarLibroPorID(id int) (Libro, error) {
	var l Libro
	query := "SELECT id, titulo, autor, descripcion, disponible FROM libros WHERE id = ?"
	err := db.DB.QueryRow(query, id).Scan(&l.ID, &l.Titulo, &l.Autor, &l.Descripcion, &l.Disponible)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return l, errors.New("libro no encontrado")
		}
		return l, err
	}
	return l, nil
}

// PrestarLibro cambia el estado del libro a no disponible
func PrestarLibro(id int) error {
	libro, err := BuscarLibroPorID(id)
	if err != nil {
		return err
	}
	if !libro.Disponible {
		return errors.New("el libro ya está prestado")
	}
	query := "UPDATE libros SET disponible = FALSE WHERE id = ?"
	_, err = db.DB.Exec(query, id)
	return err
}

// DevolverLibro cambia el estado del libro a disponible
func DevolverLibro(id int) error {
	libro, err := BuscarLibroPorID(id)
	if err != nil {
		return err
	}
	if libro.Disponible {
		return errors.New("el libro ya está disponible")
	}
	query := "UPDATE libros SET disponible = TRUE WHERE id = ?"
	_, err = db.DB.Exec(query, id)
	return err
}

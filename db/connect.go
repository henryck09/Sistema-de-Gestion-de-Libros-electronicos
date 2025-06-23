package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConexionBD() *sql.DB {
	conexion, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gestionlibros")
	if err != nil {
		log.Fatal("Error al conectar con la base de datos: ", err)
	}

	return conexion
}

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando archivo .env: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	fmt.Println("Conexión a la base de datos exitosa")

}

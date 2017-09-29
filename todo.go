// todo.go
package main

//paquetes del framework Echo
import (
	"github.com/labstack/echo"
	"database/sql"
	"go-echo-vue/handlers"
	_ "github.com/mattn/go-sqlite3"
	
)

func main(){
	//inicializa la base de datos
	db := initDB("storage.db")
	migrate(db)
	// crea una nueva instancia de Echo
	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
    e.PUT("/tasks", handlers.PutTask(db))
    e.DELETE("/tasks/:id", handlers.DeleteTask(db))


	//Corre como servidor web
	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	// Revisamos si hay errore en la db y cierra el programa
	if err != nil{
		panic(err)
	}

	//si no hay errores pero aun no podemos acceder a la db cierra el programa
	if db == nil{
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(sql)
	// cierra el programa si algo sale mal con nuestro query
	if err != nil{
		panic(err)
	}
}
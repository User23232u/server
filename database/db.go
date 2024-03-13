package database

import (
    "context"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/qiniu/qmgo"
)

var Client *qmgo.Client // Cambia 'client' a 'Client'

func init() {
    // Carga las variables de entorno desde el archivo .env
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }

    // Obtén la cadena de conexión a MongoDB de una variable de entorno
    connectionString := os.Getenv("MONGODB_URI")
    if connectionString == "" {
        log.Fatal("La variable de entorno MONGODB_URI no está configurada")
    }

    // Crea un nuevo cliente de MongoDB usando qmgo
    var err error
    Client, err = qmgo.NewClient(context.Background(), &qmgo.Config{Uri: connectionString}) // Asegúrate de que 'Client' esté en mayúscula aquí también
    if err != nil {
        log.Fatal(err)
    }

    // Imprime un mensaje si la conexión es exitosa
    log.Println("Conexión a DB exitosa")
}

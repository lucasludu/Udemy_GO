// Este package main es el paquete principal de nuestro programa, y es necesario para que se pueda ejecutar
package main

// Si no se usa fmt, se puede eliminar la importación
// Se puede importar más de un paquete a la vez, pero no es necesario
// Este fmt es un paquete que nos permite imprimir en la consola
import "fmt"

func main() {
	fmt.Printf("¡Hola, %s!\n", "Mundo")
}

package main
import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	var ruta string
	ruta = filepath.Join(".", "bbdd", "name 2")
	err := os.MkdirAll(ruta, 0755)

	if err != nil {
		// Si hay un error, imprímelo
		fmt.Println("Error al crear el directorio:", err)
	} else {
		// Si no hay error, muestra un mensaje de éxito
		fmt.Println("Directorio creado con éxito:", ruta)
	}
}
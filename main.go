package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//se crea el slice
	miLista := TaskList{}
	//intanciamos nuesto scanner para obtener los datos
	scanner := bufio.NewScanner(os.Stdin)

	//realizamos un for hasta que salgamos
	for {
		//impriumimos las opciones
		fmt.Println("\n Lista de que haceres")
		fmt.Println("1. Agregar tarea")
		fmt.Println("2. Completar tarea")
		fmt.Println("3. Listar tareas")
		fmt.Println("4. Salir")
		fmt.Print("Elige una opcion:")

		//obtenemos la opcion escrita
		scanner.Scan()
		input := scanner.Text()

		//se realiza un switch case
		switch input {
		case "1":
			//si es 1
			//pide escribir la tarea
			fmt.Print("Descripcion de la tarea: ")
			scanner.Scan()
			miLista.AddTask(scanner.Text())
		case "2":
			// si es 2
			//pide el numero de la tarea
			//pasa el texto a int y verifica si es correcto
			fmt.Print("numero de tarea a Completar: ")
			scanner.Scan()
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Por favor, ingrese un numero valido.")
				continue
			}
			//completa la tarea
			miLista.CompleteTask(num - 1)
		case "3":
			//si es 3
			//imprime la lista de tareas
			miLista.ListTasks()
		case "4":
			//si es 4
			//sale del for
			fmt.Println("Saliendo...")
			return
		default:
			//cualquier otro valor
			fmt.Println("Opcion no valida")
		}
	}
}

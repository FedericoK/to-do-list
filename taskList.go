package main

import "fmt"

//Se creo la estructura de Task
type Task struct {
	Descripcion string
	Done        bool
}

//Se creo la estructura de TaskList
//Tiene un slice de task como atributo
type TaskList struct {
	Tasks []Task
}

//Realiza un append al slice de tasklist
func (tl *TaskList) AddTask(descripcion string) {
	//al pasarle una descripcion realiza el append creando una intancia
	//de Task con la descripcion y Done en false
	tl.Tasks = append(tl.Tasks, Task{Descripcion: descripcion, Done: false})
}

//Pone en completo la tarea
func (tl *TaskList) CompleteTask(index int) {
	//Revisa si el indice no es menor que 0 o
	//que el indice no sea mayor que el length del slice
	if index < 0 || index > len(tl.Tasks) {
		fmt.Println("Indice de tarea invalido.")
		return
	}
	tl.Tasks[index].Done = true
}

//Imprime la lista creada
func (tl *TaskList) ListTasks() {
	//Revisa si el slice existe
	if len(tl.Tasks) == 0 {
		fmt.Println("No hay tareas registradas.")
		return
	}

	//recorre el slice
	for i, task := range tl.Tasks {
		//se crea el la variable status en "Pendiente"
		//porque el Done esta en false por defecto
		status := "Pendiente"
		if task.Done {
			//si Done == true, status = "Completado"
			status = "Completado"
		}

		//Imprimir con Printf para darle el formato que queremos
		fmt.Printf("%d. %s [%s]\n", i+1, task.Descripcion, status)
	}
}

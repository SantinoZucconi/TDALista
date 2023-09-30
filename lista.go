package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si no tiene elementos enlistados, falso en caso contrario
	EstaVacia() bool

	// InsertarPrimero coloca el elemento en el inicio de la lista
	InsertarPrimero(T)

	// InsertarUltimo coloca el elemento al final de la lista
	InsertarUltimo(T)

	// Si en la lista hay elementos, BorrarPrimero elimina el primero. Si esta vacia, entra en
	// panico con un mensaje "La lista esta vacia"
	BorrarPrimero() T

	// Si en la lista hay elementos, VerPrimero devuelve el primero de la lista. Si esta vacia,
	// entra en panico con un mensaje "La lista esta vacia"
	VerPrimero() T

	// Si en la lista hay elementos, VerUltimo devuelve el ultimo de la lista. Si esta vacia,
	// entra en panico con un mensaje "La lista esta vacia"
	VerUltimo() T

	// Largo devuelve la cantidad de elementos enlistados
	Largo() int

	// Iterar recorre la lista desde el inicio y evalua elemento a elemento la condicion de la funcion visitar. Si esta devuelve falso,
	// se corta la iteracion. En caso contrario, esta continua.
	Iterar(visitar func(T) bool)

	// Iterador devuelve una interfaz que permite al usuario manejar manualmente las iteraciones, pudiendo recorrer, insertar y borrar elementos
	// de la lista. Mientras se use el iterador, no es correcto utilizar InsertarPrimero, InsertarUltimo y BorrarPrimero.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// Durante la iteracion, VerActual devuelve el elemento de la lista en donde esta posicionado el iterador. Cuando llega al final,
	// entra en panico con un mensaje "El iterador termino de iterar"
	VerActual() T

	// HaySiguiente devuelve verdadero si hay un elemento mas para ver, falso en caso contrario.
	HaySiguiente() bool

	// Siguiente itera al proximo elemento. Cuando se llega al final de la lista y se quiere iterar, entra en panico con un mensaje
	// "El iterador termino de iterar"
	Siguiente()

	// Insertar agrega un elemento a la lista en la posicion anterior con respecto al elemento en donde estaba posicionado el iterador,
	// pasando a ser el actual el nuevo elemento insertado.
	Insertar(T)

	// Borrar elimina y devuelve el elemento de la lista en donde el iterador estaba posicionado, colocando el iterador en el proximo
	// elemento. Cuando se llega al final de la lista y se quiere borrar, entra en panico con un mensaje "El iterador termino de iterar"
	Borrar() T
}

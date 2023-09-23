package lista

type ListaEnlazada[T any] struct {
	inicio *nodoLista[T]
	fin    *nodoLista[T]
	largo  int
}

type Iterador[T any] struct {
	anterior *nodoLista[T]
	actual   *nodoLista[T]
}

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &ListaEnlazada[T]{}
}

func (l *ListaEnlazada[T]) EstaVacia() bool {
	return l.inicio == nil
}

func (l *ListaEnlazada[T]) InsertarPrimero(dato T) {

}

func (l *ListaEnlazada[T]) InsertarUltimo(dato T) {

}

func (l *ListaEnlazada[T]) BorrarPrimero() T {
}

func (l *ListaEnlazada[T]) VerPrimero() T {
	if !l.EstaVacia() {
		return l.inicio.dato
	} else {
		panic("La lista esta vacia")
	}
}

func (l *ListaEnlazada[T]) VerUltimo() T {
	if !l.EstaVacia() {
		return l.fin.dato
	} else {
		panic("La lista esta vacia")
	}
}

func (l *ListaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *ListaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.inicio
	for actual != nil && visitar(actual.dato) == true {
		actual = actual.siguiente
	}
}

//Iterador

func (l *ListaEnlazada[T]) Iterador() IteradorLista[T] {
	return &Iterador[T]{}
}

func (i *Iterador[T]) VerActual() T {
	if i.actual != nil {
		return i.actual.dato
	} else {
		panic("El iterador termino de iterar")
	}
}

func (i *Iterador[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *Iterador[T]) Siguiente() {
	if i.actual != nil {
		i.anterior = i.actual
		i.actual = i.actual.siguiente
	} else {
		panic("El iterador termino de iterar")
	}
}

func (i *Iterador[T]) Insertar(dato T) {
	nuevo_nodo := &nodoLista[T]{dato: dato, siguiente: i.actual}
	i.anterior.siguiente = nuevo_nodo
	i.anterior = nuevo_nodo
}

func (i *Iterador[T]) Borrar() T {
	if i.actual != nil {
		i.anterior.siguiente = i.actual.siguiente
		return i.actual.dato
	} else {
		panic("El iterador termino de iterar")
	}
}

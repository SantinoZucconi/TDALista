package lista

type listaEnlazada[T any] struct {
	inicio *nodoLista[T]
	fin    *nodoLista[T]
	largo  int
}

type iterador[T any] struct {
	anterior *nodoLista[T]
	actual   *nodoLista[T]
	lista    *listaEnlazada[T]
}

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.inicio == nil
}

func (l *listaEnlazada[T]) InsertarPrimero(valor T) {
	nuevo_nodo := nodoLista[T]{
		dato:      valor,
		siguiente: l.inicio,
	}
	if l.EstaVacia() {
		l.fin = &nuevo_nodo
	}
	l.inicio = &nuevo_nodo
	l.largo++

}

func (l *listaEnlazada[T]) InsertarUltimo(valor T) {
	nuevo_nodo := nodoLista[T]{
		dato: valor,
	}
	if l.EstaVacia() {
		l.inicio = &nuevo_nodo
	} else {
		l.fin.siguiente = &nuevo_nodo
	}
	l.fin = &nuevo_nodo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if !l.EstaVacia() {
		valor := l.inicio.dato
		l.inicio = l.inicio.siguiente
		l.largo--
		return valor
	} else {
		panic("La lista esta vacia")
	}

}

func (l *listaEnlazada[T]) VerPrimero() T {
	if !l.EstaVacia() {
		return l.inicio.dato
	} else {
		panic("La lista esta vacia")
	}
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if !l.EstaVacia() {
		return l.fin.dato
	} else {
		panic("La lista esta vacia")
	}
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.inicio
	for actual != nil && visitar(actual.dato) {
		actual = actual.siguiente
	}
}

//Iterador

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterador[T]{
		anterior: nil,
		actual:   l.inicio,
		lista:    l,
	}
}

func (i *iterador[T]) VerActual() T {
	if i.actual != nil {
		return i.actual.dato
	} else {
		panic("El iterador termino de iterar")
	}
}

func (i *iterador[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *iterador[T]) Siguiente() {
	if i.actual != nil {
		i.anterior = i.actual
		i.actual = i.actual.siguiente
	} else {
		panic("El iterador termino de iterar")
	}
}

func (i *iterador[T]) Insertar(valor T) {
	nuevo_nodo := nodoLista[T]{
		dato:      valor,
		siguiente: i.actual,
	}
	if i.lista.fin == i.anterior {
		i.lista.fin = &nuevo_nodo
	}
	if i.lista.inicio != i.actual {
		i.anterior.siguiente = &nuevo_nodo
	} else {
		i.lista.inicio = &nuevo_nodo
	}
	i.actual = &nuevo_nodo
	i.lista.largo++
}

func (i *iterador[T]) Borrar() T {
	if i.actual != nil {
		if i.lista.fin == i.actual {
			i.lista.fin = i.anterior
		}
		if i.lista.inicio == i.actual {
			i.lista.inicio = i.actual.siguiente
		} else {
			i.anterior.siguiente = i.actual.siguiente
		}
		valor_eliminado := i.actual.dato
		i.actual = i.actual.siguiente
		i.lista.largo--
		return valor_eliminado
	} else {
		panic("El iterador termino de iterar")
	}
}

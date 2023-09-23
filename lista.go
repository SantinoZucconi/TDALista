package lista

type ListaEnlazada[T any] struct {
	inicio *nodoLista[T]
	fin    *nodoLista[T]
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
}

func (l *ListaEnlazada[T]) VerUltimo() T {
}

func (l *ListaEnlazada[T]) Largo() int {
}

func (l *ListaEnlazada[T]) Iterar(visitar func(T) bool) {
}

func (l *ListaEnlazada[T]) Iterador() {
}

func (i *Iterador[T]) VerPrimero() T {
}

func (i *Iterador[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *Iterador[T]) Siguiente() {
}

func (i *Iterador[T]) Insertar(dato T) {
}

func (i *Iterador[T]) Borrar() T {
}

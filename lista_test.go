package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const tam int = 10

func TestListaVacia(t *testing.T) {
	t.Log("Comprobamos que una lista esta vacia al crearla")
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
} // :)

func TestAlgunosIntS(t *testing.T) {
	t.Log("Agregamos algunos enteros al principio y corroboramos que esten en el orden deseado")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < tam; i++ {
		lista.InsertarPrimero(i)
	}
	require.Equal(t, 9, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())
	require.Equal(t, tam, lista.Largo())
} // :)

func TestAlgunosIntE(t *testing.T) {
	t.Log("Agregamos algunos enteros al final y corroboramos que esten en el orden deseado")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < tam; i++ {
		lista.InsertarUltimo(i)
	}
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 9, lista.VerUltimo())
	require.Equal(t, tam, lista.Largo())
} // :)

func TestBorrarPrimero(t *testing.T) {
	t.Log("Probamos que se borre correctamente el primer elemento de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < tam; i++ {
		lista.InsertarUltimo(i)
	}
	valor_borrado := lista.BorrarPrimero()
	require.Equal(t, 0, valor_borrado)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 9, lista.VerUltimo())
	require.Equal(t, tam-1, lista.Largo())
} // :)

func TestInicioIterador(t *testing.T) {
	t.Log("Iniciamos el iterador externo en una lista y evaluamos que tome los valores iniciales correctos")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < tam; i++ {
		lista.InsertarUltimo(i)
	}
	i := lista.Iterador()
	for j := 0; j < tam; j++ {
		require.True(t, i.HaySiguiente())
		require.Equal(t, j, i.VerActual())
		i.Siguiente()
	}

} // :)

func TestInsertarEnIterador(t *testing.T) {
	t.Log("Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	i := lista.Iterador()
	i.Insertar(27)
	require.Equal(t, 27, lista.VerPrimero()) // 1
	require.Equal(t, 4, lista.Largo())       // 3
}

func TestInsertarFinalIterador(t *testing.T) {
	t.Log("Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	i := lista.Iterador()
	for i.HaySiguiente() {
		i.Siguiente()
	}
	i.Insertar(27)
	require.Equal(t, 27, lista.VerUltimo()) // 3
	require.Equal(t, 4, lista.Largo())      // 3
}

func TestBorrarEnIterador(t *testing.T) {
	t.Log("Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	i := lista.Iterador()
	valor_eliminado := i.Borrar()
	require.Equal(t, 1, valor_eliminado)
	require.Equal(t, 2, lista.VerPrimero()) // 1
	require.Equal(t, 2, lista.Largo())      // 3
}

func TestBorrarFinalIterador(t *testing.T) {
	t.Log("Remover el último elemento con el iterador cambia el último de la lista.")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	i := lista.Iterador()
	contador := 0
	for i.HaySiguiente() && contador < lista.Largo()-1 {
		i.Siguiente()
		contador++
	} // esta es la mejor forma de llegar al ultimo? me suena que no
	valor_eliminado := i.Borrar()
	require.Equal(t, 5, valor_eliminado)
	require.Equal(t, 4, lista.VerUltimo()) // 5
	require.Equal(t, 4, lista.Largo())     // 5
}

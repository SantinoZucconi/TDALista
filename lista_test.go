package lista_test

import (
	"tdas/lista"
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const Tam int = 10
const Vacia int = 0

// TESTS PRIMITIVAS LISTA

func TestListaVacia(t *testing.T) {
	t.Log("Comprobamos que una lista esta vacia al crearla")
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.Equal(t, Vacia, lista.Largo())
}

func TestAlgunosIntS(t *testing.T) {
	t.Log("Agregamos algunos enteros al principio y corroboramos que esten en el orden deseado")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < Tam; i++ {
		lista.InsertarPrimero(i)
	}
	require.Equal(t, 9, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())
	require.Equal(t, Tam, lista.Largo())
}

func TestAlgunosStrS(t *testing.T) {
	t.Log("Agregamos algunos strings al principio y corroboramos que esten en el orden deseado")
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("We")
	lista.InsertarUltimo("don't")
	lista.InsertarUltimo("believe")
	lista.InsertarUltimo("what's")
	lista.InsertarUltimo("on")
	lista.InsertarUltimo("TV")
	require.Equal(t, "We", lista.VerPrimero())
	require.Equal(t, "TV", lista.VerUltimo())
	require.Equal(t, 6, lista.Largo())

}

func TestAlgunosIntE(t *testing.T) {
	t.Log("Agregamos algunos enteros al final y corroboramos que esten en el orden deseado")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < Tam; i++ {
		lista.InsertarUltimo(i)
	}
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 9, lista.VerUltimo())
	require.Equal(t, Tam, lista.Largo())
}

func TestAlgunosStrE(t *testing.T) {
	t.Log("Agregamos algunos strings al final y corroboramos que esten en el orden deseado")
	lista := lista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("pista")
	lista.InsertarPrimero("la")
	lista.InsertarPrimero("explotame")
	lista.InsertarPrimero("Biza")
	lista.InsertarPrimero("dale")

	require.Equal(t, "dale", lista.VerPrimero())
	require.Equal(t, "pista", lista.VerUltimo())
	require.Equal(t, 5, lista.Largo())

}

func TestBorrarPrimero(t *testing.T) {
	t.Log("Probamos que se borre correctamente el primer elemento de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < Tam; i++ {
		lista.InsertarUltimo(i)
	}
	valor_borrado := lista.BorrarPrimero()
	require.Equal(t, 0, valor_borrado)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 9, lista.VerUltimo())
	require.Equal(t, Tam-1, lista.Largo())
}

// TESTS ITERADORES EXTERNOS

func TestInicioIterador(t *testing.T) {
	t.Log("Iniciamos el iterador externo en una lista y evaluamos que tome los valores iniciales correctos")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < Tam; i++ {
		lista.InsertarUltimo(i)
	}
	i := lista.Iterador()
	for j := 0; j < Tam; j++ {
		require.True(t, i.HaySiguiente())
		require.Equal(t, j, i.VerActual())
		i.Siguiente()
	}

}

func TestInsertarInicioIterador(t *testing.T) {
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

func TestInsertarEntrIterador(t *testing.T) {
	t.Log("Insertar un elemento en el medio se hace en la posición correcta.")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	i := lista.Iterador()
	insertado := false
	for i.HaySiguiente() && !insertado {
		if i.VerActual() == 4 {
			i.Insertar(3)
			insertado = true
		}
		i.Siguiente()
	}
	ordenado := true
	j := 1
	for k := lista.Iterador(); k.HaySiguiente(); k.Siguiente() {
		if k.VerActual() != j {
			ordenado = false
		}
		j++
	}
	require.True(t, ordenado)
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

// TESTS ITERADOR INTERNO

func TestSumaKelementos(t *testing.T) {
	t.Log("Sumamos los primeros k elementos con iterador interno")
	lista := TDALista.CrearListaEnlazada[int]()
	for k := 0; k < 2*Tam; k++ {
		lista.InsertarUltimo(k)
	}
	var suma, contador int
	elementos_a_sumar := 7
	lista.Iterar(func(x int) bool {
		contador += 1
		suma += x
		return contador < elementos_a_sumar
	})
	require.Equal(t, 21, suma)
}

func TestPosicionElemento(t *testing.T) {
	t.Log("Probamos que encontrar la posicion de un elemento con el iterador interno funcione correctamente")
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("Y")
	lista.InsertarUltimo("se")
	lista.InsertarUltimo("salieron")
	lista.InsertarUltimo("con")
	lista.InsertarUltimo("la")
	lista.InsertarUltimo("suya")
	lista.InsertarUltimo("porque")
	lista.InsertarUltimo("hay")
	lista.InsertarUltimo("ladrones")
	lista.InsertarUltimo("al")
	lista.InsertarUltimo("acecho")
	posicion := 0
	encontrado := false
	elemento_a_encontrar := "hay"
	lista.Iterar(func(s string) bool {
		if s == elemento_a_encontrar {
			encontrado = true
			posicion--
		}
		posicion++
		return !encontrado
	})

	require.Equal(t, 7, posicion)
	require.True(t, encontrado)
}

func TestSumaElementosImpares(t *testing.T) {
	t.Log("Sumamos unicamente los elementos impares de la lista de enteros")
	lista := TDALista.CrearListaEnlazada[int]()
	for k := 0; k < 2*Tam; k++ {
		lista.InsertarPrimero(k)
	}
	suma := 0
	lista.Iterar(func(i int) bool {
		if i%2 != 0 {
			suma += i
		}
		return true
	})
	require.Equal(t, 100, suma)
}

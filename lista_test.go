package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertarEnIterador(t *testing.T) {
	t.Log("Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	i := lista.Iterador()
	i.Insertar(27)
	require.Equal(t, 27, lista.VerPrimero())
}

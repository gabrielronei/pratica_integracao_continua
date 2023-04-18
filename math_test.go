package main

import "testing"

func TestSum(t *testing.T) {

     obtido := Sum(10, 10)
     esperado := 20

     if obtido != esperado {

          t.Errorf("Obteve: %d, Esperavamos: %d", obtido, esperado)
     }
}

func TestSubtract(t *testing.T) {

    obtido := Subtract(10, 10)
    esperado := 0

    if obtido != esperado {

         t.Errorf("Obteve: %d, Esperavamos: %d", obtido, esperado)
    }
}


package main

import (
	"fmt"
	"sync"
	"time"
)

/*
* sync.Mutex
*
* Hemos visto cómo los canales son excelentes para la comunicación entre gorutinas.
*
* ¿Pero qué pasa si no necesitamos comunicación? ¿Qué pasa si solo queremos asegurarnos de que solo una
* gorutina pueda acceder a una variable a la vez para evitar conflictos?
*
* Este concepto es llamado exclusión mutua, y el nombre convecional de la estructura de datos que proporciona
* es Mutex.
*
* La librería estándar de Go provee la exclusión mutua con sync.Mutex y sus dos métodos:
*
* Lock(bloqueado)
* Unlock(desbloqueado)
*
* Podemos definir un bloque de código para que se ejecute en exclusión mutua rodeándolo con una llamada
* a Lock y Unlock como se muestra en el método Inc.
*
* También podemos usar defer para asegurarnos de que el mutex se desbloquee como en el método Valor.
* */

// SafeCounter es seguro de usar simultáneamente.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc incrementa el contador para la clave dada.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Bloquear para que solo una rutia a la vez pueda acceder al mapa c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Valor devuelve el valor actual del contador para la clave dada.
func (c *SafeCounter) Valor(key string) int {
	c.mu.Lock()

	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Valor("somekey"))
}

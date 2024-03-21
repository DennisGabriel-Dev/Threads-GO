package main

import (
	"fmt"
	"net/http"
	"sync"
)

func letras_func() []string {
	return []string{
		"Caneta azul, azul caneta",
		"Caneta azul tá marcada com minha letra",
		"Caneta azul, azul caneta",
		"Caneta azul tá marcada com minha letra",

		"Todo dia eu viajo pra o colégio",
		"Com uma caneta azul e uma caneta amarela",
		"Eu perdi minha caneta e eu peço",
		"Por favor, quem encontrou, me entrega ela",

		"Caneta azul, azul caneta",
		"Caneta azul tá marcada com minha letra",

		"A professora, ela veio brigar comigo",
		"Porque eu perdi a última caneta que eu tinha",
		"Não brigue, professora, porque eu vou comprar outra canetinha",

		"Caneta azul, azul caneta",
		"Caneta azul tá marcada com minha letra",
	}
}

func acordes_func() []string {
	return []string{
		"D", "G", "A", "Em",
		"D", "G", "A", "Bm",
		"D", "G", "A", "Cm",
		"D", "G", "A",
	}
}

var (
	wg     sync.WaitGroup
	letras = letras_func()
	cifras = acordes_func()
)

func main() {
	http.HandleFunc("/", handler)

	wg.Add(len(cifras))
	for i := range letras {
		go processaLinha(i)
	}
	wg.Wait()

	fmt.Println("Todas as linhas foram processadas. Servidor iniciando...")
	http.ListenAndServe(":8080", nil)
}

func processaLinha(i int) {
	defer wg.Done()
	// time.Sleep(1 * time.Second)
	fmt.Printf("Linha %d processada: %s - %s\n", i, cifras[i%len(cifras)], letras[i])
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Caneta azul!</h1>")
	for i := range letras {
		fmt.Fprintf(w, "<p>%s - %s</p>", cifras[i%len(cifras)], letras[i])
	}
}

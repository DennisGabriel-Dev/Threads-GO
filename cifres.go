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
		"D7", "G#", "A7", "Em",
		"Em", "Am", "A", "Bm",
		"D2", "G6", "Am", "Cm",
		"D", "G7", "A",
	}
}

var (
	wg             sync.WaitGroup
	letras         = letras_func()
	cifras         = acordes_func()
	capture_cifras = make([]string, 100, 100)
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
	fmt.Printf("Linha %d processada: %s - %s\n", i, cifras[i%len(cifras)], letras[i])
	capture_cifras[i] = fmt.Sprintf("%s - %s\n", cifras[i%len(cifras)], letras[i])
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Caneta azul!</h1>")
	for _, element := range capture_cifras {
		fmt.Fprintf(w, "<p>%s</p>", element)
	}
}

package main

import "fmt"

type no struct {
	valor   int
	proximo *no
}

type lista struct {
	inicio *no
}

func adicionarInicio(l *lista, valor int) {
	novo := &no{valor: valor}

	novo.proximo = l.inicio
	l.inicio = novo
}

func adicionarFim(l *lista, valor int) {
	novo := &no{valor: valor}

	if l.inicio == nil {
		l.inicio = novo
		return
	}

	atual := l.inicio

	for atual.proximo != nil {
		atual = atual.proximo
	}

	atual.proximo = novo
}

func tamanho(l *lista) int {
	quantidade := 0
	atual := l.inicio

	for atual != nil {
		quantidade++
		atual = atual.proximo
	}

	return quantidade
}

func adicionarPosicao(l *lista, valor int, posicao int) bool {
	if posicao < 0 || posicao > tamanho(l) {
		return false
	}

	novo := &no{valor: valor}

	if posicao == 0 {
		novo.proximo = l.inicio
		l.inicio = novo
		return true
	}

	anterior := l.inicio

	for i := 0; i < posicao-1; i++ {
		anterior = anterior.proximo
	}

	novo.proximo = anterior.proximo
	anterior.proximo = novo

	return true
}

func removerInicio(l *lista) (int, bool) {
	if l.inicio == nil {
		return 0, false
	}

	valor := l.inicio.valor
	l.inicio = l.inicio.proximo

	return valor, true
}

func removerFim(l *lista) (int, bool) {
	if l.inicio == nil {
		return 0, false
	}

	if l.inicio.proximo == nil {
		valor := l.inicio.valor
		l.inicio = nil
		return valor, true
	}

	anterior := l.inicio
	atual := l.inicio.proximo

	for atual.proximo != nil {
		anterior = atual
		atual = atual.proximo
	}

	anterior.proximo = nil

	return atual.valor, true
}

func removerPosicao(l *lista, posicao int) (int, bool) {
	if posicao < 0 || posicao >= tamanho(l) {
		return 0, false
	}

	if posicao == 0 {
		return removerInicio(l)
	}

	anterior := l.inicio

	for i := 0; i < posicao-1; i++ {
		anterior = anterior.proximo
	}

	removido := anterior.proximo

	anterior.proximo = removido.proximo

	return removido.valor, true
}

func posicao(l *lista, valorProcurado int) (int, bool) {
	atual := l.inicio
	indice := 0

	for atual != nil {
		if atual.valor == valorProcurado {
			return indice, true
		}

		atual = atual.proximo
		indice++
	}

	return 0, false
}

func valorNaPosicao(l *lista, posicaoProcurada int) (int, bool) {
	if posicaoProcurada < 0 {
		return 0, false
	}

	atual := l.inicio
	indice := 0

	for atual != nil {
		if indice == posicaoProcurada {
			return atual.valor, true
		}

		atual = atual.proximo
		indice++
	}

	return 0, false
}

func imprimir(l *lista) {
	atual := l.inicio

	for atual != nil {
		fmt.Printf("%d -> ", atual.valor)
		atual = atual.proximo
	}

	fmt.Println("nil")
}

func main() {
	var l lista

	fmt.Println("EXERCÍCIO 2")

	adicionarInicio(&l, 30)
	adicionarInicio(&l, 20)
	adicionarInicio(&l, 10)

	adicionarFim(&l, 40)
	adicionarFim(&l, 50)

	imprimir(&l)

	fmt.Println("\nEXERCÍCIO 3")

	resultado := adicionarPosicao(&l, 50, 2)

	fmt.Println("Adicionou:", resultado)
	imprimir(&l)

	fmt.Println("\nEXERCÍCIO 4")

	valor, sucesso := removerInicio(&l)

	fmt.Println("Valor removido do início:", valor)
	fmt.Println("Sucesso:", sucesso)
	imprimir(&l)

	valor, sucesso = removerFim(&l)

	fmt.Println("Valor removido do fim:", valor)
	fmt.Println("Sucesso:", sucesso)
	imprimir(&l)

	fmt.Println("\nEXERCÍCIO 5")

	valor, sucesso = removerPosicao(&l, 1)

	fmt.Println("Valor removido:", valor)
	fmt.Println("Sucesso:", sucesso)
	imprimir(&l)

	fmt.Println("\nEXERCÍCIO 6")

	posicaoEncontrada, encontrado := posicao(&l, 30)

	fmt.Println("Posição do valor 30:", posicaoEncontrada)
	fmt.Println("Encontrado:", encontrado)

	fmt.Println("\nEXERCÍCIO 7")

	valor, encontrado = valorNaPosicao(&l, 1)

	fmt.Println("Valor na posição 1:", valor)
	fmt.Println("Encontrado:", encontrado)

	fmt.Println("\nEXERCÍCIO 8")

	fmt.Println("Tamanho da lista:", tamanho(&l))

	fmt.Println("Lista final:")
	imprimir(&l)
}

package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

/*
## Sobre o TDD

Extendendo a introdução do TDD, já descrita na ponderada da semana 1, a ponderada dessa semana mostra recursos adicionais para o desenvolvimento de testes. Alguns dos conceitos apresentados:

- Tests Suites: Objetos que agrupam testes, podem ser utilizados por agrupar cenários diferentes de uma mesma funcionalidade ou até mesmo funcionalidaes de um mesmo módulo
- TestMain: Função global (pode-se criar apenas 1 por pacote) onde detalhes gerais de configuração ou desmontagem dos testes podem ser configurados, como a criação de um banco de dados ou seu truncate.
- Estrutura de Projetos: O artigo mostra uma estrutura de projetos sugerida para que cada implementação seja acompanhada de seus testes.
*/

func TestMain(t *testing.T) {
    assert.Equal(t, 90, soma(45, 45), "45 + 45 deveria ser 90")
    assert.Equal(t, 1, soma(1, 0), "1 + 0 deveria ser 1")
    assert.Equal(t, 0, soma(-1, 1), "-1 + 1 deveria ser 0")
}
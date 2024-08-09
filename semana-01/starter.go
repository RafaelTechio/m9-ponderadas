package starter

import(
	"fmt"
	"math"
	"net/http"
)

// A primeira etapa do TDD é a construção de testes que falharão.
// Isso ocorre pois primeiro definimos o resultado esperado e a ação que a função executará, antes de
// definir sua implementação.

// A segunda etapa é a implementação da lógica das funções testadas até que passem nos cenários de teste.
// Dessa forma, garantimos que o código está executando da forma que gostaríamos, mesmo que não com a melhor
// implementação

// A terceira etapa é a refatoração, onde os ajustes de implementação devem ser realizados a fim de melhorar
// qualidade de código, performance, convenções etc. Isso pode gerar erros nos testes, que devem ser corrigidos.
// Dessa forma ocorre o ciclo de TDD.


// As implementações foram realizadas depois da construção inicial dos testes.
func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

func OddOrEven(num int) string {
	criteria := math.Mod(float64(num), 2)

	if criteria == 1 || criteria == -1 {
		return fmt.Sprintf("%v is an odd number", num)
	}

	return fmt.Sprintf("%v is an even number", num)
}

func Checkhealth(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "health check passed")
}
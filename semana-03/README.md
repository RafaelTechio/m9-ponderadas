# Ponderada M9 Semana 3
Rafael Techio

## Sobre o TDD

Extendendo a introdução do TDD, já descrita na ponderada da semana 1, a ponderada dessa semana mostra recursos adicionais para o desenvolvimento de testes. Alguns dos conceitos apresentados:

- **Tests Suites**: Objetos que agrupam testes, podem ser utilizados por agrupar cenários diferentes de uma mesma funcionalidade ou até mesmo funcionalidaes de um mesmo módulo:
```py
package mypackage

import (
    "testing"
    "github.com/stretchr/testify/suite"
)

// MySuite is a test suite containing multiple related tests.
type MySuite struct {
    suite.Suite
}

func (suite *MySuite) SetupTest() {
    // Setup code before each test
}

func (suite *MySuite) TestSomething() {
    // Test code
}

func TestSuite(t *testing.T) {
    // Run the test suite
    suite.Run(t, new(MySuite))
}
```

- **TestMain**: Função global (pode-se criar apenas 1 por pacote) onde detalhes gerais de configuração ou desmontagem dos testes podem ser configurados, como a criação de um banco de dados ou seu truncate:
```py
package main
import (
    "testing"
    "github.com/stretchr/testify/suite"
)

// TestMain sets up and tears down resources for all tests.
func TestMain(m *testing.M) {
    // Additional setup code here

    // Run all tests
    code := m.Run()

    // Additional teardown code here

    // Exit with the test result code
    os.Exit(code)
}
```

- **Estrutura de Projetos**: O artigo mostra uma estrutura de projetos sugerida para que cada implementação seja acompanhada de seus testes.
```
/myproject
|-- /handlers
|   |-- handler.go
|   |-- handler_test.go
|
|-- /db
|   |-- database.go
|   |-- database_test.go
|
|-- main.go
|-- go.mod
|-- go.sum
```
  - **handlers** : Manipuladores de solicitações HTTP e regras de negócio..
  - **db**: Arquivo com conexões de banco de dados.
  - **main.go** : Ponto de entrada principal para o aplicativo.
  - **go.mod** e **go.sum** : Arquivos de gerenciamento de dependências.

## Resultado de testes:

- Estrutura de Projetos:

![image](https://github.com/user-attachments/assets/3deb5285-007f-4ace-b7a7-2cb43cafd648)

- Execução de Testes:
Para testar a estrutura, uma simples função de soma foi criada:

![image](https://github.com/user-attachments/assets/c5652bac-8693-49fc-9c41-421cdecd7219)

Resultados:

![image](https://github.com/user-attachments/assets/691630c8-6625-4326-a2d6-3ce693ae4da9)

- Comentários do Código:

![image](https://github.com/user-attachments/assets/d85e98f5-b8dd-40ee-ae11-8a22a1129fd0)

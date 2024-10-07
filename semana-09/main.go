/*
Os testes de carga se relacionam com o tdd em um nível superior de maturidade, visto que seu objetivo é testar a qualidade do software, e não seu funcionamento em si. 
Dessa forma, o foco é a testagem de performance, e para isso, diversas iterações do TDD já devem ter ocorrido.
Contudo, após uma base funcional estabelicida, é totalmente possível utilizar de testes de carga no ciclo do tdd, tendo diversos benefícios. 
*/
package main

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("welcome"))
    })
    http.ListenAndServe(":3000", r)
}

package starter_test

import (
	"testing"
	"net/http/httptest"
	"io"
	"github.com/stretchr/testify/assert"
	starter "github.com/williaminfante/go_test_starter"
  )


// Testando diferentes cenários
func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("William")
	assert.Equal(t, "Hello William. Welcome!", greeting)
	
	another_greeting := starter.SayHello("asdf ghjkl")
  	assert.Equal(t, "Hello asdf ghjkl. Welcome!", another_greeting)
}

// Testando diferentes cenários em conjuntos 
func TestOddOrEven(t *testing.T) {
	t.Run("Check Non Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
		})
	t.Run("Check Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
	})
}

// Testando APIs como teste de integração 
func TestCheckhealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
	  req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
	  writer := httptest.NewRecorder()
	  starter.Checkhealth(writer, req)
	  response := writer.Result()
	  body, err := io.ReadAll(response.Body)
  
	  assert.Equal(t, "health check passed", string(body))
	  assert.Equal(t, 200, response.StatusCode)
	  assert.Equal(t,
		"text/plain; charset=utf-8",
		response.Header.Get("Content-Type"))
  
	  assert.Equal(t, nil, err)
	})
}
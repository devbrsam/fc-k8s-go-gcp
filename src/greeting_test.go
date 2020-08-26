package main
import "testing"
import "strings"

func TestGreeting( t *testing.T ) {

	deveRetornarTagBold := greeting("texto qualquer")
	
	if !strings.Contains(deveRetornarTagBold, "<b>") {
		t.Errorf("deveRetornarTagBold teste FALHOU: esperado conter a tag %v, mas retornou %v", "<b>", deveRetornarTagBold)
	} else {
		t.Logf("deveRetornarTagBold teste OK")
	}

	deveRetornarTexto := greeting("TEXTO_TESTE")
	
	if !strings.Contains(deveRetornarTexto, "TEXTO_TESTE") {
		t.Errorf("deveRetornarTexto teste FALHOU: esperado conter o texto %v, mas retornou %v", "TEXTO_TESTE", deveRetornarTexto)
	} else {
		t.Logf("deveRetornarTexto teste OK")
	}
}
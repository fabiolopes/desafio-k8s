package main

import "testing"

func TestGreeting(t *testing.T) {
	resultado := greeting()
	if resultado != "<b>Code.education Rocks!</b>" {
		t.Errorf("Texto esperado: %s, obtida: %s", "<b>Code.education Rocks!</b>", resultado)
	}
}

package properties

import "testing"

func TestFromFile(t *testing.T) {
	fn := "autoingestion.properties.example"
	p := NewPropertiesFromFile(fn)

	if p.Email != "user@example.com" {
		t.Errorf("expected %v atual %v", "user@example.com", p.Email)
	}

	if p.Password != "P@$$w0rd" {
		t.Errorf("expected %v atual %v", "P@$$w0rd", p.Email)
	}
}

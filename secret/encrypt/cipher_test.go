package encrypt

import "testing"

func TestCipher(t *testing.T) {
	key := "top secret key"
	plain := "test text"
	encrypted, err := Encrypt(key, plain)

	if err != nil {
		t.Errorf("error occurred while encyrpting text: %s", err)
	}

	decrypted, err := Decrypt(key, encrypted)

	if err != nil {
		t.Errorf("error occurred while decyrpting text: %s", err)
	}

	if decrypted != plain {
		t.Errorf("expected: %q, go %q", plain, decrypted)
	}
}

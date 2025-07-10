package hashing

import "testing"

// TestMD5 tests the MD5 hashing function.
func TestMD5(t *testing.T) {
	input := "hello world"
	expected := "5eb63bbbe01eeed093cb22bb8f5acdc3"

	md5Hash := MD5(input)

	if md5Hash != expected {
		t.Errorf("MD5 hash mismatch: got %s, want %s", md5Hash, expected)
	}
}

// TestSHA256 tests the SHA256 hashing function.
func TestSHA256(t *testing.T) {
	input := "hello world"
	expected := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"

	sha256Hash := SHA256(input)

	if sha256Hash != expected {
		t.Errorf("SHA256 hash mismatch: got %s, want %s", sha256Hash, expected)
	}
}

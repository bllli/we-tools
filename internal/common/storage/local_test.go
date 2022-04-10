package storage

import "testing"

func TestLocalStorage(t *testing.T) {
	storage, err := NewLocalStorage("/tmp/test", "http://localhost:8080")
	if err != nil {
		return
	}
	_, err = storage.GetUrl("wtf.txt")
	if err == nil {
		t.Error("Expected error")
	}
	uploadUrl, err := storage.Upload("test.txt", []byte("test"))
	if err != nil {
		t.Errorf("Upload failed: %s", err)
	}
	if uploadUrl != "http://localhost:8080/test.txt" {
		t.Errorf("Upload URL is wrong: %s", uploadUrl)
	}

	downloadUrl, err := storage.GetUrl("test.txt")
	if err != nil {
		t.Errorf("GetUrl failed: %s", err)
	}
	if downloadUrl != "http://localhost:8080/test.txt" {
		t.Errorf("Download URL is wrong: %s", downloadUrl)
	}
	data, err := storage.Get("test.txt")
	if err != nil {
		t.Errorf("Get failed: %s", err)
	}
	if string(data) != "test" {
		t.Errorf("Get data is wrong: %s", data)
	}
}

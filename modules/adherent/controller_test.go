package adherent

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUpController() {
	NewAdherent(&Adherent{Firstname: "Test1", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	NewAdherent(&Adherent{Firstname: "Test2", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	NewAdherent(&Adherent{Firstname: "Test3", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	NewAdherent(&Adherent{Firstname: "Test4", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
}

func TestAdherentsIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/adherents", nil)
	w := httptest.NewRecorder()
	AdherentsIndex(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestAdherentsCreate(t *testing.T) {
	var jsonStr = []byte(`{"firstname": "Luigi", "lastname": "Bros", "dateofbirth": "1995-04-24"}`)

	req := httptest.NewRequest("POST", "/adherents", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	AdherentsCreate(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

// func TestAdherentsShow(t *testing.T) {
// 	req, _ := http.NewRequest("Get", "/adherents/2", nil)

// 	w := httptest.NewRecorder()
// 	AdherentsShow(w, req)

// 	res := w.Result()
// 	defer res.Body.Close()
// 	_, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Errorf("expected error to be nil got %v", err)
// 	}
// }

// func TestAdherentsUpdate(t *testing.T) {

// }

// func TestAdherentsDelete(t *testing.T) {

// }

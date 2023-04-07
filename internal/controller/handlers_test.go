package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"restapi/internal/entity"
	"restapi/internal/usecase"
)

func TestFindStringHandler(t *testing.T) {
	data := []struct {
		Data     string
		Expected string
	}{
		{"asdasd", "asd"},
		{"mmmmmm", "m"},
		{"11a2bc3de", "1a2bc3de"},
		{"", ""},
		{"112233", "12"},
		{"edde", "ed"},
	}

	for _, d := range data {
		obj, err := json.Marshal(entity.StringRequest{
			Text: d.Data,
		})
		if err != nil {
			t.Errorf("error: test find string handler: marshalling: %v\n", err)
		}
		req, err := http.NewRequest(http.MethodPost, "/rest/substr/find", strings.NewReader(string(obj)))
		if err != nil {
			t.Fatalf("error while sending request: %e", err)
		}
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		h := &Handler{
			u: usecase.NewUsecase(),
		}
		h.findStringHandler(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("error: expected status code %d, but got %d", http.StatusOK, res.Code)
		}
		var result entity.StringRequest
		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			t.Errorf("error: test find string handler: marshalling: %v\n", err)
		}

		if result.Text != d.Expected {
			t.Errorf("Request body: %s \n Expected response body '%s', but got '%s'", d.Data, d.Expected, res.Body.String())
		}

	}
}

func TestFindEmailHandler(t *testing.T) {
	data := []struct {
		Data     string
		Expected string
	}{
		{"email@mail.ru", `email@mail.ru`},
		{"email@mail.ru,email@mail.ru", `email@mail.ru,email@mail.ru`},
		{"mail@inbox.ru\nexample@bk.ru", `mail@inbox.ru,example@bk.ru`},
		{"example@bk.ru example@gmail.com", `example@bk.ru,example@gmail.com`},
	}
	for _, d := range data {
		obj, err := json.Marshal(entity.EmailRequest{
			Emails: []string{d.Data},
		})
		if err != nil {
			t.Errorf("error: test find string handler: marshalling: %v\n", err)
		}
		req, err := http.NewRequest(http.MethodPost, "/rest/email/check", strings.NewReader(string(obj)))
		if err != nil {
			t.Fatalf("error while sending request: %e", err)
		}

		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		h := &Handler{
			u: usecase.NewUsecase(),
		}
		h.findEmailHandler(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("error: expected status code %d, but got %d", http.StatusOK, res.Code)
		}

		var result entity.EmailRequest
		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			t.Errorf("error: test find string handler: marshalling: %v\n", err)
		}
		results := strings.Join(result.Emails, ",")
		if results != d.Expected {
			t.Errorf("Request body: %s \n Expected response body '%s', but got '%s'", d.Data, d.Expected, res.Body.String())
		}

	}
}

func TestFindIinHandler(t *testing.T) {
	data := []struct {
		Data     uint64
		Expected string
	}{
		{9406214534244563, "940621453424"},
		{94062145123333, "940621451233"},
	}

	for _, d := range data {
		obj, err := json.Marshal(entity.IinRequest{
			IIN: d.Data,
		})
		if err != nil {
			t.Errorf("error: test find string handler: marshalling: %v\n", err)
		}
		req, err := http.NewRequest(http.MethodPost, "/rest/iin/check", strings.NewReader(string(obj)))
		if err != nil {
			t.Fatalf("error while sending request: %e", err)
		}

		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		h := &Handler{
			u: usecase.NewUsecase(),
		}
		h.findIinHandler(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("error: expected status code %d, but got %d", http.StatusOK, res.Code)
		}

		var result []string
		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			t.Errorf("error: test find string handler: marshalling: %v\n", err)
		}
		results := strings.Join(result, ",")
		if results != d.Expected {
			t.Errorf("Request body: %d \n Expected response body '%s', but got '%s'", d.Data, d.Expected, res.Body.String())
		}

	}
}

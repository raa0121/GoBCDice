package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/helpers"
)

func TestGetRoot(t *testing.T) {
	s := S{}
	s.SetUpSuite(nil)

	rec := s.PerformRequest("GET", "/", url.Values{})

	if rec.Code != http.StatusOK {
		t.Errorf("wrong code: got=%v want=%v", rec.Code, http.StatusOK)
	}

	bodyStr := rec.Body.String()
	dec := json.NewDecoder(strings.NewReader(bodyStr))

	var r helpers.ResponseMap
	err := dec.Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	expected := helpers.ResponseMap{
		"ok":      true,
		"message": "Hello. This is GoBCDiceAPI.",
	}

	if !reflect.DeepEqual(r, expected) {
		t.Errorf("wrong response: got=%+v, want=%+v", r, expected)
	}
}

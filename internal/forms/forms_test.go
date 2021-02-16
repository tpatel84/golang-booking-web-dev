package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it doesn't exist")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("Shown form doesn't have a field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non existent page")
	}

	// Code coverage for Get function errors.go
	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("Should have an error but didn't get one")
	}

	postedData := url.Values{}
	postedData.Add("some_field", "some_value")
	form = New(postedData)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows min length of 100 when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "abc123")
	form = New(postedData)

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("length of field is shorter than 1")
	}

	// Code coverage for Get function errors.go
	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("Should not have an error but got one")
	}
}

func TestForm_IsValidEmail(t *testing.T) {
	//r := httptest.NewRequest("POST", "/whatever", nil)
	//form := New(r.PostForm)

	postedData := url.Values{}
	form := New(postedData)

	form.IsValidEmail("x")
	if form.Valid() {
		t.Error("forms shows valid email for not valid field")
	}

	postedData = url.Values{}
	postedData.Add("email", "me@here.com")
	form = New(postedData)

	form.IsValidEmail("email")
	if !form.Valid() {
		t.Error("form shows non valid email for existing field")
	}

	postedData = url.Values{}
	postedData.Add("email", "x")
	form = New(postedData)

	form.IsValidEmail("email")
	if form.Valid() {
		t.Error("form shows valid email for invalid email/field")
	}

}
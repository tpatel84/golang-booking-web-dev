package renders

import (
	"github.com/tpatel84/golang-booking-web-dev/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	// http.request requires session object
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	// Add key and value into session for test
	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func getSession() (*http.Request, error) {
	// Create dummy GET request
	r, err := http.NewRequest("GET", "/some_url", nil)
	if err != nil {
		return nil, err
	}
	// Get context from http.request
	ctx := r.Context()
	// Get session data for given token
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	// Put context back into request
	r = r.WithContext(ctx)

	return r, nil
}

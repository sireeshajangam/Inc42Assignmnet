// backend/tests/unit/handlers_test.go

package unit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"example.com/myapp/backend"
)

func TestGetPostsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(backend.GetPostsHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "First Post")
	assert.Contains(t, rr.Body.String(), "Second Post")
}

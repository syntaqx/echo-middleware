package requestid

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TextRequestID(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	mw := New()

	// ensure the default values are set properly
	assert.Equal(t, mw.HeaderKey, XRequestID)
	assert.Equal(t, mw.Generate, generateID)

	mw.handleActualRequest(res, req)

	// ensure a random request id is assigned
	assert.NotEmpty(t, req.Header.Get(mw.HeaderKey))
}

func TestRequestIDCustom(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	testGenValue := "test"
	testHeaderKey := "Custom-Header-Key"
	testGenerate := func() (string, error) { return testGenValue, nil }

	mw := New()

	assert.NotEqual(t, testHeaderKey, mw.HeaderKey)

	mw.SetHeaderKey(testHeaderKey)
	mw.SetGenerate(testGenerate)

	mw.handleActualRequest(res, req)

	assert.Equal(t, testHeaderKey, mw.HeaderKey)
	assert.Equal(t, testGenValue, req.Header.Get(mw.HeaderKey))
}

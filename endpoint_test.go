package slimhttp

import (
	"net/http/httptest"
	"testing"
)

func TestTextEndpointWrapper(t *testing.T) {
	wr := httptest.NewRecorder()
	handler := endpointWrapper(newTestTextEndpoint, encodeText)
	handler(wr, nil)
	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, "Here's some text!")
}

func TestJSONEndpointWrapper(t *testing.T) {
	wr := httptest.NewRecorder()
	handler := endpointWrapper(newTestStructEndpoint, encodeJSON)
	handler(wr, nil)
	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, `{"string_key":"string-val","int_key":5,"float_key":13.37}`+"\n")
}

func TestXMLEndpointWrapper(t *testing.T) {
	wr := httptest.NewRecorder()
	handler := endpointWrapper(newTestStructEndpoint, encodeXML)
	handler(wr, nil)
	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, `<?xml version="1.0" encoding="UTF-8"?>`+"\n"+`<testStruct><stringKey>string-val</stringKey><intKey>5</intKey><floatKey>13.37</floatKey></testStruct>`)
}

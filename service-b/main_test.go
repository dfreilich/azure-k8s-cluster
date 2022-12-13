package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndexEndpoint(t *testing.T) {
	server := getHandler()
	ts := httptest.NewServer(server)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	require.NoError(t, err)

	require.Equal(t, resp.StatusCode, http.StatusOK)

	greeting, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, resp.Body.Close())
	require.NoError(t, err)

	require.Equal(t, string(greeting), "Hello World")
}

func TestHealthzEndpoint(t *testing.T) {
	server := getHandler()
	ts := httptest.NewServer(server)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/healthz")
	require.NoError(t, err)

	require.Equal(t, resp.StatusCode, http.StatusOK)

	greeting, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, resp.Body.Close())
	require.NoError(t, err)

	require.Equal(t, string(greeting), "OK")
}

func TestReadyzEndpoint(t *testing.T) {
	server := getHandler()
	ts := httptest.NewServer(server)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/readyz")
	require.NoError(t, err)

	require.Equal(t, resp.StatusCode, http.StatusOK)

	greeting, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, resp.Body.Close())
	require.NoError(t, err)

	require.Equal(t, string(greeting), "OK")
}

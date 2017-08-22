
package main

import (
"net/http"
"testing"

. "github.com/emicklei/forest"
)

var github = NewClient("http://localhost:8000", new(http.Client))

func TestForestProjectExists(t *testing.T) {
	cfg := NewConfig("/v1/user/1", "forest").Header("Accept", "application/json")
	r := github.GET(t, cfg)
	ExpectStatus(t, r, 200)
}

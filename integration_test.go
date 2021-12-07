package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	host = "https://exercise-loan-calculator.ts.r.appspot.com/"
)

func init() {
	if os.Getenv("CI") == "" {
		host = "http://localhost:8080"
	}
}

func TestServerUp(t *testing.T) {
	resp, err := http.Get(host)
	assert.NoError(t, err)
	bs, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, "Hello World!", string(bs))
}

func TestLoanCalculatorError(t *testing.T) {
	resp, err := http.Post(host+"", "", strings.NewReader("error"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest+5, resp.StatusCode)
}

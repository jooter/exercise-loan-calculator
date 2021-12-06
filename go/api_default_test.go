package swagger

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateLoan(t *testing.T) {
	reqStruct := &CalculateloanBody{LoanType: "", InterestRate: 2, LoanTerm: 30}
	reqBytes, _ := json.Marshal(reqStruct)

	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBytes))
	w := httptest.NewRecorder()

	CalculateLoan(w, r)
	r.Body.Close()

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, `{"amount_owing":[{"interest":3}]}`, w.Body.String())
}

func TestCalculateLoanError(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("error"))
	w := httptest.NewRecorder()

	CalculateLoan(w, r)
	r.Body.Close()

	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	assert.Equal(t, `{"code":"400"}`, w.Body.String())
}

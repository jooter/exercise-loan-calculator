/*
 * Loan Calculator
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CalculateLoan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dec := json.NewDecoder(r.Body)
	var req CalculateloanBody
	err := dec.Decode(&req)
	if err != nil {
		log.Println(err)
		badRequest(w)
		return
	}

	repayment := &LoanRepayments{}
	repayment.AmountOwing = append(repayment.AmountOwing, LoanRepaymentsAmountOwing{Interest: 3})
	repaymentBody, err := json.Marshal(repayment)
	if err != nil {
		log.Println(err)
		badRequest(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(repaymentBody))
}

func badRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	e := &ModelError{Code: "400"}
	b, _ := json.Marshal(e)
	fmt.Fprint(w, string(b))
}

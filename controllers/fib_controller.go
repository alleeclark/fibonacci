package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"fibonacci/common"
)

// PostFibSequenceNums returns the sequence of ints in an array
// Handler for HTTP Post /fib
func PostFibSequenceNums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in PostFibSequenceNums", r)
			common.DisplayAppError(
				w,
				errors.New("Negative int entered"),
				"Invalid data",
				500,
			)
			return
		}
	}()
	var f UserInput
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid data",
			500,
		)
		return
	}
	err = json.Unmarshal(body, &f)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	if f.Number < 0 {
		common.DisplayAppError(
			w,
			err,
			"Value should not be less than 0",
			500,
		)
		return
	} else {
		resp := GetFibSequence(f.Number)
		j, _ := json.Marshal(resp)
		w.Write(j)
	}
}

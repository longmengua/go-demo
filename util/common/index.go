package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/shopspring/decimal"
)

// add, sub, mul, div
func Math(op string, s1 string, s2 string) (string, error) {
	var toReturn = ""
	n1, e1 := decimal.NewFromString(s1)
	n2, e2 := decimal.NewFromString(s2)

	if e2 != nil {
		return "", e1
	}
	if e2 != nil {
		return "", e2
	}

	switch op {
	case "add":
		toReturn = n1.Add(n2).String()

	case "sub":
		toReturn = n1.Sub(n2).String()

	case "mul":
		toReturn = n1.Mul(n2).String()

	case "div":
		toReturn = n1.Div(n2).String()
	}
	if toReturn == "" {
		return "", errors.New("error in Math func")
	}
	return toReturn, nil
}

func HttpRequest(
	url string,
	method string,
	data interface{},
) (resp *http.Response, err error) {
	switch strings.ToLower(method) {
	case "put":
	case "delete":
	case "post":
		bytes := new(bytes.Buffer)
		json.NewEncoder(bytes).Encode(data)
		return http.Post(url, "application/json", bytes)
	case "get":
		v, _ := query.Values(data)
		u := fmt.Sprintf("%s?%s", url, v.Encode())
		return http.Get(u)
	}
	return nil, errors.New("not supported method")
}

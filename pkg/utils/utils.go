package utils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func WriteXMLResponse(w http.ResponseWriter, httpStatus int, v interface{}) error {
	body, err := xml.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal to XML: %v", err)
	}
	w.WriteHeader(httpStatus)
	_, err = w.Write(body)
	if err != nil {
		return fmt.Errorf("failed to write data to connection: %v", err)
	}
	return nil
}

func WriteJSONResponse(w http.ResponseWriter, httpStatus int, v interface{}) {
	body, err := json.Marshal(v)
	if err != nil {
		log.Printf("failed to marshal to JSON: %v", err)
		return
	}
	w.WriteHeader(httpStatus)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("failed to write data to connection: %v", err)
		return
	}
}

func CheckToken(expectedToken, Token string) bool {
	var re = regexp.MustCompile(`(?m)^\s*(Bearer)\s+(` + regexp.QuoteMeta(Token) + `)\s*$`)

	return re.MatchString(expectedToken)
}

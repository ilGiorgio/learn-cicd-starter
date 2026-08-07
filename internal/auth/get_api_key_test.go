package auth

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	var headers http.Header = make(http.Header)

	headers.Add("Authorization", "ApiKey ey10dkdktoken")
	headers.Add("Content-Type", "application/json")

	fmt.Println("Headers: ", headers)

	exp := "ey10dkdktoken"
	got, err := GetAPIKey(headers)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Gotten: ", got)

	if !reflect.DeepEqual(exp, got) {
		t.Fatalf("expected: %v, got: %v", exp, got)
	}
}

package farewell

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GreetInSpanish(ctx context.Context, name string) (string, error) {
	greeting, err := callService("get-spanish-greeting", name)
	return greeting, err
}

func FarewellInSpanish(ctx context.Context, name string) (string, error) {
	farewell, err := callService("get-spanish-farewell", name)
	return farewell, err
}

// TODO: write an Activity function that calls the microservice to
// get a farewell message in Spanish. It will be identical to the
// function above, except the first argument to the callService
// function will be "get-spanish-farewell". You can name your
// function whatever you like.

// utility function for making calls to the microservices
func callService(stem string, name string) (string, error) {
	base := "http://localhost:9999/" + stem + "?name=%s"
	url := fmt.Sprintf(base, url.QueryEscape(name))

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	translation := string(body)
	status := resp.StatusCode
	if status >= 400 {
		return "", fmt.Errorf("HTTP Error %d: %s", status, translation)
	}

	return translation, nil
}

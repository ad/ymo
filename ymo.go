package ymo

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

// Send offline conversion to metrika server
// https://yandex.ru/dev/metrika/doc/api2/management/offline_conversion/upload.html

// NewYMOClient creates a new YMOClient object
func NewYMOClient(counter, token, clientType string, debug bool) (*YMOClient, error) {
	_, err := getClientTypeHeader(clientType)
	if err != nil {
		return nil, fmt.Errorf("unknown client_type error: %s", err)
	}

	return &YMOClient{
		token:      token,
		counter:    counter,
		clientType: clientType,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		debug: debug,
	}, nil
}

// SendEvent sends one event to metrika
func (g *YMOClient) SendEvent(event Event) error {
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	fileWriter, err := writer.CreateFormFile("file", "userid.csv")
	if err != nil {
		return fmt.Errorf("writer createFromfile error: %s", err)
	}

	w := csv.NewWriter(fileWriter)

	clientTypeHeader, err := getClientTypeHeader(g.clientType)
	if err != nil {
		return fmt.Errorf("unknown client_type error: %s", err)
	}

	if event.DateTime == "" {
		event.DateTime = strconv.FormatInt(time.Now().Unix(), 10)
	}

	if event.Price == "" {
		event.Price = "0"
	}

	if event.Currency == "" {
		event.Currency = "RUB"
	}

	records := [][]string{}
	records = append(records, []string{clientTypeHeader, "Target", "DateTime", "Price", "Currency"})
	records = append(records, []string{event.ClientId, event.Target, event.DateTime, event.Price, event.Currency})

	err = w.WriteAll(records) // calls Flush internally
	if err != nil {
		return fmt.Errorf("csv.WriteAll error: %s", err)
	}

	writer.Close()

	if !g.debug {
		fmt.Printf("[DEBUG] request message: %q", body.String())
	}

	endpoint := fmt.Sprintf("https://api-metrika.yandex.net/management/v1/counter/%s/offline_conversions/upload?client_id_type=%s", g.counter, g.clientType)

	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(body.Bytes()))
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary())
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Length", fmt.Sprintf("%d", body.Len()))
	req.Header.Set("Authorization", fmt.Sprintf("OAuth %s", g.token))

	client := &http.Client{Timeout: 6 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	defer func(resp *http.Response) {
		_ = resp.Body.Close()
	}(resp)

	if !g.debug {
		return nil
	}

	b, err := io.ReadAll(resp.Body)
	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent || resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusForbidden) || (err != nil) {
		return fmt.Errorf("HTTPStatusCode: '%d'; ResponseMessage: '%s'; ErrorMessage: '%s'", resp.StatusCode, string(b), err)
	}

	fmt.Printf("[DEBUG] get metrika response %q", string(b))

	return nil
}

func getClientTypeHeader(clientType string) (string, error) {
	switch clientType {
	case "CLIENT_ID":
		return "ClientId", nil
	case "USER_ID":
		return "UserId", nil
	case "YCLID":
		return "Yclid", nil
	}

	return "", fmt.Errorf("unknown clientType: %s", clientType)
}

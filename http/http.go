package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Publisher struct {
	Client     *http.Client
	Url        string
	LogError   func(context.Context, string)
	Goroutines bool
	Retries    []time.Duration
}

func NewPublisher(client *http.Client, url string, logError func(context.Context, string), goroutines bool, retries ...time.Duration) *Publisher {
	return &Publisher{Client: client, Url: url, LogError: logError, Goroutines: goroutines, Retries: retries}
}
func (s *Publisher) Publish(ctx context.Context, data []byte) error {
	if s.Goroutines {
		go postLog(ctx, s.Client, s.Url, data, s.LogError, s.Retries...)
		return nil
	} else {
		return postLog(ctx, s.Client, s.Url, data, s.LogError, s.Retries...)
	}
}
func postLog(ctx context.Context, client *http.Client, url string, log []byte, logError func(context.Context, string), retries ...time.Duration) error {
	l := len(retries)
	if l == 0 {
		_, err := post(ctx, client, url, log)
		return err
	} else {
		return postWithRetries(ctx, client, url, log, logError, retries)
	}
}
func postWithRetries(ctx context.Context, client *http.Client, url string, log []byte, logError func(context.Context, string), retries []time.Duration) error {
	_, er1 := post(ctx, client, url, log)
	if er1 == nil {
		return er1
	}
	i := 0
	err := retry(ctx, retries, func() (err error) {
		i = i + 1
		_, er2 := post(ctx, client, url, log)
		if er2 == nil && logError != nil {
			logError(ctx, fmt.Sprintf("Send log successfully after %d retries %s", i, log))
		}
		return er2
	}, logError)
	if err != nil && logError != nil {
		logError(ctx, fmt.Sprintf("Failed to send log after %d retries: %s. Error: %s.", len(retries), log, err.Error()))
	}
	return err
}
func post(ctx context.Context, client *http.Client, url string, body []byte) (*json.Decoder, error) {
	res, er1 := do(ctx, client, url, "POST", body)
	if er1 != nil {
		return nil, er1
	}
	if res.StatusCode == 503 {
		er2 := errors.New("503 Service Unavailable")
		return nil, er2
	}
	return json.NewDecoder(res.Body), nil
}
func do(ctx context.Context, client *http.Client, url string, method string, body []byte) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	return addHeaderAndDo(client, req)
}
func addHeaderAndDo(client *http.Client, req *http.Request) (*http.Response, error) {
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	return resp, err
}

// retry Copy this code from https://stackoverflow.com/questions/47606761/repeat-code-if-an-error-occured
func retry(ctx context.Context, sleeps []time.Duration, f func() error, log func(context.Context, string)) (err error) {
	attempts := len(sleeps)
	for i := 0; ; i++ {
		err = f()
		if err == nil {
			return
		}
		if i >= (attempts - 1) {
			break
		}
		if log != nil {
			log(ctx, fmt.Sprintf("Retrying %d of %d after error: %s", i+1, attempts, err.Error()))
		}
		time.Sleep(sleeps[i])
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

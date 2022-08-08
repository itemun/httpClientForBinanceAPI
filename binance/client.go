package binance

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

/*func (c Client) GetAssets() ([]Asset, error) {

	resp, err := c.client.Get("https://api.binance.com/api/v3/ticker/24hr")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r Asset //assetsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil //r.Assets, nil

}*/

func (c Client) GetAsset(name string) (Asset, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/24hr?symbol=%s", name)
	resp, err := c.client.Get(url)
	if err != nil {
		return Asset{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Asset{}, err
	}

	var r Asset
	if err = json.Unmarshal(body, &r); err != nil {
		return Asset{}, err
	}

	return r, nil

}

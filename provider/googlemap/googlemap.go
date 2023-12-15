package googlemap

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"paradise-booking/config"
	"paradise-booking/utils"
	"time"

	"moul.io/http2curl"
)

type GoogleMapAddress struct {
	Country  string `json:"country"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
}

type GoogleMap struct {
	cfg    *config.Config
	client *http.Client
}

func NewGoogleMap(cfg *config.Config) *GoogleMap {
	return &GoogleMap{
		cfg: cfg,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (g *GoogleMap) GetGeocodeMap(ctx context.Context, lat, lng float32) (*GoogleMapAddress, error) {
	latLngValue := fmt.Sprintf("%f,%f", lat, lng)
	var address *GoogleMapAddress
	path := fmt.Sprintf("/json?latlng=%s&key=%s", latLngValue, g.cfg.GoogleMap.APIKey)
	url := utils.JoinURL(g.cfg.GoogleMap.BaseURL, path)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	body, err := g.MakeRequest(ctx, g.client, req)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &address); err != nil {
		return nil, err
	}

	return address, nil
}

func (g *GoogleMap) MakeRequest(ctx context.Context, httpClient *http.Client, req *http.Request) ([]byte, error) {

	command, _ := http2curl.GetCurlCommand(req)
	fmt.Print(command)

	resp, err := httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body. ", err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Response data: ", resp)
	fmt.Println("Body: ", string(body))
	fmt.Println("Outbound call completed with statusCode ", resp.StatusCode)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Printf("Return status is not 200. Got %d", resp.StatusCode)
		return body, fmt.Errorf("response status code is %d", resp.StatusCode)
	}

	return body, nil
}

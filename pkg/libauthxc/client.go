package libauthxc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseUrl string
}

func NewClient(host string, port int) *Client {
	return &Client{
		BaseUrl: fmt.Sprintf("%s:%d", host, port),
	}
}

func (c *Client) GetProfileByEhid(ehid string) (*GetResponseDto, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/profiles/%s", c.BaseUrl, ehid),
		nil,
	)
	if err != nil {
		return nil, err
	}

	data := GetResponseDto{}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

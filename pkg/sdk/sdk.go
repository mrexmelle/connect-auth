package libauthxc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mrexmelle/connect-authx/internal/profile"
)

type Client struct {
	BaseUrl string
}

func NewClient(host string, port int) *Client {
	return &Client{
		BaseUrl: fmt.Sprintf("%s:%d", host, port),
	}
}

func (c *Client) GetProfilesByEhid(ehid string) (*profile.GetResponseDto, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s%s/%s", c.BaseUrl, "/profiles", ehid),
		nil,
	)
	if err != nil {
		return nil, err
	}

	data := profile.GetResponseDto{}
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

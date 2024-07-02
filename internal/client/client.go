package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/AnthonyFVKT/book-illustrator-srv/internal/model"
	"io"
	"net/http"
)

const (
	endpoint        = "https://api.goapi.ai/mj/v2/imagine"
	checkEndpoint   = "https://api.midjourneyapi.xyz/mj/v2/fetch"
	upscaleEndpoint = "https://api.midjourneyapi.xyz/mj/upscale"
)

type Client struct {
	Key string `json:"key"`
}

func NewClient(key string) *Client {
	return &Client{key}
}

func (c *Client) makeRequest(ctx context.Context, method, url string, data interface{}) ([]byte, error) {
	client := http.DefaultClient

	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", c.Key)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	return io.ReadAll(resp.Body)
}

func (c *Client) RequestImage(ctx context.Context, prompt string) ([]byte, error) {
	data := model.ImageRequest{
		Prompt:          prompt,
		AspectRatio:     "4:3", // maybe will add config later
		ProcessMode:     "mixed",
		WebhookEndpoint: "",
		WebhookSecret:   "",
	}

	return c.makeRequest(ctx, http.MethodPost, endpoint, data)
}

func (c *Client) UpscaleImage(ctx context.Context, taskId, index string) ([]byte, error) {
	data := model.UpscaleRequest{
		TaskId: taskId,
		Index:  index,
	}

	return c.makeRequest(ctx, http.MethodPost, upscaleEndpoint, data)
}

func (c *Client) Check(ctx context.Context, taskId string) ([]byte, error) {
	data := model.CheckRequest{
		TaskId: taskId,
	}

	return c.makeRequest(ctx, http.MethodPost, checkEndpoint, data)
}

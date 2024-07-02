package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AnthonyFVKT/book-illustrator-srv/internal/client"
	"github.com/AnthonyFVKT/book-illustrator-srv/internal/model"
	"time"
)

const delay = time.Second * 5

type Illustrator struct {
	MJClient *client.Client
}

func NewIllustrator(client *client.Client) *Illustrator {
	return &Illustrator{MJClient: client}
}

func (s *Illustrator) MakeImageByText(ctx context.Context, prompt string, upscale bool) (string, error) {
	// request image by text
	resp, err := s.MJClient.RequestImage(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("s.MJClient.RequestImage: %w", err)
	}

	respData := &model.ImageResponse{}
	if err := json.Unmarshal(resp, respData); err != nil {
		return "", fmt.Errorf("json.Unmarshal: %w", err)
	}

	fmt.Println(string(resp)) //TODO: remove

	// wait till service complete image generation
	result, err := s.waitForTaskCompletion(ctx, respData.TaskId)
	if err != nil {
		return "", err
	}

	fmt.Println(result) //TODO: remove

	if upscale {
		fmt.Println("Upscale image...")
		// upscale final result of image
		resp, err = s.MJClient.UpscaleImage(ctx, respData.TaskId, "1") // service provide 4 variants of image, we select 1st
		if err != nil {
			return "", err
		}

		respUpscaleData := &model.UpscaleResponse{}
		if err := json.Unmarshal(resp, respUpscaleData); err != nil {
			return "", fmt.Errorf("json.Unmarshal: %w", err)
		}

		// wait till service complete image up scaling
		result, err = s.waitForTaskCompletion(ctx, respUpscaleData.TaskId)
		if err != nil {
			return "", err
		}
	}

	return result, nil
}

func (s *Illustrator) waitForTaskCompletion(ctx context.Context, taskId string) (string, error) {
	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-ticker.C:
			checkData, err := s.MJClient.Check(ctx, taskId)
			if err != nil {
				return "", err
			}

			respCheckData := &model.CheckResponse{}
			if err := json.Unmarshal(checkData, respCheckData); err != nil {
				return "", fmt.Errorf("json.Unmarshal: %w", err)
			}

			if respCheckData.Status == "finished" {
				return respCheckData.TaskResult.ImageUrl, nil
			}
		}
	}
}

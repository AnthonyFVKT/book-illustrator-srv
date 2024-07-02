package model

type ImageRequest struct {
	Prompt          string `json:"prompt"`
	AspectRatio     string `json:"aspect_ratio"`
	ProcessMode     string `json:"process_mode"`
	WebhookEndpoint string `json:"webhook_endpoint"`
	WebhookSecret   string `json:"webhook_secret"`
}

type ImageResponse struct {
	TaskId  string `json:"task_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UpscaleRequest struct {
	TaskId string `json:"origin_task_id"`
	Index  string `json:"index"`
}

type UpscaleResponse struct {
	TaskId  string `json:"task_id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CheckRequest struct {
	TaskId string `json:"task_id"`
}

type CheckResponse struct {
	Status     string `json:"status"`
	TaskResult struct {
		ImageUrl string `json:"image_url"`
	} `json:"task_result"`
}

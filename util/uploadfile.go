package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const apiEndpoint = "https://api.apilayer.com/resume_parser/upload"
const apiKey = "gNiXyflsFu3WNYCz1ZCxdWDb7oQg1Nl1"

type ResumeResponse struct {
	// Define the expected structure based on the API's response
	FilePath   string `json:"file_path"`
	Skills     string `json:"skills"`
	Education  string `json:"education"`
	Experience string `json:"experience"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

func UploadResumeToAPI(fileBytes []byte) (*ResumeResponse, error) {
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewReader(fileBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to upload resume")
	}

	var resumeResp ResumeResponse
	if err := json.NewDecoder(resp.Body).Decode(&resumeResp); err != nil {
		return nil, err
	}

	return &resumeResp, nil
}

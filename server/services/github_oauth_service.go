package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rakutentech/code-coverage-dashboard/config"
)

type GithubOAuthService struct {
}

// NewGithubOAuthService creates a new GithubOAuthService
func NewGithubOAuthService() *GithubOAuthService {
	return &GithubOAuthService{}
}

// VerifyGithubToken verifies the github token
// /repos/:owner/:repo/statuses/commit_hash
func (v *GithubOAuthService) VerifyGithubToken(api, token, orgName, repoName, commitHash string) error {
	conf := config.NewConfig()
	if conf.AppConfig.AppEnv == "local" {
		return nil
	}
	if token == "" {
		return fmt.Errorf("github token is required")
	}

	url := fmt.Sprintf("%s/repos/%s/%s/statuses/%s", api, orgName, repoName, commitHash)
	log.Print("url: ", url)

	jsonStr := `{"state":"success", "context": "code coverage dashboard", "description": "code coverage authenticated"}`
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error: %s", err.Error())
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("closing body error: %s", err.Error())
		}
	}()
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("error: %s", resp.Status)
	}

	return nil
}

func (v *GithubOAuthService) GetGithubAccessToken(code string) (string, error) {
	conf := config.NewConfig()

	// Set us the request body as JSON
	requestBodyMap := map[string]string{
		"client_id":     conf.GithubConfig.ClientID,
		"client_secret": conf.GithubConfig.ClientSecret,
		"code":          code,
	}
	requestJSON, err := json.Marshal(requestBodyMap)
	if err != nil {
		log.Println("Error: Request body marshal failed")
		return "", err
	}

	// POST request to set URL
	req, err := http.NewRequest(
		"POST",
		conf.GithubConfig.GithubURL+"/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)
	if err != nil {
		log.Println("Error: Request creation failed")
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Get the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error: Request creation failed")
		return "", err
	}

	// Response body converted to stringified JSON
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error: Response read failed")
		return "", err
	}

	// Represents the response received from Github
	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	// Convert stringified JSON to a struct object of type githubAccessTokenResponse
	var ghresp githubAccessTokenResponse
	err = json.Unmarshal(respbody, &ghresp)
	if err != nil {
		log.Println("Error: Response unmarshal failed")
		return "", err
	}

	// Return the access token (as the rest of the
	// details are relatively unnecessary for us)
	return ghresp.AccessToken, nil
}

func (v *GithubOAuthService) GetGithubUser(accessToken string) (string, error) {
	conf := config.NewConfig()
	// Get request to a set URL
	req, err := http.NewRequest(
		"GET",
		conf.GithubConfig.GithubApiURL+"/user",
		nil,
	)
	if err != nil {
		log.Println("Error: Request build failed")
		return "", err
	}

	// Set the Authorization header before sending the request
	// Authorization: token XXXXXXXXXXXXXXXXXXXXXXXXXXX
	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	// Make the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error: Request failed")
		return "", err
	}

	// Read the response as a byte slice
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error: Response read failed")
		return "", err
	}

	// Convert byte slice to string and return
	return string(respbody), nil
}

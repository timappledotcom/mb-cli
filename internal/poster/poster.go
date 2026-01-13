package poster

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"mb/internal/config"
)

func PostToMicroBlog(cfg *config.Config, text string) error {
	if cfg.MicroBlogToken == "" {
		return fmt.Errorf("Micro.blog token is missing")
	}

	data := url.Values{}
	data.Set("h", "entry")
	data.Set("content", text)

	req, err := http.NewRequest("POST", "https://micro.blog/micropub", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+cfg.MicroBlogToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("failed to post to Micro.blog: %s", resp.Status)
	}

	return nil
}


package notify

import (
	"fmt"
	"net/http"
	"net/url"
)

// SendNotification sends a Pushover notification to the user's device.
func SendNotification(token, user, message string) error {
	apiURL := "https://api.pushover.net/1/messages.json"

	data := url.Values{}
	data.Set("token", token)
	data.Set("user", user)
	data.Set("message", message)

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return fmt.Errorf("failed to send notification: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Pushover returned non-OK status: %s", resp.Status)
	}

	fmt.Println("Notification sent successfully!")
	return nil
}

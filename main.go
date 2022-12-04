package safari

import (
	"os/exec"
)

// GetCurrentSafariURL returns Safari's current URL
func GetCurrentSafariURL() (string, error) {
	out, err := exec.Command("osascript", "-e", "tell application \"Safari\" to return URL of front document").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

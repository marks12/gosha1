package updater

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/inconshreveable/go-update"
	"gosha/settings"
	"net/http"
	"runtime"
	"time"
)

const LastReleaseUrl = "https://api.github.com/repos/%s/releases/latest"

func MakeUpdate() (isRestart bool, err error) {
	res, err := http.Get(fmt.Sprintf(LastReleaseUrl, settings.RepoName))
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	data := GithubRelease{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return false, err
	}

	if data.TagName != settings.CurrentReleaseTag {
		assetName := ""
		switch runtime.GOOS {
		case "darwin":
			assetName = "gosha-mac"
		case "linux":
			assetName = "gosha"
		case "windows":
			assetName = "gosha.exe"
		default:
			return false, errors.New("Unsupported platform")
		}

		for _, asset := range data.Assets {
			if asset.Name == assetName {
				fmt.Printf("Try install update \"%s\": %s\n", data.Name, data.HTMLURL)
				err = upgrade(asset.BrowserDownloadURL)
				if err != nil {
					return false, err
				}
				fmt.Println("Updated")
				return true, nil
			}
		}
		return false, errors.New("Not found assets in last release")
	}
	return false, nil
}

func upgrade(url string) (err error) {
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return update.Apply(resp.Body, update.Options{})
}

type GithubRelease struct {
	HTMLURL string `json:"html_url"`
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

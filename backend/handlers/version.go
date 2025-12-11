package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type VersionHandler struct{}

func NewVersionHandler() *VersionHandler {
	return &VersionHandler{}
}

type VersionInfo struct {
	Version   string                   `json:"version"`
	BuildDate string                   `json:"buildDate"`
	GitCommit string                   `json:"gitCommit"`
	Changelog []VersionChangelogEntry  `json:"changelog"`
}

type VersionChangelogEntry struct {
	Version string   `json:"version"`
	Date    string   `json:"date"`
	Changes []string `json:"changes"`
}

type GitHubRelease struct {
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	PublishedAt time.Time `json:"published_at"`
	Body        string    `json:"body"`
	HTMLURL     string    `json:"html_url"`
}

// GetVersion returns the current application version
func (h *VersionHandler) GetVersion(c *gin.Context) {
	// Read version.json file
	versionFile := "version.json"

	// Check if file exists
	if _, err := os.Stat(versionFile); os.IsNotExist(err) {
		c.JSON(http.StatusOK, VersionInfo{
			Version:   "1.0.0",
			BuildDate: time.Now().Format("2006-01-02"),
			GitCommit: "unknown",
			Changelog: []VersionChangelogEntry{},
		})
		return
	}

	// Read file
	data, err := ioutil.ReadFile(versionFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read version file"})
		return
	}

	var versionInfo VersionInfo
	if err := json.Unmarshal(data, &versionInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse version file"})
		return
	}

	c.JSON(http.StatusOK, versionInfo)
}

// CheckForUpdates checks GitHub for new releases
func (h *VersionHandler) CheckForUpdates(c *gin.Context) {
	githubRepo := os.Getenv("GITHUB_REPO") // Format: "owner/repo"
	if githubRepo == "" {
		c.JSON(http.StatusOK, gin.H{
			"updateAvailable": false,
			"message":         "GitHub repository not configured",
		})
		return
	}

	// Get current version
	versionFile := "version.json"
	currentVersion := "1.0.0"

	if _, err := os.Stat(versionFile); err == nil {
		data, err := ioutil.ReadFile(versionFile)
		if err == nil {
			var versionInfo VersionInfo
			if err := json.Unmarshal(data, &versionInfo); err == nil {
				currentVersion = versionInfo.Version
			}
		}
	}

	// Fetch latest release from GitHub
	apiURL := "https://api.github.com/repos/" + githubRepo + "/releases/latest"

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Add GitHub token if available (to avoid rate limiting)
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken != "" {
		req.Header.Set("Authorization", "token "+githubToken)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"updateAvailable": false,
			"message":         "Failed to check for updates",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		c.JSON(http.StatusOK, gin.H{
			"updateAvailable": false,
			"message":         "No releases found",
		})
		return
	}

	if resp.StatusCode != 200 {
		c.JSON(http.StatusOK, gin.H{
			"updateAvailable": false,
			"message":         "Failed to fetch latest release",
		})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var latestRelease GitHubRelease
	if err := json.Unmarshal(body, &latestRelease); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse release data"})
		return
	}

	// Compare versions (simple string comparison for now)
	// For proper semantic versioning, use a library like github.com/hashicorp/go-version
	updateAvailable := latestRelease.TagName != "v"+currentVersion && latestRelease.TagName != currentVersion

	c.JSON(http.StatusOK, gin.H{
		"updateAvailable": updateAvailable,
		"currentVersion":  currentVersion,
		"latestVersion":   latestRelease.TagName,
		"releaseDate":     latestRelease.PublishedAt,
		"releaseNotes":    latestRelease.Body,
		"releaseURL":      latestRelease.HTMLURL,
	})
}

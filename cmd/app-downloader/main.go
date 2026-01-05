package main

import (
	"archive/zip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Config holds configuration parameters
type Config struct {
	AppRepo string
	Mode    string
}

// AppList matches the structure of 1panel.json
type AppList struct {
	Apps []AppDefine `json:"apps"`
}

// AppDefine matches the app definition in 1panel.json
type AppDefine struct {
	AppProperty AppProperty        `json:"additionalProperties"`
	Versions    []AppConfigVersion `json:"versions"`
	Icon        string             `json:"icon"`
}

// AppProperty holds basic app info
type AppProperty struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

// AppConfigVersion holds version info
type AppConfigVersion struct {
	Name string `json:"name"`
}

var ()

func main() {
	appRepo := flag.String("repo", "https://apps-assets.fit2cloud.com", "App Store Repository URL")
	mode := flag.String("mode", "stable", "Mode (stable/dev)")
	flag.Parse()

	config := Config{
		AppRepo: *appRepo,
		Mode:    *mode,
	}

	workDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	downloadDir := filepath.Join(workDir, config.Mode)
	if err := os.MkdirAll(downloadDir, 0755); err != nil {
		fmt.Printf("Error creating download directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Starting download from %s [%s]\n", config.AppRepo, config.Mode)
	fmt.Printf("Download directory: %s\n", downloadDir)

	// 1. Download version.txt
	versionURL := fmt.Sprintf("%s/%s/1panel.json.version.txt", config.AppRepo, config.Mode)
	versionFile := filepath.Join(downloadDir, "1panel.json.version.txt")
	if err := downloadFile(versionURL, versionFile); err != nil {
		fmt.Printf("Error downloading version.txt: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Downloaded 1panel.json.version.txt")

	// 2. Download 1panel.json.zip
	zipURL := fmt.Sprintf("%s/%s/1panel.json.zip", config.AppRepo, config.Mode)
	zipFile := filepath.Join(downloadDir, "1panel.json.zip")
	if err := downloadFile(zipURL, zipFile); err != nil {
		fmt.Printf("Error downloading 1panel.json.zip: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Downloaded 1panel.json.zip")

	// 3. Extract and parse 1panel.json
	jsonFile := filepath.Join(downloadDir, "1panel.json")
	if err := unzipFile(zipFile, jsonFile); err != nil {
		fmt.Printf("Error extracting 1panel.json: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Extracted 1panel.json")

	content, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Printf("Error reading 1panel.json: %v\n", err)
		os.Exit(1)
	}

	var list AppList
	if err := json.Unmarshal(content, &list); err != nil {
		fmt.Printf("Error parsing 1panel.json: %v\n", err)
		os.Exit(1)
	}

	// 4. Process apps
	fmt.Println("Processing apps...")
	for _, app := range list.Apps {
		appKey := app.AppProperty.Key
		appType := app.AppProperty.Type

		// Download icon
		if app.Icon != "" {
			iconURL := app.Icon
			iconName := filepath.Base(iconURL)
			if iconName == "." || iconName == "/" {
				iconName = "icon.png"
			}

			// Local path: {Mode}/1panel/{AppKey}/{iconName}
			iconDir := filepath.Join(downloadDir, "1panel", appKey)
			if err := os.MkdirAll(iconDir, 0755); err != nil {
				fmt.Printf("  Error creating directory %s: %v\n", iconDir, err)
			} else {
				iconFile := filepath.Join(iconDir, iconName)
				fmt.Printf("  Downloading icon for %s...\n", appKey)
				if err := downloadFile(iconURL, iconFile); err != nil {
					fmt.Printf("  ⚠️ Failed to download icon %s: %v\n", iconURL, err)
				} else {
					fmt.Printf("  ✓ Saved icon to %s\n", iconFile)
				}
			}
		}

		fmt.Printf("Checking app: %s (Type: %s)\n", appKey, appType)

		for _, version := range app.Versions {
			verName := version.Name

			// Local path: {Mode}/1panel/{AppKey}/{Version}
			targetDir := filepath.Join(downloadDir, "1panel", appKey, verName)
			if err := os.MkdirAll(targetDir, 0755); err != nil {
				fmt.Printf("  Error creating directory %s: %v\n", targetDir, err)
				continue
			}

			// 1. Download app package tarball
			tarballName := fmt.Sprintf("%s-%s.tar.gz", appKey, verName)
			tarballURL := fmt.Sprintf("%s/%s/1panel/%s/%s/%s",
				config.AppRepo, config.Mode, appKey, verName, tarballName)
			tarballFile := filepath.Join(targetDir, tarballName)

			fmt.Printf("  Downloading package for %s %s...\n", appKey, verName)
			if err := downloadFile(tarballURL, tarballFile); err != nil {
				fmt.Printf("  ⚠️ Failed to download %s: %v\n", tarballURL, err)
			} else {
				fmt.Printf("  ✓ Saved to %s\n", tarballFile)
			}

			// 2. Download docker-compose.yml
			composeURL := fmt.Sprintf("%s/%s/1panel/%s/%s/docker-compose.yml",
				config.AppRepo, config.Mode, appKey, verName)
			targetFile := filepath.Join(targetDir, "docker-compose.yml")
			fmt.Printf("  Downloading compose for %s %s...\n", appKey, verName)
			if err := downloadFile(composeURL, targetFile); err != nil {
				fmt.Printf("  ⚠️ Failed to download %s: %v\n", composeURL, err)
			} else {
				fmt.Printf("  ✓ Saved to %s\n", targetFile)
			}
		}
	}
	fmt.Println("\nAll operations completed.")
}

func downloadFile(url, filepath string) error {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzipFile(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name != "1panel.json" {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		outFile, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, rc)
		return err
	}
	return fmt.Errorf("1panel.json not found in zip")
}

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
	AppRepo  string
	Mode     string
	Retry    int
	Interval int
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

func fileExistsAndNotEmpty(filepath string) bool {
	info, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return !info.IsDir() && info.Size() > 0
}

func main() {
	appRepo := flag.String("repo", "https://apps-assets.fit2cloud.com", "App Store Repository URL")
	mode := flag.String("mode", "stable", "Mode (stable/dev)")
	retry := flag.Int("retry", 10, "Number of retries for download")
	interval := flag.Int("interval", 5000, "Interval between downloads in milliseconds")
	flag.Parse()

	config := Config{
		AppRepo:  *appRepo,
		Mode:     *mode,
		Retry:    *retry,
		Interval: *interval,
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

	fmt.Printf("Starting download from %s [%s] (retry=%d, interval=%dms)\n", config.AppRepo, config.Mode, config.Retry, config.Interval)
	fmt.Printf("Download directory: %s\n", downloadDir)

	// 1. Download version.txt
	versionURL := fmt.Sprintf("%s/%s/1panel.json.version.txt", config.AppRepo, config.Mode)
	versionFile := filepath.Join(downloadDir, "1panel.json.version.txt")
	if err := downloadWithRetry(versionURL, versionFile, config.Retry, config.Interval); err != nil {
		fmt.Printf("Error downloading version.txt: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Downloaded 1panel.json.version.txt")

	// 2. Download 1panel.json.zip
	zipURL := fmt.Sprintf("%s/%s/1panel.json.zip", config.AppRepo, config.Mode)
	zipFile := filepath.Join(downloadDir, "1panel.json.zip")
	if err := downloadWithRetry(zipURL, zipFile, config.Retry, config.Interval); err != nil {
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
				if fileExistsAndNotEmpty(iconFile) {
					fmt.Printf("  ✓ Skipped icon for %s (already exists)\n", appKey)
				} else {
					fmt.Printf("  Downloading icon for %s...\n", appKey)
					if err := downloadWithRetry(iconURL, iconFile, config.Retry, config.Interval); err != nil {
						fmt.Printf("  ⚠️ Failed to download icon %s: %v\n", iconURL, err)
					} else {
						fmt.Printf("  ✓ Saved icon to %s\n", iconFile)
					}
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

			if fileExistsAndNotEmpty(tarballFile) {
				fmt.Printf("  ✓ Skipped %s (already exists)\n", tarballName)
			} else {
				fmt.Printf("  Downloading package for %s %s...\n", appKey, verName)
				if err := downloadWithRetry(tarballURL, tarballFile, config.Retry, config.Interval); err != nil {
					fmt.Printf("  ⚠️ Failed to download %s: %v\n", tarballURL, err)
				} else {
					fmt.Printf("  ✓ Saved to %s\n", tarballFile)
				}
			}

			// 2. Download docker-compose.yml
			composeURL := fmt.Sprintf("%s/%s/1panel/%s/%s/docker-compose.yml",
				config.AppRepo, config.Mode, appKey, verName)
			targetFile := filepath.Join(targetDir, "docker-compose.yml")

			if fileExistsAndNotEmpty(targetFile) {
				fmt.Printf("  ✓ Skipped docker-compose.yml (already exists)\n")
			} else {
				fmt.Printf("  Downloading compose for %s %s...\n", appKey, verName)
				if err := downloadWithRetry(composeURL, targetFile, config.Retry, config.Interval); err != nil {
					fmt.Printf("  ⚠️ Failed to download %s: %v\n", composeURL, err)
				} else {
					fmt.Printf("  ✓ Saved to %s\n", targetFile)
				}
			}
		}
	}
	fmt.Println("\nAll operations completed.")
}

func downloadWithRetry(url, filepath string, retry, interval int) error {
	if interval > 0 {
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}

	var err error
	for i := 0; i <= retry; i++ {
		if i > 0 {
			fmt.Printf("  ⚠️ Download failed: %v. Retrying (%d/%d)...\n", err, i, retry)
			time.Sleep(time.Second*2 + time.Duration(interval)*2*time.Millisecond)
		}

		if err = downloadFile(url, filepath); err == nil {
			return nil
		}
	}
	return fmt.Errorf("failed after %d retries: %v", retry, err)
}

func downloadFile(url, filepath string) error {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// 模拟浏览器 User-Agent 和其他 Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="144", "Google Chrome";v="144"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// 使用临时文件下载，避免下载失败产生脏文件
	tmpFile := filepath + ".tmp"
	out, err := os.Create(tmpFile)
	if err != nil {
		return err
	}

	// 确保在函数退出时关闭文件，但在重命名之前我们需要显式关闭它
	// 这里的 defer 主要是为了处理异常情况下的关闭
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		// 下载失败，关闭并删除临时文件
		out.Close()
		os.Remove(tmpFile)
		return err
	}

	// 显式关闭文件以确保所有数据写入磁盘
	out.Close()

	// 下载成功，重命名为正式文件
	return os.Rename(tmpFile, filepath)
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

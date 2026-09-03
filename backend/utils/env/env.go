// 此文件在GPL-3.0协议下开源
// 修改者：bobwu 2026-09-01
package env

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/1Panel-dev/1Panel/backend/global"
	"github.com/joho/godotenv"
)

func Write(envMap map[string]string, filename string) error {
	content, err := Marshal(envMap)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content + "\n")
	if err != nil {
		return err
	}
	return file.Sync()
}

func Marshal(envMap map[string]string) (string, error) {
	lines := make([]string, 0, len(envMap))
	for k, v := range envMap {
		if d, err := strconv.Atoi(v); err == nil {
			lines = append(lines, fmt.Sprintf(`%s=%d`, k, d))
		} else {
			lines = append(lines, fmt.Sprintf(`%s="%s"`, k, v))
		}
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n"), nil
}

func GetEnvValueByKey(envPath, key string) (string, error) {
	envMap, err := godotenv.Read(envPath)
	if err != nil {
		return "", err
	}
	value, ok := envMap[key]
	if !ok {
		return "", fmt.Errorf("key %s not found in %s", key, envPath)
	}
	return value, nil
}

// GetCustomComposeFilePath 获取 docker-compose.yml 文件的绝对路径
// 逻辑说明：
//  1. 根据传入的默认 compose 文件路径（defaultComposePath），推导出同级目录下的 .env 文件路径。
//  2. 尝试读取并解析该 .env 文件。
//  3. 如果 .env 文件中存在名为 "CUSTOM_COMPOSE_FILE_PATH" 的变量且不为空：
//     a. 检查该路径指向的文件是否存在且大小不为 0。
//     b. 如果检查通过，将默认 .env 文件内容同步拷贝到自定义 compose 文件同级目录下（仅在内容不一致或文件不存在时写入）。
//     c. 返回该变量的值（作为自定义 compose 路径）。
//  4. 其他情况（文件不存在、解析失败、变量未设置、自定义文件无效等），均返回传入的默认路径。
//
// zhuzhiwu 20260304
func GetCustomComposeFilePath(defaultComposePath string) string {
	envPath := strings.Replace(defaultComposePath, "docker-compose.yml", ".env", 1)
	if content, err := os.ReadFile(envPath); err == nil {
		envMap, err := godotenv.UnmarshalBytes(content)
		if err == nil {
			if customComposePath, ok := envMap["CUSTOM_COMPOSE_FILE_PATH"]; ok && customComposePath != "" {
				if info, err := os.Stat(customComposePath); err == nil && info.Size() > 0 {
					global.LOG.Infof("custom compose file path: %s", customComposePath)
					customEnvPath := strings.Replace(customComposePath, "docker-compose.yml", ".env", 1)
					isWrite := true
					if existContent, err := os.ReadFile(customEnvPath); err == nil {
						if string(existContent) == string(content) {
							isWrite = false
						}
					}
					if isWrite {
						if err := os.WriteFile(customEnvPath, content, 0644); err != nil {
							global.LOG.Errorf("failed to copy %s file to %s: %v", envPath, customEnvPath, err)
						} else {
							global.LOG.Infof("successfully copied %s file to %s", envPath, customEnvPath)
						}
					}

					return customComposePath
				} else {
					global.LOG.Warnf("custom compose file path %s not exist or empty", customComposePath)
				}
			}
		}
	}
	return defaultComposePath
}

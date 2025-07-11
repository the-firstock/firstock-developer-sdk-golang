// Copyright (c) [2025] [abc]
// SPDX-License-Identifier: MIT
package Firstock

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

func encodePassword(pwd string) string {
	hash := sha256.Sum256([]byte(pwd))
	return hex.EncodeToString(hash[:])
}

func readJKeyFromConfig(userId string) (string, error) {
	configPath, err := getConfigPath()

	if err != nil {
		return "", fmt.Errorf("could not open config file: %w", err)
	}
	file, err := os.Open(configPath)
	if err != nil {
		return "", fmt.Errorf("could not open config file: %w", err)
	}
	defer file.Close()

	var config map[string]interface{}
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return "", fmt.Errorf("could not decode config JSON: %w", err)
	}

	userConfigRaw, ok := config[userId]
	if !ok {
		return "", fmt.Errorf("userId %s not found in config", userId)
	}

	userConfig, ok := userConfigRaw.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid user config format for userId %s", userId)
	}

	jkeyRaw, ok := userConfig[j_key]
	if !ok {
		return "", fmt.Errorf("jkey not found for userId %s", userId)
	}

	jkey, ok := jkeyRaw.(string)
	if !ok {
		return "", fmt.Errorf("jkey is not a string for userId %s", userId)
	}

	return jkey, nil
}

var configMu sync.Mutex

func saveJKeyToConfig(data LogoutRequest) error {
	userId := data.UserId
	jkey := data.JKey

	const configFile = "config.json"

	configMu.Lock()
	defer configMu.Unlock()

	config := map[string]map[string]string{}

	bytes, err := os.ReadFile(configFile)
	if err == nil && len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &config); err != nil {
			// If error, clear the config and write empty JSON object
			emptyConfig := map[string]map[string]string{}
			jsonBytes, _ := json.MarshalIndent(emptyConfig, "", "  ")
			_ = os.WriteFile(configFile, jsonBytes, 0644)
			config = emptyConfig
		}
	}

	if _, ok := config[userId]; !ok {
		config[userId] = map[string]string{}
	}
	config[userId]["jkey"] = jkey

	jsonBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	// Overwrite the file even if it did not exist or had read error
	return os.WriteFile(configFile, jsonBytes, 0644)
}

func removeJKeyFromConfig(userId string) error {
	const configFile = "config.json"

	configMu.Lock()
	defer configMu.Unlock()

	// Check if config file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist")
	}

	// Read and unmarshal config
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	config := map[string]map[string]string{}
	if err := json.Unmarshal(bytes, &config); err != nil {
		return err
	}

	// Remove jkey if present
	if userConfig, ok := config[userId]; ok {
		if _, exists := userConfig["jkey"]; exists {
			delete(userConfig, "jkey")
			// If userConfig is now empty, remove the userId entry
			if len(userConfig) == 0 {
				delete(config, userId)
			}
		}
	}

	// Write updated config back to file
	jsonBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, jsonBytes, 0644)
}

func getConfigPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, "config.json"), nil
}

func internalServerErrorResponse() *ErrorResponseModel {
	errDetail := ErrorDetail{
		Field:   "",
		Message: internal_server_error,
	}
	errorResponse := ErrorResponseModel{
		Code:   "500",
		Error:  errDetail,
		Name:   internal_server_error,
		Status: "failed",
	}

	return &errorResponse
}

func pleaseLoginToFirstock() *ErrorResponseModel {
	errDetail := ErrorDetail{
		Field:   "",
		Message: please_login_to_firstock,
	}
	errorResponse := ErrorResponseModel{
		Code:   "401",
		Error:  errDetail,
		Name:   please_login_to_firstock,
		Status: "failed",
	}
	return &errorResponse
}

func failureResponseStructure(data map[string]interface{}) *ErrorResponseModel {
	var errorResponse ErrorResponseModel
	var errDetail ErrorDetail

	// Extract fields from data["error"] map
	if data != nil {
		if errorMap, ok := data["error"].(map[string]interface{}); ok {
			if field, ok := errorMap["field"].(string); ok {
				errDetail.Field = field
			}
			if message, ok := errorMap["message"].(string); ok {
				errDetail.Message = message
			}
		}

		// Extract other optional fields directly from data
		code, _ := data["code"].(string)
		name, _ := data["name"].(string)
		status, _ := data["status"].(string)

		errorResponse = ErrorResponseModel{
			Code:   code,
			Error:  errDetail,
			Name:   name,
			Status: status,
		}
		return &errorResponse
	}

	// Fallback if data is nil
	errorResponse = ErrorResponseModel{
		Code: "401",
		Error: ErrorDetail{
			Field:   "",
			Message: please_login_to_firstock,
		},
		Name:   please_login_to_firstock,
		Status: "failed",
	}
	return &errorResponse
}

func successResponseStructure(data map[string]interface{}) (map[string]interface{}, *ErrorResponseModel) {
	return data, nil
}

func readJkey(userId string) (string, *ErrorResponseModel) {
	jkey, errRead := readJKeyFromConfig(userId)
	if errRead != nil {
		return jkey, pleaseLoginToFirstock()
	}
	return jkey, nil
}

func check_if_unauthorized(code string) bool {
	re := regexp.MustCompile(`4\d{2}`)
	return re.MatchString(code)
}

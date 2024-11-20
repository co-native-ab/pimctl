package credentials

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

var (
	identityServiceDirectoryName       = ".IdentityService"
	cacheNamePrefix                    = "pimctl"
	cacheFileNameSuffix                = "cache.json"
	authenticationRecordFileNameSuffix = "auth-record.json"
	configCacheFileNameSuffix          = "config-cache.json"
)

func getCacheName(profileName string) string {
	return fmt.Sprintf("%s_%s", cacheNamePrefix, profileName)
}

func getAuthenticationRecordFileName(profileName string) string {
	return fmt.Sprintf("%s_%s.%s", cacheNamePrefix, profileName, authenticationRecordFileNameSuffix)
}

func getConfigCacheFileName(profileName string) string {
	return fmt.Sprintf("%s_%s.%s", cacheNamePrefix, profileName, configCacheFileNameSuffix)
}

func getCacheFileName(profileName string) string {
	return fmt.Sprintf("%s_%s.%s", cacheNamePrefix, profileName, cacheFileNameSuffix)
}

func GetCachedAuthenticationRecord(profileName string) (azidentity.AuthenticationRecord, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return azidentity.AuthenticationRecord{}, fmt.Errorf("failed to get user home directory: %w", err)
	}

	authenticationRecordPath := filepath.Join(userHomeDir, identityServiceDirectoryName, getAuthenticationRecordFileName(profileName))
	authenticationRecord, err := retrieveCachedFile[azidentity.AuthenticationRecord](authenticationRecordPath)
	if err != nil {
		return azidentity.AuthenticationRecord{}, fmt.Errorf("failed to retrieve authentication record: %w", err)
	}

	return authenticationRecord, nil
}

func ClearCache(profileName string) error {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}

	authenticationRecordPath := filepath.Join(userHomeDir, identityServiceDirectoryName, getAuthenticationRecordFileName(profileName))
	err = os.Remove(authenticationRecordPath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove authentication record from %q: %w", authenticationRecordPath, err)
	}

	configCachePath := filepath.Join(userHomeDir, identityServiceDirectoryName, getConfigCacheFileName(profileName))
	err = os.Remove(configCachePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove config cache from %q: %w", configCachePath, err)
	}

	cachePath := filepath.Join(userHomeDir, identityServiceDirectoryName, getCacheFileName(profileName))
	err = os.Remove(cachePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove cache from %q: %w", cachePath, err)
	}

	return nil
}

type configCacheFile struct {
	CredentialMethod CredentialMethod `json:"credential_method"`
	Scopes           []string         `json:"scopes"`
}

type configCache struct {
	credentialMethod       CredentialMethod
	scopes                 []string
	azAuthenticationRecord azidentity.AuthenticationRecord
}

func (c *configCache) Store(profileName string) error {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}

	configCachePath := filepath.Join(userHomeDir, identityServiceDirectoryName, getConfigCacheFileName(profileName))
	err = storeCachedFile(configCachePath, configCacheFile{
		CredentialMethod: c.credentialMethod,
		Scopes:           c.scopes,
	})
	if err != nil {
		return fmt.Errorf("failed to store config cache: %w", err)
	}

	authenticationRecordPath := filepath.Join(userHomeDir, identityServiceDirectoryName, getAuthenticationRecordFileName(profileName))
	err = storeCachedFile(authenticationRecordPath, c.azAuthenticationRecord)
	if err != nil {
		return fmt.Errorf("failed to store authentication record: %w", err)
	}

	return nil
}

func getCacheConfig(profileName string) (*configCache, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}

	configCachePath := filepath.Join(userHomeDir, identityServiceDirectoryName, getConfigCacheFileName(profileName))
	configCacheFile, err := retrieveCachedFile[configCacheFile](configCachePath)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve config cache: %w", err)
	}

	authenticationRecordPath := filepath.Join(userHomeDir, identityServiceDirectoryName, getAuthenticationRecordFileName(profileName))
	authenticationRecord, err := retrieveCachedFile[azidentity.AuthenticationRecord](authenticationRecordPath)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve authentication record: %w", err)
	}

	return &configCache{
		credentialMethod:       configCacheFile.CredentialMethod,
		scopes:                 configCacheFile.Scopes,
		azAuthenticationRecord: authenticationRecord,
	}, nil
}

func retrieveCachedFile[T any](path string) (T, error) {
	value := *new(T)
	b, err := os.ReadFile(path)
	if err == nil {
		err = json.Unmarshal(b, &value)
	}
	if os.IsNotExist(err) {
		return *new(T), nil
	}

	return value, err
}

func storeCachedFile[T any](path string, value T) error {
	b, err := json.Marshal(value)
	if err == nil {
		err = os.WriteFile(path, b, 0600)
	}
	return err
}

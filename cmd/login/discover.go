package login

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
)

func discoverClientID(ctx context.Context, flagClientID string) (string, error) {
	errs := []error{}
	if flagClientID != "" {
		return flagClientID, nil
	}
	errs = append(errs, fmt.Errorf("no client ID provided through --client-id"))

	credentialsCache, err := cmdhelper.GetCachedCredentialAuthenticationRecord()
	if err == nil && credentialsCache != (azidentity.AuthenticationRecord{}) {
		return credentialsCache.ClientID, nil
	}
	errs = append(errs, err)

	azureCLIClientID, err := getClientIDFromAzureCLI(ctx)
	if err == nil {
		return azureCLIClientID, nil
	}
	errs = append(errs, err)

	return "", fmt.Errorf("unable to discover Client ID\nerrors: %w", errors.Join(errs...))
}

func getClientIDFromAzureCLI(ctx context.Context) (string, error) {
	stdout, err := execAzureCLI(ctx, []string{"ad", "app", "list", "--filter", "tags/any(s: s eq 'pimctl')", "--query", "[].appId", "--output", "json"})
	if err != nil {
		return "", fmt.Errorf("failed to get client ID from Azure CLI: %w", err)
	}

	var clientIDs []string
	err = json.Unmarshal(stdout, &clientIDs)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal client IDs: %w", err)
	}

	if len(clientIDs) != 1 {
		return "", fmt.Errorf("expected exactly one client ID, got %d", len(clientIDs))
	}

	return clientIDs[0], nil
}

func execAzureCLI(ctx context.Context, args []string) ([]byte, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, "az", args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		return nil, wrapCommandError(err, stderr, "failed to start Azure CLI")
	}

	err = cmd.Wait()
	if err != nil {
		return nil, wrapCommandError(err, stderr, "failed to run Azure CLI")
	}

	return stdout.Bytes(), nil
}

func wrapCommandError(err error, stderr bytes.Buffer, msg string) error {
	stderrStr := stderr.String()
	if err == nil && stderrStr != "" {
		return fmt.Errorf("%s: %s", msg, stderrStr)
	}

	if err == nil {
		return fmt.Errorf("%s", msg)
	}

	if stderrStr == "" {
		return fmt.Errorf("%s: %w", msg, err)
	}

	return fmt.Errorf("%s: %s: %w", msg, stderrStr, err)
}

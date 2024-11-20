package credentials

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CredentialMethod int

const (
	UnknownCredentialMethod CredentialMethod = iota
	DeviceCodeCredentialMethod
	InteractiveBrowserCredentialMethod
)

func (m *CredentialMethod) String() string {
	if m == nil {
		return "Unknown"
	}

	switch *m {
	case DeviceCodeCredentialMethod:
		return "DeviceCode"
	case InteractiveBrowserCredentialMethod:
		return "InteractiveBrowser"
	default:
		return "Unknown"
	}
}

func (m *CredentialMethod) Set(v string) error {
	switch v {
	case "DeviceCode":
		*m = DeviceCodeCredentialMethod
		return nil
	case "InteractiveBrowser":
		*m = InteractiveBrowserCredentialMethod
		return nil
	default:
		return fmt.Errorf("unknown credential method %q, must be one of 'DeviceCode' or 'InteractiveBrowser'", v)
	}
}

func (m *CredentialMethod) HelpText() string {
	return "Credential method to use for pimctl. Allowed: 'DeviceCode', 'InteractiveBrowser'"
}

func (m *CredentialMethod) Type() string {
	return "CredentialMethod"
}

func (m *CredentialMethod) CobraCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{
		"InteractiveBrowser\tLogin using a browser",
		"DeviceCode\tLogin using a device code",
	}, cobra.ShellCompDirectiveDefault
}

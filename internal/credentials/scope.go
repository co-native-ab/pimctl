package credentials

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type PimctlScope int

const (
	UnknownPimctlScope PimctlScope = iota
	MicrosoftGraphPimctlScope
	AzurePimctlScope
)

func (m *PimctlScope) Scopes() []string {
	switch *m {
	case MicrosoftGraphPimctlScope:
		return []string{"https://graph.microsoft.com/.default"}
	case AzurePimctlScope:
		return []string{"https://management.azure.com/.default"}
	default:
		return []string{}
	}
}

func (m *PimctlScope) String() string {
	if m == nil {
		return "Unknown"
	}

	switch *m {
	case MicrosoftGraphPimctlScope:
		return "MicrosoftGraph"
	case AzurePimctlScope:
		return "Azure"
	default:
		return "Unknown"
	}
}

func (m *PimctlScope) Set(v string) error {
	switch strings.ToLower(v) {
	case "microsoftgraph":
		*m = MicrosoftGraphPimctlScope
		return nil
	case "azure":
		*m = AzurePimctlScope
		return nil
	default:
		return fmt.Errorf("unknown scope %q, must be one of 'MicrosoftGraph' or 'Azure'", v)
	}
}

func (m *PimctlScope) HelpText() string {
	return "Scope to use for pimctl. Allowed: 'MicrosoftGraph', 'Azure'"
}

func (m *PimctlScope) Type() string {
	return "Scope"
}

func (m *PimctlScope) CobraCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{
		"MicrosoftGraph\tLogin with pimctl to Microsoft Graph",
		"Azure\tLogin with pimctl to Azure",
	}, cobra.ShellCompDirectiveDefault
}

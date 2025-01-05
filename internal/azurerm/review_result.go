package azurerm

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type ReviewResult int

const (
	UnknownReviewResult ReviewResult = iota
	ApproveReviewResult
	DenyReviewResult
)

func (r *ReviewResult) String() string {
	if r == nil {
		return "Unknown"
	}

	switch *r {
	case ApproveReviewResult:
		return "Approve"
	case DenyReviewResult:
		return "Deny"
	default:
		return "Unknown"
	}
}

func (r *ReviewResult) Set(v string) error {
	switch strings.ToLower(v) {
	case "approve":
		*r = ApproveReviewResult
		return nil
	case "deny":
		*r = DenyReviewResult
		return nil
	default:
		return fmt.Errorf("unknown review result %q, must be one of 'Approve' or 'Deny'", v)
	}
}

func (r *ReviewResult) HelpText() string {
	return "Review result to use for pending assignment approval. Allowed: 'Approve', 'Deny'"
}

func (r *ReviewResult) Type() string {
	return "ReviewResult"
}

func (r *ReviewResult) CobraCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{
		"Approve\tApprove the pending assignment",
		"Deny\tDeny the pending assignment",
	}, cobra.ShellCompDirectiveDefault
}

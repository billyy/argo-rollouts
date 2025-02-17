package cmd

import (
	"github.com/spf13/cobra"

	"github.com/argoproj/argo-rollouts/pkg/kubectl-argo-rollouts/cmd/abort"
	"github.com/argoproj/argo-rollouts/pkg/kubectl-argo-rollouts/cmd/get"
	"github.com/argoproj/argo-rollouts/pkg/kubectl-argo-rollouts/cmd/list"
	"github.com/argoproj/argo-rollouts/pkg/kubectl-argo-rollouts/cmd/pause"
	"github.com/argoproj/argo-rollouts/pkg/kubectl-argo-rollouts/cmd/promote"
	"github.com/argoproj/argo-rollouts/pkg/kubectl-argo-rollouts/cmd/retry"
	"github.com/argoproj/argo-rollouts/pkg/kubectl-argo-rollouts/cmd/set"
	"github.com/argoproj/argo-rollouts/pkg/kubectl-argo-rollouts/cmd/version"
	"github.com/argoproj/argo-rollouts/pkg/kubectl-argo-rollouts/options"
)

const (
	example = `
  # Pause the guestbook rollout
  %[1]s pause guestbook

  # Resume the guestbook rollout
  %[1]s promote guestbook
`
)

func NewCmdArgoRollouts(o *options.ArgoRolloutsOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "kubectl-argo-rollouts COMMAND",
		Short:             "Manage argo rollouts",
		Example:           o.Example(example),
		SilenceUsage:      true,
		PersistentPreRunE: o.PersistentPreRunE,
		RunE: func(c *cobra.Command, args []string) error {
			return o.UsageErr(c)
		},
	}
	cmd.AddCommand(get.NewCmdGet(o))
	cmd.AddCommand(list.NewCmdList(o))
	cmd.AddCommand(pause.NewCmdPause(o))
	cmd.AddCommand(promote.NewCmdPromote(o))
	cmd.AddCommand(version.NewCmdVersion(o))
	cmd.AddCommand(abort.NewCmdAbort(o))
	cmd.AddCommand(retry.NewCmdRetry(o))
	cmd.AddCommand(set.NewCmdSet(o))
	return cmd
}

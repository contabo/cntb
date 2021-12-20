package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var allowedShells = []string{"bash", "zsh", "fish", "powershell"}

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

	$ source <(cntb completion bash)

	# To load completions for each session, execute once:
	# Linux:
	$ cntb completion bash > /etc/bash_completion.d/cntb
	# macOS:
	$ cntb completion bash > /usr/local/etc/bash_completion.d/cntb

Zsh:

	# If shell completion is not already enabled in your environment,
	# you will need to enable it.  You can execute the following once:

	$ echo "autoload -U compinit; compinit" >> ~/.zshrc

	# To load completions for each session, execute once:
	$ cntb completion zsh > "${fpath[1]}/_cntb"

	# You will need to start a new shell for this setup to take effect.

fish:

	$ cntb completion fish | source

	# To load completions for each session, execute once:
	$ cntb completion fish > ~/.config/fish/completions/cntb.fish

PowerShell:

	PS> cntb completion powershell | Out-String | Invoke-Expression

	# To load completions for every new session, run:
	PS> cntb completion powershell > cntb.ps1
	# and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             allowedShells,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			cmd.Help()
			fmt.Printf("Requires one of %v\n", allowedShells)
			os.Exit(0)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}

package logic

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Client) RunCmdRoot(cmd *cobra.Command) error {
	err := cmd.Help()
	if err != nil {
		return fmt.Errorf("error showing help: %w", err)
	}

	return nil
}

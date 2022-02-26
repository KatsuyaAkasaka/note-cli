package adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type Command struct {
	Command string
	Desc    string
	Aliases []string
	Exec    func(ctx context.Context, args []string) error
	Timeout int64
}

func (c *Command) ToCobraCommand() *cobra.Command {
	return &cobra.Command{
		Use:     c.Command,
		Short:   c.Desc,
		Aliases: c.Aliases,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
			defer cancel()

			errCh := make(chan error)

			go func() {
				errCh <- c.Exec(ctx, args)
			}()

			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return fmt.Errorf("unknown error")
			case err := <-errCh:
				return err
			}
		},
	}
}

func (c *Command) Convert() *Command {
	if c.Command == "" {
		c.Command = "dummy"
	}
	if c.Timeout == 0 {
		c.Timeout = 5
	}
	return c
}

package builder

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openshift/origin/pkg/build/builder/cmd"
)

const longCommandSTIDesc = `
Perform a Source-to-Image Build

This command executes a Source-to-Image build using arguments passed via the environment.
It expects to be run inside of a container.
`

func NewCommandSTIBuilder(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s", name),
		Short: "Run an OpenShift Source-to-Images build",
		Long:  longCommandSTIDesc,
		Run: func(c *cobra.Command, args []string) {
			cmd.RunSTIBuild()
		},
	}

	return cmd
}

const longCommandDockerDesc = `
Perform a Docker Build

This command executes a Docker build using arguments passed via the environment.
It expects to be run inside of a container.
`

func NewCommandDockerBuilder(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s", name),
		Short: "Run an OpenShift Docker build",
		Long:  longCommandDockerDesc,
		Run: func(c *cobra.Command, args []string) {
			cmd.RunDockerBuild()
		},
	}

	return cmd
}

package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/kauefraga/pavus/internal/markdown"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func getInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:     "init",
		Aliases: []string{"i", "create", "c"},
		Short:   "Create a markdown file based in a template",
		Example: `  pavus init
  pavus i --output my-docs
  pavus create -t tooling -o my-tooling-docs`,
		Run: func(cmd *cobra.Command, args []string) {
			t, _ := cmd.Flags().GetString("template")
			if t == "" {
				prompt := promptui.Select{
					Label: "Choose a template",
					Items: markdown.GetAvailableTemplates(),
				}

				_, result, err := prompt.Run()
				if err != nil {
					color.Red("Prompt failed %s", err)
					return
				}

				t = result
			}

			content, err := markdown.GetTemplate(t)
			if err != nil {
				color.Red("Error: failed to get the template")
				os.Exit(1)
			}

			output, _ := cmd.Flags().GetString("output")
			if output == "" {
				prompt := promptui.Prompt{
					Label:   "Output file name",
					Default: "generated-by-pavus",
					Validate: func(s string) error {
						if s == "" {
							return errors.New("file name required")
						}
						return nil
					},
				}

				result, err := prompt.Run()
				if err != nil {
					color.Red("Prompt failed %s", err)
					return
				}

				output = result
			}

			if !strings.HasSuffix(output, ".md") {
				output = output + ".md"
			}

			err = os.WriteFile(output, content, 0666)
			if err != nil {
				color.Red("Error: failed to create %s", output)
				os.Exit(1)
			}

			color.Green("You're done! Take a look at the %s", output)
		},
	}

	initCmd.Flags().StringP("template", "t", "", "define the template to be used")
	initCmd.Flags().StringP("output", "o", "", "define the output markdown file")

	return initCmd
}

package list

import (
	"fmt"
	"github.com/profclems/glab/commands/cmdutils"
	"github.com/profclems/glab/commands/release/releaseutils"
	"github.com/profclems/glab/internal/utils"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

var factory *cmdutils.Factory

func NewCmdReleaseList(f *cmdutils.Factory) *cobra.Command {
	factory = f
	var releaseListCmd = &cobra.Command{
		Use:     "list [flags]",
		Short:   `List releases in a repository`,
		Long:    ``,
		Aliases: []string{"ls"},
		Args:    cobra.MaximumNArgs(3),
		RunE:    listReleases,
	}
	releaseListCmd.Flags().StringP("tag", "t", "", "Filter releases by tag <name>")
	return releaseListCmd
}

func listReleases(cmd *cobra.Command, args []string) error {

	l := &gitlab.ListReleasesOptions{}

	tag, err := cmd.Flags().GetString("tag")

	if err != nil {
		return err
	}
	if r, _ := cmd.Flags().GetString("repo"); r != "" {
		factory, err = factory.NewClient(r)
		if err != nil {
			return err
		}
	}
	gLabClient, err := factory.HttpClient()
	if err != nil {
		return err
	}
	repo, err := factory.BaseRepo()
	if err != nil {
		return err
	}

	if tag != "" {
		release, _, err := gLabClient.Releases.GetRelease(repo.FullName(), tag)
		if err != nil {
			return err
		}
		fmt.Fprintln(utils.ColorableOut(cmd), releaseutils.DisplayRelease(release))
	} else {
		releases, _, err := gLabClient.Releases.ListReleases(repo.FullName(), l)
		if err != nil {
			return err
		}
		fmt.Fprintln(utils.ColorableOut(cmd), releaseutils.DisplayAllReleases(releases, repo.FullName()))
	}
	return nil
}
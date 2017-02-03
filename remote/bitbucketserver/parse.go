package bitbucketserver

import (
	"encoding/json"
	"fmt"
	"github.com/grommet/drone/model"
	"github.com/grommet/drone/remote/bitbucketserver/internal"
	"net/http"
)

// parseHook parses a Bitbucket hook from an http.Request request and returns
// Repo and Build detail. TODO: find a way to support PR hooks
func parseHook(r *http.Request, baseURL string) (*model.Repo, *model.Build, error) {
	hook := new(internal.PostHook)
	if err := json.NewDecoder(r.Body).Decode(hook); err != nil {
		return nil, nil, err
	}
	build := convertPushHook(hook, baseURL)
	repo := &model.Repo{
		Name:     hook.Repository.Slug,
		Owner:    hook.Repository.Project.Key,
		FullName: fmt.Sprintf("%s/%s", hook.Repository.Project.Key, hook.Repository.Slug),
		Branch:   "master",
		Kind:     model.RepoGit,
	}

	return repo, build, nil
}

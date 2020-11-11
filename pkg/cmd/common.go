package cmd

import (
	"github.com/plumming/dx/pkg/api"
	"github.com/plumming/dx/pkg/kube"

	"github.com/plumming/dx/pkg/prompter"
)

type CommonOptions struct {
	prompter        prompter.Prompter
	nonGithubClient *api.Client
	githubClient    *api.Client
	kuber           kube.Kuber
}

func (c *CommonOptions) SetPrompter(p prompter.Prompter) {
	c.prompter = p
}

func (c *CommonOptions) Prompter() prompter.Prompter {
	if c.prompter == nil {
		c.prompter = prompter.NewPrompter()
	}
	return c.prompter
}

func (c *CommonOptions) SetNonGithubClient(h *api.Client) {
	c.nonGithubClient = h
}

func (c *CommonOptions) NonGithubClient() *api.Client {
	if c.nonGithubClient == nil {
		c.nonGithubClient = api.NewClient()
	}
	return c.nonGithubClient
}

func (c *CommonOptions) SetGithubClient(h *api.Client) {
	c.githubClient = h
}

func (c *CommonOptions) GithubClient() (*api.Client, error) {
	if c.githubClient == nil {
		var err error
		c.githubClient, err = api.BasicClient()
		if err != nil {
			return nil, err
		}
	}
	return c.githubClient, nil
}

func (c *CommonOptions) SetKuber(k kube.Kuber) {
	c.kuber = k
}

func (c *CommonOptions) Kuber() kube.Kuber {
	if c.kuber == nil {
		c.kuber = kube.NewKuber()
	}
	return c.kuber
}

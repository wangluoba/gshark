/*

Copyright (c) 2018 sec.xiaomi.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THEq
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

*/

package githubsearch

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"x-patrol/models"
)

var (
	GithubClients map[string]*Client
	GithubClient  *Client
)

type Client struct {
	Client *github.Client
	Token  string
}

func init() {
	GithubClients = make(map[string]*Client)
	GithubClients, _ = InitGithubClients()
}

func InitGithubClients() (map[string]*Client, error) {
	githubClients := make(map[string]*Client)
	tokens, err := models.ListValidTokens()
	if err == nil {
		for _, token := range tokens {
			githubToken := token.Token
			gitClient := &github.Client{}
			if githubToken != "" {
				ctx := context.Background()
				ts := oauth2.StaticTokenSource(
					&oauth2.Token{AccessToken: githubToken},
				)
				tc := oauth2.NewClient(ctx, ts)
				gitClient = github.NewClient(tc)
				githubClients[token.Token] = NewGitClient(gitClient, githubToken)
			}
		}
	}
	return githubClients, err
}

func GetGithubClient() (*Client, string, error) {
	var c *Client
	clients, err := InitGithubClients()
	for _, client := range clients {
		c = client
		break
	}
	return c, c.Token, err
}

func NewGitClient(GithubClient *github.Client, token string) *Client {
	return &Client{Client: GithubClient, Token: token}
}

func (c *Client) GetUserInfo(username string) (*github.User, *github.Response, error) {
	ctx := context.Background()
	return c.Client.Users.Get(ctx, username)
}

func (c *Client) GetOrgsMembers(org string) ([]*github.User, *github.Response, error) {
	ctx := context.Background()
	return c.Client.Organizations.ListMembers(ctx, org, nil)
}

func (c *Client) GetOrgsRepos(org string) ([]*github.Repository, *github.Response, error) {
	ctx := context.Background()
	return c.Client.Repositories.ListByOrg(ctx, org, nil)
}

func (c *Client) GetUserRepos(username string) ([]*github.Repository, *github.Response, error) {
	ctx := context.Background()
	return c.Client.Repositories.List(ctx, username, nil)
}

func (c *Client) GetUsersRepos(users []*github.User) map[string][]*github.Repository {
	result := make(map[string][]*github.Repository)
	for _, u := range users {
		repos, resp, _ := c.GetUserRepos(*u.Login)
		models.UpdateRate(c.Token, resp)
		result[*u.Login] = repos
	}
	return result
}

func (c *Client) GetStrUsersRepos(users []string) map[string][]*github.Repository {
	result := make(map[string][]*github.Repository)
	for _, u := range users {
		repos, resp, _ := c.GetUserRepos(u)
		models.UpdateRate(c.Token, resp)
		result[u] = repos
	}
	return result
}

func (c *Client) GetUserOrgs(username string) ([]*github.Organization, *github.Response, error) {
	ctx := context.Background()
	return c.Client.Organizations.List(ctx, username, nil)
}

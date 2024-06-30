package config

type GITHUB struct {
	GithubUserName string `yaml:"github-user-name"`
	GithubRepoName string `yaml:"github-repo-name"`

	// GitHub个人访问令牌
	GithubAccessToken string `yaml:"github-access-token"`

	// 构建GitHub API请求URL "https://api.github.com/repos/" + username + "/" + repo + "/commits"
	GithubApiUrl string `yaml:"github-api-url"`
}

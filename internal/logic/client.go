package logic

import "github.com/rewardenv/reward-greeter/internal/config"

type Client struct {
	*config.Config
}

func New(c *config.Config) *Client {
	return &Client{
		c,
	}
}

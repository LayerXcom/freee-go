package freee

// Logger generic interface for logger
type Logger interface {
	Printf(string, ...interface{})
}

func (c *Client) logf(format string, a ...interface{}) {
	if c.config.Log != nil {
		c.config.Log.Printf(format, a...)
	}
}

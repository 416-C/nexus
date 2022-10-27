package nexus

import "time"

type Option func(*Options)

type Options struct {
	BaseURL          string
	Username         string
	Password         string
	Timeout          time.Duration
	RetryCount       int
	RetryWaitTime    time.Duration
	RetryMaxWaitTime time.Duration
}

func WithBaseURL(baseURL string) Option {
	return func(o *Options) {
		o.BaseURL = baseURL
	}
}

func WithBasicAuth(username, password string) Option {
	return func(o *Options) {
		o.Username = username
		o.Password = password
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

func WithRetryCount(retryCount int) Option {
	return func(o *Options) {
		o.RetryCount = retryCount
	}
}

func WithRetryWaitTime(retryWaitTime time.Duration) Option {
	return func(o *Options) {
		o.RetryWaitTime = retryWaitTime
	}
}

func WithRetryMaxWaitTime(retryMaxWaitTime time.Duration) Option {
	return func(o *Options) {
		o.RetryMaxWaitTime = retryMaxWaitTime
	}
}

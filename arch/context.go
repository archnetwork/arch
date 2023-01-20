package arch

import "context"

type Context struct {
	context.Context
	Analysis *Analysis `json:"analysis"`
	Route    *Route    `json:"route"`
	Cache    *Cache    `json:"cache"`
}

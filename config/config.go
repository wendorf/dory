package config

import (
	"path"

	"github.com/shibukawa/configdir"
)

func SocketPath() string {
	configDirs := configdir.New("", "dory")
	cache := configDirs.QueryCacheFolder()

	cache.MkdirAll()

	return path.Join(cache.Path, "socket")
}
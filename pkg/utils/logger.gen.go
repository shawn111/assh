// Code generated by moul.io/assh/contrib/generate-loggers.sh

package utils

import "go.uber.org/zap"

func logger() *zap.Logger {
	return zap.L().Named("assh.pkg.utils")
}

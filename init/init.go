package initialization

import (
	"os"
)

func InitEnv() {

	for key, val := range EnvMap {
		os.Setenv(key, val)
	}
}

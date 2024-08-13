package helpers

import (
	"encoding/json"


	"github.com/saidamir98/udevs_pkg/logger"
)

func DataParser[t any, t2 any](src t, dst t2) {
	byData, err := json.Marshal(src)

	if err != nil {

	logger.Error(err)
		return
	}

	json.Unmarshal(byData, dst)
}
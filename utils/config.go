package utils

import (
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type Witness struct {
	ID   string
	Name string
}

func GetWitnessesList(cfg *config.Config) []Witness {
	var witnesses []Witness

	for i := 0; ; i++ {
		nameKey := fmt.Sprintf("witness[%d].name", i)
		idKey := fmt.Sprintf("witness[%d].id", i)

		witnessName := cfg.Get(nameKey)
		witnessID := cfg.Get(idKey)

		if witnessName == "" && witnessID == "" {
			break
		}

		if witnessName == "" {
			witnessName = fmt.Sprintf("unnamed-witness-%d", i)
		}

		if witnessID == "" {
			witnessID = fmt.Sprintf("unknown-id-%d", i)
		}

		witnesses = append(witnesses, Witness{
			ID:   witnessID,
			Name: witnessName,
		})
	}

	return witnesses
}

package domain

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

type Key struct {
	UniqueURL string `json:"token_id"`
}

func getID() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	id := node.Generate()
	result := id.Base36()[3:]
	return result
}

type UniqKeys struct {
	Keys []string `json:"keys"`
}

package lib

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

var (
	snowflakeNode *snowflake.Node
)

func InitSnowflakeNode(_ int64) {
	// Initialising generator for making snowflake ID
	snowflakeNode, err = snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetNewSnowflake() int64 {
	fmt.Println(snowflakeNode.Generate())
	return snowflakeNode.Generate().Int64()
}

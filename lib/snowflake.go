package lib

import "github.com/bwmarrin/snowflake"

var (
	snowflakeNode *snowflake.Node
)

func InitialiseSnowflakeNode(nodeId int64) {
	// Initialising generator for making snowflake ID
	snowflakeNode, err = snowflake.NewNode(nodeId)
}

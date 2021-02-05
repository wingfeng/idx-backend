package utils

import (
	"github.com/bwmarrin/snowflake"
)

func GeneratID() string{
	node,_:=snowflake.NewNode(1)
	return node.Generate().String()
}
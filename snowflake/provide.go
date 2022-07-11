package snowflake

var defID *Node

var (
	// Generate ... alias snowflake.Generate
	Generate = defID.Generate
)

func init() {
	defID, _ = NewNode(1)
}

// SetDefault ...set new snowflake std
func SetDefault(std *Node) {
	defID = std

	Generate = defID.Generate
}

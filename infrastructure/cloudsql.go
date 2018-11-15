package infrastructure

import (
	"fmt"
)

// CloudSQLAccessor ...
type CloudSQLAccessor struct{}

// Duck ...
func (c *CloudSQLAccessor) Duck() {
	fmt.Println("CloudSQL!")
}

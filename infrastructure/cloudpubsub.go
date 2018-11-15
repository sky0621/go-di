package infrastructure

import (
	"fmt"
)

// CloudPubSubAccessor ...
type CloudPubSubAccessor struct{}

// Duck ...
func (p *CloudPubSubAccessor) Duck() {
	fmt.Println("PubSub!")
}

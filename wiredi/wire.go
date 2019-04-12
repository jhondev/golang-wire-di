//+build wireinject

package wiredi

import "github.com/google/wire"

// Injector
func InitializeEvent() Event {
	wire.Build(NewMessage, NewEvent, NewGreeter)
	return Event{}
}

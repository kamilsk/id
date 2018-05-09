package centrifuge

import (
	"time"
)

// historyFilter allows to provide several parameters for history
// extraction.
type historyFilter struct {
	// Limit sets the max amount of messages that must
	// be returned. 0 means no limit - i.e. return all history
	// messages (limited by configured history_size). 1 means
	// last message only, 2 - two last messages etc.
	Limit int
}

// presenceStats ...
type presenceStats struct {
	NumClients int
	NumUsers   int
}

// Engine is an interface abstracting PUB/SUB mechanics and
// history/presence data manipulations.
type Engine interface {
	// Name returns a name of concrete engine implementation.
	name() string
	// Run called once on start just after engine set to node.
	run() error

	// Publish allows to send Publication into channel. This message should
	// be delivered to all clients subscribed on this channel at moment on
	// any Centrifugo node. The returned value is channel in which we will
	// send error as soon as engine finishes publish operation. Also this
	// method must maintain history for channels if enabled in channel options.
	publish(ch string, pub *Publication, opts *ChannelOptions) <-chan error
	// PublishJoin publishes Join message into channel.
	publishJoin(ch string, join *Join, opts *ChannelOptions) <-chan error
	// PublishLeave publishes Leave message into channel.
	publishLeave(ch string, leave *Leave, opts *ChannelOptions) <-chan error
	// PublishControl allows to send control command to all running nodes.
	publishControl(data []byte) <-chan error

	// Subscribe node on channel.
	subscribe(ch string) error
	// Unsubscribe node from channel.
	unsubscribe(ch string) error
	// Channels returns slice of currently active channels (with
	// one or more subscribers) on all Centrifuge nodes.
	channels() ([]string, error)

	// History returns a slice of history messages for channel.
	history(ch string, filter historyFilter) ([]*Publication, error)
	// recoverHistory allows to recover missed messages starting from last seen
	// Publication UID provided by client. This method should return as many Publications
	// as possible and boolean value indicating whether all Publications
	// were successfully restored or not. The case when publications can not be
	// fully restored can happen if old Publications already removed from history
	// due to size or lifetime limits.
	recoverHistory(ch string, lastUID string) ([]*Publication, bool, error)
	// RemoveHistory removes history from channel. This is in general not
	// needed as history expires automatically (based on history_lifetime)
	// but sometimes can be useful for application logic.
	removeHistory(ch string) error

	// Presence returns actual presence information for channel.
	presence(ch string) (map[string]*ClientInfo, error)
	// PresenseStats returns short stats of current presence data.
	presenceStats(ch string) (presenceStats, error)
	// AddPresence sets or updates presence information in channel
	// for connection with specified identifier.
	addPresence(ch string, connID string, info *ClientInfo, expire time.Duration) error
	// RemovePresence removes presence information for connection
	// with specified identifier.
	removePresence(ch string, connID string) error
}

// Package hbtype provides types for hb drivers
package hbtype

import (
	"time"

	"github.com/opensvc/om3/core/event"
	"github.com/opensvc/om3/core/node"
)

type (
	// Msg struct holds all kinds of hb message
	Msg struct {
		Kind      string                   `json:"kind" yaml:"kind"`
		Compat    uint64                   `json:"compat" yaml:"compat"`
		Gen       map[string]uint64        `json:"gen" yaml:"gen"`
		UpdatedAt time.Time                `json:"updated_at" yaml:"updated_at"`
		Ping      node.Monitor             `json:"monitor,omitempty" yaml:"monitor,omitempty"`
		Events    map[string][]event.Event `json:"events,omitempty" yaml:"events,omitempty"`
		NodeData  node.Node                `json:"node_data,omitempty" yaml:"node_data,omitempty"`
		Nodename  string                   `json:"nodename" yaml:"nodename"`
	}

	// IdStopper
	IdStopper interface {
		Id() string
		Stop() error
	}

	// Transmitter is the interface that wraps the basic methods for hb driver to send hb messages
	Transmitter interface {
		IdStopper
		Start(cmdC chan<- interface{}, dataC <-chan []byte) error
	}

	// Receiver is the interface that wraps the basic methods for hb driver to receive hb messages
	Receiver interface {
		IdStopper
		Start(cmdC chan<- interface{}, msgC chan<- *Msg) error
	}
)

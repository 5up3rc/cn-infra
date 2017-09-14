// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package messaging

import (
	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/db/keyval"
)

// Mux defines API for the plugins that use access to kafka brokers.
type Mux interface {
	// NewSyncPublisher creates a publisher that allows to publish messages
	// using synchronous API.
	NewSyncPublisher(topic string) ProtoPublisher

	// NewSyncPublisherToPartition creates a publisher that allows to publish
	// messages to selected topic/partition using synchronous API
	NewSyncPublisherToPartition(topic string, partition int32) ProtoPublisher

	// NewAsyncPublisher creates a publisher that allows to publish messages
	// using asynchronous API.
	NewAsyncPublisher(topic string, successClb func(ProtoMessage), errorClb func(err ProtoMessageErr)) ProtoPublisher

	// NewAsyncPublisherToPartition creates a publisher that allows to publish
	// messages to selected topic/partition using asynchronous API.
	NewAsyncPublisherToPartition(topic string, partition int32,
		successClb func(ProtoMessage), errorClb func(err ProtoMessageErr)) ProtoPublisher

	// NewWatcher creates a watcher that allows to start/stop consuming
	// of messaging published to selected topics/partitions.
	NewWatcher(subscriberName string) ProtoWatcher
}

// ProtoPublisher allows to publish a message of type proto.Message into
// messaging system.
type ProtoPublisher interface {
	datasync.KeyProtoValWriter
}

// ProtoWatcher allows to subscribe for receiving of messages published
// to selected topics.
type ProtoWatcher interface {
	// Watch starts consuming all selected <topics>.
	// Callback <msgCallback> is called for each delivered message.
	Watch(msgCallback func(ProtoMessage), topics ...string) error

	// WatchPartition starts consuming specific <partition> of a selected <topic>
	// from a given <offset>.
	// Callback <msgCallback> is called for each delivered message.
	WatchPartition(msgCallback func(ProtoMessage), topic string, partition int32, offset int64) error

	// StopWatch cancels the previously created subscription for consuming
	// a given <topic>.
	StopWatch(topic string) error
}

// ProtoMessage exposes parameters of a single message received from messaging
// system.
type ProtoMessage interface {
	keyval.ProtoKvPair

	// GetTopic returns the name of the topic from which the message
	// was consumed.
	GetTopic() string

	// GetTopic returns the index of the partition from which the message
	// was consumed.
	GetPartition() int32
}

// ProtoMessageErr represents a message that was not published successfully
// to a messaging system.
type ProtoMessageErr interface {
	ProtoMessage

	// Error returns an error instance describing the cause of the failed
	// delivery.
	Error() error
}

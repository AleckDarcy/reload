// Copyright 2016 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v3rpc

import (
	"github.com/AleckDarcy/reload/core/log"

	"github.com/gogo/protobuf/proto"
)

type codec struct{}

// 3milebeach todo: boundary component where server sends RPC response to the client
func (c *codec) Marshal(v interface{}) ([]byte, error) {
	log.Logger.PrintlnWithCaller("stub") // 3milebeach

	b, err := proto.Marshal(v.(proto.Message))
	sentBytes.Add(float64(len(b)))
	return b, err
}

// 3milebeach todo: boundary component where server receives RPC request from the client
func (c *codec) Unmarshal(data []byte, v interface{}) error {
	log.Logger.PrintlnWithCaller("stub") // 3milebeach

	receivedBytes.Add(float64(len(data)))
	return proto.Unmarshal(data, v.(proto.Message))
}

func (c *codec) String() string {
	return "proto"
}

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"context"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	dubbo3pb "github.com/apache/dubbo-go-samples/api"
)

func TestSayHello(t *testing.T) {
	req := &dubbo3pb.HelloRequest{
		Name: "laurence",
	}

	reply := &dubbo3pb.User{}

	ctx := context.Background()

	reply, err := greeterProvider.SayHello(ctx, req)

	assert.Nil(t, err)
	assert.Equal(t, "Hello laurence", reply.Name)
	assert.Equal(t, "12345", reply.Id)
	assert.Equal(t, int32(21), reply.Age)
}

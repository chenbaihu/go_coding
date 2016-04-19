package thriftclient

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"idl/gen-go/rt"
	"strconv"
	"time"
)

func GetRTClient(host string, port int, ttimeoutmilli int64) (client *rt.MapServiceClient, transport thrift.TTransport, err error) {
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	p := strconv.Itoa(port)
	addr := host + ":" + p
	transport, err = thrift.NewTSocketTimeout(addr, time.Duration(ttimeoutmilli)*time.Millisecond)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return nil, nil, err
	}
	transport = transportFactory.GetTransport(transport)
	//defer transport.Close()
	if err := transport.Open(); err != nil {
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
		transport, err = thrift.NewTSocketTimeout(addr, time.Duration(ttimeoutmilli)*time.Millisecond)
		if err != nil {
			fmt.Println("Error opening socket:", err)
			return nil, nil, err
		}
		transport = transportFactory.GetTransport(transport)
	}
	return rt.NewMapServiceClientFactory(transport, protocolFactory), transport, nil
}


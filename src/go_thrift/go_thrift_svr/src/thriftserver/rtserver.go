package thriftserver 

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
	_"git.apache.org/thrift.git/lib/go/thrift"
	"idl/gen-go/rt"
	_"strconv"
	_"time"
)

type MapServiceHandle struct { 
    Name string
}

func (msh* MapServiceHandle) Compute(req *rt.ComputeReq) (r *rt.ComputeResp, err error) {
    fmt.Printf("req=%v\n", req);
    rsp := &rt.ComputeResp{Version:1, TypeA1:(int32)(req.TypeA1), Status:0, MapsplitId:5, Data:"hello client"}
    return rsp, nil
}

func NewMapServiceHandle() *MapServiceHandle {
    return &MapServiceHandle{Name:"MapServiceHandle"}
}

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

package http

import (
	"encoding/base64"

	"github.com/alipay/sofa-mosn/pkg/istio/utils"
	"github.com/alipay/sofa-mosn/pkg/types"
)

// CheckData extract HTTP data for Mixer check
type CheckData struct {
	reqHeaders  types.HeaderMap
	requestInfo types.RequestInfo
	connection  types.Connection
}

// NewCheckData return CheckData
func NewCheckData(reqHeaders types.HeaderMap, requestInfo types.RequestInfo, connection types.Connection) *CheckData {
	return &CheckData{
		reqHeaders:  reqHeaders,
		requestInfo: requestInfo,
		connection:  connection,
	}
}

// ExtractIstioAttributes Find "x-istio-attributes" HTTP header.
// If found, base64 decode its value,  pass it out
func (c *CheckData) ExtractIstioAttributes() (data string, ret bool) {
	val, ret := c.reqHeaders.Get(utils.KIstioAttributeHeader)
	if ret {
		d, _ := base64.StdEncoding.DecodeString(val)
		data = string(d)
	}
	return
}

// GetSourceIPPort get downstream tcp connection ip and port.
func (c *CheckData) GetSourceIPPort() (ip string, port int32, ret bool) {
	if c.connection != nil {
		ip, port, ret = utils.GetIPPort(c.connection.RemoteAddr())
		return
	}

	ret = false
	return
}
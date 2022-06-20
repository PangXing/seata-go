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

package codec

import (
	"github.com/seata/seata-go/pkg/common/binary"
	transaction2 "github.com/seata/seata-go/pkg/common/error"
	"github.com/seata/seata-go/pkg/protocol/message"
)

type CommonGlobalEndResponseCodec struct {
}

func (c *CommonGlobalEndResponseCodec) Encode(in interface{}) []byte {
	buf := binary.NewByteBuf(0)
	resp := in.(message.AbstractGlobalEndResponse)

	buf.WriteByte(byte(resp.ResultCode))
	if resp.ResultCode == message.ResultCodeFailed {
		var msg string
		if len(resp.Msg) > 128 {
			msg = resp.Msg[:128]
		} else {
			msg = resp.Msg
		}
		Write16String(msg, buf)
	}
	buf.WriteByte(byte(resp.TransactionExceptionCode))
	buf.WriteByte(byte(resp.GlobalStatus))

	return buf.RawBuf()
}

func (c *CommonGlobalEndResponseCodec) Decode(in []byte) interface{} {
	buf := binary.NewByteBuf(len(in))
	buf.Write(in)

	msg := message.AbstractGlobalEndResponse{}

	resultCode := ReadByte(buf)
	msg.ResultCode = message.ResultCode(resultCode)
	if msg.ResultCode == message.ResultCodeFailed {
		length := ReadUInt16(buf)
		if length > 0 {
			bytes := make([]byte, length)
			msg.Msg = string(Read(buf, bytes))
		}
	}

	exceptionCode := ReadByte(buf)
	msg.TransactionExceptionCode = transaction2.TransactionExceptionCode(exceptionCode)

	globalStatus := ReadByte(buf)
	msg.GlobalStatus = message.GlobalStatus(globalStatus)

	return msg
}
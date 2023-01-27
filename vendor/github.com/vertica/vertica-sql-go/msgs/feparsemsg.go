package msgs

// Copyright (c) 2019-2023 Micro Focus or one of its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import "fmt"

// FEParseMsg docs
type FEParseMsg struct {
	PreparedName string
	Command      string
	NumArgs      uint16
}

// Flatten docs
func (m *FEParseMsg) Flatten() ([]byte, byte) {

	buf := newMsgBuffer()

	buf.appendString(m.PreparedName)
	buf.appendString(m.Command)
	buf.appendUint16(m.NumArgs)

	for i := uint16(0); i < m.NumArgs; i++ {
		buf.appendUint32(0)
	}

	return buf.bytes(), 'P'
}

func (m *FEParseMsg) String() string {
	return fmt.Sprintf(
		"Parse: PreparedName='%s', Command='%s', NumArgs=%d",
		m.PreparedName,
		m.Command,
		m.NumArgs)
}

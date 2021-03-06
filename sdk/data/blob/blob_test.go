// Copyright 2018 The Containerfs Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package blob

import (
	"testing"
	"time"

	"fmt"
	"github.com/tiglabs/containerfs/util/log"
)

const (
	TestVolName    = "blob"
	TestMasterAddr = "10.196.31.173:8001,10.196.31.141:8001,10.196.30.200:8001"
	TestLogPath    = "testlog"
)

var gBlobClient *BlobClient

func init() {
	_, err := log.InitLog(TestLogPath, "Blob_UT", log.DebugLevel)
	if err != nil {
		panic(err)
	}

	bc, err := NewBlobClient(TestVolName, TestMasterAddr)
	if err != nil {
		panic(err)
	}
	gBlobClient = bc
}

func TestWrite(t *testing.T) {
	data := []byte("1234")
	key, err := gBlobClient.Write(data)
	if err != nil {
		t.Fatalf("Write: data(%v) err(%v)", string(data), err)
	}
	fmt.Println(fmt.Sprintf("Write key(%v) success", key))
	rdata, rerr := gBlobClient.Read(key)
	if rerr != nil {
		t.Fatalf("Read: key (%v) data(%v) err(%v)", key, string(rdata), rerr)
	}
	fmt.Println(fmt.Sprintf("Read key(%v) success", key))
	derr := gBlobClient.Delete(key)
	if derr != nil {
		t.Fatalf("deleteKey: key(%v) err(%v)", key, derr)
	}
	fmt.Println(fmt.Sprintf("Delete key(%v) success", key))

	time.Sleep(2 * time.Second)
}

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package netchan

import (
    "fmt"
    "testing"
    "netchan"
)

type value struct {
    i int
    s string
}

const count = 10

func exportSend(exp *Exporter, t *testing.T) {
    ch := make(chan value)
    err := exp.Export("exportedSend", ch, Send, new(value))
    if err != nil {
        t.Fatal("exportSend:", err)
    }
    for i := 0; i < count; i++ {
        ch <- value{23 + i, "hello"}
    }
}

func exportReceive(exp *Exporter, t *testing.T) {
    ch := make(chan value)
    err := exp.Export("exportedRecv", ch, Recv, new(value))
    if err != nil {
        t.Fatal("exportReceive:", err)
    }
    for i := 0; i < count; i++ {
        v := <-ch
        fmt.Printf("%v\n", v)
        if v.i != 45+i || v.s != "hello" {
            t.Errorf("export Receive: bad value: expected 4%d, hello; got %+v", 45+i, v)
        }
    }
}

func importReceive(imp *Importer, t *testing.T) {
    ch := make(chan value)
    err := imp.ImportNValues("exportedSend", ch, Recv, new(value), count)
    if err != nil {
        t.Fatal("importReceive:", err)
    }
    for i := 0; i < count; i++ {
        v := <-ch
        fmt.Printf("%v\n", v)
        if v.i != 23+i || v.s != "hello" {
            t.Errorf("importReceive: bad value: expected %d, hello; got %+v", 23+i, v)
        }
    }
}

func importSend(imp *Importer, t *testing.T) {
    ch := make(chan value)
    err := imp.ImportNValues("exportedRecv", ch, Send, new(value), count)
    if err != nil {
        t.Fatal("importSend:", err)
    }
    for i := 0; i < count; i++ {
        ch <- value{45 + i, "hello"}
    }
}

func TestExportSendImportReceive(t *testing.T) {
    exp, err := NewExporter("tcp", ":0")
    if err != nil {
        t.Fatal("new exporter:", err)
    }
    imp, err := NewImporter("tcp", exp.Addr().String())
    if err != nil {
        t.Fatal("new importer:", err)
    }
    go exportSend(exp, t)
    importReceive(imp, t)
}

func TestExportReceiveImportSend(t *testing.T) {
    exp, err := NewExporter("tcp", ":0")
    if err != nil {
        t.Fatal("new exporter:", err)
    }
    imp, err := NewImporter("tcp", exp.Addr().String())
    if err != nil {
        t.Fatal("new importer:", err)
    }
    go importSend(imp, t)
    exportReceive(exp, t)
}


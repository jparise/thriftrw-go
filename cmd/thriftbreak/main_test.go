// Copyright (c) 2021 Uber Technologies, Inc.
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

package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/thriftrw/internal/breaktest"
	"go.uber.org/thriftrw/internal/compare"
)

func TestThriftBreakIntegration(t *testing.T) {
	tests := []struct {
		desc     string
		want     string
		extraCmd string
	}{
		{
			desc: "output",
			want: `c.thrift:deleting service "Baz"` + "\n" +
				`d.thrift:deleting service "Qux"` + "\n" +
				`v2.thrift:deleting service "Bar"` + "\n" +
				`v1.thrift:removing method "methodA" in service "Foo"` + "\n" +
				`v1.thrift:adding a required field "C" to "AddedRequiredField"` + "\n",
		},
		{
			desc: "json output",
			want: `{"FilePath":"c.thrift","Message":"deleting service \"Baz\""}` + "\n" +
				`{"FilePath":"d.thrift","Message":"deleting service \"Qux\""}` + "\n" +
				`{"FilePath":"v2.thrift","Message":"deleting service \"Bar\""}` + "\n" +
				`{"FilePath":"v1.thrift","Message":"removing method \"methodA\" in service \"Foo\""}` + "\n" +
				`{"FilePath":"v1.thrift","Message":"adding a required field \"C\" to \"AddedRequiredField\""}` + "\n",
			extraCmd: "--json",
		},
	}
	from := map[string]string{
		"v1.thrift": "namespace rb v1\n" +
			"struct AddedRequiredField {\n" +
			"    1: optional string A\n" +
			"    2: optional string B\n" +
			"}\n" +
			"\nservice Foo {\n    void methodA()\n}",
		"test/v2.thrift": `service Bar {}`,
		"test/c.thrift":  `service Baz {}`,
		"test/d.thrift": `include "../v1.thrift"
		service Qux {}`, // d.thrift will be deleted below.
		"somefile.go": `service Quux{}`, // a .go file, not a .thrift.
	}
	// For c.thrift we are also checking to make sure includes work as expected.
	to := map[string]string{
		"v1.thrift": "namespace rb v1\n" +
			"struct AddedRequiredField {\n" +
			"    1: optional string A\n" +
			"    2: optional string B\n" +
			"    3: required string C\n}\n" +
			"service Foo {}",
		"test/v2.thrift": `service Foo {}`,
		"test/c.thrift": `include "../v1.thrift"
		service Bar {}`,
		"somefile.go": `service Qux{}`,
	}
	remove := []string{"test/d.thrift"}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tmpDir := t.TempDir()
			breaktest.CreateRepoAndCommit(t, tmpDir, from, to, remove)

			f, err := ioutil.TempFile(tmpDir, "stdout")
			require.NoError(t, err, "create temporary file")
			defer func(oldStdout *os.File) {
				assert.NoError(t, f.Close())
				os.Stdout = oldStdout
			}(os.Stdout)
			os.Stdout = f

			err = run([]string{"-C=" + tmpDir, tt.extraCmd})

			require.Error(t, err, "expected an error with Thrift backwards incompatible changes")
			assert.EqualError(t, err, "found 5 issues")

			stderr, err := ioutil.ReadFile(f.Name())
			require.NoError(t, err)

			out := string(stderr)
			assert.Equal(t, tt.want, out)
		})
	}
}

func TestDiagnosticPrinters(t *testing.T) {
	t.Parallel()
	tests := []struct {
		desc   string
		want   string
		writer func(io.Writer) func(compare.Diagnostic) error
	}{
		{
			desc:   "json writer",
			want:   `{"FilePath":"foo.thrift","Message":"error"}` + "\n",
			writer: jsonOutput,
		},
		{
			desc:   "readable writer",
			want:   "foo.thrift:error\n",
			writer: readableOutput,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			var b bytes.Buffer
			w := tt.writer(&b)
			err := w(compare.Diagnostic{
				FilePath: "foo.thrift",
				Message:  "error",
			})
			require.NoError(t, err)
			assert.Equal(t, tt.want, b.String())
		})
	}
}

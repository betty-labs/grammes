// Copyright (c) 2018 Northwestern Mutual.
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

package manager

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/betty-labs/grammes/logging"
)

func TestDropVertexLabel(t *testing.T) {
	Convey("Given a string executor and drop query manager", t, func() {
		execute := func(string) ([][]byte, error) { return nil, nil }
		dm := newDropQueryManager(logging.NewNilLogger(), execute)
		Convey("When DropVertexLabel is called", func() {
			err := dm.DropVertexLabel("testlabel")
			Convey("The return error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestDropVertexLabelError(t *testing.T) {
	Convey("Given a string executor and drop query manager", t, func() {
		execute := func(string) ([][]byte, error) { return nil, errors.New("ERROR") }
		dm := newDropQueryManager(logging.NewNilLogger(), execute)
		Convey("When DropVertexLabel is called and encounters an error", func() {
			err := dm.DropVertexLabel("testlabel")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestDropVertexByID(t *testing.T) {
	Convey("Given a string executor and drop query manager", t, func() {
		execute := func(string) ([][]byte, error) { return nil, nil }
		dm := newDropQueryManager(logging.NewNilLogger(), execute)
		Convey("When DropVertexByID is called", func() {
			err := dm.DropVertexByID(1234)
			Convey("Then the return error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestDropVertexByIDError(t *testing.T) {
	Convey("Given a string executor and drop query manager", t, func() {
		execute := func(string) ([][]byte, error) { return nil, errors.New("ERROR") }
		dm := newDropQueryManager(logging.NewNilLogger(), execute)
		Convey("When DropVertexByID is called and encounters an error", func() {
			err := dm.DropVertexByID(1234)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestDropVerticesByString(t *testing.T) {
	Convey("Given a string executor and drop query manager", t, func() {
		execute := func(string) ([][]byte, error) { return nil, nil }
		dm := newDropQueryManager(logging.NewNilLogger(), execute)
		Convey("When DropVertexByString is called", func() {
			err := dm.DropVerticesByString("testquery")
			Convey("Then the return error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestDropVerticesByStringError(t *testing.T) {
	Convey("Given a string executor and drop query manager", t, func() {
		execute := func(string) ([][]byte, error) { return nil, errors.New("ERROR") }
		dm := newDropQueryManager(logging.NewNilLogger(), execute)
		Convey("When DropVertexByString is called and encounters an error", func() {
			err := dm.DropVerticesByString("testquery")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestDropVerticesByQuery(t *testing.T) {
	Convey("Given a string executor and drop query manager", t, func() {
		execute := func(string) ([][]byte, error) { return nil, nil }
		dm := newDropQueryManager(logging.NewNilLogger(), execute)
		Convey("When DropVertexByQuery is called", func() {
			var q mockQuery
			err := dm.DropVerticesByQuery(q)
			Convey("Then the return error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestDropVerticesByQueryError(t *testing.T) {
	Convey("Given a string executor and drop query manager", t, func() {
		execute := func(string) ([][]byte, error) { return nil, errors.New("ERROR") }
		dm := newDropQueryManager(logging.NewNilLogger(), execute)
		Convey("When DropVertexByQuery is called and encounters an error", func() {
			var q mockQuery
			err := dm.DropVerticesByQuery(q)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

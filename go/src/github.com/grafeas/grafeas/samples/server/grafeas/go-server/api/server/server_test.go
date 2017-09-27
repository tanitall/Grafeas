// Copyright 2017 The Grafeas Authors. All rights reserved.
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

package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"

	"errors"
	"fmt"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/storage"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/testing"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/v1alpha1"
)

func TestHandler_CreateNote(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{storage.NewMemStore()}}
	n := testutil.Note()
	if err := createNote(n, h); err != nil {
		t.Errorf("%v", err)
	}
}

func TestHandler_CreateOccurrence(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{storage.NewMemStore()}}
	n := testutil.Note()
	if err := createNote(n, h); err != nil {
		t.Fatalf("Error creating note: %v", err)
	}
	o := testutil.Occurrence(n.Name)
	if err := createOccurrence(o, h); err != nil {
		t.Errorf("%v", err)
	}
}

func TestHandler_CreateOperation(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{storage.NewMemStore()}}
	o := testutil.Operation()
	if err := createOperation(o, h); err != nil {
		t.Errorf("%v", err)
	}
}

func createOccurrence(o swagger.Occurrence, g Handler) error {
	rawOcc, err := json.Marshal(&o)
	reader := bytes.NewReader(rawOcc)
	if err != nil {
		return errors.New(fmt.Sprintf("error marshalling json: %v", err))
	}
	r, err := http.NewRequest("POST",
		"/v1alpha1/projects/test-project/occurrences", reader)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating http request %v", err))
	}
	w := httptest.NewRecorder()
	g.CreateOccurrence(w, r)
	if w.Code != 200 {
		return errors.New(fmt.Sprintf("CreateOccurrence(%v) got %v want 200", o, w.Code))
	}
	return nil
}

func createNote(n swagger.Note, g Handler) error {
	rawNote, err := json.Marshal(&n)
	reader := bytes.NewReader(rawNote)
	if err != nil {
		return errors.New(fmt.Sprintf("error marshalling json: %v", err))
	}
	r, err := http.NewRequest("POST",
		"/v1alpha1/projects/vulnerability-scanner-a/notes?note_id=CVE-1999-0710", reader)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating http request %v", err))
	}
	w := httptest.NewRecorder()
	g.CreateNote(w, r)
	if w.Code != 200 {
		return errors.New(fmt.Sprintf("CreateNote(%v) got %v want 200", n, w.Code))
	}
	return nil
}

func createOperation(o swagger.Operation, g Handler) error {
	rawOp, err := json.Marshal(&o)
	reader := bytes.NewReader(rawOp)
	if err != nil {
		return errors.New(fmt.Sprintf("error marshalling json: %v", err))
	}
	r, err := http.NewRequest("POST",
		"/v1alpha1/projects/vulnerability-scanner-a/operations", reader)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating http request %v", err))
	}
	w := httptest.NewRecorder()
	g.CreateOperation(w, r)
	if w.Code != 200 {
		return errors.New(fmt.Sprintf("CreateNote(%v) got %v want 200", o, w.Code))
	}
	return nil
}

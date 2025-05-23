/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package jwk implements did:jwk method
package jwk

import (
	"errors"

	diddoc "github.com/trustbloc/did-go/doc/did"
	vdrapi "github.com/trustbloc/did-go/vdr/api"
)

const (
	// DIDMethod did method.
	DIDMethod = "jwk"
)

// VDR implements did:jwk method support.
type VDR struct{}

// New returns new instance of VDR that works with did:jwk method.
func New() *VDR {
	return &VDR{}
}

// Accept accepts did:jwk method.
func (v *VDR) Accept(method string, opts ...vdrapi.DIDMethodOption) bool {
	return method == DIDMethod
}

// Close frees resources being maintained by VDR.
func (v *VDR) Close() error {
	return nil
}

// Update did doc.
func (v *VDR) Update(didDoc *diddoc.Doc, opts ...vdrapi.DIDMethodOption) error {
	return errors.New("not supported")
}

// Deactivate did doc.
func (v *VDR) Deactivate(didID string, opts ...vdrapi.DIDMethodOption) error {
	return errors.New("not supported")
}

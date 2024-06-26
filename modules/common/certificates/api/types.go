// Copyright 2017 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import "crypto/tls"

const (
	// DashboardCertName is the certificate file names that will be generated by Dashboard
	DashboardCertName = "dashboard.crt"
	// DashboardKeyName is the key file names that will be generated by Dashboard
	DashboardKeyName = "dashboard.key"
)

// Manager is responsible for generating and storing self-signed certificates that can be used by Dashboard
// to serve over HTTPS.
type Manager interface {
	// GetCertificates loads existing certificates or generates self-signed certificates.
	GetCertificates() ([]tls.Certificate, error)

	GetCertificatePaths() (string, string, error)
}

// Creator is responsible for preparing and generating certificates.
type Creator interface {
	// GenerateKey generates certificate key
	GenerateKey() interface{}
	// GenerateCertificate generates certificate
	GenerateCertificate(key interface{}) []byte
	// StoreCertificates saves certificates in a given path
	StoreCertificates(path string, key interface{}, certBytes []byte) (string, string)
	// KeyCertPEMBytes converts the key and cert to PEM format
	KeyCertPEMBytes(key interface{}, certBytes []byte) (keyPEM []byte, certPEM []byte, err error)
	// GetKeyFileName returns certificate key file name
	GetKeyFileName() string
	// GetCertFileName returns certificate file name
	GetCertFileName() string
}

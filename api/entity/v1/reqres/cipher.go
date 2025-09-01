//    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import "github.com/spiffe/spike-sdk-go/api/entity/data"

// CipherEncryptRequest for encrypting data
type CipherEncryptRequest struct {
	// Plaintext data to encrypt
	Plaintext []byte `json:"plaintext"`
	// Optional: specify encryption algorithm/version
	Algorithm string `json:"algorithm,omitempty"`
}

// CipherEncryptResponse contains encrypted data
type CipherEncryptResponse struct {
	// Version byte for future compatibility
	Version byte `json:"version"`
	// Nonce used for encryption
	Nonce []byte `json:"nonce"`
	// Encrypted ciphertext
	Ciphertext []byte `json:"ciphertext"`
	// Error code if operation failed
	Err data.ErrorCode `json:"err,omitempty"`
}

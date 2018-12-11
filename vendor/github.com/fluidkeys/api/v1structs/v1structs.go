package v1structs

import (
	"time"
)

// GetPublicKeyResponse is the JSON structure returned by the get public key
// API endpoint. See:
// https://github.com/fluidkeys/api/blob/master/README.md#get-a-public-key
type GetPublicKeyResponse struct {
	// ArmoredPublicKey is the ASCII-armored OpenPGP public key.
	ArmoredPublicKey string `json:"armoredPublicKey"`
}

// UpsertPublicKeyRequest is a request to create or update a public key.
type UpsertPublicKeyRequest struct {
	// ArmoredPublicKey is the public key to be created or updated
	ArmoredPublicKey string `json:"armoredPublicKey"`

	// ArmoredSignedJSON is an ASCII-armored message, decoding to a JSON
	// message which decodes as an UpsertPublicKeySignedData
	ArmoredSignedJSON string `json:"armoredSignedJSON"`
}

// UpsertPublicKeySignedData is data self-signed by the public key being
// upserted to ensure that only the owner of a public key can upload it
// (a third party can't generate a valid signature)
type UpsertPublicKeySignedData struct {
	// The client's current time which must be within 24 hours of the
	// server's timestamp
	Timestamp time.Time `json:"timestamp"`

	// SingleUseUUID is a random UUID that is used once and must not be used
	// again. Its purpose is to prevent replay attacks where signed JSON
	// is re-sent to the API at a later date (possibly at a different
	// endpoint)
	SingleUseUUID string `json:"singleUseUuid"`

	// PublicKeySHA256 is the SHA256 hash of the ArmoredPublicKey in the
	// outer request
	PublicKeySHA256 string `json:"publicKeySha256"`
}

// UpsertPublicKeyResponse is the JSON response returned from the upsert public
// key endpoint.
type UpsertPublicKeyResponse struct {

	// ArmoredEncryptedBasicAuthPassword, when decrypted, contains a
	// system-generated password that can be used to authenticate
	// subsequent API calls using HTTP basic auth.
	ArmoredEncryptedBasicAuthPassword string `json:"armoredEncryptedBasicAuthPassword"`
}

// SendSecretRequest is the JSON structure used for requests to the send secret
// API endpoint. See:
// https://github.com/fluidkeys/api/blob/master/README.md#send-a-secret-to-a-public-key
type SendSecretRequest struct {
	RecipientFingerprint   string `json:"recipientFingerprint"`
	ArmoredEncryptedSecret string `json:"armoredEncryptedSecret"`
}

// ListSecretsResponse is the JSON structure returned by the list secrets
// API endpoint. See:
// https://github.com/fluidkeys/api/blob/master/README.md#list-your-secrets
type ListSecretsResponse struct {
	Secrets []Secret `json:"secrets"`
}

// Secret is the JSON structure containing the metadata and content for an
// encrypted secret returned by the list secrets API endpoint.
type Secret struct {
	// EncryptedMetadata is an ASCII-armored encrypted PGP message which
	// decrypts to a `SecretMetadata` JSON structure.
	EncryptedMetadata string `json:"encryptedMetadata"`

	// EncryptedContent is an ASCII-armored encrypted PGP message
	// containing the actual content of the secret.
	EncryptedContent string `json:"encryptedContent"`
}

// SecretMetadata contains non-content information about an encrypted secret.
type SecretMetadata struct {
	// SecretUUID uniquely identifies the secret to the API
	SecretUUID string `json:"secretUuid"`
}

// ErrorResponse is the JSON structure returned when the API encounters an
// error.
type ErrorResponse struct {
	// Detail is a human-readable string describing the error.
	Detail string `json:"detail"`
}

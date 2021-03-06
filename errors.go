package dtls

import "errors"

// Typed errors
var (
	ErrConnClosed = errors.New("dtls: conn is closed")

	errBufferTooSmall                    = errors.New("dtls: buffer is too small")
	errCertificateUnset                  = errors.New("dtls: handshakeMessageCertificate can not be marshalled without a certificate")
	errClientCertificateRequired         = errors.New("dtls: server required client verification, but got none")
	errClientCertificateNotVerified      = errors.New("dtls: client sent certificate but did not verify it")
	errCertificateVerifyNoCertificate    = errors.New("dtls: client sent certificate verify but we have no certificate to verify")
	errCipherSuiteNoIntersection         = errors.New("dtls: Client+Server do not support any shared cipher suites")
	errCipherSuiteUnset                  = errors.New("dtls: server hello can not be created without a cipher suite")
	errCompressionmethodUnset            = errors.New("dtls: server hello can not be created without a compression method")
	errContextUnsupported                = errors.New("dtls: context is not supported for ExportKeyingMaterial")
	errCookieMismatch                    = errors.New("dtls: Client+Server cookie does not match")
	errCookieTooLong                     = errors.New("dtls: cookie must not be longer then 255 bytes")
	errDTLSPacketInvalidLength           = errors.New("dtls: packet is too short")
	errHandshakeInProgress               = errors.New("dtls: Handshake is in progress")
	errHandshakeMessageUnset             = errors.New("dtls: handshake message unset, unable to marshal")
	errInvalidCipherSpec                 = errors.New("dtls: cipher spec invalid")
	errInvalidCipherSuite                = errors.New("dtls: invalid or unknown cipher suite")
	errInvalidCompressionMethod          = errors.New("dtls: invalid or unknown compression method")
	errInvalidContentType                = errors.New("dtls: invalid content type")
	errInvalidECDSASignature             = errors.New("dtls: ECDSA signature contained zero or negative values")
	errInvalidEllipticCurveType          = errors.New("dtls: invalid or unknown elliptic curve type")
	errInvalidExtensionType              = errors.New("dtls: invalid extension type")
	errInvalidHashAlgorithm              = errors.New("dtls: invalid hash algorithm")
	errInvalidMAC                        = errors.New("dtls: invalid mac")
	errInvalidNamedCurve                 = errors.New("dtls: invalid named curve")
	errInvalidPrivateKey                 = errors.New("dtls: invalid private key type")
	errInvalidSignatureAlgorithm         = errors.New("dtls: invalid signature algorithm")
	errKeySignatureGenerateUnimplemented = errors.New("dtls: Unable to generate key signature, unimplemented")
	errKeySignatureMismatch              = errors.New("dtls: Expected and actual key signature do not match")
	errKeySignatureVerifyUnimplemented   = errors.New("dtls: Unable to verify key signature, unimplemented")
	errLengthMismatch                    = errors.New("dtls: data length and declared length do not match")
	errNilNextConn                       = errors.New("dtls: Conn can not be created with a nil nextConn")
	errNotEnoughRoomForNonce             = errors.New("dtls: Buffer not long enough to contain nonce")
	errNotImplemented                    = errors.New("dtls: feature has not been implemented yet")
	errReservedExportKeyingMaterial      = errors.New("dtls: ExportKeyingMaterial can not be used with a reserved label")
	errSequenceNumberOverflow            = errors.New("dtls: sequence number overflow")
	errServerMustHaveCertificate         = errors.New("dtls: Certificate is mandatory for server")
	errUnableToMarshalFragmented         = errors.New("dtls: unable to marshal fragmented handshakes")
	errVerifyDataMismatch                = errors.New("dtls: Expected and actual verify data does not match")
	errNoConfigProvided                  = errors.New("dtls: No config provided")
	errPSKAndCertificate                 = errors.New("dtls: Certificate and PSK or PSK Identity Hint provided")
	errNoAvailableCipherSuites           = errors.New("dtls: Connection can not be created, no CipherSuites satisfy this Config")
)

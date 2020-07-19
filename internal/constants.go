package internal

const (
	// DefaultServicePort is the port on which the svc is going to Listen and serve
	DefaultServicePort = 3333

	// Shift is the number of shift offset to be used to encrypt the string in encode endpoint
	Shift = 32

	// CeaserCipherEncode will be used to trigger the encoding operation for ceaser cipher using the helper function
	CeaserCipherEncode = 1

	// CeaserCipherDecode will be used to trigger the decoding operation for ceaser cipher using the helper function
	CeaserCipherDecode = -1

	// AlphabetsLength is the length of alphabets
	AlphabetsLength = 26
)

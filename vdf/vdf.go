package vdf

// VDF A Verifiable Delay Function (VDF).  VDFs are problems that require a
// certain amount of time to solve, even on a parallel machine, but can be
// validated much more easily.
type VDF interface {
	Solve(challenge []byte, difficulty uint64) ([]byte, error)
	Verify(challenge []byte, difficulty uint64, alleged_solution []byte) error
}

type WesolowskiVDF struct {
	intSizeBits int
}

func NewWesolowskiVDF(intSizeBits int) *WesolowskiVDF {
	return &WesolowskiVDF{
		intSizeBits: intSizeBits,
	}
}

func (v *WesolowskiVDF) Solve(challenge []byte, difficulty int) ([]byte, error) {
	yBuf, proofBuf := GenerateVDF(challenge[:], difficulty, v.intSizeBits)
	elementLength := 2 * ((v.intSizeBits + 16) >> 4)
	proofLenInBytes := 2 * elementLength
	output := make([]byte, proofLenInBytes)

	copy(output[:elementLength], yBuf)
	copy(output[elementLength:], proofBuf)

	return output, nil
}

func (v *WesolowskiVDF) Verify(challenge []byte, difficulty int, alleged_solution []byte) bool {
	return VerifyVDF(challenge[:], alleged_solution[:], difficulty, v.intSizeBits)
}

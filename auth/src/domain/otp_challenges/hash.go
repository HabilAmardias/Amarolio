package otpchallenges

type HasherItf interface {
	Hash(plainText string) (string, error)
	Validate(encodedHash, plainText string) (bool, error)
}

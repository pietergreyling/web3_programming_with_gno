// guess.gno (realm)

// see: https://docs.onbloc.xyz/docs/building-a-realm/realm-examples/hello.gno

package guess

var (
	min, max = 1, 100
	tries    = 0
	correct  = false
	secret   = rand.Intn(max-min) + min
)

func GetSecret() {
	return secret
}

func SetSecret(secret int) {
	secret = secret
}

func SetRandomSecret(secret int) {
	SetSecret(rand.Intn(max-min) + min)
}

func Render(path string) string {
	return "-- Render(path string) called with\n# " + path
}

package otherpkg

import "github.com/mitar/accurate-test-coverage/pkg"

func SayCodeCoverage() string {
	return pkg.SayCode() + " " + pkg.SayCoverage()
}

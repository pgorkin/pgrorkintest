package documented

// PublicMethodDoesNothing just to check linting
func PublicMethodDoesNothing() {

}

type privateStruct struct {
	privateOne string
	privateTwo string
}

// PublicStructt balbla
type PublicStruct struct {
	// PublicOne 1
	PublicOne string
	// PublicTwo 2
	PublicTwo string
}

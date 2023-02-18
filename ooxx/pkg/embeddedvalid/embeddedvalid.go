package embeddedvalid

type EmbeddedValid struct {
	isValid bool
}

func (ev *EmbeddedValid) PassValid() {
	ev.isValid = true
}

func (ev *EmbeddedValid) IsValid() bool {
	return ev.isValid
}

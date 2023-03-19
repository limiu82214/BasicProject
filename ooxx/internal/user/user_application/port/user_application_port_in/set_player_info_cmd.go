package user_application_port_in

import "github.com/pkg/errors"

var errFailNicknameLen = errors.New("nickname length wrong")

type SetPlayerInfoCmd struct {
	Nickname string
	isValid  bool
}

func NewSetPlayerInfoCmd(nickname string) (*SetPlayerInfoCmd, error) {
	l := len(nickname)
	if l < 1 || l > 3 {
		return nil, errFailNicknameLen
	}

	return &SetPlayerInfoCmd{
		Nickname: nickname,
		isValid:  true,
	}, nil
}

func (spic *SetPlayerInfoCmd) IsValid() bool {
	return spic.isValid
}

package game_domain

import "github.com/pkg/errors"

var ErrGetEmpty = errors.New("get empty")
var ErrSetState = errors.New("to SetState in domain board")

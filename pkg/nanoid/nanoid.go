package nanoid

import (
	"sync"

	"github.com/jaevor/go-nanoid"
)

var (
	_once      sync.Once
	_generator func() string
)

func initGenerator() {
	var err error
	_generator, err = nanoid.Canonic()
	if err != nil {
		panic(err)
	}
}

func Gen() string {
	_once.Do(initGenerator)
	return _generator()
}

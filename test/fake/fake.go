package fake

import (
	"errors"
)

var (
	ErrActual       = errors.New("actual")
	ErrElemExsit    = errors.New("elem already exist")
	ErrElemNotExsit = errors.New("elem not exist")
)

func Exec(cmd string, args ...string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func Belong(points string, lines []string) bool { _ = "STUB: not implemented"; return false }

type Slice []int

func NewSlice() Slice { _ = "STUB: not implemented"; return *new(Slice) }

func (this *Slice) Add(elem int) error { _ = "STUB: not implemented"; return nil }

func (this *Slice) Remove(elem int) error { _ = "STUB: not implemented"; return nil }

func (this *Slice) Append(elems ...int) int { _ = "STUB: not implemented"; return 0 }

func ReadLeaf(url string) (string, error) { _ = "STUB: not implemented"; return "", nil }

type Etcd struct {
}

func (this *Etcd) Retrieve(url string) (string, error) { _ = "STUB: not implemented"; return "", nil }

var Marshal = func(v interface{}) ([]byte, error) {
	return nil, nil
}

type Db interface {
	Retrieve(url string) (string, error)
}

type Mysql struct {
}

func (this *Mysql) Retrieve(url string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func NewDb(style string) Db { _ = "STUB: not implemented"; return *new(Db) }

type PrivateMethodStruct struct {
}

func (this *PrivateMethodStruct) ok() bool { _ = "STUB: not implemented"; return false }

func (this *PrivateMethodStruct) Happy() string { _ = "STUB: not implemented"; return "" }

func (this PrivateMethodStruct) haveEaten() bool { _ = "STUB: not implemented"; return false }

func (this PrivateMethodStruct) AreYouHungry() string { _ = "STUB: not implemented"; return "" }

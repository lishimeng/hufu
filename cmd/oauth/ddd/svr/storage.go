package svr

import (
	"errors"
	"fmt"
)

type Storage interface {
	RemoveCode(code string) error

	RemoveAccess(access string) error

	RemoveRefresh(refresh string) error

	GetCode(code string) (clientId string, err error)

	GetAccess(access string) (clientId string, err error)

	GetRefresh(refresh string) (clientId string, err error)

	SaveCode(code string, clientId string) (err error)

	SaveAccess(access string, clientId string) (err error)

	SaveRefresh(refresh string, clientId string) (err error)
}

const (
	codeKey    = "oauth_c_%s"
	accessKey  = "oauth_a_%s"
	refreshKey = "oauth_r_%s"
)

var (
	ErrNotFound = errors.New("not found")
)

type MemoryStorage struct {
	mem map[string]string
}

func NewMemStorage() (s Storage) {
	ms := MemoryStorage{mem: make(map[string]string)}
	s = &ms
	return
}
func (ms *MemoryStorage) RemoveCode(code string) error {
	key := fmt.Sprintf(codeKey, code)
	delete(ms.mem, key)
	return nil
}

func (ms *MemoryStorage) RemoveAccess(access string) error {
	key := fmt.Sprintf(accessKey, access)
	delete(ms.mem, key)
	return nil
}

func (ms *MemoryStorage) RemoveRefresh(refresh string) error {
	key := fmt.Sprintf(refreshKey, refresh)
	delete(ms.mem, key)
	return nil
}

func (ms *MemoryStorage) GetCode(code string) (clientId string, err error) {
	key := fmt.Sprintf(codeKey, code)
	clientId, ok := ms.mem[key]
	if !ok {
		err = ErrNotFound
		return
	}
	return
}

func (ms *MemoryStorage) GetAccess(access string) (clientId string, err error) {
	key := fmt.Sprintf(accessKey, access)
	clientId, ok := ms.mem[key]
	if !ok {
		err = ErrNotFound
		return
	}
	return
}

func (ms *MemoryStorage) GetRefresh(refresh string) (clientId string, err error) {
	key := fmt.Sprintf(refreshKey, refresh)
	clientId, ok := ms.mem[key]
	if !ok {
		err = ErrNotFound
		return
	}
	return
}

func (ms *MemoryStorage) SaveCode(code string, clientId string) (err error) {
	key := fmt.Sprintf(codeKey, code)
	ms.mem[key] = clientId
	return
}

func (ms *MemoryStorage) SaveAccess(access string, clientId string) (err error) {
	key := fmt.Sprintf(accessKey, access)
	ms.mem[key] = clientId
	return
}

func (ms *MemoryStorage) SaveRefresh(refresh string, clientId string) (err error) {
	key := fmt.Sprintf(refreshKey, refresh)
	ms.mem[key] = clientId
	return
}

package files

import (
	"TestGOBot/lib/e"
	"TestGOBot/storage"
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("cannot save page", err) }()

	filePath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}
	filePath = filepath.Join(filePath, fName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}
	return nil
}

func (s Storage) PickRandom(UserName string) (page *storage.Page, err error) {
	defer func() { err = e.WrapIfErr("cannot pick random", err) }()

	filePath := filepath.Join(s.basePath, page.UserName)

	files, err := os.ReadDir(filePath)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, storage.ErrNoSavedPages
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(filePath, file.Name()))
}

func (s Storage) Remove(p *storage.Page) error {
	fileName, err := fileName(p)
	if err != nil {
		return e.Wrap("cannot remove file", err)
	}
	path := filepath.Join(s.basePath, p.UserName, fileName)
	if err := os.Remove(path); err != nil {
		msg := fmt.Sprintf("cannot remove file %s", path)

		return e.Wrap(msg, err)
	}
	return nil
}

func (s Storage) IsExists(page *storage.Page) (bool, error) {
	fileName, err := fileName(page)
	if err != nil {
		return false, e.Wrap("cannot check if file exists", err)
	}
	path := filepath.Join(s.basePath, page.UserName, fileName)

	switch _, err := os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("cannot check if file %s exists", path)

		return false, e.Wrap(msg, err)
	}
	return true, nil
}

func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, e.Wrap("cannot decode page", err)
	}
	defer func() { _ = f.Close() }()

	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, e.Wrap("cannot decode page", err)
	}
	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}

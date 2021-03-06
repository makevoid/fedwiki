package pagestore

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/egonelbre/fedwiki"
)

// Load loads page from `filename` and adds `slug` to it
func Load(filename string, slug fedwiki.Slug) (*fedwiki.Page, error) {
	data, err := ioutil.ReadFile(filename)
	err = ConvertOSError(err)
	if err != nil {
		return nil, err
	}

	page := &fedwiki.Page{}
	err = json.Unmarshal(data, page)
	if err != nil {
		return nil, err
	}

	if page.PageHeader.Date.IsZero() {
		if info, err := os.Stat(filename); err == nil {
			page.PageHeader.Date = fedwiki.Date{info.ModTime()}
		} else {
			page.PageHeader.Date = fedwiki.Date{page.LastModified()}
		}
	}

	page.PageHeader.Slug = slug
	return page, nil
}

// LoadHeader loads header from `filename` and adds `slug` to it
func LoadHeader(filename string, slug fedwiki.Slug) (*fedwiki.PageHeader, error) {
	data, err := ioutil.ReadFile(filename)
	err = ConvertOSError(err)
	if err != nil {
		return nil, err
	}

	header := &fedwiki.PageHeader{}
	err = json.Unmarshal(data, header)
	if err != nil {
		return nil, err
	}

	if header.Date.IsZero() {
		if info, err := os.Stat(filename); err == nil {
			header.Date = fedwiki.Date{info.ModTime()}
		}
	}

	header.Slug = slug
	return header, nil
}

// Create creates a new federated wiki page on disk with `filename`
func Create(page *fedwiki.Page, filename string) error {
	data, err := json.Marshal(page)
	if err != nil {
		return err
	}

	//TODO: handle case when it exists
	return ConvertOSError(ioutil.WriteFile(filename, data, 0755))
}

// Save updates the page on disk
func Save(page *fedwiki.Page, filename string) error {
	data, err := json.Marshal(page)
	if err != nil {
		return err
	}

	return ConvertOSError(ioutil.WriteFile(filename, data, 0755))
}

package server

import "fmt"

func (c *connection) FileCount() (int64, error) {
	return c.File.Query().Count()
}

func (c *connection) FileByPath(path string) (*File, error) {
	list, err := c.File.Query().Where("path", path).Run()
	if err != nil {
		return nil, err
	}

	if len(list) > 1 {
		return nil, fmt.Errorf("more than one file found for path: %s", path)
	}

	if len(list) == 0 {
		return &File{Path: path}, nil
	}

	return list[0], nil
}

// Scry Info.  All rights reserved.
// license that can be found in the license file.

package binary

import "github.com/scryinfo/dot/dot"

type Binary struct {
}

func NewBinary() *Binary {
	return &Binary{}
}

func (c *Binary) Create(l dot.Line) error {
	return nil
}

func (c *Binary) Start(ignore bool) error {
	return nil
}

func (c *Binary) Stop(ignore bool) error {
	return nil
}

func (c *Binary) Destroy(ignore bool) error {
	return nil
}

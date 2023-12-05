// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redblacktree

import (
	"encoding/json"
	"github.com/vuuvv/errors"
	"github.com/vuuvv/gone/utils"
)

// ToJSON outputs the JSON representation of the tree.
func (tree *Tree[K, V]) ToJSON() ([]byte, error) {
	elements := make(map[string]V)
	it := tree.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	ret, err := json.Marshal(&elements)
	return ret, errors.WithStack(err)
}

// FromJSON populates the tree from the input JSON representation.
func (tree *Tree[K, V]) FromJSON(data []byte) error {
	var tmp any = tree
	stringTree, ok := tmp.(*Tree[string, V])
	if !ok {
		return errors.New("tree must be of type *Tree[string, V]")
	}
	elements := make(map[string]V)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		stringTree.Clear()
		for key, value := range elements {
			stringTree.Put(key, value)
		}
	}
	return errors.WithStack(err)
}

// UnmarshalJSON @implements json.Unmarshaler
func (tree *Tree[K, V]) UnmarshalJSON(bytes []byte) error {
	return tree.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (tree *Tree[K, V]) MarshalJSON() ([]byte, error) {
	return tree.ToJSON()
}

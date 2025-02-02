// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	// "time"

	commonmodels "github.com/goharbor/harbor/src/common/models"
)

// User ...
type User = commonmodels.User

// Users the collection for User
type Users []*User

// MapByUserID returns map which key is UserID of the user and value is the user itself
func (users Users) MapByUserID() map[int]*User {
	m := map[int]*User{}
	for _, user := range users {
		m[user.UserID] = user
	}

	return m
}

// UserTable is the name of table in DB that holds the user object
const UserTable = "harbor_user"

// Option ...
type Option func(*Options)

// Options ...
type Options struct {
	IncludeDefaultAdmin bool
}

// WithDefaultAdmin set the IncludeAdmin = true
func WithDefaultAdmin() Option {
	return func(o *Options) {
		o.IncludeDefaultAdmin = true
	}
}

// NewOptions ...
func NewOptions(options ...Option) *Options {
	opts := &Options{}
	for _, f := range options {
		f(opts)
	}
	return opts
}

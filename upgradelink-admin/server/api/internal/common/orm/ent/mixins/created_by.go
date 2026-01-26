// Copyright 2023 The Ryan SU Authors (https://github.com/suyuan32). All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/gofrs/uuid/v5"
)

// CreatedByMixin for embedding the created user's uuid info in different schemas.
type CreatedByMixin struct {
	mixin.Schema
}

// Fields for all schemas that embed CreatedByMixin.
func (CreatedByMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("created_by", uuid.UUID{}).
			Optional().
			Comment("Created user's UUID | 创建者 UUID"),
	}
}

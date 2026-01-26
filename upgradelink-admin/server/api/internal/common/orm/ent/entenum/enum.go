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

package entenum

const (
	// TenantDefaultId is the default id of tenant
	TenantDefaultId uint64 = 1
)

const (
	// DataPermAll is the data permission of all data
	DataPermAll    = 1
	DataPermAllStr = "1"

	// DataPermCustomDept is the data permission of custom department data
	DataPermCustomDept    = 2
	DataPermCustomDeptStr = "2"

	// DataPermOwnDeptAndSub is the data permission of users's own department and sub departments data
	DataPermOwnDeptAndSub    = 3
	DataPermOwnDeptAndSubStr = "3"

	// DataPermOwnDept is the data permission of users's own department data
	DataPermOwnDept    = 4
	DataPermOwnDeptStr = "4"

	// DataPermSelf is the data permission of your own data
	DataPermSelf    = 5
	DataPermSelfStr = "5"
)

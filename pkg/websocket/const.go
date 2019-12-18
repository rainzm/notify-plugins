// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package websocket

const (
	AUTH_URI          = "auth_uri"
	ADMIN_USER        = "admin_user"
	ADMIN_PASSWORD    = "admin_password"
	ADMIN_TENANT_NAME = "admin_tenant_name"

	NOTINIT = "Send service hasn't been init"
)

var (
	FAIL_KEY = []string{"失败", "fail ", " failed"}
)

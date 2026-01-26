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

package captcha

// Conf is the captcha configuration structure
type Conf struct {
	KeyLong   int    `json:",optional,default=5,env=CAPTCHA_KEY_LONG"`                                       // captcha length
	ImgWidth  int    `json:",optional,default=240,env=CAPTCHA_IMG_WIDTH"`                                    // captcha width
	ImgHeight int    `json:",optional,default=80,env=CAPTCHA_IMG_HEIGHT"`                                    // captcha height
	Driver    string `json:",optional,default=digit,options=[digit,string,math,chinese],env=CAPTCHA_DRIVER"` // captcha type
}

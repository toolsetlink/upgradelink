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

import (
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"
	"image/color"
)

const zhDict = "的一是在不了有和人这中大为上个国我以要他时来用们生到作地于出就分对成会可主发年动同工也能下过子说产种面而方后多定行学法所民得经十三之进着等部度家电力里如水化高自二理起小物现实加量都两体制机当使点从业本去把性好应开它合还因由其些然前外天政四日那社义事平形相全表间样与关各重新线内数正心反你明看原又么利比或但质气第向道命此变条只没结解问意建月公无系军很情者最立代想已通并提直题党程展五果料象员革位入常文总次品式活设及管特件长求老头基资边流路级少图山统接知较将组见计别她手角期根论运农指几九区强放决西被干做必战先回则任取据处队南给色光门即保治北造百规热领七海口东导器压志世金增争济阶油思术极交受联什认六共权收证改清己美再采转更单风切打白教速花带安场身车例真务具万每目至达走积示议声报斗完类八离华名确才科张信马节话米整空元况今集温传土许步群广石记需段研界拉林律叫且究观越织装影算低持音众书布复容儿须际商非验连断深难近矿千周委素技备半办青省列习响约支般史感劳便团往酸历市克何除消构府称太准精值号率族维划选标写存候毛亲快效斯院查江型眼王按格养易置派层片始却专状育厂京识适属圆包火住调满县局照参红细引听该铁价严"

// MustNewRedisCaptcha returns the captcha using redis, it will exit when error occur
func MustNewRedisCaptcha(c Conf, r *redis2.Redis) *base64Captcha.Captcha {
	driver := NewDriver(c)

	store := NewRedisStore(r)

	return base64Captcha.NewCaptcha(driver, store)
}

// MustNewOriginalRedisCaptcha returns the captcha using original go redis, it will exit when error occur
func MustNewOriginalRedisCaptcha(c Conf, r redis.UniversalClient) *base64Captcha.Captcha {
	driver := NewDriver(c)

	store := NewOriginalRedisStore(r)

	return base64Captcha.NewCaptcha(driver, store)
}

func NewDriver(c Conf) base64Captcha.Driver {
	var driver base64Captcha.Driver

	bgColor := &color.RGBA{
		R: 254,
		G: 254,
		B: 254,
		A: 254,
	}

	fonts := []string{
		"ApothecaryFont.ttf",
		"DENNEthree-dee.ttf",
		"Flim-Flam.ttf",
		"RitaSmith.ttf",
		"actionj.ttf",
		"chromohv.ttf",
	}

	switch c.Driver {
	case "digit":
		driver = base64Captcha.NewDriverDigit(c.ImgHeight, c.ImgWidth,
			c.KeyLong, 0.7, 80)
	case "string":
		driver = base64Captcha.NewDriverString(c.ImgHeight, c.ImgWidth, 12, 3, c.KeyLong,
			"qwertyupasdfghjkzxcvbnm23456789",
			bgColor, nil, fonts)
	case "math":
		driver = base64Captcha.NewDriverMath(c.ImgHeight, c.ImgWidth, 12, 3, bgColor,
			nil, fonts)
	case "chinese":
		driver = base64Captcha.NewDriverChinese(c.ImgHeight, c.ImgWidth, 10, 4, c.KeyLong,
			zhDict, bgColor, nil,
			[]string{"wqy-microhei.ttc"})
	default:
		driver = base64Captcha.NewDriverDigit(c.ImgHeight, c.ImgWidth,
			c.KeyLong, 0.7, 80)
	}

	return driver
}

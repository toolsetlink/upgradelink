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

package i18n

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"upgradelink-api/server/api/internal/common/utils/parse"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"
)

//go:embed locale/*.json
var LocaleFS embed.FS

// Translator is a struct storing translating data.
type Translator struct {
	bundle       *i18n.Bundle
	localizer    map[language.Tag]*i18n.Localizer
	supportLangs []language.Tag
}

// AddBundleFromEmbeddedFS adds new bundle into translator from embedded file system
func (l *Translator) AddBundleFromEmbeddedFS(file embed.FS, path string) error {
	if _, err := l.bundle.LoadMessageFileFS(file, path); err != nil {
		return err
	}
	return nil
}

// AddBundleFromFile adds new bundle into translator from file path.
func (l *Translator) AddBundleFromFile(path string) error {
	if _, err := l.bundle.LoadMessageFile(path); err != nil {
		return err
	}
	return nil
}

// AddLanguageSupport adds supports for new language
func (l *Translator) AddLanguageSupport(lang language.Tag) {
	l.supportLangs = append(l.supportLangs, lang)
	// Store the original tag
	l.localizer[lang] = i18n.NewLocalizer(l.bundle, lang.String())
	// Also store with just the language code (e.g., "ja" instead of "ja-JP")
	langStr := lang.String()
	// Extract base language code (first part before '-')
	baseLangStr := langStr
	if idx := strings.Index(langStr, "-"); idx > 0 {
		baseLangStr = langStr[:idx]
	}
	// If base language is different from full language, store it
	if baseLangStr != langStr {
		baseLangTag := language.Make(baseLangStr)
		l.localizer[baseLangTag] = i18n.NewLocalizer(l.bundle, baseLangStr)
	}
}

// Trans used to translate any i18n string.
func (l *Translator) Trans(ctx context.Context, msgId string) string {

	lang := language.Chinese.String()
	clang := ctx.Value("lang")
	if clang != nil {
		lang = clang.(string)
	}

	message, err := l.MatchLocalizer(lang).LocalizeMessage(&i18n.Message{ID: msgId})
	if err != nil {
		return msgId
	}

	if message == "" {
		return msgId
	}

	return message
}

// MatchLocalizer used to matcher the localizer in map
func (l *Translator) MatchLocalizer(lang string) *i18n.Localizer {
	// First try to parse Accept-Language header
	tags := parse.ParseTags(lang)
	for _, v := range tags {
		if val, ok := l.localizer[v]; ok {
			return val
		}
	}

	// If not found, try direct language tag matching
	if lang != "" {
		tag := language.Make(lang)
		if val, ok := l.localizer[tag]; ok {
			return val
		}
	}

	return l.localizer[language.Chinese]
}

// NewTranslator returns a translator by embedded FS.
// It loads translation files from embedded file system.
// e.g. trans = i18n.NewTranslator(i18n2.LocaleFS)
func NewTranslator(efs embed.FS) *Translator {
	trans := &Translator{}
	trans.localizer = make(map[language.Tag]*i18n.Localizer)
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	trans.bundle = bundle

	var files []string
	if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d == nil {
			logx.Must(fmt.Errorf("wrong directory path"))
		}
		if !d.IsDir() {
			files = append(files, path)
		}

		return err
	}); err != nil {
		logx.Must(fmt.Errorf("failed to get any files from embedded FS, error: %v", err))
	}

	for _, v := range files {
		languageName := strings.TrimSuffix(filepath.Base(v), ".json")
		// Use language.Make to parse single language code directly
		langTag := language.Make(languageName)
		trans.AddLanguageSupport(langTag)
		err := trans.AddBundleFromEmbeddedFS(efs, v)
		if err != nil {
			logx.Must(fmt.Errorf("failed to load files from %s for i18n, error: %s", v, err.Error()))
		}
	}

	return trans
}

package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./template/*.tmpl" ./schema --feature sql/execquery,intercept,sql/modifier

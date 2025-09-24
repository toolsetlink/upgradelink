// Package http_handlers
//
//	@Author: Shuangpeng.Guo
//	@Date: 2023/8/3
//	@Desc:
//	@Version: V1.0
//	@Resume:
package http_handlers

import (
	"testing"
)

func TestErrMsg2Code(t *testing.T) {
	//type args struct {
	//	m string
	//}
	//tests := []struct {
	//	name         string
	//	args         args
	//	wantHttpCode int
	//	wantCode     int
	//	wantMsg      string
	//}{
	//	{
	//		name: "test leErrMsg",
	//		args: args{
	//			NewLinkErr(context.Background(), ErrParamMiss, "field \"Signature\" is not set").Error(),
	//		},
	//		wantHttpCode: ErrParamMiss / 1000,
	//		wantCode:     ErrParamMiss,
	//		wantMsg:      "field \"Signature\" is not set",
	//	},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		gotHttpCode, gotCode, gotMsg, _ := ErrMsg2Code(tt.args.m)
	//		if gotHttpCode != tt.wantHttpCode {
	//			t.Errorf("ErrMsg2Code() gotHttpCode = %v, wantHttpCode %v", gotHttpCode, tt.wantHttpCode)
	//		}
	//		if gotCode != tt.wantCode {
	//			t.Errorf("ErrMsg2Code() gotCode = %v, wantCode %v", gotCode, tt.wantCode)
	//		}
	//		if gotMsg != tt.wantMsg {
	//			t.Errorf("ErrMsg2Code() gotMsg = %v, wantMsg %v", gotMsg, tt.wantMsg)
	//		}
	//	})
	//}
}

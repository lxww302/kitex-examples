// Code generated by Kitex v0.0.8. DO NOT EDIT.

package bizservice

import (
	"context"
	"github.com/cloudwego/kitex-examples/generic/kitex_gen/http"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return bizServiceServiceInfo
}

var bizServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "BizService"
	handlerType := (*http.BizService)(nil)
	methods := map[string]kitex.MethodInfo{
		"BizMethod1": kitex.NewMethodInfo(bizMethod1Handler, newBizServiceBizMethod1Args, newBizServiceBizMethod1Result, false),
		"BizMethod2": kitex.NewMethodInfo(bizMethod2Handler, newBizServiceBizMethod2Args, newBizServiceBizMethod2Result, false),
		"BizMethod3": kitex.NewMethodInfo(bizMethod3Handler, newBizServiceBizMethod3Args, newBizServiceBizMethod3Result, false),
	}
	extra := map[string]interface{}{
		"PackageName": "http",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.0.8",
		Extra:           extra,
	}
	return svcInfo
}

func bizMethod1Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*http.BizServiceBizMethod1Args)
	realResult := result.(*http.BizServiceBizMethod1Result)
	success, err := handler.(http.BizService).BizMethod1(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBizServiceBizMethod1Args() interface{} {
	return http.NewBizServiceBizMethod1Args()
}

func newBizServiceBizMethod1Result() interface{} {
	return http.NewBizServiceBizMethod1Result()
}

func bizMethod2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*http.BizServiceBizMethod2Args)
	realResult := result.(*http.BizServiceBizMethod2Result)
	success, err := handler.(http.BizService).BizMethod2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBizServiceBizMethod2Args() interface{} {
	return http.NewBizServiceBizMethod2Args()
}

func newBizServiceBizMethod2Result() interface{} {
	return http.NewBizServiceBizMethod2Result()
}

func bizMethod3Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*http.BizServiceBizMethod3Args)
	realResult := result.(*http.BizServiceBizMethod3Result)
	success, err := handler.(http.BizService).BizMethod3(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBizServiceBizMethod3Args() interface{} {
	return http.NewBizServiceBizMethod3Args()
}

func newBizServiceBizMethod3Result() interface{} {
	return http.NewBizServiceBizMethod3Result()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) BizMethod1(ctx context.Context, req *http.BizRequest) (r *http.BizResponse, err error) {
	var _args http.BizServiceBizMethod1Args
	_args.Req = req
	var _result http.BizServiceBizMethod1Result
	if err = p.c.Call(ctx, "BizMethod1", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BizMethod2(ctx context.Context, req *http.BizRequest) (r *http.BizResponse, err error) {
	var _args http.BizServiceBizMethod2Args
	_args.Req = req
	var _result http.BizServiceBizMethod2Result
	if err = p.c.Call(ctx, "BizMethod2", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BizMethod3(ctx context.Context, req *http.BizRequest) (r *http.BizResponse, err error) {
	var _args http.BizServiceBizMethod3Args
	_args.Req = req
	var _result http.BizServiceBizMethod3Result
	if err = p.c.Call(ctx, "BizMethod3", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

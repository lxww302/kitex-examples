// Code generated by thriftgo (0.3.13). DO NOT EDIT.

package account

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type AccountService interface {
	Deduct(ctx context.Context, userId string, money int32) (err error)
}

type AccountServiceClient struct {
	c thrift.TClient
}

func NewAccountServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AccountServiceClient {
	return &AccountServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewAccountServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AccountServiceClient {
	return &AccountServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewAccountServiceClient(c thrift.TClient) *AccountServiceClient {
	return &AccountServiceClient{
		c: c,
	}
}

func (p *AccountServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *AccountServiceClient) Deduct(ctx context.Context, userId string, money int32) (err error) {
	var _args AccountServiceDeductArgs
	_args.UserId = userId
	_args.Money = money
	var _result AccountServiceDeductResult
	if err = p.Client_().Call(ctx, "Deduct", &_args, &_result); err != nil {
		return
	}
	return nil
}

type AccountServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      AccountService
}

func (p *AccountServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *AccountServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *AccountServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewAccountServiceProcessor(handler AccountService) *AccountServiceProcessor {
	self := &AccountServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("Deduct", &accountServiceProcessorDeduct{handler: handler})
	return self
}
func (p *AccountServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type accountServiceProcessorDeduct struct {
	handler AccountService
}

func (p *accountServiceProcessorDeduct) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AccountServiceDeductArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Deduct", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := AccountServiceDeductResult{}
	if err2 = p.handler.Deduct(ctx, args.UserId, args.Money); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Deduct: "+err2.Error())
		oprot.WriteMessageBegin("Deduct", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("Deduct", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type AccountServiceDeductArgs struct {
	UserId string `thrift:"userId,1" frugal:"1,default,string" json:"userId"`
	Money  int32  `thrift:"money,2" frugal:"2,default,i32" json:"money"`
}

func NewAccountServiceDeductArgs() *AccountServiceDeductArgs {
	return &AccountServiceDeductArgs{}
}

func (p *AccountServiceDeductArgs) InitDefault() {
}

func (p *AccountServiceDeductArgs) GetUserId() (v string) {
	return p.UserId
}

func (p *AccountServiceDeductArgs) GetMoney() (v int32) {
	return p.Money
}
func (p *AccountServiceDeductArgs) SetUserId(val string) {
	p.UserId = val
}
func (p *AccountServiceDeductArgs) SetMoney(val int32) {
	p.Money = val
}

var fieldIDToName_AccountServiceDeductArgs = map[int16]string{
	1: "userId",
	2: "money",
}

func (p *AccountServiceDeductArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AccountServiceDeductArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *AccountServiceDeductArgs) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.UserId = _field
	return nil
}
func (p *AccountServiceDeductArgs) ReadField2(iprot thrift.TProtocol) error {

	var _field int32
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Money = _field
	return nil
}

func (p *AccountServiceDeductArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Deduct_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *AccountServiceDeductArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("userId", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.UserId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *AccountServiceDeductArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("money", thrift.I32, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Money); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *AccountServiceDeductArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountServiceDeductArgs(%+v)", *p)

}

func (p *AccountServiceDeductArgs) DeepEqual(ano *AccountServiceDeductArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.UserId) {
		return false
	}
	if !p.Field2DeepEqual(ano.Money) {
		return false
	}
	return true
}

func (p *AccountServiceDeductArgs) Field1DeepEqual(src string) bool {

	if strings.Compare(p.UserId, src) != 0 {
		return false
	}
	return true
}
func (p *AccountServiceDeductArgs) Field2DeepEqual(src int32) bool {

	if p.Money != src {
		return false
	}
	return true
}

type AccountServiceDeductResult struct {
}

func NewAccountServiceDeductResult() *AccountServiceDeductResult {
	return &AccountServiceDeductResult{}
}

func (p *AccountServiceDeductResult) InitDefault() {
}

var fieldIDToName_AccountServiceDeductResult = map[int16]string{}

func (p *AccountServiceDeductResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *AccountServiceDeductResult) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("Deduct_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *AccountServiceDeductResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountServiceDeductResult(%+v)", *p)

}

func (p *AccountServiceDeductResult) DeepEqual(ano *AccountServiceDeductResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	return true
}

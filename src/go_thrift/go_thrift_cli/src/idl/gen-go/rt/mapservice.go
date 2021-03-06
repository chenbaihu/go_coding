// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package rt

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type MapService interface {
	// Parameters:
	//  - Req
	Compute(req *ComputeReq) (r *ComputeResp, err error)
}

type MapServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewMapServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *MapServiceClient {
	return &MapServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewMapServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *MapServiceClient {
	return &MapServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Req
func (p *MapServiceClient) Compute(req *ComputeReq) (r *ComputeResp, err error) {
	if err = p.sendCompute(req); err != nil {
		return
	}
	return p.recvCompute()
}

func (p *MapServiceClient) sendCompute(req *ComputeReq) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("compute", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ComputeArgs{
		Req: req,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *MapServiceClient) recvCompute() (value *ComputeResp, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "compute failed: out of sequence response")
		return
	}
	result := ComputeResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Se != nil {
		err = result.Se
		return
	}
	value = result.GetSuccess()
	return
}

type MapServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      MapService
}

func (p *MapServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *MapServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *MapServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewMapServiceProcessor(handler MapService) *MapServiceProcessor {

	self4 := &MapServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self4.processorMap["compute"] = &mapServiceProcessorCompute{handler: handler}
	return self4
}

func (p *MapServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x5

}

type mapServiceProcessorCompute struct {
	handler MapService
}

func (p *mapServiceProcessorCompute) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ComputeArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("compute", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ComputeResult{}
	var retval *ComputeResp
	var err2 error
	if retval, err2 = p.handler.Compute(args.Req); err2 != nil {
		switch v := err2.(type) {
		case *ServiceException:
			result.Se = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing compute: "+err2.Error())
			oprot.WriteMessageBegin("compute", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("compute", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type ComputeArgs struct {
	Req *ComputeReq `thrift:"req,1" json:"req"`
}

func NewComputeArgs() *ComputeArgs {
	return &ComputeArgs{}
}

var ComputeArgs_Req_DEFAULT *ComputeReq

func (p *ComputeArgs) GetReq() *ComputeReq {
	if !p.IsSetReq() {
		return ComputeArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *ComputeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ComputeArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ComputeArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = &ComputeReq{}
	if err := p.Req.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Req, err)
	}
	return nil
}

func (p *ComputeArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("compute_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *ComputeArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:req: %s", p, err)
	}
	if err := p.Req.Write(oprot); err != nil {
		return fmt.Errorf("%T error writing struct: %s", p.Req, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:req: %s", p, err)
	}
	return err
}

func (p *ComputeArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ComputeArgs(%+v)", *p)
}

type ComputeResult struct {
	Success *ComputeResp      `thrift:"success,0" json:"success"`
	Se      *ServiceException `thrift:"se,1" json:"se"`
}

func NewComputeResult() *ComputeResult {
	return &ComputeResult{}
}

var ComputeResult_Success_DEFAULT *ComputeResp

func (p *ComputeResult) GetSuccess() *ComputeResp {
	if !p.IsSetSuccess() {
		return ComputeResult_Success_DEFAULT
	}
	return p.Success
}

var ComputeResult_Se_DEFAULT *ServiceException

func (p *ComputeResult) GetSe() *ServiceException {
	if !p.IsSetSe() {
		return ComputeResult_Se_DEFAULT
	}
	return p.Se
}
func (p *ComputeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ComputeResult) IsSetSe() bool {
	return p.Se != nil
}

func (p *ComputeResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ComputeResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &ComputeResp{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *ComputeResult) ReadField1(iprot thrift.TProtocol) error {
	p.Se = &ServiceException{}
	if err := p.Se.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Se, err)
	}
	return nil
}

func (p *ComputeResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("compute_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *ComputeResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *ComputeResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetSe() {
		if err := oprot.WriteFieldBegin("se", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:se: %s", p, err)
		}
		if err := p.Se.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Se, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:se: %s", p, err)
		}
	}
	return err
}

func (p *ComputeResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ComputeResult(%+v)", *p)
}

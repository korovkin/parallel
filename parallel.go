// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package parallel

import(
	"bytes"
	"context"
	"reflect"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type Result_ int64
const (
  Result__OK Result_ = 0
  Result__ERROR Result_ = 2
)

func (p Result_) String() string {
  switch p {
  case Result__OK: return "OK"
  case Result__ERROR: return "ERROR"
  }
  return "<UNSET>"
}

func Result_FromString(s string) (Result_, error) {
  switch s {
  case "OK": return Result__OK, nil 
  case "ERROR": return Result__ERROR, nil 
  }
  return Result_(0), fmt.Errorf("not a valid Result_ string")
}


func Result_Ptr(v Result_) *Result_ { return &v }

func (p Result_) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *Result_) UnmarshalText(text []byte) error {
q, err := Result_FromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *Result_) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = Result_(v)
return nil
}

func (p * Result_) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
// Attributes:
//  - CmdLine
//  - Ticket
type Cmd struct {
  CmdLine string `thrift:"cmdLine,1" db:"cmdLine" json:"cmdLine"`
  Ticket int64 `thrift:"ticket,2" db:"ticket" json:"ticket"`
}

func NewCmd() *Cmd {
  return &Cmd{}
}


func (p *Cmd) GetCmdLine() string {
  return p.CmdLine
}

func (p *Cmd) GetTicket() int64 {
  return p.Ticket
}
func (p *Cmd) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Cmd)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.CmdLine = v
}
  return nil
}

func (p *Cmd)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Ticket = v
}
  return nil
}

func (p *Cmd) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Cmd"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Cmd) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("cmdLine", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:cmdLine: ", p), err) }
  if err := oprot.WriteString(string(p.CmdLine)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.cmdLine (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:cmdLine: ", p), err) }
  return err
}

func (p *Cmd) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("ticket", thrift.I64, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:ticket: ", p), err) }
  if err := oprot.WriteI64(int64(p.Ticket)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.ticket (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:ticket: ", p), err) }
  return err
}

func (p *Cmd) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Cmd(%+v)", *p)
}

// Attributes:
//  - Stdout
//  - Stderr
//  - Tags
type Output struct {
  Stdout string `thrift:"stdout,1" db:"stdout" json:"stdout"`
  Stderr string `thrift:"stderr,2" db:"stderr" json:"stderr"`
  Tags map[string]string `thrift:"tags,3" db:"tags" json:"tags"`
}

func NewOutput() *Output {
  return &Output{}
}


func (p *Output) GetStdout() string {
  return p.Stdout
}

func (p *Output) GetStderr() string {
  return p.Stderr
}

func (p *Output) GetTags() map[string]string {
  return p.Tags
}
func (p *Output) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.MAP {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Output)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Stdout = v
}
  return nil
}

func (p *Output)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Stderr = v
}
  return nil
}

func (p *Output)  ReadField3(iprot thrift.TProtocol) error {
  _, _, size, err := iprot.ReadMapBegin()
  if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
  }
  tMap := make(map[string]string, size)
  p.Tags =  tMap
  for i := 0; i < size; i ++ {
var _key0 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _key0 = v
}
var _val1 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _val1 = v
}
    p.Tags[_key0] = _val1
  }
  if err := iprot.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
  }
  return nil
}

func (p *Output) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Output"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Output) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("stdout", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:stdout: ", p), err) }
  if err := oprot.WriteString(string(p.Stdout)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.stdout (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:stdout: ", p), err) }
  return err
}

func (p *Output) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("stderr", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:stderr: ", p), err) }
  if err := oprot.WriteString(string(p.Stderr)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.stderr (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:stderr: ", p), err) }
  return err
}

func (p *Output) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("tags", thrift.MAP, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:tags: ", p), err) }
  if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Tags)); err != nil {
    return thrift.PrependError("error writing map begin: ", err)
  }
  for k, v := range p.Tags {
    if err := oprot.WriteString(string(k)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    if err := oprot.WriteString(string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteMapEnd(); err != nil {
    return thrift.PrependError("error writing map end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:tags: ", p), err) }
  return err
}

func (p *Output) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Output(%+v)", *p)
}

// Attributes:
//  - What
//  - Output
type ExecuteException struct {
  What string `thrift:"what,1" db:"what" json:"what"`
  Output *Output `thrift:"output,2" db:"output" json:"output"`
}

func NewExecuteException() *ExecuteException {
  return &ExecuteException{}
}


func (p *ExecuteException) GetWhat() string {
  return p.What
}
var ExecuteException_Output_DEFAULT *Output
func (p *ExecuteException) GetOutput() *Output {
  if !p.IsSetOutput() {
    return ExecuteException_Output_DEFAULT
  }
return p.Output
}
func (p *ExecuteException) IsSetOutput() bool {
  return p.Output != nil
}

func (p *ExecuteException) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ExecuteException)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.What = v
}
  return nil
}

func (p *ExecuteException)  ReadField2(iprot thrift.TProtocol) error {
  p.Output = &Output{}
  if err := p.Output.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Output), err)
  }
  return nil
}

func (p *ExecuteException) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("ExecuteException"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ExecuteException) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("what", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:what: ", p), err) }
  if err := oprot.WriteString(string(p.What)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.what (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:what: ", p), err) }
  return err
}

func (p *ExecuteException) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("output", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:output: ", p), err) }
  if err := p.Output.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Output), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:output: ", p), err) }
  return err
}

func (p *ExecuteException) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ExecuteException(%+v)", *p)
}

func (p *ExecuteException) Error() string {
  return p.String()
}

type Parallel interface {
  Ping(ctx context.Context) (r string, err error)
  // Parameters:
  //  - Command
  Execute(ctx context.Context, command *Cmd) (r *Output, err error)
}

type ParallelClient struct {
  c thrift.TClient
}

func NewParallelClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ParallelClient {
  return &ParallelClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewParallelClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ParallelClient {
  return &ParallelClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewParallelClient(c thrift.TClient) *ParallelClient {
  return &ParallelClient{
    c: c,
  }
}

func (p *ParallelClient) Client_() thrift.TClient {
  return p.c
}
func (p *ParallelClient) Ping(ctx context.Context) (r string, err error) {
  var _args2 ParallelPingArgs
  var _result3 ParallelPingResult
  if err = p.Client_().Call(ctx, "Ping", &_args2, &_result3); err != nil {
    return
  }
  return _result3.GetSuccess(), nil
}

// Parameters:
//  - Command
func (p *ParallelClient) Execute(ctx context.Context, command *Cmd) (r *Output, err error) {
  var _args4 ParallelExecuteArgs
  _args4.Command = command
  var _result5 ParallelExecuteResult
  if err = p.Client_().Call(ctx, "Execute", &_args4, &_result5); err != nil {
    return
  }
  switch {
  case _result5.E!= nil:
    return r, _result5.E
  }

  return _result5.GetSuccess(), nil
}

type ParallelProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Parallel
}

func (p *ParallelProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *ParallelProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *ParallelProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewParallelProcessor(handler Parallel) *ParallelProcessor {

  self6 := &ParallelProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self6.processorMap["Ping"] = &parallelProcessorPing{handler:handler}
  self6.processorMap["Execute"] = &parallelProcessorExecute{handler:handler}
return self6
}

func (p *ParallelProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x7.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x7

}

type parallelProcessorPing struct {
  handler Parallel
}

func (p *parallelProcessorPing) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := ParallelPingArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Ping", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := ParallelPingResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.Ping(ctx); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Ping: " + err2.Error())
    oprot.WriteMessageBegin("Ping", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("Ping", thrift.REPLY, seqId); err2 != nil {
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

type parallelProcessorExecute struct {
  handler Parallel
}

func (p *parallelProcessorExecute) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := ParallelExecuteArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Execute", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := ParallelExecuteResult{}
var retval *Output
  var err2 error
  if retval, err2 = p.handler.Execute(ctx, args.Command); err2 != nil {
  switch v := err2.(type) {
    case *ExecuteException:
  result.E = v
    default:
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Execute: " + err2.Error())
    oprot.WriteMessageBegin("Execute", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  }
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("Execute", thrift.REPLY, seqId); err2 != nil {
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


// HELPER FUNCTIONS AND STRUCTURES

type ParallelPingArgs struct {
}

func NewParallelPingArgs() *ParallelPingArgs {
  return &ParallelPingArgs{}
}

func (p *ParallelPingArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ParallelPingArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Ping_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ParallelPingArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ParallelPingArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ParallelPingResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewParallelPingResult() *ParallelPingResult {
  return &ParallelPingResult{}
}

var ParallelPingResult_Success_DEFAULT string
func (p *ParallelPingResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return ParallelPingResult_Success_DEFAULT
  }
return *p.Success
}
func (p *ParallelPingResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *ParallelPingResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ParallelPingResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *ParallelPingResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Ping_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ParallelPingResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *ParallelPingResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ParallelPingResult(%+v)", *p)
}

// Attributes:
//  - Command
type ParallelExecuteArgs struct {
  Command *Cmd `thrift:"command,1" db:"command" json:"command"`
}

func NewParallelExecuteArgs() *ParallelExecuteArgs {
  return &ParallelExecuteArgs{}
}

var ParallelExecuteArgs_Command_DEFAULT *Cmd
func (p *ParallelExecuteArgs) GetCommand() *Cmd {
  if !p.IsSetCommand() {
    return ParallelExecuteArgs_Command_DEFAULT
  }
return p.Command
}
func (p *ParallelExecuteArgs) IsSetCommand() bool {
  return p.Command != nil
}

func (p *ParallelExecuteArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ParallelExecuteArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.Command = &Cmd{}
  if err := p.Command.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Command), err)
  }
  return nil
}

func (p *ParallelExecuteArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Execute_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ParallelExecuteArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("command", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:command: ", p), err) }
  if err := p.Command.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Command), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:command: ", p), err) }
  return err
}

func (p *ParallelExecuteArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ParallelExecuteArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type ParallelExecuteResult struct {
  Success *Output `thrift:"success,0" db:"success" json:"success,omitempty"`
  E *ExecuteException `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewParallelExecuteResult() *ParallelExecuteResult {
  return &ParallelExecuteResult{}
}

var ParallelExecuteResult_Success_DEFAULT *Output
func (p *ParallelExecuteResult) GetSuccess() *Output {
  if !p.IsSetSuccess() {
    return ParallelExecuteResult_Success_DEFAULT
  }
return p.Success
}
var ParallelExecuteResult_E_DEFAULT *ExecuteException
func (p *ParallelExecuteResult) GetE() *ExecuteException {
  if !p.IsSetE() {
    return ParallelExecuteResult_E_DEFAULT
  }
return p.E
}
func (p *ParallelExecuteResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *ParallelExecuteResult) IsSetE() bool {
  return p.E != nil
}

func (p *ParallelExecuteResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ParallelExecuteResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &Output{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *ParallelExecuteResult)  ReadField1(iprot thrift.TProtocol) error {
  p.E = &ExecuteException{}
  if err := p.E.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
  }
  return nil
}

func (p *ParallelExecuteResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Execute_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ParallelExecuteResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *ParallelExecuteResult) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetE() {
    if err := oprot.WriteFieldBegin("e", thrift.STRUCT, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err) }
    if err := p.E.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err) }
  }
  return err
}

func (p *ParallelExecuteResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ParallelExecuteResult(%+v)", *p)
}


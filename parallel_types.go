// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package parallel

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Result_ int64

const (
	Result__OK    Result_ = 0
	Result__ERROR Result_ = 2
)

func (p Result_) String() string {
	switch p {
	case Result__OK:
		return "OK"
	case Result__ERROR:
		return "ERROR"
	}
	return "<UNSET>"
}

func Result_FromString(s string) (Result_, error) {
	switch s {
	case "OK":
		return Result__OK, nil
	case "ERROR":
		return Result__ERROR, nil
	}
	return Result_(0), fmt.Errorf("not a valid Result_ string")
}

func Result_Ptr(v Result_) *Result_ { return &v }

func (p Result_) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *Result_) UnmarshalText(text []byte) error {
	q, err := Result_FromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//  - CmdLine
//  - Ticket
type Cmd struct {
	CmdLine string `thrift:"cmdLine,1" json:"cmdLine"`
	Ticket  int64  `thrift:"ticket,2" json:"ticket"`
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
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
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
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Cmd) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.CmdLine = v
	}
	return nil
}

func (p *Cmd) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Ticket = v
	}
	return nil
}

func (p *Cmd) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Cmd"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Cmd) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("cmdLine", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:cmdLine: ", p), err)
	}
	if err := oprot.WriteString(string(p.CmdLine)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.cmdLine (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:cmdLine: ", p), err)
	}
	return err
}

func (p *Cmd) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ticket", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:ticket: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Ticket)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ticket (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:ticket: ", p), err)
	}
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
	Stdout string            `thrift:"stdout,1" json:"stdout"`
	Stderr string            `thrift:"stderr,2" json:"stderr"`
	Tags   map[string]string `thrift:"tags,3" json:"tags"`
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
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
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
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Output) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Stdout = v
	}
	return nil
}

func (p *Output) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Stderr = v
	}
	return nil
}

func (p *Output) readField3(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.Tags = tMap
	for i := 0; i < size; i++ {
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
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Output) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("stdout", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:stdout: ", p), err)
	}
	if err := oprot.WriteString(string(p.Stdout)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.stdout (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:stdout: ", p), err)
	}
	return err
}

func (p *Output) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("stderr", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:stderr: ", p), err)
	}
	if err := oprot.WriteString(string(p.Stderr)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.stderr (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:stderr: ", p), err)
	}
	return err
}

func (p *Output) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tags", thrift.MAP, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:tags: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Tags)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Tags {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:tags: ", p), err)
	}
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
	What   string  `thrift:"what,1" json:"what"`
	Output *Output `thrift:"output,2" json:"output"`
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
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
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
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ExecuteException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.What = v
	}
	return nil
}

func (p *ExecuteException) readField2(iprot thrift.TProtocol) error {
	p.Output = &Output{}
	if err := p.Output.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Output), err)
	}
	return nil
}

func (p *ExecuteException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ExecuteException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ExecuteException) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("what", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:what: ", p), err)
	}
	if err := oprot.WriteString(string(p.What)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.what (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:what: ", p), err)
	}
	return err
}

func (p *ExecuteException) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("output", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:output: ", p), err)
	}
	if err := p.Output.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Output), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:output: ", p), err)
	}
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

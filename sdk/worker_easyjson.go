// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package sdk

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson82a45abeDecodeGithubComOvhCdsSdk(in *jlexer.Lexer, out *WorkerTakeForm) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "BookedJobID":
			out.BookedJobID = int64(in.Int64())
		case "Time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Time).UnmarshalJSON(data))
			}
		case "OS":
			out.OS = string(in.String())
		case "Arch":
			out.Arch = string(in.String())
		case "Version":
			out.Version = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a45abeEncodeGithubComOvhCdsSdk(out *jwriter.Writer, in WorkerTakeForm) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"BookedJobID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.BookedJobID))
	}
	{
		const prefix string = ",\"Time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Time).MarshalJSON())
	}
	{
		const prefix string = ",\"OS\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OS))
	}
	{
		const prefix string = ",\"Arch\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Arch))
	}
	{
		const prefix string = ",\"Version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Version))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WorkerTakeForm) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82a45abeEncodeGithubComOvhCdsSdk(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WorkerTakeForm) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82a45abeEncodeGithubComOvhCdsSdk(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WorkerTakeForm) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82a45abeDecodeGithubComOvhCdsSdk(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WorkerTakeForm) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82a45abeDecodeGithubComOvhCdsSdk(l, v)
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk1(in *jlexer.Lexer, out *WorkerRegistrationForm) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Name":
			out.Name = string(in.String())
		case "Token":
			out.Token = string(in.String())
		case "ModelID":
			out.ModelID = int64(in.Int64())
		case "Hatchery":
			out.Hatchery = int64(in.Int64())
		case "HatcheryName":
			out.HatcheryName = string(in.String())
		case "BinaryCapabilities":
			if in.IsNull() {
				in.Skip()
				out.BinaryCapabilities = nil
			} else {
				in.Delim('[')
				if out.BinaryCapabilities == nil {
					if !in.IsDelim(']') {
						out.BinaryCapabilities = make([]string, 0, 4)
					} else {
						out.BinaryCapabilities = []string{}
					}
				} else {
					out.BinaryCapabilities = (out.BinaryCapabilities)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.BinaryCapabilities = append(out.BinaryCapabilities, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Version":
			out.Version = string(in.String())
		case "OS":
			out.OS = string(in.String())
		case "Arch":
			out.Arch = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a45abeEncodeGithubComOvhCdsSdk1(out *jwriter.Writer, in WorkerRegistrationForm) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Token\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Token))
	}
	{
		const prefix string = ",\"ModelID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ModelID))
	}
	{
		const prefix string = ",\"Hatchery\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Hatchery))
	}
	{
		const prefix string = ",\"HatcheryName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.HatcheryName))
	}
	{
		const prefix string = ",\"BinaryCapabilities\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.BinaryCapabilities == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.BinaryCapabilities {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"OS\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OS))
	}
	{
		const prefix string = ",\"Arch\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Arch))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WorkerRegistrationForm) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82a45abeEncodeGithubComOvhCdsSdk1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WorkerRegistrationForm) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82a45abeEncodeGithubComOvhCdsSdk1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WorkerRegistrationForm) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82a45abeDecodeGithubComOvhCdsSdk1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WorkerRegistrationForm) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82a45abeDecodeGithubComOvhCdsSdk1(l, v)
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk2(in *jlexer.Lexer, out *Worker) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "group_id":
			out.GroupID = int64(in.Int64())
		case "model_id":
			out.ModelID = int64(in.Int64())
		case "action_build_id":
			out.ActionBuildID = int64(in.Int64())
		case "model":
			if in.IsNull() {
				in.Skip()
				out.Model = nil
			} else {
				if out.Model == nil {
					out.Model = new(Model)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Model).UnmarshalJSON(data))
				}
			}
		case "hatchery_id":
			out.HatcheryID = int64(in.Int64())
		case "hatchery_name":
			out.HatcheryName = string(in.String())
		case "job_type":
			out.JobType = string(in.String())
		case "status":
			out.Status = Status(in.String())
		case "up_to_date":
			out.Uptodate = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a45abeEncodeGithubComOvhCdsSdk2(out *jwriter.Writer, in Worker) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"group_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.GroupID))
	}
	{
		const prefix string = ",\"model_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ModelID))
	}
	{
		const prefix string = ",\"action_build_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ActionBuildID))
	}
	{
		const prefix string = ",\"model\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Model == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Model).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"hatchery_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.HatcheryID))
	}
	{
		const prefix string = ",\"hatchery_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.HatcheryName))
	}
	{
		const prefix string = ",\"job_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.JobType))
	}
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"up_to_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Uptodate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Worker) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82a45abeEncodeGithubComOvhCdsSdk2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Worker) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82a45abeEncodeGithubComOvhCdsSdk2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Worker) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82a45abeDecodeGithubComOvhCdsSdk2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Worker) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82a45abeDecodeGithubComOvhCdsSdk2(l, v)
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk3(in *jlexer.Lexer, out *SpawnErrorForm) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Error":
			out.Error = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a45abeEncodeGithubComOvhCdsSdk3(out *jwriter.Writer, in SpawnErrorForm) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Error))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SpawnErrorForm) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82a45abeEncodeGithubComOvhCdsSdk3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SpawnErrorForm) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82a45abeEncodeGithubComOvhCdsSdk3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SpawnErrorForm) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82a45abeDecodeGithubComOvhCdsSdk3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SpawnErrorForm) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82a45abeDecodeGithubComOvhCdsSdk3(l, v)
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk4(in *jlexer.Lexer, out *ModelVirtualMachine) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "image":
			out.Image = string(in.String())
		case "flavor":
			out.Flavor = string(in.String())
		case "user_data":
			out.UserData = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a45abeEncodeGithubComOvhCdsSdk4(out *jwriter.Writer, in ModelVirtualMachine) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"image\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Image))
	}
	if in.Flavor != "" {
		const prefix string = ",\"flavor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Flavor))
	}
	{
		const prefix string = ",\"user_data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserData))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ModelVirtualMachine) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82a45abeEncodeGithubComOvhCdsSdk4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ModelVirtualMachine) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82a45abeEncodeGithubComOvhCdsSdk4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ModelVirtualMachine) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82a45abeDecodeGithubComOvhCdsSdk4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ModelVirtualMachine) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82a45abeDecodeGithubComOvhCdsSdk4(l, v)
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk5(in *jlexer.Lexer, out *ModelDocker) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "image":
			out.Image = string(in.String())
		case "memory":
			out.Memory = int64(in.Int64())
		case "cmd":
			out.Cmd = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a45abeEncodeGithubComOvhCdsSdk5(out *jwriter.Writer, in ModelDocker) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"image\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"memory\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Memory))
	}
	{
		const prefix string = ",\"cmd\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Cmd))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ModelDocker) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82a45abeEncodeGithubComOvhCdsSdk5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ModelDocker) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82a45abeEncodeGithubComOvhCdsSdk5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ModelDocker) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82a45abeDecodeGithubComOvhCdsSdk5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ModelDocker) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82a45abeDecodeGithubComOvhCdsSdk5(l, v)
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk6(in *jlexer.Lexer, out *Model) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "model_virtual_machine":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ModelVirtualMachine).UnmarshalJSON(data))
			}
		case "model_docker":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ModelDocker).UnmarshalJSON(data))
			}
		case "communication":
			out.Communication = string(in.String())
		case "disabled":
			out.Disabled = bool(in.Bool())
		case "restricted":
			out.Restricted = bool(in.Bool())
		case "registered_capabilities":
			if in.IsNull() {
				in.Skip()
				out.RegisteredCapabilities = nil
			} else {
				in.Delim('[')
				if out.RegisteredCapabilities == nil {
					if !in.IsDelim(']') {
						out.RegisteredCapabilities = make([]Requirement, 0, 1)
					} else {
						out.RegisteredCapabilities = []Requirement{}
					}
				} else {
					out.RegisteredCapabilities = (out.RegisteredCapabilities)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Requirement
					if data := in.Raw(); in.Ok() {
						in.AddError((v4).UnmarshalJSON(data))
					}
					out.RegisteredCapabilities = append(out.RegisteredCapabilities, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "registered_os":
			out.RegisteredOS = string(in.String())
		case "registered_arch":
			out.RegisteredArch = string(in.String())
		case "need_registration":
			out.NeedRegistration = bool(in.Bool())
		case "last_registration":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastRegistration).UnmarshalJSON(data))
			}
		case "user_last_modified":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UserLastModified).UnmarshalJSON(data))
			}
		case "created_by":
			easyjson82a45abeDecodeGithubComOvhCdsSdk7(in, &out.CreatedBy)
		case "provision":
			out.Provision = int64(in.Int64())
		case "group_id":
			out.GroupID = int64(in.Int64())
		case "group":
			easyjson82a45abeDecodeGithubComOvhCdsSdk8(in, &out.Group)
		case "nb_spawn_err":
			out.NbSpawnErr = int64(in.Int64())
		case "last_spawn_err":
			out.LastSpawnErr = string(in.String())
		case "date_last_spawn_err":
			if in.IsNull() {
				in.Skip()
				out.DateLastSpawnErr = nil
			} else {
				if out.DateLastSpawnErr == nil {
					out.DateLastSpawnErr = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.DateLastSpawnErr).UnmarshalJSON(data))
				}
			}
		case "is_deprecated":
			out.IsDeprecated = bool(in.Bool())
		case "is_official":
			out.IsOfficial = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a45abeEncodeGithubComOvhCdsSdk6(out *jwriter.Writer, in Model) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"model_virtual_machine\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ModelVirtualMachine).MarshalJSON())
	}
	{
		const prefix string = ",\"model_docker\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ModelDocker).MarshalJSON())
	}
	{
		const prefix string = ",\"communication\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Communication))
	}
	{
		const prefix string = ",\"disabled\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Disabled))
	}
	{
		const prefix string = ",\"restricted\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Restricted))
	}
	{
		const prefix string = ",\"registered_capabilities\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.RegisteredCapabilities == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.RegisteredCapabilities {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Raw((v6).MarshalJSON())
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"registered_os\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RegisteredOS))
	}
	{
		const prefix string = ",\"registered_arch\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RegisteredArch))
	}
	{
		const prefix string = ",\"need_registration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.NeedRegistration))
	}
	{
		const prefix string = ",\"last_registration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LastRegistration).MarshalJSON())
	}
	{
		const prefix string = ",\"user_last_modified\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.UserLastModified).MarshalJSON())
	}
	{
		const prefix string = ",\"created_by\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson82a45abeEncodeGithubComOvhCdsSdk7(out, in.CreatedBy)
	}
	{
		const prefix string = ",\"provision\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Provision))
	}
	{
		const prefix string = ",\"group_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.GroupID))
	}
	{
		const prefix string = ",\"group\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson82a45abeEncodeGithubComOvhCdsSdk8(out, in.Group)
	}
	{
		const prefix string = ",\"nb_spawn_err\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.NbSpawnErr))
	}
	{
		const prefix string = ",\"last_spawn_err\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LastSpawnErr))
	}
	{
		const prefix string = ",\"date_last_spawn_err\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.DateLastSpawnErr == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.DateLastSpawnErr).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"is_deprecated\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsDeprecated))
	}
	{
		const prefix string = ",\"is_official\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsOfficial))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Model) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82a45abeEncodeGithubComOvhCdsSdk6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Model) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82a45abeEncodeGithubComOvhCdsSdk6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Model) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82a45abeDecodeGithubComOvhCdsSdk6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Model) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82a45abeDecodeGithubComOvhCdsSdk6(l, v)
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk8(in *jlexer.Lexer, out *Group) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "admins":
			if in.IsNull() {
				in.Skip()
				out.Admins = nil
			} else {
				in.Delim('[')
				if out.Admins == nil {
					if !in.IsDelim(']') {
						out.Admins = make([]User, 0, 1)
					} else {
						out.Admins = []User{}
					}
				} else {
					out.Admins = (out.Admins)[:0]
				}
				for !in.IsDelim(']') {
					var v7 User
					easyjson82a45abeDecodeGithubComOvhCdsSdk7(in, &v7)
					out.Admins = append(out.Admins, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]User, 0, 1)
					} else {
						out.Users = []User{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v8 User
					easyjson82a45abeDecodeGithubComOvhCdsSdk7(in, &v8)
					out.Users = append(out.Users, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tokens":
			if in.IsNull() {
				in.Skip()
				out.Tokens = nil
			} else {
				in.Delim('[')
				if out.Tokens == nil {
					if !in.IsDelim(']') {
						out.Tokens = make([]Token, 0, 1)
					} else {
						out.Tokens = []Token{}
					}
				} else {
					out.Tokens = (out.Tokens)[:0]
				}
				for !in.IsDelim(']') {
					var v9 Token
					easyjson82a45abeDecodeGithubComOvhCdsSdk9(in, &v9)
					out.Tokens = append(out.Tokens, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a45abeEncodeGithubComOvhCdsSdk8(out *jwriter.Writer, in Group) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if len(in.Admins) != 0 {
		const prefix string = ",\"admins\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v10, v11 := range in.Admins {
				if v10 > 0 {
					out.RawByte(',')
				}
				easyjson82a45abeEncodeGithubComOvhCdsSdk7(out, v11)
			}
			out.RawByte(']')
		}
	}
	if len(in.Users) != 0 {
		const prefix string = ",\"users\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v12, v13 := range in.Users {
				if v12 > 0 {
					out.RawByte(',')
				}
				easyjson82a45abeEncodeGithubComOvhCdsSdk7(out, v13)
			}
			out.RawByte(']')
		}
	}
	if len(in.Tokens) != 0 {
		const prefix string = ",\"tokens\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.Tokens {
				if v14 > 0 {
					out.RawByte(',')
				}
				easyjson82a45abeEncodeGithubComOvhCdsSdk9(out, v15)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk9(in *jlexer.Lexer, out *Token) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "group_id":
			out.GroupID = int64(in.Int64())
		case "group_name":
			out.GroupName = string(in.String())
		case "token":
			out.Token = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "creator":
			out.Creator = string(in.String())
		case "expiration":
			out.Expiration = Expiration(in.Int())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a45abeEncodeGithubComOvhCdsSdk9(out *jwriter.Writer, in Token) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"group_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.GroupID))
	}
	{
		const prefix string = ",\"group_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.GroupName))
	}
	{
		const prefix string = ",\"token\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Token))
	}
	{
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"creator\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Creator))
	}
	{
		const prefix string = ",\"expiration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Expiration))
	}
	{
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Created).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk7(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "username":
			out.Username = string(in.String())
		case "fullname":
			out.Fullname = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "admin":
			out.Admin = bool(in.Bool())
		case "groups":
			if in.IsNull() {
				in.Skip()
				out.Groups = nil
			} else {
				in.Delim('[')
				if out.Groups == nil {
					if !in.IsDelim(']') {
						out.Groups = make([]Group, 0, 1)
					} else {
						out.Groups = []Group{}
					}
				} else {
					out.Groups = (out.Groups)[:0]
				}
				for !in.IsDelim(']') {
					var v16 Group
					easyjson82a45abeDecodeGithubComOvhCdsSdk8(in, &v16)
					out.Groups = append(out.Groups, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "origin":
			out.Origin = string(in.String())
		case "favorites":
			if in.IsNull() {
				in.Skip()
				out.Favorites = nil
			} else {
				in.Delim('[')
				if out.Favorites == nil {
					if !in.IsDelim(']') {
						out.Favorites = make([]Favorite, 0, 1)
					} else {
						out.Favorites = []Favorite{}
					}
				} else {
					out.Favorites = (out.Favorites)[:0]
				}
				for !in.IsDelim(']') {
					var v17 Favorite
					easyjson82a45abeDecodeGithubComOvhCdsSdk10(in, &v17)
					out.Favorites = append(out.Favorites, v17)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "permissions":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Permissions).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a45abeEncodeGithubComOvhCdsSdk7(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"fullname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Fullname))
	}
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"admin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Admin))
	}
	if len(in.Groups) != 0 {
		const prefix string = ",\"groups\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v18, v19 := range in.Groups {
				if v18 > 0 {
					out.RawByte(',')
				}
				easyjson82a45abeEncodeGithubComOvhCdsSdk8(out, v19)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"origin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Origin))
	}
	{
		const prefix string = ",\"favorites\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Favorites == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.Favorites {
				if v20 > 0 {
					out.RawByte(',')
				}
				easyjson82a45abeEncodeGithubComOvhCdsSdk10(out, v21)
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"permissions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Permissions).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk10(in *jlexer.Lexer, out *Favorite) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "project_ids":
			if in.IsNull() {
				in.Skip()
				out.ProjectIDs = nil
			} else {
				in.Delim('[')
				if out.ProjectIDs == nil {
					if !in.IsDelim(']') {
						out.ProjectIDs = make([]int64, 0, 8)
					} else {
						out.ProjectIDs = []int64{}
					}
				} else {
					out.ProjectIDs = (out.ProjectIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v22 int64
					v22 = int64(in.Int64())
					out.ProjectIDs = append(out.ProjectIDs, v22)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "workflow_ids":
			if in.IsNull() {
				in.Skip()
				out.WorkflowIDs = nil
			} else {
				in.Delim('[')
				if out.WorkflowIDs == nil {
					if !in.IsDelim(']') {
						out.WorkflowIDs = make([]int64, 0, 8)
					} else {
						out.WorkflowIDs = []int64{}
					}
				} else {
					out.WorkflowIDs = (out.WorkflowIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v23 int64
					v23 = int64(in.Int64())
					out.WorkflowIDs = append(out.WorkflowIDs, v23)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a45abeEncodeGithubComOvhCdsSdk10(out *jwriter.Writer, in Favorite) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"project_ids\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.ProjectIDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v24, v25 := range in.ProjectIDs {
				if v24 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v25))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"workflow_ids\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.WorkflowIDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v26, v27 := range in.WorkflowIDs {
				if v26 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v27))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

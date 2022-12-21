// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonF50c8087DecodeHeadHunterInternalEntityModels(in *jlexer.Lexer, out *EducationDetail) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "resume_id":
			out.ResumeId = uint(in.Uint())
		case "certificate_degree_name":
			out.CertificateDegreeName = string(in.String())
		case "major":
			out.Major = string(in.String())
		case "university_name":
			out.UniversityName = string(in.String())
		case "starting_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartingDate).UnmarshalJSON(data))
			}
		case "completion_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CompletionDate).UnmarshalJSON(data))
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
func easyjsonF50c8087EncodeHeadHunterInternalEntityModels(out *jwriter.Writer, in EducationDetail) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"resume_id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ResumeId))
	}
	{
		const prefix string = ",\"certificate_degree_name\":"
		out.RawString(prefix)
		out.String(string(in.CertificateDegreeName))
	}
	{
		const prefix string = ",\"major\":"
		out.RawString(prefix)
		out.String(string(in.Major))
	}
	{
		const prefix string = ",\"university_name\":"
		out.RawString(prefix)
		out.String(string(in.UniversityName))
	}
	{
		const prefix string = ",\"starting_date\":"
		out.RawString(prefix)
		out.Raw((in.StartingDate).MarshalJSON())
	}
	{
		const prefix string = ",\"completion_date\":"
		out.RawString(prefix)
		out.Raw((in.CompletionDate).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EducationDetail) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF50c8087EncodeHeadHunterInternalEntityModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EducationDetail) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF50c8087EncodeHeadHunterInternalEntityModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EducationDetail) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF50c8087DecodeHeadHunterInternalEntityModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EducationDetail) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF50c8087DecodeHeadHunterInternalEntityModels(l, v)
}

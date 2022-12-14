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

func easyjsonFc5c8ee5DecodeHeadHunterInternalEntityModels(in *jlexer.Lexer, out *Skill) {
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
		case "id":
			out.ID = uint(in.Uint())
		case "skillSetName":
			out.SkillName = string(in.String())
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
func easyjsonFc5c8ee5EncodeHeadHunterInternalEntityModels(out *jwriter.Writer, in Skill) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"skillSetName\":"
		out.RawString(prefix)
		out.String(string(in.SkillName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Skill) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFc5c8ee5EncodeHeadHunterInternalEntityModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Skill) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFc5c8ee5EncodeHeadHunterInternalEntityModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Skill) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFc5c8ee5DecodeHeadHunterInternalEntityModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Skill) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFc5c8ee5DecodeHeadHunterInternalEntityModels(l, v)
}

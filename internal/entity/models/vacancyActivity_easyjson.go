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

func easyjson4038ac0aDecodeHeadHunterInternalEntityModels(in *jlexer.Lexer, out *VacancyActivityPreview) {
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
		case "user_account_id":
			out.UserAccountId = uint(in.Uint())
		case "id":
			out.ResumeId = uint(in.Uint())
		case "vacancy_id":
			out.VacancyId = uint(in.Uint())
		case "applicant_name":
			out.ApplicantName = string(in.String())
		case "applicant_surname":
			out.ApplicantSurname = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "image":
			out.Image = string(in.String())
		case "created_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ApplyDate).UnmarshalJSON(data))
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
func easyjson4038ac0aEncodeHeadHunterInternalEntityModels(out *jwriter.Writer, in VacancyActivityPreview) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_account_id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.UserAccountId))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Uint(uint(in.ResumeId))
	}
	{
		const prefix string = ",\"vacancy_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.VacancyId))
	}
	{
		const prefix string = ",\"applicant_name\":"
		out.RawString(prefix)
		out.String(string(in.ApplicantName))
	}
	{
		const prefix string = ",\"applicant_surname\":"
		out.RawString(prefix)
		out.String(string(in.ApplicantSurname))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix)
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"created_date\":"
		out.RawString(prefix)
		out.Raw((in.ApplyDate).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VacancyActivityPreview) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4038ac0aEncodeHeadHunterInternalEntityModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VacancyActivityPreview) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4038ac0aEncodeHeadHunterInternalEntityModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VacancyActivityPreview) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4038ac0aDecodeHeadHunterInternalEntityModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VacancyActivityPreview) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4038ac0aDecodeHeadHunterInternalEntityModels(l, v)
}
func easyjson4038ac0aDecodeHeadHunterInternalEntityModels1(in *jlexer.Lexer, out *VacancyActivity) {
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
		case "user_account_id":
			out.UserAccountId = uint(in.Uint())
		case "id":
			out.ResumeId = uint(in.Uint())
		case "vacancy_id":
			out.VacancyId = uint(in.Uint())
		case "created_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ApplyDate).UnmarshalJSON(data))
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
func easyjson4038ac0aEncodeHeadHunterInternalEntityModels1(out *jwriter.Writer, in VacancyActivity) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_account_id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.UserAccountId))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Uint(uint(in.ResumeId))
	}
	{
		const prefix string = ",\"vacancy_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.VacancyId))
	}
	{
		const prefix string = ",\"created_date\":"
		out.RawString(prefix)
		out.Raw((in.ApplyDate).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VacancyActivity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4038ac0aEncodeHeadHunterInternalEntityModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VacancyActivity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4038ac0aEncodeHeadHunterInternalEntityModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VacancyActivity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4038ac0aDecodeHeadHunterInternalEntityModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VacancyActivity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4038ac0aDecodeHeadHunterInternalEntityModels1(l, v)
}
func easyjson4038ac0aDecodeHeadHunterInternalEntityModels2(in *jlexer.Lexer, out *GetAllAppliesResponce) {
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
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]*VacancyActivityPreview, 0, 8)
					} else {
						out.Data = []*VacancyActivityPreview{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *VacancyActivityPreview
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(VacancyActivityPreview)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Data = append(out.Data, v1)
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
func easyjson4038ac0aEncodeHeadHunterInternalEntityModels2(out *jwriter.Writer, in GetAllAppliesResponce) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix[1:])
		if in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Data {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetAllAppliesResponce) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4038ac0aEncodeHeadHunterInternalEntityModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAllAppliesResponce) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4038ac0aEncodeHeadHunterInternalEntityModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAllAppliesResponce) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4038ac0aDecodeHeadHunterInternalEntityModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAllAppliesResponce) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4038ac0aDecodeHeadHunterInternalEntityModels2(l, v)
}

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

func easyjson39b3a2f5DecodeHeadHunterInternalEntityModels(in *jlexer.Lexer, out *Resumes) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Resumes, 0, 8)
			} else {
				*out = Resumes{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *Resume
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(Resume)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson39b3a2f5EncodeHeadHunterInternalEntityModels(out *jwriter.Writer, in Resumes) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
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

// MarshalJSON supports json.Marshaler interface
func (v Resumes) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Resumes) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Resumes) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Resumes) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels(l, v)
}
func easyjson39b3a2f5DecodeHeadHunterInternalEntityModels1(in *jlexer.Lexer, out *ResumePreviews) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ResumePreviews, 0, 8)
			} else {
				*out = ResumePreviews{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 *ResumePreview
			if in.IsNull() {
				in.Skip()
				v4 = nil
			} else {
				if v4 == nil {
					v4 = new(ResumePreview)
				}
				(*v4).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson39b3a2f5EncodeHeadHunterInternalEntityModels1(out *jwriter.Writer, in ResumePreviews) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			if v6 == nil {
				out.RawString("null")
			} else {
				(*v6).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ResumePreviews) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResumePreviews) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResumePreviews) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResumePreviews) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels1(l, v)
}
func easyjson39b3a2f5DecodeHeadHunterInternalEntityModels2(in *jlexer.Lexer, out *ResumePreview) {
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
		case "image":
			out.Image = string(in.String())
		case "applicant_name":
			out.ApplicantName = string(in.String())
		case "applicant_surname":
			out.ApplicantSurname = string(in.String())
		case "id":
			out.Id = uint(in.Uint())
		case "title":
			out.Title = string(in.String())
		case "created_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedTime).UnmarshalJSON(data))
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
func easyjson39b3a2f5EncodeHeadHunterInternalEntityModels2(out *jwriter.Writer, in ResumePreview) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix[1:])
		out.String(string(in.Image))
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
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Uint(uint(in.Id))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"created_date\":"
		out.RawString(prefix)
		out.Raw((in.CreatedTime).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResumePreview) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResumePreview) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResumePreview) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResumePreview) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels2(l, v)
}
func easyjson39b3a2f5DecodeHeadHunterInternalEntityModels3(in *jlexer.Lexer, out *ResumeFilter) {
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
		case "Title":
			out.Title = string(in.String())
		case "Location":
			out.Location = string(in.String())
		case "ExperienceInYears":
			out.ExperienceInYears = string(in.String())
		case "FirstSalaryValue":
			out.FirstSalaryValue = string(in.String())
		case "SecondSalaryValue":
			out.SecondSalaryValue = string(in.String())
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
func easyjson39b3a2f5EncodeHeadHunterInternalEntityModels3(out *jwriter.Writer, in ResumeFilter) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"Location\":"
		out.RawString(prefix)
		out.String(string(in.Location))
	}
	{
		const prefix string = ",\"ExperienceInYears\":"
		out.RawString(prefix)
		out.String(string(in.ExperienceInYears))
	}
	{
		const prefix string = ",\"FirstSalaryValue\":"
		out.RawString(prefix)
		out.String(string(in.FirstSalaryValue))
	}
	{
		const prefix string = ",\"SecondSalaryValue\":"
		out.RawString(prefix)
		out.String(string(in.SecondSalaryValue))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResumeFilter) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResumeFilter) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResumeFilter) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResumeFilter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels3(l, v)
}
func easyjson39b3a2f5DecodeHeadHunterInternalEntityModels4(in *jlexer.Lexer, out *Resume) {
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
		case "user_account_id":
			out.UserAccountId = uint(in.Uint())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "created_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedTime).UnmarshalJSON(data))
			}
		case "location":
			out.Location = string(in.String())
		case "experience":
			out.ExperienceInYears = string(in.String())
		case "salary":
			out.Salary = uint(in.Uint())
		case "education_detail":
			(out.EducationDetail).UnmarshalEasyJSON(in)
		case "experience_detail":
			(out.ExperienceDetail).UnmarshalEasyJSON(in)
		case "applicant_skills":
			if in.IsNull() {
				in.Skip()
				out.ApplicantSkills = nil
			} else {
				in.Delim('[')
				if out.ApplicantSkills == nil {
					if !in.IsDelim(']') {
						out.ApplicantSkills = make([]Skill, 0, 2)
					} else {
						out.ApplicantSkills = []Skill{}
					}
				} else {
					out.ApplicantSkills = (out.ApplicantSkills)[:0]
				}
				for !in.IsDelim(']') {
					var v7 Skill
					(v7).UnmarshalEasyJSON(in)
					out.ApplicantSkills = append(out.ApplicantSkills, v7)
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
func easyjson39b3a2f5EncodeHeadHunterInternalEntityModels4(out *jwriter.Writer, in Resume) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"user_account_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.UserAccountId))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"created_date\":"
		out.RawString(prefix)
		out.Raw((in.CreatedTime).MarshalJSON())
	}
	if in.Location != "" {
		const prefix string = ",\"location\":"
		out.RawString(prefix)
		out.String(string(in.Location))
	}
	if in.ExperienceInYears != "" {
		const prefix string = ",\"experience\":"
		out.RawString(prefix)
		out.String(string(in.ExperienceInYears))
	}
	if in.Salary != 0 {
		const prefix string = ",\"salary\":"
		out.RawString(prefix)
		out.Uint(uint(in.Salary))
	}
	{
		const prefix string = ",\"education_detail\":"
		out.RawString(prefix)
		(in.EducationDetail).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"experience_detail\":"
		out.RawString(prefix)
		(in.ExperienceDetail).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"applicant_skills\":"
		out.RawString(prefix)
		if in.ApplicantSkills == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.ApplicantSkills {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Resume) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Resume) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Resume) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Resume) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels4(l, v)
}
func easyjson39b3a2f5DecodeHeadHunterInternalEntityModels5(in *jlexer.Lexer, out *Response) {
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
func easyjson39b3a2f5EncodeHeadHunterInternalEntityModels5(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels5(l, v)
}
func easyjson39b3a2f5DecodeHeadHunterInternalEntityModels6(in *jlexer.Lexer, out *GetAllResumesResponcePointer) {
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
						out.Data = make([]*Resume, 0, 8)
					} else {
						out.Data = []*Resume{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v10 *Resume
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						if v10 == nil {
							v10 = new(Resume)
						}
						(*v10).UnmarshalEasyJSON(in)
					}
					out.Data = append(out.Data, v10)
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
func easyjson39b3a2f5EncodeHeadHunterInternalEntityModels6(out *jwriter.Writer, in GetAllResumesResponcePointer) {
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
			for v11, v12 := range in.Data {
				if v11 > 0 {
					out.RawByte(',')
				}
				if v12 == nil {
					out.RawString("null")
				} else {
					(*v12).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetAllResumesResponcePointer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAllResumesResponcePointer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson39b3a2f5EncodeHeadHunterInternalEntityModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAllResumesResponcePointer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAllResumesResponcePointer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson39b3a2f5DecodeHeadHunterInternalEntityModels6(l, v)
}

package domain

import tjson "github.com/husamettinarabaci/go-pdftojpeg/tool/json"

type Response struct {
	Files [][]byte `json:"files"`
}

func (o Response) ToJson() string {
	return tjson.ToJson(o)
}

func (a Response) FromJson(i string) Response {
	return tjson.FromJson[Response](i)
}

func NewResponse(files [][]byte) Response {
	return Response{
		Files: files,
	}
}

func (o Response) IsEmpty() bool {
	return o.ToJson() == Response{}.ToJson()
}

func (o Response) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o Response) IsEqual(i Response) bool {
	return o.ToJson() == i.ToJson()
}

func (o Response) IsValid() error {
	if o.IsEmpty() {
		return ErrInvalidInput
	}
	if o.Files == nil {
		return ErrInvalidInput
	}
	if len(o.Files) == 0 {
		return ErrInvalidInput
	}
	return nil
}

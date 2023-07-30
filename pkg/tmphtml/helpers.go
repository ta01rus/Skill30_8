package tmphtml

import (
	"bytes"
	"encoding/json"
	"html/template"
	"path"
)

type KeyFn func(...any) *template.HTML

func (f *FileHtmlTemplate) Include(fl string, args any) template.HTML {

	var buf bytes.Buffer

	file := path.Join(f.Dir, fl)

	t, err := template.ParseFiles(file)
	if err != nil {
		return template.HTML(err.Error())
	}

	err = t.Execute(&buf, args)
	if err != nil {
		return template.HTML(err.Error())
	}

	return template.HTML(buf.Bytes())

}

func (f *FileHtmlTemplate) ToJson(obj any) template.HTML {
	b, _ := json.Marshal(obj)
	return template.HTML(string(b))
}

func (f *FileHtmlTemplate) SetOptions(key string, obj any) string {
	f.Options[key] = obj
	return ""
}

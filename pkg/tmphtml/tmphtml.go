package tmphtml

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"path/filepath"
)

// Компанент генерации html
type FileHtmlTemplate struct {
	// имя компанента
	Name string
	// директория компанента
	Dir string

	Files map[string]string

	Tmp *template.Template

	Options map[string]any
}

func New(name, dir string, fs []string, fncs map[string]KeyFn) (*FileHtmlTemplate, error) {
	tmp := &FileHtmlTemplate{
		Name:    name,
		Dir:     dir,
		Options: make(map[string]any),
	}
	fn, err := load(tmp.Dir)
	if err != nil {
		return nil, err
	}
	files := []string{}
	files = append(files, fn["index.html"])
	for _, n := range fs {
		nm, ok := fn[n]
		if !ok {
			return nil, fmt.Errorf("шаблон не найден")
		}
		files = append(files, nm)
	}
	funcMap := template.FuncMap{
		"include": tmp.Include,
		"json":    tmp.ToJson,
		"options": tmp.SetOptions,
	}
	for key, f := range fncs {
		funcMap[key] = f
	}
	ts, err := template.New(name).Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	tmp.Tmp = ts
	// предворительно для теста
	_, err = tmp.Render(nil)
	if err != nil {
		return nil, err
	}

	return tmp, nil
}

// загрузка шаблонов
func load(dir string) (map[string]string, error) {
	ret := make(map[string]string)
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ret[info.Name()] = path
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (f *FileHtmlTemplate) Render(args any) (*template.HTML, error) {
	var b bytes.Buffer

	err := f.Tmp.Execute(&b, args)
	if err != nil {
		return nil, err
	}

	ret := template.HTML(b.String())

	return &ret, nil
}

package helper

import (
	"time"
)

type Shader struct {
	id                 ProgramId
	vertPath           string
	fragPath           string
	vert_modified_time time.Time
	frag_modified_time time.Time
}

func NewShader(vert_path string, frag_path string) (*Shader, error) {
	id, err := CreateProgram(vert_path, frag_path)
	if err != nil {
		return nil, err
	}
	frag_mod_time, frag_time_err := get_modified_time(frag_path)
	if frag_time_err != nil {
		return nil, frag_time_err
	}
	vert_mod_time, vert_time_err := get_modified_time(frag_path)
	if vert_time_err != nil {
		return nil, vert_time_err
	}
	result := &Shader{id, vert_path, frag_path, vert_mod_time, frag_mod_time}
	return result, nil
}

func (shader *Shader) Use() {
	UseProgram(shader.id)
}

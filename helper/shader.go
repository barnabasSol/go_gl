package helper

import (
	"time"
)

type Shader struct {
	id            ProgramId
	vertPath      string
	fragPath      string
	modified_time time.Time
}

func new_shader(vert_path string, frag_path string) (*Shader, error) {
	id, err := CreateProgram(vert_path, frag_path)
	if err != nil {
		return nil, err
	}
	return &Shader{id, vert_path, frag_path, get_modified_time(frag_path)}, nil
}

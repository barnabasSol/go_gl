package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func (shader *Shader) CheckForShaderChanges() error {
	vert_mod_time, vert_err := get_modified_time(shader.vertPath)
	if vert_err != nil {
		return vert_err
	}
	frag_mod_time, frag_err := get_modified_time(shader.fragPath)
	if frag_err != nil {
		return frag_err
	}
	if !vert_mod_time.Equal(shader.vert_modified_time) || !frag_mod_time.Equal(shader.frag_modified_time) {
		new_id, err := CreateProgram(shader.vertPath, shader.fragPath)
		if err != nil {
			fmt.Println(err)
		} else {
			gl.DeleteProgram(uint32(shader.id))
			shader.id = new_id
		}

	}
	return nil
}

func LoadShader(path string, shaderType uint32) (ShaderId, error) {
	shaderFile, os_err := os.ReadFile(path)
	if os_err != nil {
		return 0, os_err
	}
	shaderFileStr := string(shaderFile)
	shader_id, s_err := createShader(shaderFileStr, shaderType)
	if s_err != nil {
		return 0, os_err
	}
	return shader_id, nil
}

func get_modified_time(file_path string) (time.Time, error) {
	file, err := os.Stat(file_path)
	if err != nil {
		return time.Time{}, err
	}
	return file.ModTime(), nil
}

package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
)

var loadedShaders = make(map[ProgramId]*Shader)

func CheckForShaderChanges() {
	for _, shader := range loadedShaders {
		vert_mod_time := get_modified_time(shader.vertPath)
		frag_mod_time := get_modified_time(shader.fragPath)
		if !vert_mod_time.Equal(shader.modified_time) || !frag_mod_time.Equal(shader.modified_time) {
			new_id, err := CreateProgram(shader.vertPath, shader.fragPath)
			if err != nil {
				fmt.Println(err)
			} else {
				gl.DeleteProgram(uint32(shader.id))
				shader.id = new_id
			}

		}

	}
}

func LoadShader(path string, shaderType uint32) (ShaderId, error) {
	shaderFile, os_err := os.ReadFile(path)
	if os_err != nil {
		panic(os_err)
	}
	shaderFileStr := string(shaderFile)
	shader_id, s_err := createShader(shaderFileStr, shaderType)
	if s_err != nil {
		return 0, os_err
	}
	return shader_id, nil
}

func get_modified_time(file_path string) time.Time {
	file, stat_err := os.Stat(file_path)
	_ = file
	if stat_err != nil {
		panic(stat_err)
	}
	return file.ModTime()
}

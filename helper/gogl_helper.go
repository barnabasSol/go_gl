package helper

import (
	"errors"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type ShaderId uint32
type ProgramId uint32
type TextureId uint32
type BufferId uint32

func GetVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

func createShader(shader_content string, shader_type uint32) (ShaderId, error) {
	shader_id := gl.CreateShader(shader_type)
	shader_content = shader_content + "\x00"
	csource, free := gl.Strs(shader_content)
	gl.ShaderSource(shader_id, 1, csource, nil)
	free()
	gl.CompileShader(shader_id)
	var status int32
	gl.GetShaderiv(shader_id, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var log_length int32
		gl.GetShaderiv(shader_id, gl.INFO_LOG_LENGTH, &log_length)
		log := strings.Repeat("\x00", int(log_length+1))
		gl.GetShaderInfoLog(shader_id, log_length, nil, gl.Str(log))
		switch shader_type {
		case gl.VERTEX_SHADER:
			return 0, errors.New("vertex shader compilation faild sadly, here's why:\n" + log)
		case gl.FRAGMENT_SHADER:
			return 0, errors.New("fragment shader compilation faild sadly, here's why:\n" + log)
		}
	}
	return ShaderId(shader_id), nil
}

func CreateProgram(vert_path string, frag_path string) (ProgramId, error) {
	vert, err := LoadShader(vert_path, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}
	frag, err := LoadShader(frag_path, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	shader_program := gl.CreateProgram()
	gl.AttachShader(shader_program, uint32(vert))
	gl.AttachShader(shader_program, uint32(frag))
	gl.LinkProgram(shader_program)

	var success int32
	gl.GetProgramiv(shader_program, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var log_length int32
		gl.GetShaderiv(shader_program, gl.INFO_LOG_LENGTH, &log_length)
		log := strings.Repeat("\x00", int(log_length+1))
		gl.GetProgramInfoLog(shader_program, log_length, nil, gl.Str(log))
		return 0, errors.New("failed to link sadly\n" + log)
	}

	gl.DeleteShader(uint32(vert))
	gl.DeleteShader(uint32(frag))

	return ProgramId(shader_program), nil
}

func GenBindBuffer(target uint32, location int32) BufferId {
	var buffer uint32
	gl.GenBuffers(location, &buffer)
	gl.BindBuffer(target, buffer)
	return BufferId(buffer)
}

func GenBindVertexArray(location int32) BufferId {
	var VAO uint32
	gl.GenVertexArrays(location, &VAO)
	gl.BindVertexArray(VAO)
	return BufferId(VAO)
}

func BufferDataFloat(target uint32, data []float32, usage uint32) {
	gl.BufferData(target, len(data)*4, gl.Ptr(data), usage)
}

func BufferDataInt(target uint32, data []uint32, usage uint32) {
	gl.BufferData(target, len(data)*4, gl.Ptr(data), usage)
}
func UnbindVertexArray() {
	gl.BindVertexArray(0)
}

func UseProgram(shader_program ProgramId) {
	gl.UseProgram(uint32(shader_program))
}
func BindVertextArray(vaoID BufferId) {
	gl.BindVertexArray(uint32(vaoID))
}

func GenEBO() BufferId {
	var EBO uint32
	gl.GenBuffers(1, &EBO)
	return BufferId(EBO)
}

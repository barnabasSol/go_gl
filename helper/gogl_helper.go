package helper

import (
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type ShaderId uint32
type ProgramId uint32
type VBOID uint32
type VAOID uint32

func GetVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

func CreateShader(shader_src string, shader_type uint32) ShaderId {
	shader_id := gl.CreateShader(shader_type)
	shader_src = shader_src + "\x00"
	csource, free := gl.Strs(shader_src)
	gl.ShaderSource(shader_id, 1, csource, nil)
	free()
	gl.CompileShader(shader_id)
	var vert_status int32
	gl.GetShaderiv(shader_id, gl.COMPILE_STATUS, &vert_status)
	if vert_status == gl.FALSE {
		var log_length int32
		gl.GetShaderiv(shader_id, gl.INFO_LOG_LENGTH, &log_length)
		log := strings.Repeat("\x00", int(log_length+1))
		gl.GetShaderInfoLog(shader_id, log_length, nil, gl.Str(log))
		switch shader_type {
		case gl.VERTEX_SHADER:
			panic("failed vertex shader \n" + log + "\n")
		case gl.FRAGMENT_SHADER:
			panic("failed fragment shader \n" + log + "\n")
		}
	}
	return ShaderId(shader_id)
}

func CreateProgram(vert ShaderId, frag ShaderId) ProgramId {
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
		gl.GetShaderInfoLog(shader_program, log_length, nil, gl.Str(log))
		panic("failed link program shader \n" + log)
	}
	gl.DeleteShader(uint32(vert))
	gl.DeleteShader(uint32(frag))
	return ProgramId(shader_program)
}

func GenBindBuffer(target uint32) VBOID {
	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(target, VBO)
	return VBOID(VBO)
}

func GenBindVertexArray() VAOID {
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)
	return VAOID(VAO)
}

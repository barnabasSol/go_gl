#version 330 core

layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTexCoord;


uniform float xv;
uniform float yv;

out vec2 TexCoord;

void main()
{
    gl_Position = vec4(aPos.x + xv, aPos.y + yv, aPos.z, 1.0);
    TexCoord = aTexCoord;
}
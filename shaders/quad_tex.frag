#version 330 core

out vec4 fragColor;
in vec2 TexCoord;

uniform sampler2D texture1;

void main()
{
    vec4 tex = texture(texture1, TexCoord);
    fragColor = vec4(tex.x, tex.y, tex.z, 1.0);
}

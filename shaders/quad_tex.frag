#version 330 core

out vec4 fragColor;
in vec2 TexCoord;

uniform float x;
uniform float y;

uniform sampler2D texture1;

void main()
{
    vec4 tex = texture(texture1, TexCoord);
    fragColor = vec4(tex.x + x, tex.y + x, 0.0 + x, 0.0 + x);
}

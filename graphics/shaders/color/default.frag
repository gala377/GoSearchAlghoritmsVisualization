#version 410

out vec4 frag_colour;

uniform vec4 color;

void main()
{
    frag_colour = color;
}
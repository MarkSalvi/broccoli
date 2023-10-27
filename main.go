package main

import (
	"fmt"
	"log"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/MarkSalvi/glHelper"
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
)

type ProgramID uint32


var vertices[] float32 = []float32{
	-0.5,-0.5,0.0,
	0.5,-0.5,0.0,
	0.0,0.5,0.0,
}

func init() {
	runtime.LockOSThread()
}

func main(){
	//sdl initialization
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if(err != nil){
		log.Fatalln("Sdl Init Error: ",err)
	}
	defer sdl.Quit()

	var flags uint32
	flags = /*sdl.WINDOW_INPUT_GRABBED |*/ sdl.WINDOW_OPENGL
	window, err := sdl.CreateWindow("Test",0,0,1080,720,flags)
	if(err != nil){
		log.Fatalln("Sdl window creation Error: ",err)
	}
	context , err := window.GLCreateContext()
	if err != nil {
		log.Fatalln("GL context creation Error: ",err)
	}
	err = window.GLMakeCurrent(context)
	if err != nil {
		log.Fatalln("GL make current Error: ",err)
	}
	defer window.Destroy()

	gl.Init()
	fmt.Println(glHelper.GetVersion())
	
	shaderProgram,err := glHelper.NewShader("./shaders/vertex.frag","./shaders/frag.frag")
	if (err != nil){
		log.Fatalln("shaderProgram Error: ",err)
	}

	glHelper.GenBindBuffer(gl.ARRAY_BUFFER)
	triangleVAO := glHelper.GenBindVertexArray()
	glHelper.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.Ptr(nil))
	gl.EnableVertexAttribArray(0)
	glHelper.BindVertexArray(0)



	
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.2, 0.2, 0.3, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	
	for{

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		gl.ClearColor(0.2, 0.2, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		shaderProgram.Use()
		glHelper.BindVertexArray(triangleVAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)


	}

}


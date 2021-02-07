package main

import (
	"time"

	"github.com/mitchellh/go-glint"
	gc "github.com/mitchellh/go-glint/components"
)

func main() {
	d := glint.New()
	d.Append(
		glint.Style(
			glint.Layout(
				gc.Spinner(),
				glint.Layout(glint.Text("Build site and validate links...")).MarginLeft(1),
				glint.Layout(gc.Stopwatch(time.Now())).MarginLeft(1),
			).Row(),
			glint.Color("green"),
		),
		glint.Layout(
			gc.Spinner(),
			glint.Layout(glint.Text("Preparing execution environment...")).MarginLeft(1),
			glint.Layout(gc.Stopwatch(time.Now())).MarginLeft(1),
		).MarginLeft(2).Row(),
		glint.Layout(
			glint.Text("Preparing volume to work with..."),
		).MarginLeft(4),
		glint.Text("\nWaiting..."),
	)
	for {
		d.RenderFrame()
		time.Sleep(time.Second)
	}

	//d.Render(context.Background())
}
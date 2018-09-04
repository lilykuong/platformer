package main

import (
	//for sprite
	"image"
	"os"
	_ "image/png" //for other images just import their packages

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
	// opens file using os package
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// decode image using image package
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	// converts image.Image to PictureData, return type is pixel.Picture
	return pixel.PictureDataFromImage(img), nil
}


func run() {
	// Sets the Window Bounds and settings
	cfg := pixelgl.WindowConfig{
		Title:  "Corgi",
		Bounds: pixel.R(0, 0, 768, 768),
		// updates the window according to how fast your computer updates the screen
		VSync: true,
		Resizable: true,
		//Icon []pixel.Picture
	}
	// Opens the window
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPicture("corgi.png")
	if err != nil {
		panic(err)
	}
	// sprite are anchored by centered of pic
	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(colornames.Blue)
	// Moved can do rotations, scaling, movement
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		win.Update()
	}
}

func main(){
	pixelgl.Run(run)
}

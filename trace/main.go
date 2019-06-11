package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	http.Handle("/hello", http.HandlerFunc(helloHandler))

	http.ListenAndServe("localhost:8181", http.DefaultServeMux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func watermark(w http.ResponseWriter, r *http.Request) {
	image1, err := os.Open("jellyfish.jpg")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	first, err := jpeg.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()

	image2, err := os.Open("pokeball.png")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	second, err := png.Decode(image2)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image2.Close()

	offset := image.Pt(300, 200)
	b := first.Bounds()

	min := image.Point{X: 0, Y: 0}
	max := image.Point{X: 1024, Y: 1024}
	rectange := image.Rectangle{Min: min, Max: max}

	image := image.NewRGBA(rectange)
	draw.Draw(image, rectange, second, image.ZP, draw.Over)

	third, err := os.Create("result.jpg")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, image, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()

}

package main

import (
	"bytes"
	"github.com/wuyuanyi135/benchmark-grpc-image-sequence/protogen"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"
)


type ImageServiceImpl struct {
	cache []*ImageCache
}
type ImageCache struct {
	Name     string
	Blob     []byte
	MIMEType string
	Decoded  []uint8
	Height   uint64
	Width    uint64
}

func NewImageServiceImpl() *ImageServiceImpl {
	files, err := ioutil.ReadDir("images")
	if err != nil {
		panic(err)
	}
	// memory cache: do not read too many files
	var cache []*ImageCache
	for _, file := range files {
		ext := filepath.Ext(file.Name())

		if !file.IsDir() {
			var mime string
			switch ext {
			case ".png":
				mime = "image/png"
			case ".jpg":
			case ".jpeg":
				mime = "image/jpeg"
			default:
				continue
			}
			filePath := filepath.Join("images", file.Name())
			var f []byte
			f, err = ioutil.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

			reader := bytes.NewReader(f)
			img, _, err := image.Decode(reader)
			if err != nil {
				panic(err)
			}
			m := image.NewRGBA(img.Bounds())
			draw.Draw(m, m.Bounds(), img, image.ZP, draw.Src)

			height := img.Bounds().Dy()
			width := img.Bounds().Dx()

			cache = append(cache, &ImageCache{
				Name:     filePath,
				Blob:     f,
				MIMEType: mime,
				Decoded: m.Pix,
				Height: uint64(height),
				Width: uint64(width),
			})
		}
	}

	return &ImageServiceImpl{cache: cache}
}


func (s *ImageServiceImpl) ImageStreaming(req *protogen.ImageStreamRequest, srv protogen.ImageSequenceService_ImageStreamingServer) (err error) {

	interval := req.IntervalMs;
	for {
		select {
		case <-time.After(time.Duration(interval) * time.Millisecond):
			// read a random image from 'images' folder
			n := rand.Intn(len(s.cache))

			cache := s.cache[n]
			if req.Bitmap {
				// send decoded file
				err = srv.Send(&protogen.ImageStream{
					Image:  cache.Decoded,
					Type:   "image/x-rgba",
					Height: cache.Height,
					Width: cache.Width,
				})
			} else {
				// send raw file
				err = srv.Send(&protogen.ImageStream{
					Image: cache.Blob,
					Type:  cache.MIMEType,
					Height: cache.Height,
					Width: cache.Width,
				})
				if err != nil {
					return
				}
			}

			break
		case <-srv.Context().Done():
			return nil
		}
	}
}

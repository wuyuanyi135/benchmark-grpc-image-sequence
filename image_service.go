package main

import (
	"github.com/wuyuanyi135/benchmark-grpc-image-sequence/protogen"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"
)

type ImageServiceImpl struct {
	cache []*ImageCache
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
			cache = append(cache, &ImageCache{
				Name:     filePath,
				Blob:     f,
				MIMEType: mime,
			})
		}
	}

	return &ImageServiceImpl{cache: cache}
}

type ImageCache struct {
	Name     string
	Blob     []byte
	MIMEType string
}

func (s *ImageServiceImpl) ImageStreaming(req *protogen.ImageStreamRequest, srv protogen.ImageSequenceService_ImageStreamingServer) (err error) {

	interval := req.IntervalMs;
	for {
		select {
		case <-time.After(time.Duration(interval) * time.Millisecond):
			// read a random image from 'images' folder
			n := rand.Intn(len(s.cache))
			err = srv.Send(&protogen.ImageStream{
				Image: s.cache[n].Blob,
				Type:  s.cache[n].MIMEType,
			})
			if err != nil {
				return
			}
			break
		case <-srv.Context().Done():
			return nil
		}
	}
}

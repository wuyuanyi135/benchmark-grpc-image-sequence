Compare 3 binary image display techniques for Web development.
[StackOverflow](https://stackoverflow.com/questions/55608249/is-it-possible-to-render-a-png-jpeg-byte-array-to-an-element-without-creating-a)


# The three possible solutions

- Transfer jpeg/png encoded data, Use <img> to load the blob url.

- Transfer raw RGBA encoded data, use <canvas>.putImageData to update image.

- Transfer jpeg/png encoded data, *decode it into RGBA*, then use <canvas>.putImageData to update image. (Not implemented because I don't know how to decode a file into RGBA buffer in javascript)


From my test, the first method should be most efficient despite the annoying fake network request in Chrome DevTools. Note that attached CDT will drastically slow down the image update.


# Build
```
cd ui
npm install

cd ..
python3 protoc-generation.py

cd ui
npm run build

cd ..
go build .
./benchmark-grpc-image-sequence
```

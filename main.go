package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
	"github.com/wuyuanyi135/benchmark-grpc-image-sequence/protogen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net/http"
	"path"
)

func main() {
	dev := flag.Bool("Development", false, "use development forwarding")
	flag.Parse()
	if !*dev {
		gin.SetMode(gin.ReleaseMode)
	}
	server := grpc.NewServer()

	imageService := NewImageServiceImpl()

	protogen.RegisterImageSequenceServiceServer(server, imageService);
	wrappedGrpc := grpcweb.WrapServer(server)

	router := gin.Default()
	router.Use(GinGrpcWebMiddleware(wrappedGrpc))
	if *dev {
		fwd, _ := forward.New()
		router.GET("/*proxy", func(context *gin.Context) {
			context.Request.URL = testutils.ParseURI("http://localhost:4200")
			fwd.ServeHTTP(context.Writer, context.Request)
		})
		router.POST("/*proxy", func(context *gin.Context) {
			context.Request.URL = testutils.ParseURI("http://localhost:4200")
			fwd.ServeHTTP(context.Writer, context.Request)
		})
	} else {
		fmt.Println("Using dist files")
		router.Static("/", path.Join("ui", "dist","ui"))
	}

	if err := http.ListenAndServe(":8088", router); err != nil {
		grpclog.Fatalf("failed starting http2 server: %v", err)
	}
}

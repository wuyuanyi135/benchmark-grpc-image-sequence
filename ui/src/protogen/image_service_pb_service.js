// package: image_service
// file: image_service.proto

var image_service_pb = require("./image_service_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ImageSequenceService = (function () {
  function ImageSequenceService() {}
  ImageSequenceService.serviceName = "image_service.ImageSequenceService";
  return ImageSequenceService;
}());

ImageSequenceService.ImageStreaming = {
  methodName: "ImageStreaming",
  service: ImageSequenceService,
  requestStream: false,
  responseStream: true,
  requestType: image_service_pb.ImageStreamRequest,
  responseType: image_service_pb.ImageStream
};

exports.ImageSequenceService = ImageSequenceService;

function ImageSequenceServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ImageSequenceServiceClient.prototype.imageStreaming = function imageStreaming(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(ImageSequenceService.ImageStreaming, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.end.forEach(function (handler) {
        handler();
      });
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.ImageSequenceServiceClient = ImageSequenceServiceClient;


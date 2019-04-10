// package: image_service
// file: image_service.proto

import * as jspb from "google-protobuf";

export class ImageStreamRequest extends jspb.Message {
  getIntervalMs(): number;
  setIntervalMs(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ImageStreamRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ImageStreamRequest): ImageStreamRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ImageStreamRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ImageStreamRequest;
  static deserializeBinaryFromReader(message: ImageStreamRequest, reader: jspb.BinaryReader): ImageStreamRequest;
}

export namespace ImageStreamRequest {
  export type AsObject = {
    intervalMs: number,
  }
}

export class ImageStream extends jspb.Message {
  getImage(): Uint8Array | string;
  getImage_asU8(): Uint8Array;
  getImage_asB64(): string;
  setImage(value: Uint8Array | string): void;

  getType(): string;
  setType(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ImageStream.AsObject;
  static toObject(includeInstance: boolean, msg: ImageStream): ImageStream.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ImageStream, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ImageStream;
  static deserializeBinaryFromReader(message: ImageStream, reader: jspb.BinaryReader): ImageStream;
}

export namespace ImageStream {
  export type AsObject = {
    image: Uint8Array | string,
    type: string,
  }
}


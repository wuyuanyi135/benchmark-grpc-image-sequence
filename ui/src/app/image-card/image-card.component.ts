import {Component, Input, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {DomSanitizer, SafeUrl} from '@angular/platform-browser';
import {grpc} from '@improbable-eng/grpc-web';
import {ImageSequenceService, ImageSequenceServiceImageStreaming} from '../../protogen/image_service_pb_service';
import {ImageStream, ImageStreamRequest} from '../../protogen/image_service_pb';
import {BehaviorSubject, Subject} from 'rxjs';
import {bufferCount, map, takeUntil} from 'rxjs/operators';

@Component({
  selector: 'app-image-card',
  templateUrl: './image-card.component.html',
  styleUrls: ['./image-card.component.css']
})
export class ImageCardComponent implements OnInit, OnDestroy {
  @Input() title: string;
  @Input() interval: number;

  private _streaming = new BehaviorSubject(false);
  private _streamClient: grpc.Client<grpc.ProtobufMessage, grpc.ProtobufMessage>;
  @Input() set streaming(value: boolean) {
    this._streaming.next(value);
  }

  @ViewChild(HTMLImageElement) imgRef: HTMLImageElement;
  public ImageSrc: string;
  public end = false;

  private actualInterval = new Subject<number>();
  private messageReceived = new Subject<number>();
  private destory = new Subject();
  constructor(private sanitizer: DomSanitizer) {

  }

  ngOnInit() {
    this._streaming.pipe(takeUntil(this.destory)).subscribe(value => {
      if (value) {
        this.createStream();
      } else {
        if (this._streamClient) {
          this._streamClient.close();
        }
      }
    });

    this.messageReceived
      .pipe(bufferCount(2, 1)) // two successive value
      .pipe(map(value => value[1] - value[0])) // ms difference
      .pipe(bufferCount(4, 1))
      .pipe(map(value => value.reduce((p, c) => p + c) / (value.length)))
      .pipe(takeUntil(this.destory)).subscribe(this.actualInterval);
  }

  createStream() {
    const request = new ImageStreamRequest();
    request.setIntervalMs(this.interval);
    const client = grpc.client(ImageSequenceService.ImageStreaming, {
      host: window.location.origin,
      debug: false,
      transport: grpc.WebsocketTransport(),
    });
    client.onMessage((res: ImageStream) => {
      URL.revokeObjectURL(this.ImageSrc);
      const blob = new Blob([res.getImage_asU8()], {type: res.getType()});
      this.ImageSrc = URL.createObjectURL(blob);
      // this.ImageSrc = `data:${res.getType()};base64, ${res.getImage_asB64()}`;
      this.messageReceived.next(Date.now());
    });
    client.onEnd((code, message) => {
      this.end = true;
    });
    this._streamClient = client;
    client.start();
    client.send(request);
  }

  ngOnDestroy(): void {
    this.destory.complete();
    if (this._streamClient) {
      this._streamClient.close();
    }
  }
}

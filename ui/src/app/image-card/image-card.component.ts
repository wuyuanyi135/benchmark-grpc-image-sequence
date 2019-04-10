import {Component, Input, OnInit, ViewChild} from '@angular/core';
import {DomSanitizer, SafeUrl} from '@angular/platform-browser';
import {grpc} from '@improbable-eng/grpc-web';
import {ImageSequenceService, ImageSequenceServiceImageStreaming} from '../../protogen/image_service_pb_service';
import {ImageStream, ImageStreamRequest} from '../../protogen/image_service_pb';
import {BehaviorSubject, Subject} from 'rxjs';

@Component({
  selector: 'app-image-card',
  templateUrl: './image-card.component.html',
  styleUrls: ['./image-card.component.css']
})
export class ImageCardComponent implements OnInit {
  @Input() title: string;
  @Input() interval: number;
  @ViewChild(HTMLImageElement) imgRef: HTMLImageElement;
  public ImageSrc: string;
  public end = false;

  constructor(private sanitizer: DomSanitizer) {

  }

  ngOnInit() {
    const request = new ImageStreamRequest();
    request.setIntervalMs(this.interval);
    grpc.invoke(
      ImageSequenceService.ImageStreaming,
      {
        host: window.location.origin,
        request,
        debug: false,
        onMessage: (res: ImageStream) => {
          URL.revokeObjectURL(this.ImageSrc);
          const blob = new Blob([res.getImage_asU8()], {type: res.getType()});
          this.ImageSrc = URL.createObjectURL(blob);
          // this.ImageSrc = `data:${res.getType()};base64, ${res.getImage_asB64()}`;
        },
        onEnd: (code, message) => {
          this.end = true;
        }
      }
    );
  }

}

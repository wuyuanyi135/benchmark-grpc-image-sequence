import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'ui';
  cardNumber = 1;
  streamState = false;
  interval = 1000;

  buildArray(cardNumber: number): Array<any> {
    return Array(cardNumber);
  }
}

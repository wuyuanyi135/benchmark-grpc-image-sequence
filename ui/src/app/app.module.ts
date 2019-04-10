import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatCardModule} from '@angular/material';
import { ImageCardComponent } from './image-card/image-card.component';
import { SanitizeUrlPipe } from './sanitize-url.pipe';

@NgModule({
  declarations: [
    AppComponent,
    ImageCardComponent,
    SanitizeUrlPipe
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatCardModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

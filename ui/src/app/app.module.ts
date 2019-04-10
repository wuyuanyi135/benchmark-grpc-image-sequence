import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
  MatButtonModule,
  MatButtonToggleModule,
  MatCardModule,
  MatCheckboxModule,
  MatGridListModule,
  MatInputModule, MatSelectModule
} from '@angular/material';
import { ImageCardComponent } from './image-card/image-card.component';
import { SanitizeUrlPipe } from './sanitize-url.pipe';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import { ImageCardServerRgbaComponent } from './image-card-server-rgba/image-card-server-rgba.component';

@NgModule({
  declarations: [
    AppComponent,
    ImageCardComponent,
    SanitizeUrlPipe,
    ImageCardServerRgbaComponent,
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    FormsModule,
    ReactiveFormsModule,
    MatInputModule,
    MatCardModule,
    MatButtonModule,
    MatGridListModule,
    MatButtonToggleModule,
    MatCheckboxModule,
    MatSelectModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { OverviewModule } from './overview/overview.module';

import {
  ApiSensorRepository,
  ApiReportRepository,
} from './services/api';

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    OverviewModule,
  ],
  providers: [
    ApiSensorRepository,
    ApiReportRepository,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

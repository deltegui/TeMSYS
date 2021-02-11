import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { OverviewModule } from './overview/overview.module';
import { LoginModule } from './login/login.module';

import {
  ApiSensorRepository,
  ApiReportRepository,
  ApiUserRepository,
} from './services/api';

import { LocalStorageTokenRepository } from './services/localstoragetokenrepo';

import {
  UserService
} from './services/user.service';

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    OverviewModule,
    LoginModule,
  ],
  providers: [
    ApiSensorRepository,
    ApiReportRepository,
    ApiUserRepository,
    LocalStorageTokenRepository,
    UserService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

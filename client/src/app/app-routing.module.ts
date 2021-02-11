import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { OverviewPageComponent } from './overview/overview-page/overview-page.component';
import { LoginPageComponent } from './login/login-page/login-page.component';
import { PanelPageComponent } from './panel/panel-page/panel-page.component';

const routes: Routes = [
  { path: '', component: OverviewPageComponent },
  { path: 'login', component: LoginPageComponent },
  { path: 'panel', component: PanelPageComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule { }

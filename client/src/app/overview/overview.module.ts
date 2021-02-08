import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { GeneralCardComponent } from './general-card/general-card.component';
import { OverviewPageComponent } from './overview-page/overview-page.component';
import { SensorCardComponent } from './sensor-card/sensor-card.component';

@NgModule({
  declarations: [
    GeneralCardComponent,
    OverviewPageComponent,
    SensorCardComponent,
  ],
  imports: [
    CommonModule
  ]
})
export class OverviewModule { }

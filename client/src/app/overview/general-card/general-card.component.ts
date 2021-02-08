import { Component, OnInit } from '@angular/core';
import { ApiSensorRepository } from 'src/app/services/api';
import { Report } from '../../services/models';

function getCurrentSeasonName(): string {
  const now = new Date();
  const month = now.getMonth();
  if (month >= 2 && month < 5) {
    return 'spring';
  }
  if (month >= 5 && month < 8) {
    return 'summer';
  }
  if (month >= 8 && month < 11) {
    return 'autumn';
  }
  if (month === 11 || (month >= 0 && month < 2)) {
    return 'winter';
  }
  return 'spring';
}

@Component({
  selector: 'app-general-card',
  templateUrl: './general-card.component.html',
  styleUrls: ['./general-card.component.css']
})
export class GeneralCardComponent implements OnInit {
  temperature: number | null;
  humidity: number | null;

  constructor(private sensorRepo: ApiSensorRepository) {
    this.temperature = null;
    this.humidity = null
  }

  ngOnInit(): void {
    this.setSeasonImage();
    this.sensorRepo.getCurrentAverageState()
      .then((reports: Report[]) => reports.forEach((report) => {
        if (report.type === 'temperature') {
          this.temperature = report.value;
        }
        if (report.type === 'humidity') {
          this.humidity = report.value;
        }
      }));
  }

  private setSeasonImage(): void {
    const season = getCurrentSeasonName();
    const element = document.getElementsByClassName('general-card')[0];
    element.setAttribute(
      'style',
      `background-image: linear-gradient(rgba(26, 26, 29, 0.8) 70%, var(--bg-main-color)), url("/assets/${season}.jpg")`,
    );
  }

}

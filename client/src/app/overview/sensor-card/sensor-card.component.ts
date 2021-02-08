import {
  Component,
  OnInit,
  Input,
} from '@angular/core';
import { ApiSensorRepository } from 'src/app/services/api';
import { Report } from '../../services/models';

@Component({
  selector: 'app-sensor-card',
  templateUrl: './sensor-card.component.html',
  styleUrls: ['./sensor-card.component.css'],
})
export class SensorCardComponent implements OnInit {
  @Input() name: string;
  temperature: number | null;
  humidity: number | null;
  enabled: boolean;

  constructor(private sensorRepo: ApiSensorRepository) {
    this.name = '';
    this.temperature = null;
    this.humidity = null
    this.enabled = false;
  }

  ngOnInit(): void {
    this.sensorRepo.getCurrentStateByName(this.name)
      .then((reports: Report[]) => reports.forEach((report) => {
        this.enabled = true;
        if (report.type === 'temperature') {
          this.temperature = report.value;
        }
        if (report.type === 'humidity') {
          this.humidity = report.value;
        }
      }));
  }
}

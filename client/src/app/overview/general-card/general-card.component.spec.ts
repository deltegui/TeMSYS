import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GeneralCardComponent } from './general-card.component';

describe('GeneralCardComponent', () => {
  let component: GeneralCardComponent;
  let fixture: ComponentFixture<GeneralCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GeneralCardComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(GeneralCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

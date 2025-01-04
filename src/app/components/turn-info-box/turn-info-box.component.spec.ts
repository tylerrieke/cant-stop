import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TurnInfoBoxComponent } from './turn-info-box.component';

describe('TurnInfoBoxComponent', () => {
  let component: TurnInfoBoxComponent;
  let fixture: ComponentFixture<TurnInfoBoxComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TurnInfoBoxComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TurnInfoBoxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

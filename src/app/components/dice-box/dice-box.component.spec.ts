import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DiceBoxComponent } from './dice-box.component';

describe('DiceBoxComponent', () => {
  let component: DiceBoxComponent;
  let fixture: ComponentFixture<DiceBoxComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DiceBoxComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DiceBoxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

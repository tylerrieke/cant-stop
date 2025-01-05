import {Component, Input} from '@angular/core';
import {ChipComponent} from '../chip/chip.component';

@Component({
  selector: '[app-turn-info-box]',
  imports: [ChipComponent],
  templateUrl: './turn-info-box.component.html',
  styleUrl: './turn-info-box.component.css'
})
export class TurnInfoBoxComponent {
  @Input() color: string = "blue"
  @Input() turn: number = 0
}

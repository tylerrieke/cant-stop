import {Component, Input} from '@angular/core';
import {DiceBoxComponent} from '../dice-box/dice-box.component';
import {NgIf} from '@angular/common';

@Component({
  selector: '[app-option-box]',
  imports: [DiceBoxComponent, NgIf],
  templateUrl: './option-box.component.html',
  styleUrl: './option-box.component.css'
})
export class OptionBoxComponent {
  @Input() dicePairs: any[][] = []

  pairTotal(pair: any[]) {
    return pair[0] + pair[1]
  }
}

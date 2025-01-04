import {Component, Input} from '@angular/core';
import {DieComponent} from '../die/die.component';
import {NgFor} from '@angular/common';
import {getRandomInt} from '../utils';

@Component({
  selector: '[app-dice-box]',
  imports: [DieComponent, NgFor],
  templateUrl: './dice-box.component.html',
  styleUrl: './dice-box.component.css'
})
export class DiceBoxComponent {
  @Input() dice= [1, 1, 1, 1]
  @Input() diceSize: string = "medium"
}

import {Component, Input} from '@angular/core';
import {NgClass, NgFor, NgIf} from '@angular/common';
import {ChipComponent} from '../chip/chip.component';

@Component({
  selector: '[app-square]',
  imports: [NgFor, NgIf, NgClass, ChipComponent],
  templateUrl: './square.component.html',
  styleUrl: './square.component.css'
})
export class SquareComponent {
  @Input() title: string = '';
  @Input() numValue: number = 0;
  @Input() chips: any = [];
  @Input() winningColor: string = '';

  isTop() {
    return this.numValue > 0
  }
}

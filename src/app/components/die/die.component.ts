import {Component, Input} from '@angular/core';
import {NgClass, NgIf} from '@angular/common';
import {getRandomInt} from '../utils';

@Component({
  selector: '[app-die]',
  imports: [NgIf, NgClass],
  templateUrl: './die.component.html',
  styleUrl: './die.component.css'
})
export class DieComponent {
  @Input() value: number = getRandomInt(6) + 1
  @Input() size: string = "large"

  getSize() {
    switch (this.size) {
      case "small":
        return "sm";
      case "medium":
        return "md";
      default:
        return "lg";
    }
  }
}

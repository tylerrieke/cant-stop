import {Component, Input} from '@angular/core';
import {NgClass, NgStyle} from '@angular/common';

@Component({
  selector: '[app-chip]',
  imports: [NgClass, NgStyle],
  templateUrl: './chip.component.html',
  styleUrl: './chip.component.css'
})
export class ChipComponent {
  @Input() index: number = 0;
  @Input() numberOf: number = 0;
  @Input() color: string = "crimson"
  @Input() size: string = "medium"

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

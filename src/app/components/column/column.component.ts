import {Component, Input} from '@angular/core';
import {NgFor, NgIf} from '@angular/common';
import {SquareComponent} from '../square/square.component';
import {reverseMapping} from '../utils'

@Component({
  selector: '[app-column]',
  imports: [NgFor, NgIf, SquareComponent],
  templateUrl: './column.component.html',
  styleUrl: './column.component.css'
})
export class ColumnComponent {
  squares: any = [{}, {},{},{}]

  @Input() colNum: number = 0
  @Input() colorChips: any = {}

  ngOnInit() {
    let numSquares = 13 - (Math.abs(7 - this.colNum) * 2)
    let numberChips: {[x: number]: [string]} = reverseMapping(this.colorChips)
    let winningChipColor = (numberChips[numSquares] || [""])[0]
    this.squares = Array.from({length: numSquares}, (_, i) => {
      let chipColors = numberChips[numSquares - i]
      return chipColors ? {chips: chipColors, winningChipColor: winningChipColor} : {winningChipColor: winningChipColor}
    })
  }
}

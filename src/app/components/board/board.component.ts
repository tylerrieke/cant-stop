import {Component, Input} from '@angular/core';
import {NgFor, NgIf} from '@angular/common';
import {ColumnComponent} from '../column/column.component';

@Component({
  selector: '[app-board]',
  imports: [ColumnComponent, NgFor, NgIf],
  templateUrl: './board.component.html',
  styleUrl: './board.component.css'
})

export class BoardComponent {
  @Input() board: any = {}
  @Input() numPlayers: number = 4

  getColumns() {
    const columns: any = [
      { value: 2, chips: {}},
      { value: 3, chips: {}},
      { value: 4, chips: {}},
      { value: 5, chips: {}},
      { value: 6, chips: {}},
      { value: 7, chips: {}},
      { value: 8, chips: {}},
      { value: 9, chips: {}},
      { value: 10, chips: {}},
      { value: 11, chips: {}},
      { value: 12, chips: {}}
    ];

    (this.board.Players || []).forEach((player: any, index: number) => {
      if (index != this.board.CurrentPlayerIndex) {
        for (const key in player.ColumnValues) {
          if (Object.prototype.hasOwnProperty.call(player.ColumnValues, key)) {
            columns[Number(key) - 2].chips[player.Color] = player.ColumnValues[key];
          }
        }
      }
    })

    if (this.board.Players) {
      const player = this.board.Players[this.board.CurrentPlayerIndex]
      for (const key in player.ColumnValues) {
        if (Object.prototype.hasOwnProperty.call(player.ColumnValues, key)) {
          columns[Number(key) - 2].chips[player.Color] = player.ColumnValues[key];
        }
      }
    }

    const currentTurnValues = this.board.CurrentTurnColumns || {}
    for (const key in currentTurnValues) {
      if (Object.prototype.hasOwnProperty.call(currentTurnValues, key)) {
        columns[Number(key)-2].chips["black"] = currentTurnValues[key];
      }
    }
    return columns;
  }
}

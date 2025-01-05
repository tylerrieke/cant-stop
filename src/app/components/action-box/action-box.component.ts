import {Component, EventEmitter, Input, Output} from '@angular/core';
import {DiceBoxComponent} from '../dice-box/dice-box.component';
import {OptionBoxComponent} from '../option-box/option-box.component';
import {NgClass, NgFor, NgIf} from '@angular/common';
import {TurnInfoBoxComponent} from '../turn-info-box/turn-info-box.component';
import {CsServiceService} from '../../services/cs-service.service';
import {ChipComponent} from '../chip/chip.component';

@Component({
  selector: '[app-action-box]',
  imports: [TurnInfoBoxComponent, DiceBoxComponent, OptionBoxComponent, ChipComponent, NgFor, NgIf, NgClass],
  templateUrl: './action-box.component.html',
  styleUrl: './action-box.component.css'
})
export class ActionBoxComponent {
  @Input() dice :any = []
  @Input() phase :string = "TURN_START"
  @Input() options :any[][][] = []
  @Input() currentPlayerColor :string = "blue"
  @Input() turnNumber :number = 0

  @Output() boardUpdated = new EventEmitter<any>();

  constructor(private css: CsServiceService) {}

  canRoll() {
    return this.phase == "TURN_START" || this.phase == "ROLL_OR_STOP"
  }

  rollDice() {
    this.css.roll().subscribe((data) => {
      this.handleBoardUpdate(data)
    })
  }

  canStop() {
    return this.phase == "ROLL_OR_STOP"
  }

  stopTurn() {
    this.css.stop().subscribe((data) => {
      this.handleBoardUpdate(data)
    })
  }

  selectOption(option: any) {
    this.css.selectOption(option).subscribe((data) => {
      this.handleBoardUpdate(data)
    })
  }

  canSwitchPlayers() {
    return this.phase == "BUST"
  }

  switchPlayers() {
    this.css.nextPlayer().subscribe((data) => {
      this.handleBoardUpdate(data)
    })
  }

  isBusted() {
    return this.phase == "BUST"
  }

  isGameOver() {
    return this.phase == "GAME_OVER" || !this.dice || this.dice.length == 0
  }

  createGame(numPlayers: any) {
    this.css.startGame(numPlayers).subscribe((data) => {
      this.handleBoardUpdate(data)
    })
  }

  handleBoardUpdate(boardUpdate: any) {
    this.boardUpdated.emit(boardUpdate);
  }
}

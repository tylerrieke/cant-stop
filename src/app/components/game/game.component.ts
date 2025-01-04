import {Component} from '@angular/core';
import {BoardComponent} from '../board/board.component';
import {ActionBoxComponent} from '../action-box/action-box.component';
import {CsServiceService} from '../../services/cs-service.service';

@Component({
  selector: 'app-game',
  imports: [BoardComponent, ActionBoxComponent],
  templateUrl: './game.component.html',
  styleUrl: './game.component.css'
})
export class GameComponent {

  board: any = {}

  constructor(private css: CsServiceService) {}

  handleBoardUpdate(boardUpdate: any) {
    this.board = boardUpdate
  }

  ngOnInit() {
    this.css.getGame().subscribe(data => {
      if (data) {
        this.board = data
      } else {
        this.css.startGame(4).subscribe(data => this.board = data)
      }
    })
  }
}

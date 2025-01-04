import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {environment} from "../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class CsServiceService {

  constructor(private http: HttpClient) { }

  getTitle() {
    return this.http.get<any>(`${environment.serverUrl}/hello-world`)
  }

  getGame() {
    return this.http.get<any>(`${environment.serverUrl}/game`)
  }

  startGame(numPlayers: number) {
    return this.http.post<any>(`${environment.serverUrl}/game/start?numPlayers=${numPlayers}`, {})
  }

  roll() {
    return this.http.post<any>(`${environment.serverUrl}/game/roll`, {})
  }

  selectOption(option: any) {
    return this.http.post<any>(`${environment.serverUrl}/game/select`, option)
  }

  stop() {
    return this.http.post<any>(`${environment.serverUrl}/game/stop`, {})
  }

  nextPlayer() {
    return this.http.post<any>(`${environment.serverUrl}/game/player`, {})
  }
}

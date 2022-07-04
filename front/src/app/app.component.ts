import { Component } from '@angular/core';
import { NgForm } from '@angular/forms';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  constructor(private http: HttpClient) {}

  postLogin(username: string, password: string) {
      var body = JSON.stringify({username: username, password: password});
      var header = new HttpHeaders()
        .set("content-type", "application/json");

      return this.http.post<any>('/', body, {headers: header});
  }

  handleLoginErrror(error: HttpErrorResponse) {
    console.log(error.status);

    return throwError(() => new Error("shit"));
  }

  onSubmit(loginForm: NgForm) {

    this.postLogin(loginForm.value.username, loginForm.value.username)
        .subscribe({
          next(r) {
            console.log(r);
          },
          error(e) {
            console.log(e.status);
          },
          complete() {

          }
        });
  }

}
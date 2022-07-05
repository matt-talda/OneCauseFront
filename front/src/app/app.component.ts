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

  postLogin(username: string, password: string, token: string) {
      let body = JSON.stringify({username: username, password: password, token: token});
      let header = new HttpHeaders()
        .set("content-type", "application/json");

      return this.http.post<any>('/auth/login', body, {headers: header});
  }

  handleLoginErrror(error: HttpErrorResponse) {
    console.log(error.status);

    return throwError(() => new Error("shit"));
  }

  onSubmit(loginForm: NgForm) {

    this.postLogin(loginForm.value.username, loginForm.value.password, loginForm.value.token)
        .subscribe({
          next(r) {
            window.location.href = 'http://onecause.com';
          },
          error(e) {
            console.log(e.status);
          },
          complete() {

          }
        });
  }

}
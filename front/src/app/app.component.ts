import { Component, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  @ViewChild('loginForm') ngForm!: NgForm

  formChangesSubscription: any;
  errorMsg: string|null = null;

  constructor(private http: HttpClient) {}

  postLogin(username: string, password: string, token: string) {
      let body = JSON.stringify({username: username, password: password, token: token});
      let header = new HttpHeaders()
        .set("content-type", "application/json");

      return this.http.post<any>('/auth/login', body, {headers: header});
  }

  onSubmit(loginForm: NgForm) {

    if (loginForm.invalid) {
      this.errorMsg = 'Error: invalid username, password, or token'
      return;
    }

    const auth: any = this.postLogin(loginForm.value.username, loginForm.value.password, loginForm.value.token)
        .subscribe({
          next: r => window.location.href = 'http://onecause.com',
          error: e => this.errorMsg = `ERROR: ${e.error.error}`,
          complete: () => auth.unsubscribe()
        });
  }

  ngAfterViewInit() {
    this.formChangesSubscription = this.ngForm.form.valueChanges.subscribe(() => { this.errorMsg = null})
  }

  ngOnDestroy() {
    this.formChangesSubscription.unsubscribe();
  }
}
import { Component, OnInit } from '@angular/core';
import { UserService } from 'src/app/services/user.service';
import { FormControl } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {
  username = new FormControl('');
  password = new FormControl('');
  loginErr: string = '';

  constructor(
    private userService: UserService,
    private router: Router,
  ) {}

  ngOnInit(): void {}

  onSend() {
    this.loginErr = '';
    this.userService.login({
      name: this.username.value,
      password: this.password.value,
    }).then(() => this.router.navigate(['/panel']))
      .catch((err) => {
        this.loginErr = err.reason;
      });
  }

}

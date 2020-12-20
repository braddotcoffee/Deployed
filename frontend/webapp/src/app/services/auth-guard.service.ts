import { Injectable } from '@angular/core';
import { Router, CanActivate } from '@angular/router';
import { AngularFireAuth } from '@angular/fire/auth';
import * as firebase from 'firebase/app';
import 'firebase/auth';
import { from, Observable } from 'rxjs';
import { concatMap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthGuardService {

  constructor(private router: Router, private auth: AngularFireAuth) { }

  getAuthenticationToken(): Observable<string> {
    return from(this.auth.currentUser).pipe(
      concatMap(user => {
        if (user) { return user.getIdToken(); }
        return '';
      })
    );
  }

  login(): void {
    this.auth.signInWithPopup(new firebase.default.auth.GoogleAuthProvider()).then(() => {
      this.router.navigate(['/preview-deployments']);
    });
  }
}

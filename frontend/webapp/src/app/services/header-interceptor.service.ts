import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { concatMap } from 'rxjs/operators';
import { AuthGuardService } from './auth-guard.service';

@Injectable({
  providedIn: 'root'
})
export class HeaderInterceptorService implements HttpInterceptor {

  constructor(private auth: AuthGuardService) { }
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return this.auth.getAuthenticationToken().pipe(
      concatMap(token => {
        const clonedRequest = req.clone({
          headers: req.headers.set('Authorization', token)
        });
        return next.handle(clonedRequest);
      })
    );
  }
}

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HelloWorldService {

  constructor(private http: HttpClient) { }

  getTitle(): Observable<any> {
    return this.http.get(`${environment.backendUrl}/hello-world`);
  }
}

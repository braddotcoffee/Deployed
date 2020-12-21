import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class NetworkConfigService {

  constructor(private http: HttpClient) { }

  updateNetworkConfigs(): void {
    this.http.get(`${environment.backendUrl}/update-network-config`).subscribe();
  }
}

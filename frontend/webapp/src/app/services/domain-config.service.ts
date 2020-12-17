import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { DomainConfiguration } from 'src/types/domainConfiguration_pb';

@Injectable({
  providedIn: 'root'
})
export class DomainConfigService {

  constructor(private http: HttpClient) { }

  addDomainConfig(name: string, domainConfig: DomainConfiguration): void {
    this.http.post(`${environment.backendUrl}/add-domain-config?name=${name}`, domainConfig.toObject()).subscribe();
  }
}

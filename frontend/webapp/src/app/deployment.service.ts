import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Deployment } from 'src/types/deployment_pb';
import { environment } from '../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class DeploymentService {

  constructor(private http: HttpClient) { }

  addDeployment(deployment: Deployment): void {
    console.log(deployment.serializeBinary());
    this.http.post(`${environment.backendUrl}/add-deployment`, deployment.serializeBinary()).subscribe();
  }
}

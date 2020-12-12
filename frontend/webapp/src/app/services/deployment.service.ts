import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { Deployment } from 'src/types/deployment_pb';
import { environment } from '../../environments/environment';
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';
import { readableStatus } from '../../types/utils/deployment-utils';
import DeployStatus = Deployment.DeployStatus;

@Injectable({
  providedIn: 'root'
})
export class DeploymentService {

  constructor(private http: HttpClient) { }

  addDeployment(deployment: Deployment): void {
    this.http.post(`${environment.backendUrl}/add-deployment`, deployment.toObject()).subscribe();
  }

  getDeployments(): Observable<Deployment[]> {
    return this.http.get<any[]>(`${environment.backendUrl}/get-deployments`).pipe(
      map(objects => {
        return objects.map(this.deploymentFromObject);
      }));
  }

  getDeployment(name: string): Observable<Deployment> {
    return this.http.get<any>(`${environment.backendUrl}/get-deployment?name=${name}`).pipe(
      map(this.deploymentFromObject)
    );
  }

  deployNewVersion(name: string): Observable<any> {
    return this.http.get(`${environment.backendUrl}/deploy-new-version?name=${name}`);
  }

  readableStatus(deployment: Deployment): string {
    return readableStatus(deployment.getStatus());
  }

  getGithubUrl(deployment: Deployment): string {
    const repo = deployment.getRepository();
    const sliceStart = repo.indexOf(':') + 1;
    return 'https://github.com/' + repo.slice(sliceStart, repo.length - 4);
  }

  getIcon(deployment: Deployment): string {
    const status = deployment.getStatus();
    if (status === DeployStatus.COMPLETE) {
      return 'done';
    }

    if (status === DeployStatus.ERROR) {
      return 'error';
    }

    if (status === DeployStatus.IN_PROGRESS) {
      return 'sync';
    }

    if (status === DeployStatus.NOT_STARTED) {
      return 'pause';
    }
    return 'forward';
  }

  private deploymentFromObject(object: any): Deployment {
    const deployment = new Deployment();
    if (object.last_deploy !== undefined) {
      const lastDeploy = new Timestamp();
      lastDeploy.setSeconds(object.last_deploy.seconds);
      lastDeploy.setNanos(object.last_deploy.nanos);
      deployment.setLastDeploy(lastDeploy);
    }
    deployment.setName(object.name);
    deployment.setRepository(object.repository);
    deployment.setDockerfile(object.dockerfile);
    deployment.setDomain(object.domain);
    deployment.setStatus(object.status);
    deployment.setCommit(object.commit);
    return deployment;
  }
}

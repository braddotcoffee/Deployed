import { Component, OnInit } from '@angular/core';
import { Deployment } from 'src/types/deployment_pb';
import { DeploymentService } from './deployment.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent implements OnInit {
  title = 'webapp';
  constructor(private deploymentService: DeploymentService) { }

  ngOnInit(): void {
    const deployment = new Deployment();

    deployment.setName('Test Deployment From Angular');
    this.deploymentService.addDeployment(deployment);
  }
}

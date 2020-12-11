import { Component, OnInit } from '@angular/core';
import { Deployment } from 'src/types/deployment_pb';
import { DeploymentService } from '../services/deployment.service';
import { readableStatus } from '../../types/utils/deployment-utils';
import DeployStatus = Deployment.DeployStatus;
import { map } from 'rxjs/operators';

@Component({
  selector: 'app-preview-deployments',
  templateUrl: './preview-deployments.component.html',
  styleUrls: ['./preview-deployments.component.scss']
})
export class PreviewDeploymentsComponent implements OnInit {
  deployments: Deployment[] = [];
  githubUrls: Map<string, string> = new Map();
  icons: Map<string, string> = new Map();

  constructor(private deploymentService: DeploymentService) { }

  ngOnInit(): void {
    this.deploymentService.getDeployments().subscribe((deployments: Deployment[]) => {
      this.deployments = deployments;
      this.deployments.forEach(deployment => {
        this.githubUrls.set(deployment.getName(), this.deploymentService.getGithubUrl(deployment));
        this.icons.set(deployment.getName(), this.deploymentService.getIcon(deployment));
      });
    });
  }
}

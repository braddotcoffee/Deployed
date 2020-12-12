import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Deployment } from 'src/types/deployment_pb';
import { DeploymentService } from '../services/deployment.service';

@Component({
  selector: 'app-deployment-detail',
  templateUrl: './deployment-detail.component.html',
  styleUrls: ['./deployment-detail.component.scss']
})
export class DeploymentDetailComponent implements OnInit {
  deployment: Deployment | undefined;
  githubUrl = '';
  icon = '';
  complete = false;
  errorMessage = '';

  constructor(private deploymentService: DeploymentService, private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.getDeployment();
  }

  getDeployment(): void {
    const name = this.route.snapshot.paramMap.get('name');
    if (name === null) {
      this.complete = true;
      return;
    }
    this.deploymentService.getDeployment(name).subscribe(deployment => {
      this.deployment = deployment;
      this.githubUrl = this.deploymentService.getGithubUrl(deployment);
      this.icon = this.deploymentService.getIcon(deployment);
      this.complete = true;
    }, _ => {
      this.complete = true;
      this.errorMessage = 'Unable to find deployment';
    });
  }

  deployNewVersion(): void {
    if (this.deployment === undefined) { return; }
    this.deploymentService.deployNewVersion(this.deployment.getName()).subscribe(
      _ => { },
      err => {
        console.log(err);
        this.errorMessage = 'Unable to deploy new version';
      });
  }
}

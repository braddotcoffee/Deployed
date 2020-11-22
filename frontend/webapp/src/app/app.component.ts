import { Component, OnInit } from '@angular/core';
import { Deployment } from 'src/types/deployment_pb';
import { DeploymentService } from './deployment.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'webapp';
  constructor(private deploymentService: DeploymentService) { }

  ngOnInit(): void {
  }
}

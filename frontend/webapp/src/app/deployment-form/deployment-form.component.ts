import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { StepperFormInput, StepperFormStep } from 'src/app/forms/stepper-forms/stepper-form-types';
import { Deployment } from 'src/types/deployment_pb';
import { DeploymentService } from '../services/deployment.service';

@Component({
  selector: 'app-deployment-form',
  templateUrl: './deployment-form.component.html',
  styleUrls: ['./deployment-form.component.scss']
})
export class DeploymentFormComponent implements OnInit {
  @Input() deployment: Deployment = new Deployment();
  nameFormGroup: FormGroup = new FormGroup({});
  repositoryFormGroup: FormGroup = new FormGroup({});
  dockerfileFormGroup: FormGroup = new FormGroup({});
  domainFormGroup: FormGroup = new FormGroup({});

  steps: StepperFormStep[] = [];

  constructor(private formBuilder: FormBuilder, private deploymentService: DeploymentService, private router: Router) { }

  ngOnInit(): void {
    this.nameFormGroup = this.formBuilder.group({
      name: [this.deployment.getName(), Validators.required]
    });
    this.repositoryFormGroup = this.formBuilder.group({
      repository: [this.deployment.getRepository(), Validators.required]
    });
    this.dockerfileFormGroup = this.formBuilder.group({
      dockerfile: [this.deployment.getDockerfile(), Validators.required]
    });
    this.domainFormGroup = this.formBuilder.group({
      domain: [this.deployment.getDomain(), Validators.required]
    });
    this.steps = [
      {
        formGroup: this.nameFormGroup,
        stepLabel: 'Name your application',
        inputs: [{
          label: 'Application Name',
          controlName: 'name',
          placeholder: 'AwesomeNewApp'
        }]
      },
      {
        formGroup: this.repositoryFormGroup,
        stepLabel: 'Link your repository',
        inputs: [{
          label: 'Repository Link',
          controlName: 'repository',
          placeholder: 'https://your-repository.com'
        }]
      },
      {
        formGroup: this.dockerfileFormGroup,
        stepLabel: 'Find your Dockerfile',
        inputs: [{
          label: 'Dockerfile path',
          controlName: 'dockerfile',
          placeholder: '/path/to/Dockerfile'
        }]
      },
      {
        formGroup: this.domainFormGroup,
        stepLabel: 'Host your project',
        inputs: [{
          label: 'Host domain',
          controlName: 'domain',
          placeholder: 'https://some.domain.com'
        }]
      }
    ];
  }

  onSubmit(): void {
    this.deployment.setName(this.nameFormGroup.get('name')?.value);
    this.deployment.setRepository(this.repositoryFormGroup.get('repository')?.value);
    this.deployment.setDockerfile(this.dockerfileFormGroup.get('dockerfile')?.value);
    this.deployment.setDomain(this.domainFormGroup.get('domain')?.value);
    this.deploymentService.addDeployment(this.deployment);
    this.router.navigate(['/preview-deployments']);
  }
}

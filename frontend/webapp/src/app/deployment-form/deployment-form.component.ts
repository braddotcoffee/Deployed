import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { StepperFormInput, StepperFormStep } from 'src/app/forms/stepper-forms/stepper-form-types';

@Component({
  selector: 'app-deployment-form',
  templateUrl: './deployment-form.component.html',
  styleUrls: ['./deployment-form.component.scss']
})
export class DeploymentFormComponent implements OnInit {
  nameFormGroup: FormGroup = new FormGroup({});
  repositoryFormGroup: FormGroup = new FormGroup({});
  dockerfileFormGroup: FormGroup = new FormGroup({});

  steps: StepperFormStep[] = [];

  constructor(private formBuilder: FormBuilder) { }

  ngOnInit(): void {
    this.nameFormGroup = this.formBuilder.group({
      name: ['', Validators.required]
    });
    this.repositoryFormGroup = this.formBuilder.group({
      repository: ['', Validators.required]
    });
    this.dockerfileFormGroup = this.formBuilder.group({
      dockerfile: ['', Validators.required]
    });
    this.steps = [
      {
        formGroup: this.nameFormGroup,
        stepLabel: 'Name your application',
        inputs: [{
          label: 'Application Name',
          controlName: 'name'
        }]
      },
      {
        formGroup: this.repositoryFormGroup,
        stepLabel: 'Link your repository',
        inputs: [{
          label: 'Repository Link',
          controlName: 'repository'
        }]
      },
      {
        formGroup: this.dockerfileFormGroup,
        stepLabel: 'Find your Dockerfile',
        inputs: [{
          label: 'Dockerfile path',
          controlName: 'dockerfile'
        }]
      }
    ];
  }

  onSubmit(): void {
    console.log(this.nameFormGroup.get('name')?.value);
  }
}

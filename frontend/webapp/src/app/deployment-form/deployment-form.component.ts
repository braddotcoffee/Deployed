import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-deployment-form',
  templateUrl: './deployment-form.component.html',
  styleUrls: ['./deployment-form.component.scss']
})
export class DeploymentFormComponent implements OnInit {
  nameFormGroup: FormGroup = new FormGroup({});
  repositoryFormGroup: FormGroup = new FormGroup({});
  dockerfileFormGroup: FormGroup = new FormGroup({});

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
  }

  onSubmit(): void {

  }

}

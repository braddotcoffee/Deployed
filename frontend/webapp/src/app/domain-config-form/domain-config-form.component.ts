import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { DomainConfiguration } from 'src/types/domainConfiguration_pb';
import { StepperFormStep } from '../forms/stepper-forms/stepper-form-types';
import { DomainConfigService } from '../services/domain-config.service';

@Component({
  selector: 'app-domain-config-form',
  templateUrl: './domain-config-form.component.html',
  styleUrls: ['./domain-config-form.component.scss']
})
export class DomainConfigFormComponent implements OnInit {
  @Input() name = '';
  @Input() domainConfig: DomainConfiguration = new DomainConfiguration();
  nameFormGroup: FormGroup = new FormGroup({});
  domainFormGroup: FormGroup = new FormGroup({});
  portFormGroup: FormGroup = new FormGroup({});

  steps: StepperFormStep[] = [];

  constructor(private formBuilder: FormBuilder, private router: Router, private domainConfigService: DomainConfigService) { }

  ngOnInit(): void {
    this.nameFormGroup = this.formBuilder.group({
      name: [this.name, Validators.required]
    })
    this.domainFormGroup = this.formBuilder.group({
      domain: [this.domainConfig.getDomain(), Validators.required]
    });
    this.portFormGroup = this.formBuilder.group({
      port: [this.domainConfig.getPort(), Validators.required]
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
        formGroup: this.domainFormGroup,
        stepLabel: 'Host your project',
        inputs: [{
          label: 'Host domain',
          controlName: 'domain',
          placeholder: 'https://some.domain.com'
        }]
      },
      {
        formGroup: this.portFormGroup,
        stepLabel: 'Port forward your project',
        inputs: [{
          label: 'Host port',
          controlName: 'port',
          placeholder: '3000'
        }]
      }
    ];
  }

  onSubmit(): void {
    this.domainConfig.setDomain(this.domainFormGroup.get('domain')?.value);
    this.domainConfig.setPort(this.portFormGroup.get('port')?.value);
    this.name = this.nameFormGroup.get('name')?.value;
    this.domainConfigService.addDomainConfig(this.name, this.domainConfig);
    this.router.navigate(['/preview-deployments']);
  }

}

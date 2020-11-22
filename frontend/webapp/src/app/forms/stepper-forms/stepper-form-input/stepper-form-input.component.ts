import { Component, Input, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';

@Component({
  selector: 'app-stepper-form-input',
  templateUrl: './stepper-form-input.component.html',
  styleUrls: ['./stepper-form-input.component.scss']
})
export class StepperFormInputComponent implements OnInit {
  @Input() controlName: string | undefined;
  @Input() label: string | undefined;
  @Input() formGroup: FormGroup = new FormGroup({});

  constructor() { }

  ngOnInit(): void {
  }

}

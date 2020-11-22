import { Output } from '@angular/core';
import { Component, EventEmitter, Input, OnInit } from '@angular/core';
import { StepperFormStep } from '../stepper-form-types';

@Component({
  selector: 'app-stepper-form',
  templateUrl: './stepper-form.component.html',
  styleUrls: ['./stepper-form.component.scss']
})
export class StepperFormComponent implements OnInit {
  @Input() steps: StepperFormStep[] = [];
  @Output() submitEvent: EventEmitter<void> = new EventEmitter();

  constructor() { }

  ngOnInit(): void {
  }

  emitSubmit(): void {
    this.submitEvent.emit();
  }
}

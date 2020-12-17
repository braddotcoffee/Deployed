import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DeploymentFormComponent } from './deployment-form/deployment-form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatStepperModule } from '@angular/material/stepper';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDividerModule } from '@angular/material/divider';
import { StepperFormInputComponent } from './forms/stepper-forms/stepper-form-input/stepper-form-input.component';
import { StepperFormStepComponent } from './forms/stepper-forms/stepper-form-step/stepper-form-step.component';
import { StepperFormComponent } from './forms/stepper-forms/stepper-form/stepper-form.component';
import { MatIconModule } from '@angular/material/icon';
import { MatTableModule } from '@angular/material/table';
import { ToolbarComponent } from './toolbar/toolbar.component';
import { MatToolbarModule } from '@angular/material/toolbar';
import { PreviewDeploymentsComponent } from './preview-deployments/preview-deployments.component';
import { DeploymentDetailComponent } from './deployment-detail/deployment-detail.component';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatMenuModule } from '@angular/material/menu';
import { DomainConfigFormComponent } from './domain-config-form/domain-config-form.component';



@NgModule({
  declarations: [
    AppComponent,
    DeploymentFormComponent,
    StepperFormInputComponent,
    StepperFormStepComponent,
    StepperFormComponent,
    ToolbarComponent,
    PreviewDeploymentsComponent,
    DeploymentDetailComponent,
    DomainConfigFormComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    ReactiveFormsModule,
    BrowserAnimationsModule,
    MatInputModule,
    MatFormFieldModule,
    MatStepperModule,
    MatButtonModule,
    MatCardModule,
    MatDividerModule,
    MatIconModule,
    MatTableModule,
    MatToolbarModule,
    MatProgressSpinnerModule,
    MatMenuModule,
  ],
  providers: [HttpClientModule],
  bootstrap: [AppComponent]
})
export class AppModule { }

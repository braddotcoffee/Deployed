import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DeploymentFormComponent } from './deployment-form/deployment-form.component';

const routes: Routes = [
  { path: 'new-deployment', component: DeploymentFormComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

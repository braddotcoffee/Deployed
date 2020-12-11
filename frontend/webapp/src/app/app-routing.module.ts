import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DeploymentDetailComponent } from './deployment-detail/deployment-detail.component';
import { DeploymentFormComponent } from './deployment-form/deployment-form.component';
import { PreviewDeploymentsComponent } from './preview-deployments/preview-deployments.component';

const routes: Routes = [
  { path: 'new-deployment', component: DeploymentFormComponent },
  { path: 'preview-deployments', component: PreviewDeploymentsComponent },
  { path: 'deployment/:name', component: DeploymentDetailComponent },
  { path: '', redirectTo: '/preview-deployments', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

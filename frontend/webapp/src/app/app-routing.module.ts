import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DeploymentDetailComponent } from './deployment-detail/deployment-detail.component';
import { DeploymentFormComponent } from './deployment-form/deployment-form.component';
import { DomainConfigFormComponent } from './domain-config-form/domain-config-form.component';
import { PreviewDeploymentsComponent } from './preview-deployments/preview-deployments.component';
import { AuthGuardService } from './services/auth-guard.service';
import { AuthPipeGenerator, canActivate, redirectUnauthorizedTo } from '@angular/fire/auth-guard';
import { LoginComponent } from './login/login.component';

const redirectUnauthorizedToLogin: AuthPipeGenerator = () => redirectUnauthorizedTo(['login']);

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'new-deployment', component: DeploymentFormComponent, ...canActivate(redirectUnauthorizedToLogin) },
  { path: 'new-domain-config', component: DomainConfigFormComponent, ...canActivate(redirectUnauthorizedToLogin) },
  { path: 'preview-deployments', component: PreviewDeploymentsComponent, ...canActivate(redirectUnauthorizedToLogin) },
  { path: 'deployment/:name', component: DeploymentDetailComponent, ...canActivate(redirectUnauthorizedToLogin) },
  { path: '', redirectTo: '/preview-deployments', pathMatch: 'full', ...canActivate(redirectUnauthorizedToLogin) }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

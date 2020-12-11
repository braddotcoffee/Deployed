import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PreviewDeploymentsComponent } from './preview-deployments.component';

describe('PreviewDeploymentsComponent', () => {
  let component: PreviewDeploymentsComponent;
  let fixture: ComponentFixture<PreviewDeploymentsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PreviewDeploymentsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PreviewDeploymentsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

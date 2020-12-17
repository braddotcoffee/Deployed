import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DomainConfigFormComponent } from './domain-config-form.component';

describe('DomainConfigFormComponent', () => {
  let component: DomainConfigFormComponent;
  let fixture: ComponentFixture<DomainConfigFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DomainConfigFormComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DomainConfigFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

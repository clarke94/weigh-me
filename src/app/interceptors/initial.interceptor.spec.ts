import { TestBed } from '@angular/core/testing';

import { InitialInterceptor } from './initial.interceptor';

describe('InitialInterceptor', () => {
  beforeEach(() => TestBed.configureTestingModule({
    providers: [
      InitialInterceptor
      ]
  }));

  it('should be created', () => {
    const interceptor: InitialInterceptor = TestBed.inject(InitialInterceptor);
    expect(interceptor).toBeTruthy();
  });
});

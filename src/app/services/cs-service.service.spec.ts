import { TestBed } from '@angular/core/testing';

import { CsServiceService } from './cs-service.service';

describe('CsServiceService', () => {
  let service: CsServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CsServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

import { Injectable } from '@angular/core';
import {
    HttpRequest,
    HttpHandler,
    HttpEvent,
    HttpInterceptor
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { Store } from '@ngrx/store';
import { active, inactive } from '../store/actions/initial.actions';
import { finalize, debounce } from 'rxjs/operators';

@Injectable()
export class InitialInterceptor implements HttpInterceptor {

    constructor(
        private store: Store<{ initial: boolean }>
    ) { }

    intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
        this.store.dispatch(active());
        return next.handle(request).pipe(
            finalize(() => this.store.dispatch(inactive()))
        );
    }
}

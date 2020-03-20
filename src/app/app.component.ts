import { Component, OnInit } from '@angular/core';
import { Store, select } from '@ngrx/store';
import { Observable } from 'rxjs';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
    loading: Observable<boolean>;

    constructor(
        private store: Store<{ initial: boolean }>
    ) { }

    ngOnInit(): void {
        this.loading = this.store.pipe(select('initial'));
    }
}

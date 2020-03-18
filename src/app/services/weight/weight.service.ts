import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class WeightService {

    constructor(
        private http: HttpClient
    ) { }

    getAllWeights(): Observable<any> {
        return this.http.get('/api/v1/weights');
    }

    addWeight(data): Observable<any> {
        return this.http.post('/api/v1/weight', data);
    }
}

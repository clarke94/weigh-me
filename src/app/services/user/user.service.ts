import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
    providedIn: 'root'
})
export class UserService {

    constructor(
        private http: HttpClient
    ) { }

    getAllWeights(): Observable<any> {
        return this.http.get('/weights').pipe(
            map(res => {
                let series: any[] = [];
                for (const [_, val] of Object.entries(res)) {
                    series.push(val);
                }

                const data = [{
                    "name": "Weights",
                    "series": series
                }];
                return data;
            })
        );
    }

    addWeight(data): Observable<any> {
        return this.http.post('/weights', data);
    }

    updateWeight(data): Observable<any> {
        return this.http.put(`/weights/${data.id}`, data);
    }

    deleteWeight(id): Observable<any> {
        return this.http.delete(`/weights/${id}`);
    }
}

import { Injectable } from '@angular/core';

import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class SelectoroperatorsService {

  constructor(private http: HttpClient) { }

  // Method using subscribe
  querySelectorOperators(featurelayer:string,attribute:string): Observable<any> {
    return this.http.post<any>('http://goapp.local:80/map/featureoperatures',{ 
      selectedLayer: featurelayer,
      selectedAttribute:attribute
    },);
  }
}

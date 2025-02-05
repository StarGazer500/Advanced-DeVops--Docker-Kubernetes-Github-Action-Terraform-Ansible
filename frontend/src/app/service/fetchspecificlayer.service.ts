import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class FetchspecificlayerService {

  constructor(private http: HttpClient) { }

  // Method using subscribe
  querySpecficFeatureLayer(featurelayer:string): Observable<any> {
    return this.http.post<any>('http://goapp.local:80/map/searchbyfeaturelayer',{ 
      selectedLayer: featurelayer
    },);
  }
}


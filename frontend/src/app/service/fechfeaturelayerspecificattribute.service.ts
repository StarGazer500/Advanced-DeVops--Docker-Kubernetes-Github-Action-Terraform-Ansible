import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FechfeaturelayerspecificattributeService {

  constructor(private http: HttpClient) { }

  // Method using subscribe
  querySpecficFeatureLayerbyAttribute(featurelayer:string,attribute:string): Observable<any> {
    return this.http.post<any>('http://goapp.local:80/map/searchbycolumn',{ 
      selectedLayer: featurelayer,
      selectedAttribute:attribute

    },);
  }
}

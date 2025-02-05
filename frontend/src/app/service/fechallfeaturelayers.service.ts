import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FechallfeaturelayersService {

  constructor(private http: HttpClient) { }

  // Method using subscribe
  queryAllFeatureLayers(): Observable<any> {
    return this.http.post<any>('http://goapp.local:80/map/searchallfeaturelayersdata',{});
  }
}

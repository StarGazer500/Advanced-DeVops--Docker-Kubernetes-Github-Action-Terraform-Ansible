import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SimplesearchService {

  constructor(private http: HttpClient) { }

  // Method using subscribe
  querySimpleSearchValue(value:string): Observable<any> {
    return this.http.post<any>('http://goapp.local:80/map/simplesearch',{ 
      searchValue: value
    },);
  }
}

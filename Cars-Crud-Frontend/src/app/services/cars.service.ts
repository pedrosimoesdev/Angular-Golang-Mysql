import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class CarsService {


  private SERVER = "http://localhost:8080/";

  constructor(private httpClient: HttpClient) { }


  public getCars() { 
    return this.httpClient.get(this.SERVER); 
    } 

    // public post(url: string, data: any, options?: any) { 
    //   return this.http.post(url, data, options); 
    //   } 
}

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { catchError, map, Observable, retry, throwError } from 'rxjs';
import { Cars } from '../model/cars'



@Injectable({
  providedIn: 'root'
})
export class CarsService {


  private SERVER = "http://localhost:8080/";

  constructor(private httpClient: HttpClient) { }


    getCars(): Observable<Cars> {
      return this.httpClient
        .get<Cars>(this.SERVER )
        .pipe(retry(1), catchError(this.handleError));
    }

  // Error handling
  handleError(error: any) {
    let errorMessage = '';
    if (error.error instanceof ErrorEvent) {
      // Get client-side error
      errorMessage = error.error.message;
    } else {
      // Get server-side error
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    window.alert(errorMessage);
    return throwError(() => {
      return errorMessage;
    });
  }

    // public post(url: string, data: any, options?: any) { 
    //   return this.http.post(url, data, options); 
    //   } 
}

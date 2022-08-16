import { NgModule,OnInit } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CarsComponent } from './cars/cars.component';
import { CarsEditComponent } from './cars/cars-edit/cars-edit.component';
import { CarsViewComponent } from './cars/cars-view/cars-view.component';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    CarsComponent,
    AppComponent,
    CarsEditComponent,
    CarsViewComponent,
 
  
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [CarsComponent],
  bootstrap: [AppComponent]
})

export class AppModule { 
}




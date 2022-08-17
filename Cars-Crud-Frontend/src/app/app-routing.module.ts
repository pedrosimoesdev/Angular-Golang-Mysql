import { NgModule,OnInit } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CarsComponent } from './cars/cars.component';
import { CarsCreateComponent } from './cars/cars-create/cars-create.component';

const routes: Routes = [ 
  { path: "cars", component: CarsComponent },
  { path: "cars/create", component: CarsCreateComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

import { Component, OnInit } from '@angular/core';
import {CarsService} from '../services/cars.service'


@Component({
  selector: 'app-cars',
  templateUrl: './cars.component.html',
  styleUrls: ['./cars.component.scss']
})
export class CarsComponent implements OnInit {

  constructor(private carService: CarsService){
   
  }
  title = 'Show Cars'
  cars = []



  ngOnInit(): void {
   
      this.getCars()
    

  }

    getCars(){
   
  this.carService.getCars().subscribe(res => {
    console.log(res);

    });
}



}

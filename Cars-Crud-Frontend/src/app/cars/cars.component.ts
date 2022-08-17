import { Component, OnInit } from '@angular/core';
import { map } from 'rxjs';
import {CarsService} from '../services/cars.service'
import {Cars} from '../model/cars'



@Component({
  selector: 'app-cars',
  templateUrl: './cars.component.html',
  styleUrls: ['./cars.component.scss']
})
export class CarsComponent implements OnInit {


  constructor(private carService: CarsService){
   
  }
  title = 'Show Cars'
  Cars : any;
  
  ngOnInit(): void {
   
     this.getCars()
    
  }

  getCars(){
   this.carService.getCars().subscribe(result => {
     console.log(result)
    this.Cars=result;
   })
  }
}

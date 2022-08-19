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
    //call services to get all records
   this.carService.getCars().subscribe(result => {
    this.Cars=result;
   })
  }

  DeleteRecords(id: any){
    console.log(id);
    if(confirm("Are you sure to delete " )) {
      //call services to get all records
     this.carService.deleteCar(id).subscribe(result => {
      alert(result)
       this.getCars()
     })
    }
  }
}

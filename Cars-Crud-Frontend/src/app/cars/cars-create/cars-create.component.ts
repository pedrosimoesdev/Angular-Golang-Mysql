import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import {CarsService} from '../../services/cars.service'


@Component({
  selector: 'app-cars-create',
  templateUrl: './cars-create.component.html',
  styleUrls: ['./cars-create.component.scss']
})
export class CarsCreateComponent implements OnInit {

 
  checkoutForm = this.formBuilder.group({
    Name: ['', Validators.required],
    Model: ['', Validators.required],
    Year: ['', Validators.required],
    
  });


  constructor(
    private formBuilder: FormBuilder,
    private carService: CarsService
  ) { 
   
  }

  ngOnInit(): void {
 
  }

  onSubmit(){

    let name = this.checkoutForm.value.Name;
    let model = this.checkoutForm.value.Model;
    let year = this.checkoutForm.value.Year;

    let car = [name, model, year];
    

    if(name === '' || model === '' || year ===  '' ){
      alert('Please fill all inputs')
      return;
    } 

  
    this.carService.createCar(car).subscribe(result => {
      console.log(result)
     
    })

   }
  
}

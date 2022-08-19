import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
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
    private carService: CarsService,
    private router: Router
  ) { 
   
  }

  ngOnInit(): void {
 
  }

  onSubmit(){

    let name = this.checkoutForm.value.Name;
    let model = this.checkoutForm.value.Model;
    let year = this.checkoutForm.value.Year;

    console.log(name)

    let car = { name, model, year};
    

    if(name === '' || model === '' || year ===  '' ){
      alert('Please fill all inputs')
      return;
    } 

  
    this.carService.createCar(car).subscribe(result => {
     alert(result)
    this.checkoutForm.reset();
    this.router.navigate(['/cars']);
     
    })

   }
  
}

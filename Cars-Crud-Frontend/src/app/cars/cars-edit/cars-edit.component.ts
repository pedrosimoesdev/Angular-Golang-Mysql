import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { FormBuilder, Validators } from '@angular/forms';
import {CarsService} from '../../services/cars.service'


@Component({
  selector: 'app-cars-edit',
  templateUrl: './cars-edit.component.html',
  styleUrls: ['./cars-edit.component.scss']
})
export class CarsEditComponent implements OnInit {

  id: any;
  name: any;
  model: any;
  year: any;

  checkoutForm = this.formBuilder.group({
    Id: ['', Validators.required],
    Name: ['', Validators.required],
    Model: ['', Validators.required],
    Year: ['', Validators.required],
    
  });
  constructor(private route: ActivatedRoute,
    private formBuilder: FormBuilder,
    private carService: CarsService,
    private router: Router) { }

  ngOnInit(): void {
   this.id = this.route.snapshot.paramMap.get('id')
   this.name = this.route.snapshot.paramMap.get('name')
   this.model = this.route.snapshot.paramMap.get('model')
   this.year = this.route.snapshot.paramMap.get('Year')

   

   this.checkoutForm.controls['Id'].patchValue(this.id);
   this.checkoutForm.controls['Name'].patchValue(this.name);
   this.checkoutForm.controls['Model'].patchValue(this.model);
   this.checkoutForm.controls['Year'].patchValue(this.year);
   
  }

  onSubmit(){

    let name = this.checkoutForm.value.Name;
    let model = this.checkoutForm.value.Model;
    let year = this.checkoutForm.value.Year;
    let id = this.checkoutForm.value.Id

    let car = { id ,name, model, year};

    console.log(car);
    
    if(name === '' || model === '' || year ===  '' ){
      alert('Please fill all inputs')
      return;
    } 
    //call service to insert valus of database
    this.carService.editCar(car).subscribe(result => {
     alert(result)
    this.checkoutForm.reset();
    this.router.navigate(['/cars']);
     
    })

   }

}

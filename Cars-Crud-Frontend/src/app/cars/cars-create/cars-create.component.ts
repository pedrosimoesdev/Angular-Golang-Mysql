import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';


@Component({
  selector: 'app-cars-create',
  templateUrl: './cars-create.component.html',
  styleUrls: ['./cars-create.component.scss']
})
export class CarsCreateComponent implements OnInit {

  checkoutForm = this.formBuilder.group({
    name: '',
    model: '',
    year: '',
  });


  constructor(
    private formBuilder: FormBuilder,
  ) { }

  ngOnInit(): void {
  }

  onSubmit(){
    console.log('teasting');
  }

}

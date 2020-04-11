import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {
    personalInfoForm: FormGroup;
    fitnessForm: FormGroup;

    constructor(
        private formBuilder: FormBuilder
    ) { }

    ngOnInit(): void {
        this.InitialisePersonalInfoForm();
        this.InitialiseFitnessForm();
    }

    InitialisePersonalInfoForm(): void {
        this.personalInfoForm = this.formBuilder.group({
            displayName: ['', Validators.compose([
                Validators.required,
                Validators.minLength(3),
            ])],
            gender: ['', Validators.compose([
                Validators.required,
                Validators.minLength(3),
            ])],
            dob: ['', Validators.required]
        });
    }

    InitialiseFitnessForm(): void {
        this.fitnessForm = this.formBuilder.group({
            height: ['', Validators.compose([
                Validators.required,
                Validators.min(59),
                Validators.max(300),
            ])],
        });
    }
}

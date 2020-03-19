import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, AbstractControl } from '@angular/forms';
import { WeightService } from 'src/app/services/weight/weight.service';

@Component({
    selector: 'app-weight',
    templateUrl: './weight.component.html',
    styleUrls: ['./weight.component.scss']
})
export class WeightComponent implements OnInit {
    measurements = [
        {
            value: 'lb',
            viewValue: 'pounds (lb)'
        },
        {
            value: 'kg',
            viewValue: 'Kilograms (kg)'
        }
    ];
    selected: any = '';
    weightForm: FormGroup;
    maxDate: Date;

    constructor(
        private formBuilder: FormBuilder,
        private weight: WeightService
    ) { }

    ngOnInit(): void {
        this.setSelectedValue();
        this.setMaxDate();
        this.setWeightForm();
    }

    setSelectedValue(): void {
        this.selected = this.measurements[1].value;
    }

    setMaxDate(): void {
        this.maxDate = new Date();
    }

    setWeightForm(): void {
        this.weightForm = this.formBuilder.group({
            weight: ['', Validators.compose([
                Validators.required,
                Validators.min(0),
                Validators.pattern(/^[+]?([0-9]+(?:[\.][0-9]*)?|\.[0-9]+)$/)
            ])],
            measurement: [this.selected, Validators.required],
            date: [new Date(), Validators.compose([
                Validators.required
            ])],
        });
    }

    get formControls(): { [key: string]: AbstractControl; } {
        return this.weightForm.controls;
    }

    weightSubmit(): void {
        if (this.weightForm.invalid) return;

        const weight = this.formControls.weight.value;
        const measurement = this.formControls.measurement.value;
        const dateInput = this.formControls.date.value;

        console.log('submitted', { weight, measurement, dateInput });
        const data = {
            "name": dateInput,
            "value": weight
        }
        this.weight.addWeight(data).subscribe(
            res => {
                console.log(res);
            }
        )
    }

}

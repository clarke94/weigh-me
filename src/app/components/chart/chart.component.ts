import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { WeightService } from 'src/app/services/weight/weight.service';
import { DatePipe } from '@angular/common';
import { FormGroup, FormBuilder, Validators, AbstractControl } from '@angular/forms';

@Component({
    selector: 'app-chart',
    templateUrl: './chart.component.html',
    styleUrls: ['./chart.component.scss'],
    encapsulation: ViewEncapsulation.None
})
export class ChartComponent implements OnInit {
    view: any[] = [0, 300];
    data: any;
    measurements: any[] = [];
    selected: any = '';
    weightForm: FormGroup;
    maxDate: Date;

    // options
    legend: boolean = false;
    showLabels: boolean = true;
    animations: boolean = true;
    xAxis: boolean = true;
    yAxis: boolean = true;
    showYAxisLabel: boolean = true;
    showXAxisLabel: boolean = true;
    xAxisLabel: string = 'Date';
    yAxisLabel: string = '';
    timeline: boolean = true;

    colorScheme = {
        domain: ['#5AA454', '#E44D25', '#CFC0BB', '#7aa3e5', '#a8385d', '#aae3f5']
    };

    constructor(
        private weight: WeightService,
        private datePipe: DatePipe,
        private formBuilder: FormBuilder
    ) { }

    ngOnInit(): void {
        this.setInitialChartData();
        this.setMaxDate();
        this.setWeightForm();
        this.setMeasurements();
    }

    setInitialChartData(): void {
        let dates: any[] = [];
        for (var i = 0; i < 7; i++) {
            var d = new Date();
            d.setDate(d.getDate() - i);
            dates.push({ name: this.datePipe.transform(d), value: Math.floor(Math.random() * (95 * 100 - 70 * 100) + 70 * 100) / (100) })
        }

        const defaultData = [
            {
                "name": "Weight",
                "series": dates
            }
        ];

        this.data = defaultData;
    }

    updatedMeasurement(): void {
        const measurement = this.formControls.measurement.value;
        const newValues = [];
        if (measurement === 'Kilograms (kg)') {
            for (var i = 0; i < this.data[0].series.length; i++) {
                this.data[0].series[i].value /= 2.2046;
                newValues.push(this.data[0].series[i]);
            }
        } else {
            for (var i = 0; i < this.data[0].series.length; i++) {
                this.data[0].series[i].value *= 2.2046;
                newValues.push(this.data[0].series[i]);
            }
        }
        const updatedDate = [
            {
                "name": "Weight",
                "series": newValues
            }
        ]
        this.data = updatedDate;
    }

    setMeasurements(): void {
        this.measurements = [
            'pounds (lb)',
            'Kilograms (kg)'
        ]
        this.yAxisLabel = this.measurements[1];
        this.formControls.measurement.setValue(this.measurements[1]);
    }

    getChartData(): void { // Currently not in use
        this.weight.getAllWeights().subscribe(data => this.data = data);
    }

    onSelect(data): void {
        console.log('Item clicked', JSON.parse(JSON.stringify(data)));
        this.formControls.weight.setValue(data.value);
        this.formControls.date.setValue(new Date(data.name));
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
            measurement: ['', Validators.required],
            date: [new Date(), Validators.compose([
                Validators.required,
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

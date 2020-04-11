import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/user/user.service';

@Component({
    selector: 'app-profile',
    templateUrl: './profile.component.html',
    styleUrls: ['./profile.component.scss']
})
export class ProfileComponent implements OnInit {

    constructor(
        public auth: AuthService
    ) { }

    ngOnInit(): void {
        this.auth.userProfile$.subscribe(
            data => {
                console.log('user', data);
            }
        )
    }
}

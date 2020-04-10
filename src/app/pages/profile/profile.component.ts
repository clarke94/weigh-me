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
        public auth: AuthService,
        private userService: UserService
    ) { }

    ngOnInit(): void {
        this.getUserProfile();
    }

    getUserProfile() {
        this.userService.getUserByAuthId("1").subscribe(
            data => {
                console.log(data);
            },
            err => {
                console.log('error', err);
            }
        );
    }

}

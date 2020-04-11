import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ProfileComponent } from './profile.component';
import { RegisterGuard } from 'src/app/guards/register/register.guard';


const routes: Routes = [
    {
        path: '',
        component: ProfileComponent,
        canActivate: [RegisterGuard]
    }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class ProfileRoutingModule { }

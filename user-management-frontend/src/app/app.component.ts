import { Component } from '@angular/core';
import { HomeComponent } from './home/home.component';
import {MatDividerModule} from '@angular/material/divider';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [HomeComponent, MatDividerModule, CommonModule],
  template: `
    <main>
      <h1>
        User Management System
      </h1>
      <mat-divider></mat-divider>
      <section class="content">
        <app-home></app-home>
      </section>
    </main>
  `,
  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  title = 'User Management System';
}

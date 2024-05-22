import { Component } from '@angular/core';
import { HomeComponent } from './home/home.component';
import { TestComponent } from './test/test.component';
import {MatDividerModule} from '@angular/material/divider';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [HomeComponent, TestComponent, MatDividerModule],
  template: `
    <main>
      <h1>
        Benju Koirala Project
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
  title = 'Benju Project';
}

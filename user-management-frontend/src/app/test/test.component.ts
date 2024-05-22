import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-test',
  template: `<button (click)="testConnection()">Test Connection</button>`,
  standalone: true,
  imports: []
})
export class TestComponent {
  private baseUrl = 'http://localhost:8080/api/users';

  constructor(private http: HttpClient) { }

  testConnection() {
    this.http.get(this.baseUrl).subscribe(response => {
      console.log('Success:', response);
    }, error => {
      console.error('Error:', error);
    });
  }
}

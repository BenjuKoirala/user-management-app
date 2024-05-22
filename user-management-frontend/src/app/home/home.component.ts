import { Component, OnInit, ViewChild } from '@angular/core';
import {MatTableModule, MatTableDataSource } from '@angular/material/table';
import {MatPaginator, MatPaginatorModule} from '@angular/material/paginator';
import {MatButtonModule} from '@angular/material/button';
import { CreateUserDialog } from '../dialogues/create-user.component';
import { MatDialog } from '@angular/material/dialog';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';


export interface UserDetails {
  id: number,
  name: string;
  email: string;
}

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [MatTableModule, MatButtonModule, MatPaginatorModule],
  providers: [],
  animations: [],
  templateUrl: './home-component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit {
  private baseUrl = 'http://localhost:8080/api/users'; // Backend URL
  displayedColumns: string[] = ['id', 'name', 'email'];
  dataSource = new MatTableDataSource<UserDetails>([]);
  clickedRow: UserDetails | undefined = undefined;
  createdUser: UserDetails | undefined = undefined;
  
  constructor(private http: HttpClient, public dialog: MatDialog) { }

  ngOnInit() {
    this.loadUsers();
  }

  loadUsers() {
    this.getUsers().subscribe(users => {
      console.log('Fetched users:', users);
      this.dataSource.data = users;
    });
  }

  getUsers(): Observable<UserDetails[]> {
    // here
    return this.http.get<UserDetails[]>(this.baseUrl);
  }

  getUser(id: number): Observable<UserDetails> {
    return this.http.get<UserDetails>(`${this.baseUrl}/${id}`);
  }

  createUser(user :UserDetails): Observable<UserDetails> {
    return this.http.post<UserDetails>(this.baseUrl, user)
  }

  updateUser(user: UserDetails): Observable<UserDetails> {
    return this.http.put<UserDetails>(`${this.baseUrl}/${user.id}`, user);
  }

  deleteUser(id: number): Observable<void> {
    return this.http.delete<void>(`${this.baseUrl}/${id}`);
  }

  //@ViewChild(MatPaginator) paginator: MatPaginator;

  ngAfterViewInit() {
    //this.dataSource.paginator = this.paginator;
  }

  setSelectedRow(row: UserDetails) {
    this.clickedRow = row;
    console.log(this.clickedRow);
  }

  isRowSelected(row: UserDetails) {
    return this.clickedRow == row;
  }

  onUpdateUser() {
    if (this.clickedRow) {
      const userId = this.clickedRow.id;
      const dialogRef = this.dialog.open(CreateUserDialog, {
        data: { name: this.clickedRow.name, email: this.clickedRow.email }
      });

      dialogRef.afterClosed().subscribe(result => {
        if (result) {
          const updatedUser: UserDetails = { id: userId, name: result.name, email: result.email };
          this.updateUser(updatedUser).subscribe(() => {
            console.log(`Updated user ${updatedUser.id}`);
            this.loadUsers();
          });
        }
      });
    }
  }

  onDeleteUser() {
    console.log(`Deleting user ${this.clickedRow}`)
    if (this.clickedRow) {
      this.deleteUser(this.clickedRow.id).subscribe(() => {
        console.log(`Deleted user ${this.clickedRow}`);
        this.loadUsers();
      });
    }
  }

  onCreateUser() {
    console.log("Creating new user");
    
    const dialogRef = this.dialog.open(CreateUserDialog, {
      data: { name: '', email: '' }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        const newUser: UserDetails = { id: 0, name: result.name, email: result.email };
        this.createUser(newUser).subscribe(createdUser => {
          console.log("Created new user", createdUser);
          this.loadUsers();
        });
      }
    });
  }

  hasNoSelection() {
    this.clickedRow == undefined;
  }
}

import { Component, ViewChild } from '@angular/core';
import {MatTableModule, MatTableDataSource } from '@angular/material/table';
import {MatPaginator, MatPaginatorModule} from '@angular/material/paginator';
import {MatButtonModule} from '@angular/material/button';
import { CreateUserDialog } from '../dialogues/create-user.component';
import { MatDialog } from '@angular/material/dialog';

export interface UserDetails {
  id: number,
  name: string;
  email: string;
}

const ELEMENT_DATA: UserDetails[] = [
  {id: 1, name: 'Benju', email: 'benju@Koirala.np'},
  {id: 2, name: 'Ashok', email: 'ashok@Adhikari.np'},
  {id: 3, name: 'Aiwen', email: 'aiwen@Adhikari.np'},
  {id: 4, name: 'Arlin', email: 'arlin@Adhikari.np'},
];

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [MatTableModule, MatButtonModule, MatPaginatorModule],
  providers: [],
  animations: [],
  templateUrl: 'home-component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  displayedColumns: string[] = ['id', 'name', 'email'];
  users = ELEMENT_DATA;
  dataSource = new MatTableDataSource<UserDetails>(this.users);
  clickedRow: undefined | UserDetails = undefined;
  createdUser: UserDetails | undefined = undefined;
  
  constructor(public dialog: MatDialog) {}

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

  updateUser() {
    console.log(`Updating user ${this.clickedRow}`)
    const dialogRef = this.dialog.open(CreateUserDialog, {
      data: {name: "", email: ""}
    });

    dialogRef.afterClosed().subscribe(result => {
      this.createdUser = {id: 0, name: result.name, email: result.email};
      console.log(this.createdUser)
      // connect to backend
    });
  }

  deleteUser() {
    console.log(`Deleting user ${this.clickedRow}`)
    // connect to backend and delete
    // after success remove from table UI
    // remove from this.users
    // this.dataSource = new MatTableDataSource(this.users);
  }

  createUser() {
    console.log("Creating new user");
    
    const dialogRef = this.dialog.open(CreateUserDialog, {
      data: {name: "", email: ""}
    });

    dialogRef.afterClosed().subscribe(result => {
      this.createdUser = {id: 0, name: result.name, email: result.email};

      // connect to backend and if success append in the table

      this.users.push(this.createdUser);
      this.dataSource = new MatTableDataSource(this.users);
      console.log(this.createdUser)
    });
  }

  hasNoSelection() {
    this.clickedRow == undefined;
  }
}

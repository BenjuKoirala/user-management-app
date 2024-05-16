import { Component, ViewChild } from '@angular/core';
import {MatTableModule, MatTableDataSource } from '@angular/material/table';
import {MatPaginator, MatPaginatorModule} from '@angular/material/paginator';
import {MatButtonModule} from '@angular/material/button';

export interface UserDetails {
  id: number,
  fname: string;
  lname: string;
  phone: string;
}

const ELEMENT_DATA: UserDetails[] = [
  {id: 1, fname: 'Benju', lname: 'Koirala', phone: '702-666-888'},
  {id: 2, fname: 'Ashok', lname: 'Adhikari', phone: '702-666-888'},
  {id: 3, fname: 'Aiwen', lname: 'Adhikari', phone: '702-666-888'},
  {id: 4, fname: 'Arlin', lname: 'Adhikari', phone: '702-666-888'},
];

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [MatTableModule, MatButtonModule, MatPaginatorModule],
  providers: [],
  templateUrl: 'home-component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  displayedColumns: string[] = ['id', 'fname', 'lname', 'phone'];
  users = ELEMENT_DATA;
  dataSource = new MatTableDataSource<UserDetails>(this.users);
  clickedRow: undefined | UserDetails = undefined;

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
  }

  deleteUser() {
    console.log(`Deleting user ${this.clickedRow}`)
  }

  createUser() {
    console.log("Creating new user");
    this.users.push({id: 6, fname: 'Aarya', lname: 'Manandhar', phone: '702-666-999'});
    this.dataSource = new MatTableDataSource(this.users);
  }

  hasNoSelection() {
    this.clickedRow == undefined;
  }
}

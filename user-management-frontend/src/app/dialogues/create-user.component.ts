import { Component, ViewChild, Inject } from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {FormsModule} from '@angular/forms';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {UserDetails} from '../home/home.component';
import {
  MatDialog,
  MAT_DIALOG_DATA,
  MatDialogRef,
  MatDialogTitle,
  MatDialogContent,
  MatDialogActions,
  MatDialogClose,
} from '@angular/material/dialog';


@Component({
  selector: 'create-user-dialog',
  standalone: true,
  imports: [MatFormFieldModule, MatInputModule, FormsModule, MatButtonModule, MatDialogContent, MatDialogActions, MatDialogTitle, MatDialogClose],
  templateUrl: 'user-inputs.html',
  styleUrl: '../home/home.component.css'
})
export class CreateUserDialog {
  dialogTitle: string = "Create";
  constructor(
    public dialogRef: MatDialogRef<CreateUserDialog>,
    @Inject(MAT_DIALOG_DATA) public data: {name: string, email: string},
  ) {}

  onNoClick(): void {
    this.dialogRef.close();
  }
}
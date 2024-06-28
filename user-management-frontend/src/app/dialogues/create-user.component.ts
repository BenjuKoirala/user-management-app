import { Component, ViewChild, Inject } from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {FormsModule} from '@angular/forms';
import { CommonModule } from '@angular/common';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {UserDetails} from '../home/home.component';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import {
  MatDialog,
  MAT_DIALOG_DATA,
  MatDialogRef,
  MatDialogTitle,
  MatDialogContent,
  MatDialogActions,
  MatDialogClose,
} from '@angular/material/dialog';

interface UserData {
  name: string;
  email: string;
}

@Component({
  selector: 'create-user-dialog',
  standalone: true,
  imports: [MatFormFieldModule, MatInputModule, FormsModule, CommonModule, MatButtonModule, MatDialogContent, MatDialogActions, MatDialogTitle, MatDialogClose],
  templateUrl: 'create-user.component.html',
  styleUrl: '../home/home.component.css'
})
export class CreateUserDialog {
  dialogTitle: string = "Create";
  data: { name: string, email: string };
  constructor(
    public dialogRef: MatDialogRef<CreateUserDialog>,
    @Inject(MAT_DIALOG_DATA) public inputData: { name: string, email: string }
  ) {
    this.data = { name: inputData.name, email: inputData.email };
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  formInvalid(): boolean {
    return !this.data.name || !this.data.email;
  }
}
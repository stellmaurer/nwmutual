/*************************************************************
 *
 * Code written by: Stephen Ellmaurer <stellmaurer@gmail.com>
 *
 ************************************************************/
import { Component, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'text-input-component',
  templateUrl: './text-input.component.html',
  styleUrls: ['./text-input.component.css']
})
export class TextInputComponent {

  @Output() whatToSearchFor = new EventEmitter();
  public textToSearchFor : string;

  constructor() {
    this.textToSearchFor = "";
    this.emit();
  }

  emit(){
    this.whatToSearchFor.emit(this.textToSearchFor);
  }
}

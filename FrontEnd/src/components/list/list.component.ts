/*************************************************************
 *
 * Code written by: Stephen Ellmaurer <stellmaurer@gmail.com>
 *
 ************************************************************/
import { Component } from '@angular/core';
import { Data } from '../../models/Data';

@Component({
  selector: 'list-component',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent {

  public resultsPerPage : number;
  // input
  public textToSearchFor : string;

  constructor(public data : Data) {
    this.resultsPerPage = 50;
    this.textToSearchFor = "";
    data.getJobs(this.resultsPerPage, 1, this.textToSearchFor);
  }

  goToPage(page : number){
    this.data.getJobs(this.resultsPerPage, page, this.textToSearchFor);
  }
}

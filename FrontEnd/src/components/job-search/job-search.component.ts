/*************************************************************
 *
 * Code written by: Stephen Ellmaurer <stellmaurer@gmail.com>
 *
 ************************************************************/
import { Component } from '@angular/core';
import { Data } from '../../models/Data';
import { ListComponent } from '../list/list.component';
import { TextInputComponent } from '../text-input/text-input.component'

@Component({
  selector: 'job-search-component',
  templateUrl: './job-search.component.html',
  styleUrls: ['./job-search.component.css']
})
export class JobSearchComponent {

  public resultsPerPage : number;
  public initialPageNumber : number;
  public initialTitleToSearchBy : string;
  public textInput : string;
  public filteredJobs : string[];
  public error : boolean;

  constructor(public data : Data) {
    this.resultsPerPage = 10;
    this.initialPageNumber = 1;
    this.initialTitleToSearchBy = "";
    this.error = false;
    data.getJobs()
    .then((res) => {
        this.filteredJobs = data.jobs;
        this.error = false;
    })
    .catch((err) => {
        this.error = true;
    });
  }

  whatToSearchForChanged(textInput: string) {
    this.textInput = textInput;
    this.filterJobs();
  }

  filterJobs(){
    let a = new Array<string>();
    for(let i = 0; i < this.data.jobs.length; i++){
      if(this.data.jobs[i].toLowerCase().includes(this.textInput.toLowerCase())){
        a.push(this.data.jobs[i]);
      }
    }
    this.filteredJobs = a;
  }


}

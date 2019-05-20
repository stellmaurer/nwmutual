/*************************************************************
 *
 * Code written by: Stephen Ellmaurer <stellmaurer@gmail.com>
 *
 ************************************************************/
import { Component, Input } from '@angular/core';
import { Data } from '../../models/Data';

@Component({
  selector: 'list-component',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent {
  // Inputs
  @Input() data : string[];
  @Input() resultsPerPage : number;

  // other member variables
  public page : number;
  public pages : number[];
  public pageData : string[];

  constructor() {
    this.page = 1;
    this.pages = new Array<number>();
  }

  ngOnChanges(){
    if((this.data != undefined) && (this.resultsPerPage != undefined)){
         this.pages = new Array<number>(Math.ceil(this.data.length / this.resultsPerPage));
         this.page = 1;
         this.updatePageData();
    }
  }

  goToPage(page : number){
    this.page = page;
    this.updatePageData();
  }

  updatePageData(){
    let startingIndex = (this.page - 1) * this.resultsPerPage;
    let endingIndex = 0;
    if(this.page == this.pages.length){
      endingIndex = this.data.length;
    }else{
      endingIndex = this.page * this.resultsPerPage;
    }
    this.pageData = this.data.slice(startingIndex, endingIndex);
  }
}

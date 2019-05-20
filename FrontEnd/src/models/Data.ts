/*************************************************************
 *
 * Code written by: Stephen Ellmaurer <stellmaurer@gmail.com>
 *
 ************************************************************/
import { Queries } from "./Queries";
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { NgZone } from '@angular/core';

@Injectable()
export class Data{

    public pages : number[];
    public jobs : string[];

    constructor(private zone: NgZone, private http : HttpClient) {
        this.pages = new Array<number>();
        this.jobs = new Array<string>();
    }

    public getJobs(resultsPerPage : number, page : number, textToSearchFor : string){
      var tempThis = this;
      return new Promise((resolve, reject) => {
          var queries = new Queries(tempThis, tempThis.http);
          queries.getJobs(resultsPerPage, page, textToSearchFor)
          .then((res) => {
              resolve("getJobs query succeeded.");
          })
          .catch((err) => {
              reject(err);
          });
      });
    }

}

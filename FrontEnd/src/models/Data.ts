/*************************************************************
 *
 * Code written by: Stephen Ellmaurer <stellmaurer@gmail.com>
 *
 ************************************************************/
import { Job } from './Job';
import { Queries } from "./Queries";
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable()
export class Data{
    public page : number;
    public totalItemCount : number;
    public jobs : Job[];
    constructor(private http : HttpClient) {
        this.page = -1;
        this.totalItemCount = -1;
        this.jobs = new Array<Job>();
    }

    public getJobs(resultsPerPage : number, page : number){
      var tempThis = this;
      return new Promise((resolve, reject) => {
          var queries = new Queries(tempThis, tempThis.http);
          queries.getJobs(resultsPerPage, page)
          .then((res) => {
              resolve("getJobs query succeeded.");
          })
          .catch((err) => {
              reject(err);
          });
      });
    }

}

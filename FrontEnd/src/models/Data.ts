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

    public jobs : string[];

    constructor(private zone: NgZone, private http : HttpClient) {
        this.jobs = new Array<string>();
    }

    public getJobs(){
      var tempThis = this;
      return new Promise((resolve, reject) => {
          var queries = new Queries(tempThis, tempThis.http);
          this.zone.run(() => {
            queries.getJobs()
            .then((res) => {
                resolve("getJobs query succeeded.");
            })
            .catch((err) => {
                reject(err);
            });
          });
      });
    }

}

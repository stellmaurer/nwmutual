/*************************************************************
 *
 * Code written by: Stephen Ellmaurer <stellmaurer@gmail.com>
 *
 ************************************************************/
import { Job } from './Job';
import { Data } from './Data';
import { HttpClient, HttpErrorResponse, HttpHeaders, HttpParams } from '@angular/common/http';

interface GetJobsResponse {
  data : JobResponse;
}
interface JobResponse {
  jobs : JobData[];
  page : number;
  totalItemCount : number;
}
interface JobData {
  title : string;
}
interface Success {
  message : string;
}

const options = {
  headers: new HttpHeaders({
    'Content-Type':  'application/x-www-form-urlencoded',
    "Access-Control-Allow-Origin": "*"
  })
};

export class Queries{

  constructor(private appData : Data, private http: HttpClient){

  }

  public getJobs(resultsPerPage : number, page : number){
    return new Promise((resolve, reject) => {
        let url = "http://localhost:5000/jobs?city=Milwaukee&state=Wisconsin" +
                  "&resultsPerPage=" + encodeURIComponent(String(resultsPerPage)) +
                  "&page=" + encodeURIComponent(String(page));
        let body = "";
        this.http.get<GetJobsResponse>(url, options).subscribe(
          data => {
            console.log(data);
            this.appData.page = data.data.page;
            this.appData.totalItemCount = data.data.totalItemCount;
            this.appData.jobs = Array.from(data.data.jobs);
          },
          (err: HttpErrorResponse) => {
            if (err.error instanceof Error) {
              console.log("Client-side error occured.");
            } else {
              console.log("Server-side error occured.");
            }
          }
        );
    });
  }

}

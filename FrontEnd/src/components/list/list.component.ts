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

  title = 'the nw mutual app';

  constructor(private data : Data) {
    data.getJobs(10, 0);
  }
}

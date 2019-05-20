import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { Data } from '../models/Data';
import { ListComponent } from '../components/list/list.component';
import { JobSearchComponent } from '../components/job-search/job-search.component';
import { TextInputComponent } from '../components/text-input/text-input.component';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    AppComponent,
    ListComponent,
    JobSearchComponent,
    TextInputComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [Data],
  bootstrap: [AppComponent]
})
export class AppModule { }

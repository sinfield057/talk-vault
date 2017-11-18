import { NotFoundComponent } from './not-found/not-found.component';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    NotFoundComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot( [
      { path: '', component: HomeComponent },
      { path: '**', component: NotFoundComponent },
    ] ),
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

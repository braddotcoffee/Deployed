import { Component, OnInit } from '@angular/core';
import { HelloWorldService } from './hello-world.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent implements OnInit {
  title = 'webapp';
  constructor(private helloWorld: HelloWorldService) { }

  ngOnInit(): void {
    this.helloWorld.getTitle()
      .subscribe(data => {
        this.title = data.title;
        console.log(this.title);
      });

    console.log(this.title);
  }
}

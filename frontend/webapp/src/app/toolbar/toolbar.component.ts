import { Component, OnInit } from '@angular/core';
import { NetworkConfigService } from '../services/network-config.service';

@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.scss']
})
export class ToolbarComponent implements OnInit {

  constructor(private networkConfig: NetworkConfigService) { }

  ngOnInit(): void {
  }

  refreshNetworkConfig(): void {
    this.networkConfig.updateNetworkConfigs();
  }

}

import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TranslateService } from '@ngx-translate/core';
import { take } from 'rxjs/operators';
import { Type } from 'src/app/proto/generated/zitadel/user_pb';
import { Breadcrumb, BreadcrumbService, BreadcrumbType } from 'src/app/services/breadcrumb.service';

@Component({
  selector: 'cnsl-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.scss'],
})
export class UserListComponent {
  public Type: any = Type;
  public type: Type = Type.TYPE_HUMAN;

  constructor(public translate: TranslateService, activatedRoute: ActivatedRoute, breadcrumbService: BreadcrumbService) {
    activatedRoute.data.pipe(take(1)).subscribe((params) => {
      const { type } = params;
      this.type = type;
    });
    const iambread = new Breadcrumb({
      type: BreadcrumbType.IAM,
      name: 'IAM',
      routerLink: ['/system'],
    });
    const bread: Breadcrumb = {
      type: BreadcrumbType.ORG,
      routerLink: ['/org'],
    };
    breadcrumbService.setBreadcrumb([iambread, bread]);
  }
}

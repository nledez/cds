import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {AuthentificationStore} from '../../../../service/auth/authentification.store';
import {WorkerModel, ModelPattern} from '../../../../model/worker-model.model';
import {Group} from '../../../../model/group.model';
import {WorkerModelService} from '../../../../service/worker-model/worker-model.service';
import {GroupService} from '../../../../service/group/group.service';
import {SharedService} from '../../../../shared/shared.service';
import {ToastService} from '../../../../shared/toast/ToastService';
import {TranslateService} from '@ngx-translate/core';
import {User} from '../../../../model/user.model';
import {finalize} from 'rxjs/operators';

@Component({
    selector: 'app-worker-model-edit',
    templateUrl: './worker-model.edit.html',
    styleUrls: ['./worker-model.edit.scss']
})
export class WorkerModelEditComponent implements OnInit {
    loading = false;
    deleteLoading = false;
    workerModel: WorkerModel;
    workerModelTypes: Array<string>;
    workerModelCommunications: Array<string>;
    workerModelGroups: Array<Group>;
    workerModelPatterns: Array<ModelPattern> = [];
    workerModelPatternsFiltered: Array<ModelPattern> = [];
    patternSelected: ModelPattern;
    currentUser: User;
    canEdit = false;

    private workerModelNamePattern: RegExp = new RegExp('^[a-zA-Z0-9._-]{1,}$');
    private workerModelPatternError = false;

    constructor(private sharedService: SharedService,
                private _workerModelService: WorkerModelService, private _groupService: GroupService,
                private _toast: ToastService, private _translate: TranslateService,
                private _route: ActivatedRoute, private _router: Router,
                private _authentificationStore: AuthentificationStore) {
        this.currentUser = this._authentificationStore.getUser();
        this._groupService.getGroups(true).subscribe( groups => {
            this.workerModelGroups = groups;
        });
        this.loading = true;
        this._workerModelService.getWorkerModelPatterns()
          .pipe(finalize(() => this.loading = false))
          .subscribe((patterns) => {
              this.workerModelPatterns = patterns;
              if (this.workerModel) {
                  this.workerModelPatternsFiltered = patterns.filter((wmp) => wmp.type === this.workerModel.type);
              } else {
                  this.workerModelPatternsFiltered = patterns;
              }
          });
    }

    ngOnInit() {
        this._route.params.subscribe(params => {
            this._workerModelService.getWorkerModelTypes().subscribe( wmt => {
                this.workerModelTypes = wmt;
            });
            this._workerModelService.getWorkerModelCommunications().subscribe( wmc => {
                this.workerModelCommunications = wmc;
            });

            if (params['workerModelName'] !== 'add') {
                this.reloadData(params['workerModelName']);
            } else {
                this.workerModel = new WorkerModel();
            }
        });
    }

    reloadData(workerModelName: string): void {
      this._workerModelService.getWorkerModelByName(workerModelName).subscribe( wm => {
          this.workerModel = wm;
          this.workerModelPatternsFiltered = this.workerModelPatterns.filter((wmp) => wmp.type === wm.type);
          if (this.currentUser.admin) {
              this.canEdit = true;
              return;
          }


          // here, check if user is admin of worker model group
          this._groupService.getGroupByName(wm.group.name).subscribe( gr => {
              if (gr.admins) {
                for (let i = 0; i < gr.admins.length; i++) {
                    if (gr.admins[i].username === this.currentUser.username) {
                      this.canEdit = true;
                      break;
                    };
                }
              }
          });
      });
    }

    clickDeleteButton(): void {
      this.deleteLoading = true;
      this._workerModelService.deleteWorkerModel(this.workerModel).subscribe( wm => {
          this.deleteLoading = false;
          this._toast.success('', this._translate.instant('worker_model_deleted'));
          this._router.navigate(['../'], { relativeTo: this._route });
      }, () => {
          this.loading = false;
      });
    }

    clickSaveButton(): void {
      if (!this.workerModel.name) {
          return;
      }

      if (!this.workerModelNamePattern.test(this.workerModel.name)) {
          this.workerModelPatternError = true;
          return;
      }

      // cast to int
      this.workerModel.group_id = Number(this.workerModel.group_id);
      this.workerModelGroups.forEach( g => {
        if (this.workerModel.group_id === g.id) {
          this.workerModel.group = g;
          return;
        }
      });

      this.loading = true;
      if (this.workerModel.id > 0) {
        this._workerModelService.updateWorkerModel(this.workerModel).subscribe( wm => {
            this.loading = false;
            this._toast.success('', this._translate.instant('worker_model_saved'));
            this._router.navigate(['settings', 'worker-model', this.workerModel.name]);
        }, () => {
            this.loading = false;
        });
      } else {
        this._workerModelService.createWorkerModel(this.workerModel).subscribe( wm => {
            this.loading = false;
            this._toast.success('', this._translate.instant('worker_model_saved'));
            this._router.navigate(['settings', 'worker-model', this.workerModel.name]);
        }, () => {
            this.loading = false;
        });
      }
    }

    getDescriptionHeight(): number {
        return this.sharedService.getTextAreaheight(this.workerModel.description);
    }

    filterPatterns(type: string) {
        this.patternSelected = null;
        this.workerModelPatternsFiltered = this.workerModelPatterns.filter((wmp) => wmp.type === type);
    }

    preFillModel(pattern: ModelPattern) {
        if (!this.workerModel || !this.workerModel.type || !pattern) {
            return;
        }
        switch (this.workerModel.type) {
            case 'docker':
                this.workerModel.model_docker.cmd = pattern.model.cmd;
                break
            default:
                this.workerModel.model_virtual_machine.pre_cmd = pattern.model.pre_cmd;
                this.workerModel.model_virtual_machine.cmd = pattern.model.cmd;
                this.workerModel.model_virtual_machine.post_cmd = pattern.model.post_cmd;
        }
    }
}

import {Component} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {WorkflowNodeRun, WorkflowRun} from '../../../../model/workflow.run.model';
import {AutoUnsubscribe} from '../../../../shared/decorator/autoUnsubscribe';
import {PipelineStatus} from '../../../../model/pipeline.model';
import {Project} from '../../../../model/project.model';
import {RouterService} from '../../../../service/router/router.service';
import {WorkflowRunService} from '../../../../service/workflow/run/workflow.run.service';
import {DurationService} from '../../../../shared/duration/duration.service';
import {first} from 'rxjs/operators';
import {EventSubscription} from '../../../../model/event.model';
import {WorkflowEventStore} from '../../../../service/workflow/workflow.event.store';
import {EventStore} from '../../../../service/event/event.store';
import {Subscription} from 'rxjs/Subscription';

@Component({
    selector: 'app-workflow-run-node',
    templateUrl: './node.html',
    styleUrls: ['./node.scss']
})
@AutoUnsubscribe()
export class WorkflowNodeRunComponent {

    nodeRun: WorkflowNodeRun;
    subNodeRun: Subscription;

    // Context info
    project: Project;
    workflowName: string;
    duration: string;

    workflowRun: WorkflowRun;

    // History
    nodeRunsHistory = new Array<WorkflowNodeRun>();

    selectedTab: string;

    constructor(private _activatedRoute: ActivatedRoute,
                private _router: Router, private _routerService: RouterService, private _workflowRunService: WorkflowRunService,
                private _durationService: DurationService,
                private _workflowEventStore: WorkflowEventStore, private _eventStore: EventStore) {

        this._activatedRoute.data.subscribe(datas => {
            this.project = datas['project'];
        });

        // Tab selection
        this._activatedRoute.queryParams.subscribe(q => {
            if (q['tab']) {
                this.selectedTab = q['tab'];
            } else {
                this.selectedTab = 'pipeline';
            }
        });

        // Get workflow name
        this.workflowName = this._routerService.getRouteSnapshotParams({}, this._router.routerState.snapshot.root)['workflowName'];

        this._activatedRoute.params.subscribe(params => {
            this.nodeRun = null;
            let number = params['number'];
            let nodeRunId = params['nodeId'];

            if (this.project && this.project.key && this.workflowName && number && nodeRunId) {
                // Get workflow Run
                this.initWorkflowNodeRun(number, nodeRunId);
            }
        });
    }

    initWorkflowNodeRun(number, nodeRunId): void {
        this._workflowRunService.getWorkflowRun(this.project.key, this.workflowName, number).pipe(first()).subscribe(wr => {
            this.workflowRun = wr;
            this._workflowEventStore.setSelectedRun(this.workflowRun);

            // subscribe to run event
            let s = new EventSubscription();
            s.key = this.project.key;
            s.workflow_name = this.workflowName;
            s.runs = true;
            s.num = wr.num;
            this._eventStore.changeFilter(s);

            let historyChecked = false;
            this._workflowRunService.getWorkflowNodeRun(this.project.key, this.workflowName, number, nodeRunId)
                .pipe(first()).subscribe(nodeRun => {
                this.nodeRun = nodeRun;
                if (!historyChecked) {
                    historyChecked = true;
                    this._workflowRunService.nodeRunHistory(
                        this.project.key, this.workflowName,
                        number, this.nodeRun.workflow_node_id)
                        .pipe(first())
                        .subscribe(nrs => this.nodeRunsHistory = nrs);
                }

                if (this.nodeRun && !PipelineStatus.isActive(this.nodeRun.status)) {
                    this.duration = this._durationService.duration(new Date(this.nodeRun.start), new Date(this.nodeRun.done));
                }
                this._workflowEventStore.setSelectedNodeRun(this.nodeRun);
                this.subNodeRun = this._workflowEventStore.selectedNodeRun().subscribe(wnr => {
                    this.nodeRun = wnr;
                });
                let f = new EventSubscription();
                f.key = this.project.key;
                f.workflow_name = this.workflowName;
                f.num = number;
                f.runs = true;
                this._eventStore.changeFilter(f);
            });
        });
    }

    showTab(tab: string): void {
        this._router.navigateByUrl('/project/' + this.project.key +
            '/workflow/' + this.workflowName +
            '/run/' + this.nodeRun.num +
            '/node/' + this.nodeRun.id +
            '?&tab=' + tab);
    }
}

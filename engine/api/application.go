package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/go-gorp/gorp"
	"github.com/gorilla/mux"

	"github.com/ovh/cds/engine/api/application"
	"github.com/ovh/cds/engine/api/businesscontext"
	"github.com/ovh/cds/engine/api/environment"
	"github.com/ovh/cds/engine/api/group"
	"github.com/ovh/cds/engine/api/permission"
	"github.com/ovh/cds/engine/api/pipeline"
	"github.com/ovh/cds/engine/api/poller"
	"github.com/ovh/cds/engine/api/project"
	"github.com/ovh/cds/engine/api/repositoriesmanager"
	"github.com/ovh/cds/engine/api/sanity"
	"github.com/ovh/cds/engine/api/scheduler"
	"github.com/ovh/cds/engine/api/trigger"
	"github.com/ovh/cds/engine/api/workflow"
	"github.com/ovh/cds/sdk"
	"github.com/ovh/cds/sdk/log"
)

func getApplicationsHandler(w http.ResponseWriter, r *http.Request, db *gorp.DbMap, c *businesscontext.Ctx) error {
	return WriteJSON(w, r, c.Project.Applications, http.StatusOK)
}

func getApplicationTreeHandler(w http.ResponseWriter, r *http.Request, db *gorp.DbMap, c *businesscontext.Ctx) error {
	vars := mux.Vars(r)
	projectKey := vars["key"]
	applicationName := vars["permApplicationName"]

	tree, err := workflow.LoadCDTree(db, projectKey, applicationName, c.User, "", 0)
	if err != nil {
		log.Warning("getApplicationTreeHandler: Cannot load CD Tree for applications %s: %s\n", applicationName, err)
		return err
	}

	return WriteJSON(w, r, tree, http.StatusOK)
}

func getPipelineBuildBranchHistoryHandler(w http.ResponseWriter, r *http.Request, db *gorp.DbMap, c *businesscontext.Ctx) error {
	err := r.ParseForm()
	if err != nil {
		log.Warning("getPipelineBranchHistoryHandler> Cannot parse form: %s\n", err)
		return sdk.ErrUnknownError
	}

	pageString := r.Form.Get("page")
	nbPerPageString := r.Form.Get("perPage")

	var nbPerPage int
	if nbPerPageString != "" {
		nbPerPage, err = strconv.Atoi(nbPerPageString)
		if err != nil {
			return err
		}
	} else {
		nbPerPage = 20
	}

	var page int
	if pageString != "" {
		page, err = strconv.Atoi(pageString)
		if err != nil {
			return err
		}
	} else {
		nbPerPage = 0
	}

	pbs, err := pipeline.GetBranchHistory(db, c.Project.Key, c.Application.Name, page, nbPerPage)
	if err != nil {
		log.Warning("getPipelineBranchHistoryHandler> Cannot get history by branch: %s", err)
		return fmt.Errorf("Cannot load pipeline branch history: %s", err)
	}

	return WriteJSON(w, r, pbs, http.StatusOK)
}

func getApplicationDeployHistoryHandler(w http.ResponseWriter, r *http.Request, db *gorp.DbMap, c *businesscontext.Ctx) error {
	pbs, err := pipeline.GetDeploymentHistory(db, c.Project.Key, c.Application.Name)
	if err != nil {
		log.Warning("getPipelineDeployHistoryHandler> Cannot get history by env: %s", err)
		return fmt.Errorf("Cannot load pipeline deployment history: %s", err)
	}

	return WriteJSON(w, r, pbs, http.StatusOK)
}

func getApplicationBranchVersionHandler(w http.ResponseWriter, r *http.Request, db *gorp.DbMap, c *businesscontext.Ctx) error {
	vars := mux.Vars(r)
	projectKey := vars["key"]
	applicationName := vars["permApplicationName"]

	branch := r.FormValue("branch")

	app, err := application.LoadByName(db, projectKey, applicationName, c.User, application.LoadOptions.WithTriggers)
	if err != nil {
		log.Warning("getApplicationBranchVersionHandler: Cannot load application %s for project %s from db: %s\n", applicationName, projectKey, err)
		return err
	}

	versions, err := pipeline.GetVersions(db, app, branch)
	if err != nil {
		log.Warning("getApplicationBranchVersionHandler: Cannot load version for application %s on branch %s: %s\n", applicationName, branch, err)
		return err
	}

	return WriteJSON(w, r, versions, http.StatusOK)
}

func getApplicationTreeStatusHandler(w http.ResponseWriter, r *http.Request, db *gorp.DbMap, c *businesscontext.Ctx) error {
	branchName := r.FormValue("branchName")
	versionString := r.FormValue("version")

	var version int64
	var errV error
	if versionString != "" {
		version, errV = strconv.ParseInt(versionString, 10, 64)
		if errV != nil {
			return sdk.WrapError(errV, "getApplicationTreeStatusHandler>Cannot cast version %s into int", versionString)
		}
	}

	pbs, schedulers, pollers, hooks, errPB := workflow.GetWorkflowStatus(db, c.Project.Key, c.Application.Name, c.User, branchName, version)
	if errPB != nil {
		return sdk.WrapError(errPB, "getApplicationHandler> Cannot load CD Tree status")
	}

	response := struct {
		Builds     []sdk.PipelineBuild     `json:"builds"`
		Schedulers []sdk.PipelineScheduler `json:"schedulers"`
		Pollers    []sdk.RepositoryPoller  `json:"pollers"`
		Hooks      []sdk.Hook              `json:"hooks"`
	}{
		pbs,
		schedulers,
		pollers,
		hooks,
	}

	return WriteJSON(w, r, response, http.StatusOK)
}

func getApplicationHandler(w http.ResponseWriter, r *http.Request, db *gorp.DbMap, c *businesscontext.Ctx) error {
	vars := mux.Vars(r)
	projectKey := vars["key"]
	applicationName := vars["permApplicationName"]

	applicationStatus := FormBool(r, "applicationStatus")
	withPollers := FormBool(r, "withPollers")
	withHooks := FormBool(r, "withHooks")
	withNotifs := FormBool(r, "withNotifs")
	withWorkflow := FormBool(r, "withWorkflow")
	withRepoManager := FormBool(r, "withRepoMan")
	withTriggers := FormBool(r, "withTriggers")
	withSchedulers := FormBool(r, "withSchedulers")
	branchName := r.FormValue("branchName")
	versionString := r.FormValue("version")

	loadOptions := []application.LoadOptionFunc{
		application.LoadOptions.WithVariables,
		application.LoadOptions.WithRepositoryManager,
		application.LoadOptions.WithVariables,
		application.LoadOptions.WithPipelines,
	}
	if withHooks {
		loadOptions = append(loadOptions, application.LoadOptions.WithHooks)
	}
	if withTriggers {
		loadOptions = append(loadOptions, application.LoadOptions.WithTriggers)
	}
	if withNotifs {
		loadOptions = append(loadOptions, application.LoadOptions.WithNotifs)
	}

	app, errApp := application.LoadByName(db, projectKey, applicationName, c.User, loadOptions...)
	if errApp != nil {
		log.Warning("getApplicationHandler: Cannot load application %s for project %s from db: %s\n", applicationName, projectKey, errApp)
		return errApp
	}

	if err := application.LoadGroupByApplication(db, app); err != nil {
		return sdk.WrapError(err, "getApplicationHandler> Unable to load groups by application")
	}

	if withPollers {
		var errPoller error
		app.RepositoryPollers, errPoller = poller.LoadByApplication(db, app.ID)
		if errPoller != nil {
			log.Warning("getApplicationHandler: Cannot load pollers for application %s: %s\n", applicationName, errPoller)
			return errPoller
		}
	}

	if withSchedulers {
		var errScheduler error
		app.Schedulers, errScheduler = scheduler.GetByApplication(db, app)
		if errScheduler != nil {
			log.Warning("getApplicationHandler: Cannot load schedulers for application %s: %s\n", applicationName, errScheduler)
			return errScheduler
		}
	}

	if withWorkflow {
		var errWorflow error
		app.Workflows, errWorflow = workflow.LoadCDTree(db, projectKey, applicationName, c.User, "", 0)
		if errWorflow != nil {
			log.Warning("getApplicationHandler: Cannot load CD Tree for applications %s: %s\n", app.Name, errWorflow)
			return errWorflow
		}
	}

	if withRepoManager {
		var errRepo error
		_, app.RepositoriesManager, errRepo = repositoriesmanager.LoadFromApplicationByID(db, app.ID)
		if errRepo != nil {
			return sdk.WrapError(errRepo, "getApplicationHandler: Cannot load repo manager for application %s", app.Name)
		}
	}

	if applicationStatus {
		var pipelineBuilds []sdk.PipelineBuild

		version := 0
		if versionString != "" {
			var errStatus error
			version, errStatus = strconv.Atoi(versionString)
			if errStatus != nil {
				log.Warning("getApplicationHandler: Version %s is not an integer: %s\n", versionString, errStatus)
				return errStatus
			}
		}

		if version == 0 {
			var errBuilds error
			pipelineBuilds, errBuilds = pipeline.GetAllLastBuildByApplication(db, app.ID, branchName, 0)
			if errBuilds != nil {
				log.Warning("getApplicationHandler: Cannot load app status: %s\n", errBuilds)
				return errBuilds
			}
		} else {
			if branchName == "" {
				log.Warning("getApplicationHandler: branchName must be provided with version param\n")
				return sdk.ErrBranchNameNotProvided
			}
			var errPipBuilds error
			pipelineBuilds, errPipBuilds = pipeline.GetAllLastBuildByApplication(db, app.ID, branchName, version)
			if errPipBuilds != nil {
				log.Warning("getApplicationHandler: Cannot load app status by version: %s\n", errPipBuilds)
				return errPipBuilds
			}
		}
		al := r.Header.Get("Accept-Language")
		for _, p := range pipelineBuilds {
			p.Translate(al)
		}
		app.PipelinesBuild = pipelineBuilds
	}

	app.Permission = permission.ApplicationPermission(app.ID, c.User)

	return WriteJSON(w, r, app, http.StatusOK)
}

func getApplicationBranchHandler(w http.ResponseWriter, r *http.Request, db *gorp.DbMap, c *businesscontext.Ctx) error {
	var branches []sdk.VCSBranch
	if c.Application.RepositoryFullname != "" && c.Application.RepositoriesManager != nil {
		client, erra := repositoriesmanager.AuthorizedClient(db, c.Project.Key, c.Application.RepositoriesManager.Name)
		if erra != nil {
			log.Warning("getApplicationBranchHandler> Cannot get client got %s %s : %s", c.Project.Key, c.Application.RepositoriesManager.Name, erra)
			return sdk.ErrNoReposManagerClientAuth
		}
		var errb error
		branches, errb = client.Branches(c.Application.RepositoryFullname)
		if errb != nil {
			log.Warning("getApplicationBranchHandler> Cannot get branches from repository %s: %s", c.Application.RepositoryFullname, errb)
			return sdk.ErrNoReposManagerClientAuth
		}
	} else {
		var errg error
		branches, errg = pipeline.GetBranches(db, c.Application)
		if errg != nil {
			log.Warning("getApplicationBranchHandler> Cannot get branches from builds: %s", errg)
			return errg
		}
	}

	//Yo analyze branch and delete pipeline_build for old branches...

	return WriteJSON(w, r, branches, http.StatusOK)
}

func addApplicationHandler(w http.ResponseWriter, r *http.Request, db *gorp.DbMap, c *businesscontext.Ctx) error {
	var app sdk.Application
	if err := UnmarshalBody(r, &app); err != nil {
		return err
	}

	// check application name pattern
	regexp := regexp.MustCompile(sdk.NamePattern)
	if !regexp.MatchString(app.Name) {
		log.Warning("addApplicationHandler: Application name %s do not respect pattern %s", app.Name, sdk.NamePattern)
		return sdk.ErrInvalidApplicationPattern
	}

	tx, err := db.Begin()
	if err != nil {
		log.Warning("addApplicationHandler> Cannot start transaction: %s\n", err)
		return err
	}

	defer tx.Rollback()

	if err := application.Insert(tx, c.Project, &app, c.User); err != nil {
		log.Warning("addApplicationHandler> Cannot insert pipeline: %s\n", err)
		return err
	}

	if err := group.LoadGroupByProject(tx, c.Project); err != nil {
		log.Warning("addApplicationHandler> Cannot load group from project: %s\n", err)
		return err
	}

	if err := application.AddGroup(tx, c.Project, &app, c.User, c.Project.ProjectGroups...); err != nil {
		log.Warning("addApplicationHandler> Cannot add groups on application: %s\n", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Warning("addApplicationHandler> Cannot commit transaction: %s\n", err)
		return err
	}
	return nil
}

func deleteApplicationHandler(w http.ResponseWriter, r *http.Request, db *gorp.DbMap, c *businesscontext.Ctx) error {
	nb, errNb := pipeline.CountBuildingPipelineByApplication(db, c.Application.ID)
	if errNb != nil {
		log.Warning("deleteApplicationHandler> Cannot count pipeline build for application %d: %s\n", c.Application.ID, errNb)
		return errNb
	}

	if nb > 0 {
		log.Warning("deleteApplicationHandler> Cannot delete application [%d], there are building pipelines: %d\n", c.Application.ID, nb)
		return sdk.ErrAppBuildingPipelines
	}

	tx, err := db.Begin()
	if err != nil {
		log.Warning("deleteApplicationHandler> Cannot begin transaction: %s\n", err)
		return err
	}
	defer tx.Rollback()

	err = application.DeleteApplication(tx, c.Application.ID)
	if err != nil {
		log.Warning("deleteApplicationHandler> Cannot delete application: %s\n", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Warning("deleteApplicationHandler> Cannot commit transaction: %s\n", err)
		return err
	}
	return nil
}

func cloneApplicationHandler(w http.ResponseWriter, r *http.Request, db *gorp.DbMap, c *businesscontext.Ctx) error {
	envs, errE := environment.LoadEnvironments(db, c.Project.Key, true, c.User)
	if errE != nil {
		log.Warning("cloneApplicationHandler> Cannot load Environments %s: %s\n", c.Project.Key, errE)
		return errE

	}
	c.Project.Environments = envs

	var newApp sdk.Application
	if err := UnmarshalBody(r, &newApp); err != nil {
		return err
	}

	appToClone, errApp := application.LoadByName(db, c.Project.Key, c.Application.Name, c.User, application.LoadOptions.Default, application.LoadOptions.WithGroups)
	if errApp != nil {
		log.Warning("cloneApplicationHandler> Cannot load application %s: %s\n", c.Application.Name, errApp)
		return errApp
	}

	tx, errBegin := db.Begin()
	if errBegin != nil {
		log.Warning("cloneApplicationHandler> Cannot start transaction : %s\n", errBegin)
		return errBegin
	}
	defer tx.Rollback()

	if err := cloneApplication(tx, c.Project, &newApp, appToClone, c.User); err != nil {
		log.Warning("cloneApplicationHandler> Cannot insert new application %s: %s\n", newApp.Name, err)
		return err
	}

	if err := project.UpdateLastModified(tx, c.User, c.Project); err != nil {
		log.Warning("cloneApplicationHandler: Cannot update last modified date: %s\n", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Warning("cloneApplicationHandler> Cannot commit transaction : %s\n", err)
		return err
	}
	return WriteJSON(w, r, newApp, http.StatusOK)
}

// cloneApplication Clone an application with all her dependencies: pipelines, permissions, triggers
func cloneApplication(db gorp.SqlExecutor, proj *sdk.Project, newApp *sdk.Application, appToClone *sdk.Application, u *sdk.User) error {
	newApp.Pipelines = appToClone.Pipelines
	newApp.ApplicationGroups = appToClone.ApplicationGroups

	// Create Application
	if err := application.Insert(db, proj, newApp, u); err != nil {
		return err
	}

	var variablesToDelete []string
	for _, v := range newApp.Variable {
		if v.Type == sdk.KeyVariable {
			variablesToDelete = append(variablesToDelete, fmt.Sprintf("%s.pub", v.Name))
		}
	}

	for _, vToDelete := range variablesToDelete {
		for i := range newApp.Variable {
			if vToDelete == newApp.Variable[i].Name {
				newApp.Variable = append(newApp.Variable[:i], newApp.Variable[i+1:]...)
				break
			}
		}
	}

	// Insert variable
	for _, v := range newApp.Variable {
		var errVar error
		// If variable is a key variable, generate a new one for this application
		if v.Type == sdk.KeyVariable {
			errVar = application.AddKeyPairToApplication(db, newApp, v.Name, u)
		} else {
			errVar = application.InsertVariable(db, newApp, v, u)
		}
		if errVar != nil {
			return errVar
		}
	}

	// Attach pipeline + Set pipeline parameters
	for _, appPip := range newApp.Pipelines {
		if _, err := application.AttachPipeline(db, newApp.ID, appPip.Pipeline.ID); err != nil {
			return err
		}

		if err := application.UpdatePipelineApplication(db, newApp, appPip.Pipeline.ID, appPip.Parameters, u); err != nil {
			return err
		}
	}

	// Load trigger to clone
	triggers, err := trigger.LoadTriggerByApp(db, appToClone.ID)
	if err != nil {
		return err
	}

	// Clone trigger
	for _, t := range triggers {
		// Insert new trigger
		if t.DestApplication.ID == appToClone.ID {
			t.DestApplication = *newApp
		}
		t.SrcApplication = *newApp
		if err := trigger.InsertTrigger(db, &t); err != nil {
			return err
		}
	}

	//Reload trigger
	for i := range newApp.Pipelines {
		appPip := &newApp.Pipelines[i]
		var errTrig error
		appPip.Triggers, errTrig = trigger.LoadTriggersByAppAndPipeline(db, newApp.ID, appPip.Pipeline.ID)
		if errTrig != nil {
			log.Warning("cloneApplication> Cannot load triggers: %s\n", errTrig)
			return errTrig
		}
	}

	// Insert Permission
	if err := application.AddGroup(db, proj, newApp, u, newApp.ApplicationGroups...); err != nil {
		return err
	}

	if err := sanity.CheckApplication(db, proj, newApp); err != nil {
		log.Warning("cloneApplication> Cannot check application sanity: %s\n", err)
		return err
	}

	return nil
}

func updateApplicationHandler(w http.ResponseWriter, r *http.Request, db *gorp.DbMap, c *businesscontext.Ctx) error {
	envs, errloadenv := environment.LoadEnvironments(db, c.Project.Key, true, c.User)
	if errloadenv != nil {
		log.Warning("updateApplicationHandler> Cannot load environments %s: %s\n", c.Project.Key, errloadenv)
		return errloadenv
	}
	c.Project.Environments = envs

	var appPost sdk.Application
	if err := UnmarshalBody(r, &appPost); err != nil {
		return err
	}

	// check application name pattern
	regexp := regexp.MustCompile(sdk.NamePattern)
	if !regexp.MatchString(appPost.Name) {
		log.Warning("updateApplicationHandler: Application name %s do not respect pattern %s", appPost.Name, sdk.NamePattern)
		return sdk.ErrInvalidApplicationPattern
	}

	//Update name and Metadata
	c.Application.Name = appPost.Name
	c.Application.Metadata = appPost.Metadata

	tx, err := db.Begin()
	if err != nil {
		log.Warning("updateApplicationHandler> Cannot start transaction: %s\n", err)
		return err
	}
	defer tx.Rollback()
	if err := application.Update(tx, c.Application, c.User); err != nil {
		log.Warning("updateApplicationHandler> Cannot delete application %s: %s\n", c.Application.Name, err)
		return err
	}

	if err := sanity.CheckApplication(tx, c.Project, c.Application); err != nil {
		log.Warning("updateApplicationHandler: Cannot check application sanity: %s\n", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Warning("updateApplicationHandler> Cannot commit transaction: %s\n", err)
		return err
	}

	return WriteJSON(w, r, c.Application, http.StatusOK)

}
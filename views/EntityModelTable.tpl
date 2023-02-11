[[define "entityModelList"]]
<table class="table table-hover rounded tableFixHead">
    <thead>
        <tr>
            <th scope="col" data-sortable="true">Name</th>
            <th scope="col" data-sortable="true">Start Date</th>
            <th scope="col" data-sortable="true">End Date</th>
            <th scope="col" data-sortable="true">Version</th>
        </tr>
    </thead>
    <tbody>
        [[range .models]]
            <tr ng-click="viewEntityModel('[[.Name]]')">
                <th style="min-width: 250px; position: sticky; Left: 0px;" scope="row">[[.Name]]</th>
                <th style="min-width: 250px; position: sticky; Left: 0px;" scope="row">[[.StartDate.MonthName]] [[.StartDate.Year]]</th>
                <th style="min-width: 250px; position: sticky; Left: 0px;" scope="row">[[.SalesDate.MonthName]] [[.SalesDate.Year]]</th>
                <th style="min-width: 250px; position: sticky; Left: 0px;" scope="row">[[.Version]]</th>
            </tr>
        [[end]]
    </tbody>
</table>
[[end]]

[[template "entityModelList" .]]

[[define "entitySelection"]]
<table class="table table-hover rounded tableFixHead">
    <thead>
        <tr>
            <th scope="col" data-sortable="true">Entity</th>
            <th scope="col" data-sortable="true">Type</th>
        </tr>
    </thead>
    <tbody>
        [[range .]]
            <tr ng-click="getModels('[[.Name]]')">
                <th style="min-width: 250px; position: sticky; Left: 0px;" scope="row">[[.Name]]</th>
                <th style="min-width: 250px; position: sticky; Left: 0px;" scope="row">[[.EntityType]]</th>
            </tr>
        [[end]]
    </tbody>
</table>
[[end]]

{{template "header" .}}
<body>
    <br>
    <br>
    <div class="container-fluid">
        {{template "EntitySettings" .}}
    </div>
    {{if .entity.Table}}
        <div class="container-fluid">
            {{template "CFTable" .entity}}
        </div>
    {{end}}
    {{if .entity.ChildEntities}}
        <div class="container-fluid">
            {{template "EntityTable" .entity.ChildEntities}}
        </div>
    {{end}}
    {{if not .entity.ChildUnits}}
        <div class="container-fluid">
            {{template "AddChildEntityModal" .entity}}
        </div>
    {{end}}
    {{if .entity.ChildUnits}}
        <div class="container-fluid">
            {{template "UnitTable" .entity}}
        </div>
    {{end}}
    {{if not .entity.ChildEntities}}
        <div class="container-fluid">
            {{template "AddChildUnitModal" .entity.Name}}
        </div>
    {{end}}
    {{template "footer"}}
</body>

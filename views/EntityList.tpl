{{template "header"}}
<body>
    <br>
    <br>
    <div class="container-fluid">
        {{template "EntityTable" .entitymap}}
    </div>
    <div class="container-fluid">
        {{template "CreateEntityModal" .parent}}
    </div>
</body>
{{template "footer"}}
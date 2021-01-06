<!-- both the CreateEntityModal and AddChildEntityModal use the EntityModal, and should be used in their proper contexts -->

{{define "CreateEntityModal"}}
    <form id="modelform" class="form-horizontal form-well" style="color: #006A4D;" role="form" method="post" action="/CreateEntity">
        {{template "EntityModal"}}
    </form>
    <div class="text-center">
        <a href="" class="btn btn-default btn-rounded my-3" data-toggle="modal" data-target="#modalLRForm">Add new Investment</a>
    </div>
{{end}}


{{define "AddChildEntityModal"}}
    <form id="modelform" class="form-horizontal form-well" style="color: #006A4D;" role="form" method="post" action="/AddChildEntity">
        {{template "EntityModal" .}}
    </form>
    <div class="text-center">
        <a href="" class="btn btn-default btn-rounded my-3" data-toggle="modal" data-target="#modalLRForm">Add new Investment</a>
    </div>
{{end}}


{{define "EntityModal"}}
    <div class="modal fade" id="modalLRForm" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
        <div class="modal-dialog cascading-modal" role="document">
            <!--Content-->
            <div class="modal-content">
                <!--Modal cascading tabs-->
                <div class="modal-c-tabs">
                    <!-- Nav tabs -->
                    <ul class="nav nav-tabs md-tabs tabs-2 light-blue darken-3" role="tablist">
                    <li class="nav-item">
                        <a class="nav-link active" data-toggle="tab" href="#panel8" role="tab"><i class="fas fa-user mr-1"></i>
                        General</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" data-toggle="tab" href="#panel9" role="tab"><i class="fas fa-user-plus mr-1"></i>
                        ERV</a>
                    </li>
                    </ul>
                    <!-- Tab panels -->
                    <div class="tab-content">
                        <!-- General tab -->
                        <div class="tab-pane fade in show active" id="panel8" role="tabpanel">
                            <div class="modal-body mb-1">
                                <div class="form-group">
                                    <div class="form-group">
                                        <label for="contract">Parent</label>
                                        <input type="text" readonly class="form-control" id="parent" name="parent" Value="{{.Name}}">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <div class="form-group">
                                        <label for="name">Name</label>
                                        <input type="text" class="form-control" id="name" name="name">
                                    </div>
                                </div>
                            </div>
                        </div>
                        <!--ERV tab-->
                        <div class="tab-pane fade" id="panel9" role="tabpanel">
                            <div class="modal-body">
                            </div>
                        </div>
                    <!--/.Panel 8-->
                    </div>
                </div>
                <div class="d-flex justify-content-around">
                    <div class="float-left">
                        <button type="submit" class="btn" style="background-color: #006A4D; color:white">Submit</button>
                    </div>
                    <div class="float-right">
                        <button type="button" class="btn" data-dismiss="modal">Close</button>
                    </div>
                </div>
                <br>
            </div>
        </div>
    </div>
{{end}}
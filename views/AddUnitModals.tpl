[[define "AddChildUnitModal"]]
    <form id="unitmodal" class="form-horizontal form-well" style="color: #006A4D;" role="form" method="post" action="/AddChildUnit">
        <div class="modal fade" id="unitmodalLRForm" tabindex="-1" role="dialog" aria-labelledby="myModalLabelunit" aria-hidden="true">
            <div class="modal-dialog cascading-modal" role="document">
                <!--Content-->
                <div class="modal-content">
                    <!--Modal cascading tabs-->
                    <div class="modal-c-tabs">
                        <!-- Nav tabs -->
                        <ul class="nav nav-tabs md-tabs tabs-2 light-blue darken-3" role="tablist">
                        <li class="nav-item">
                            <a class="nav-link active" data-toggle="tab" href="#panel7" role="tab"><i class="fas fa-user mr-1"></i>
                            General</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" data-toggle="tab" href="#panelerv" role="tab"><i class="fas fa-user-plus mr-1"></i>
                            ERV</a>
                        </li>
                        </ul>
                        <!-- Tab panels -->
                        <div class="tab-content">
                            <!-- General tab -->
                            <div class="tab-pane fade in show active" id="panel7" role="tabpanel">
                                <div class="modal-body mb-1">
                                    <div class="form-group">
                                        <div class="form-group">
                                            <label for="parent">Parent</label>
                                            <input type="text" readonly class="form-control" id="parent" name="parent" Value="[[.]]">
                                        </div>
                                    </div>
                                    <div class="form-group">
                                        <div class="form-group">
                                            <label for="unitname">Unit Name</label>
                                            <input type="text" class="form-control" id="unitname" name="unitname">
                                        </div>
                                    </div>
                                    <div class="form-group col-sm-1 d-flex align-items-center">
                                        <div class="btn-group btn-group-toggle" data-toggle="buttons">
                                            <label class="btn btn-secondary active" id="occupiedlabel" onclick="statusOccupied()">
                                                <input type="radio" name="status" id="occupied" value="Occupied">Occupied
                                            </label>
                                            <label class="btn btn-secondary" id="vacantlabel" onclick="statusVacant()">
                                                <input type="radio" name="status" id="vacant" value="Vacant">Vacant
                                            </label>
                                        </div>
                                    </div>
                                    <div class="form-row">&nbsp
                                        <div class="form-group">
                                            <div class="form-group">
                                                <label for="tenant">Tenant</label>
                                                <input type="text" class="form-control" id="tenant" name="tenant">
                                            </div>
                                        </div>&nbsp &nbsp &nbsp &nbsp &nbsp   
                                        <div class="form-group">
                                            <div class="form-group">
                                                <label for="rent">Passing Rent</label>
                                                <input type="text" class="form-control" id="rent" name="rent">
                                            </div>
                                        </div>
                                    </div>
                                    <div class="form-row">&nbsp
                                        <div class="form-group">
                                            <div class="form-group">
                                                <label for="startmonth">Lease Start Month</label>
                                                <input type="text" class="form-control" id="leasestartmonth" name="startmonth">
                                            </div>
                                        </div>&nbsp &nbsp &nbsp &nbsp &nbsp
                                        <div class="form-group">
                                            <div class="form-group">
                                                <label for="startyear">Lease Start Year</label>
                                                <input type="text" class="form-control" id="leasestartyear" name="startyear">
                                            </div>
                                        </div>								
                                    </div>
                                    <div class="form-row">&nbsp			
                                        <div class="form-group">
                                            <div class="form-group">
                                                <label for="expirymonth">Lease Expiry Month</label>
                                                <input type="text" class="form-control" id="expirymonth" name="expirymonth">
                                            </div>
                                        </div>&nbsp &nbsp &nbsp &nbsp &nbsp    
                                        <div class="form-group">
                                            <div class="form-group">
                                                <label for="expiryyear">Lease Expiry Year</label>
                                                <input type="text" class="form-control" id="expiryyear" name="expiryyear">
                                            </div>
                                        </div>								
                                    </div>
                                </div>
                            </div>
                            <!--ERV tab-->
                            <div class="tab-pane fade" id="panelerv" role="tabpanel">
                                <div class="modal-body">
                                    <div class="form-row">			
                                        <div class="form-group">
                                            <div class="form-group">
                                                <label for="retail">Amount</label>
                                                <input type="text" class="form-control" id="amount" name="amount">
                                            </div>
                                        </div>&nbsp &nbsp &nbsp &nbsp &nbsp
                                        <div class="form-group">
                                            <div class="form-group">
                                                <label for="retailarea">Area</label>
                                                <input type="text" class="form-control" id="area" name="area">
                                            </div>
                                        </div>
                                        <div class="form-group">
                                            <div class="form-group">
                                                <label for="type">Type</label>
                                                <input type="text" class="form-control" id="type" name="type">
                                            </div>
                                        </div>								
                                    </div>
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
    </form>
    <div class="text-center">
        <a href="" class="btn btn-default btn-rounded my-3" data-toggle="modal" data-target="#unitmodalLRForm">Add new Unit</a>
    </div>
    <script>
        function statusOccupied() {
            document.getElementById('tenant').readOnly = false;
            document.getElementById('rent').readOnly = false;
            document.getElementById('leasestartmonth').readOnly = false;            
            document.getElementById('leasestartyear').readOnly = false;
            document.getElementById('expirymonth').readOnly = false;            
            document.getElementById('expiryyear').readOnly = false;
        }

        function statusVacant() {
            document.getElementById('tenant').readOnly = true;
            document.getElementById('tenant').value = '';
            document.getElementById('rent').readOnly = true;
            document.getElementById('rent').value = '';
            document.getElementById('leasestartmonth').readOnly = true;
            document.getElementById('leasestartmonth').value = '';
            document.getElementById('leasestartyear').readOnly = true;
            document.getElementById('leasestartyear').value = '';
            document.getElementById('expirymonth').readOnly = true;
            document.getElementById('expirymonth').value = '';
            document.getElementById('expiryyear').readOnly = true;
            document.getElementById('expiryyear').value = '';
        }
    </script>
[[end]]
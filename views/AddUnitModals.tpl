[[define "AddChildUnitModal"]]
    <div class="jw-modal" id="addUnitModal">
        <div class="jw-modal-body container">
            <div class="form-group">
                <div class="form-group">
                    <input type="text" hidden class="form-control" id="new_unit_parent" name="parent" Value="[[.]]">
                </div>
            </div>
            <div class="form-group">
                <div class="form-group">
                    <label for="unitname">Unit Name</label>
                    <input type="text" class="form-control" id="new_unit_unitname" name="unitname">
                </div>
            </div>
            <div class="btn-group" role="group" aria-label="Basic example">
                <button type="button" class="btn secondary-bg" id="new_unit_occupied" onclick="statusOccupied()">Occupied</button>
                <button type="button" class="btn secondary-bg" id="new_unit_vacant" onclick="statusVacant()">Vacant</button>
            </div>
            <div class="form-row">&nbsp
                <div class="form-group">
                    <div class="form-group">
                        <label for="tenant">Tenant</label>
                        <input type="text" class="form-control" id="new_unit_tenant" name="tenant">
                    </div>
                </div>&nbsp &nbsp &nbsp &nbsp &nbsp   
                <div class="form-group">
                    <div class="form-group">
                        <label for="rent">Annual Passing Rent</label>
                        <input type="text" class="form-control" id="new_unit_rent" name="rent">
                    </div>
                </div>
            </div>
            <div class="form-row">&nbsp
                <div class="form-group">
                    <div class="form-group">
                        <label for="startmonth">Lease Start Month</label>
                        <input type="text" class="form-control" id="new_unit_startmonth" name="startmonth">
                    </div>
                </div>&nbsp &nbsp &nbsp &nbsp &nbsp
                <div class="form-group">
                    <div class="form-group">
                        <label for="startyear">Lease Start Year</label>
                        <input type="text" class="form-control" id="new_unit_startyear" name="startyear">
                    </div>
                </div>								
            </div>
            <div class="form-row">&nbsp			
                <div class="form-group">
                    <div class="form-group">
                        <label for="expirymonth">Lease Expiry Month</label>
                        <input type="text" class="form-control" id="new_unit_expirymonth" name="expirymonth">
                    </div>
                </div>&nbsp &nbsp &nbsp &nbsp &nbsp    
                <div class="form-group">
                    <div class="form-group">
                        <label for="expiryyear">Lease Expiry Year</label>
                        <input type="text" class="form-control" id="new_unit_expiryyear" name="expiryyear">
                    </div>
                </div>								
            </div>
            <div class="form-row">			
                <div class="form-group">
                    <div class="form-group">
                        <label for="retail">Amount</label>
                        <input type="text" class="form-control" id="new_unit_amount" name="amount">
                    </div>
                </div>&nbsp &nbsp &nbsp &nbsp &nbsp
                <div class="form-group">
                    <div class="form-group">
                        <label for="retailarea">Area</label>
                        <input type="text" class="form-control" id="new_unit_area" name="area">
                    </div>
                </div>
            </div>        
            <div class="d-flex justify-content-around">
                <div class="float-left">
                    <button class="btn" type="button" id="submitUnit" ng-click="addUnit()">Submit</button>
                </div>
                <div class="float-right">
                    <button class="btn" type="button" id="closeSubmitUnit" onclick="closeModal()">Close</button>
                </div>
            </div>
        </div>
    </div>
    <div class="text-center">
        <a href="" class="btn btn-default btn-rounded my-3" onclick="openModal('addUnitModal')">Add new Unit</a>
    </div>
    <script>

        function openModal(id) {
            document.getElementById(id).classList.add('open');
            document.body.classList.add('jw-modal-open');
        }

        // close currently open modal
        function closeModal() {
            document.querySelector('.jw-modal.open').classList.remove('open');
            document.body.classList.remove('jw-modal-open');
        }

        window.addEventListener('load', function() {
        // close modals on background click
            document.addEventListener('click', event => {
                if (event.target.classList.contains('jw-modal')) {
                    closeModal();
                }
            });
        });

        function statusOccupied() {
            document.getElementById('new_unit_vacant').classList.add("btn-off");
            document.getElementById('new_unit_occupied').classList.remove("btn-off");
            document.getElementById('new_unit_tenant').readOnly = false;
            document.getElementById('new_unit_rent').readOnly = false;
            document.getElementById('new_unit_startmonth').readOnly = false;            
            document.getElementById('new_unit_startyear').readOnly = false;
            document.getElementById('new_unit_expirymonth').readOnly = false;            
            document.getElementById('new_unit_expiryyear').readOnly = false;
        }

        function statusVacant() {
            document.getElementById('new_unit_occupied').classList.add("btn-off");
            document.getElementById('new_unit_vacant').classList.remove("btn-off");
            document.getElementById('new_unit_tenant').readOnly = true;
            document.getElementById('new_unit_tenant').value = '';
            document.getElementById('new_unit_rent').readOnly = true;
            document.getElementById('new_unit_rent').value = '';
            document.getElementById('new_unit_startmonth').readOnly = true;
            document.getElementById('new_unit_startmonth').value = '';
            document.getElementById('new_unit_startyear').readOnly = true;
            document.getElementById('new_unit_startyear').value = '';
            document.getElementById('new_unit_expirymonth').readOnly = true;
            document.getElementById('new_unit_expirymonth').value = '';
            document.getElementById('new_unit_expiryyear').readOnly = true;
            document.getElementById('new_unit_expiryyear').value = '';
        }
    </script>
[[end]]
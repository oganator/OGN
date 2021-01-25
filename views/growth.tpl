[[define "growthinput"]]
    <br>
    <div class="container-fluid">
        <table class="table table-hover rounded tableFixHead" id="growth">
            <thead>
                <tr>
                    <th scope="col">Series</th>
                    <th scope="col">Is ERV</th>
                    [[$header := .Model.TableHeader.Blended.Yearly]]
                    [[range $header]]
                    <th scope="col">&nbsp&nbsp&nbsp&nbsp[[.Year]]</th>
                    [[end]]
                </tr>
            </thead>
            <tbody id="growthbody">
                [[$growth := .Model.GrowthInput]]
                [[$data := .Model.GrowthInputData]]
                [[range $key, $value := $growth]]
                [[$iserv := index $data $key ]]
                <tr>
                    <td>
                        <input readonly type ="text" class="form-control" id="[[$key]]" name="growthitem" value="[[$key]]" >
                    </td>
                    <td>
                        <input type ="text" class="form-control" id="[[$key]]iserv" name="[[$key]]iserv" value="[[$iserv.IsERV]]" >
                    </td>
                    [[range $header]]
                        <td>
                            <input type="text" class="form-control" id="[[$key]][[.Year]]" name="[[$key]][[.Year]]" value=[[index $value.Value .Year]]>
                        </td>
                    [[end]]
                </tr>
                [[end]]
            </tbody>
        </table>
    </div>
    <div class="modal fade" id="growthmodal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
        <div class="modal-dialog cascading-modal" role="document">
            <div class="modal-content">
                <div class="modal-c-tabs">
                    <div class="tab-pane fade in show active" id="panel8" role="tabpanel">
                        <div class="modal-body mb-1">
                            <div class="form-row">
                                <div class="form-group">
                                    <label for="name">Name</label>
                                    <input type="text" class="form-control" id="growthitem" name="growthitem">
                                </div>&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp
                                <div class="form-group align-self-center">
                                    <div class="btn-group btn-group-toggle" id="isERVdiv" data-toggle="buttons">
                                        <label class="btn btn-secondary active" id="ervyes">
                                            <input type="radio" name="isERV" id="ervyesinput" checked value="yes">ERV 
                                        </label>
                                        <label class="btn btn-secondary" id="ervno">
                                            <input type="radio" name="isERV" id="ervnoinput" value="no">Not ERV
                                        </label>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="tab-pane fade" id="panel9" role="tabpanel">
                        <div class="modal-body">
                        </div>
                    </div>
                </div>
                <div class="d-flex justify-content-around">
                    <div class="float-left">
                        <button type="button" class="btn" data-dismiss="modal" style="background-color: #006A4D; color:white" onclick="addGrowthRow()">Submit</button>
                    </div>
                    <div class="float-right">
                        <button type="button" class="btn" data-dismiss="modal">Close</button>
                    </div>
                </div>
                <br>
            </div>
        </div>
    </div>
    <div class="text-center">
        <a href="" class="btn btn-default btn-rounded my-3" data-toggle="modal" data-target="#growthmodal">Add new Growth Item</a>
    </div>
    <script>
        window.onload = updateRowName()

        function updateRowName() {
            var table = document.getElementById('growth');
            for (var i = 1, row; row = table.rows[i]; i++) {
                var x = row.children[0];
                var y = x.children[0];
                var name = 'growthitem';
                name = name.concat(i);
                y.setAttribute('name',name);
            }
        }

        function addGrowthRow() {
            var header = [[.Model.TableHeader.Blended.Yearly]];
            var len = header.length;
            var growthitem = document.getElementById('growthitem').value;
            var iserv = false;
            if (document.getElementById('ervyesinput').checked){
                iserv = true;
            }
            var table = document.getElementById('growth');
            var rowCnt = table.rows.length;
            var tr = table.insertRow(rowCnt);
            //
            for (var i = 0; i < len; i++){
                var td = document.createElement('td');
                td = tr.insertCell(i);
                var ele = document.createElement('input');
                ele.setAttribute('type', 'text');
                var nameid = growthitem;
                nameid = nameid.concat(header[i]);
                ele.setAttribute('class', 'form-control');
                ele.setAttribute('name', nameid);
                ele.setAttribute('id', nameid);
                ele.setAttribute('value', '0.0');
                td.appendChild(ele);
            }
            // isERV
            var td = document.createElement('td');
            td = tr.insertCell(0);
            var ele = document.createElement('input');
            ele.setAttribute('type', 'text');
            ele.readOnly = true;
            var name = growthitem;
            name = name.concat('iserv');
            ele.setAttribute('class', 'form-control');
            ele.setAttribute('name', name);
            ele.setAttribute('id', name);
            ele.setAttribute('value', iserv);
            td.appendChild(ele);
            // series
            var td = document.createElement('td');
            td = tr.insertCell(0);
            var ele = document.createElement('input');
            ele.setAttribute('type', 'text');
            ele.readOnly = true;
            var name = 'growthitem';
            name = name.concat(rowCnt);
            ele.setAttribute('class', 'form-control');
            ele.setAttribute('name', name);
            ele.setAttribute('id', growthitem);
            ele.setAttribute('value', growthitem);
            td.appendChild(ele);
        }
    </script>
[[end]]
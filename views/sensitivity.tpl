[[define "sensitivity"]]
<div class="container-fluid" style="width: 95%;">
    <div class="row">
        <div class="container-fluid" style="width: 50%;">
        </div>
        <div class="container-fluid" style="width: 50%;">
            <table class="table table-hover rounded tableFixHead" id="sensitivityGrid">
                <tbody>
                [[range .grid]]
                    [[$row := .]]
                    <tr>
                    [[range $row]]
                        <td>{{[[.]] | number:3}}</td>
                    [[end]]
                    </tr>
                [[end]]
                </tbody>
            </table>
        </div>
    </div>
</div>
[[end]]

[[template "sensitivity" .]]
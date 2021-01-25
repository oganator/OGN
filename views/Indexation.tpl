[[define "indexation"]]
    <div class="container-fluid shadow-lg rounded" style="width: 95%">
        <table class="table table-hover tableFixHead rounded">
            <thead>
                <tr>
                    <th scope="col"></th>
                    [[range .]]
                    <th scope="col">Extension [[.EXTNumber]]
                    [[end]]
                </tr>
            </thead>
            <tbody>
                <tr>
                    <th scope="row">Start Date</th>
                    [[range .]]
                    [[range .RenewIndex]]
                    <td>[[.]]</td>
                    [[end]]
                    [[end]]
                </tr>
                <tr>
                    <th scope="row">Start Date</th>
                    [[range .]]
                    [[range .RotateIndex]]
                    <td>[[.]]</td>
                    [[end]]
                    [[end]]
                </tr>
            </tbody>
        </table>
    </div>
[[end]]
[[template "header"]]
    <body>
        <br>
        <br>
        <div class="container-fluid">
            [[template "UnitSettings" .unit]]
        </div>
        <div class="tabs-wrapper">
            <ul class="nav nav-tabs tabPinned" role="tablist">
                <li class="nav-item">
                    <a class="nav-link waves-light active" data-toggle="tab" href="#cf" role="tab">Cash Flow</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link waves-light" data-toggle="tab" href="#rs" role="tab">Rent Schedule</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link waves-light" data-toggle="tab" href="#ind" role="tab">Indexation</a>
                </li>                
            </ul>
        </div>
        <div class="tab-content">
            <div class="tab-pane fade in show active" id="cf" role="tabpanel">
                [[template "CFTable" .unit.Model]]
            </div>
            <div class="tab-pane fade" id="rs" role="tabpanel">
                [[template "rentschedule" .unit.RentSchedule]]
            </div>
            <div class="tab-pane fade" id="ind" role="tabpanel">
                [[template "indexation" .unit.RentSchedule]]
            </div>
        </div>
[[template "footer"]]
</body>


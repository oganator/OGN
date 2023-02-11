[[template "header" .]]

    <br>
    <div class="row">
        <div>
        </div>
        <div class="row container-fluid " style="width: 95%;" id="body-Home">
            <div class="container-fluid shadow-lg rounded" style="width: 25%;">
                [[template "entitySelection" .entityMap]]
            </div>
            <div class="container-fluid shadow-lg rounded" style="width: 70%;">
                <div id="entityModelTable"></div>
                    <div bind-html-compile = entityModelTableResponse></div>			
                </div>
            </div>
        </div>
        <div bind-html-compile = viewEntityResponse></div>
    </div>


[[template "footer" .]]
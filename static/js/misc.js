$(document).ready(function() {
    var counter = 1;
    $("#addrow").on("click", function() {
        var newRow = $("<tr>");
        var cols = "";
        cols += '<td><input type="text" class="form-control" name="series' + counter + '" id="series' + counter + '"/></td>';
        cols += '<td><input type="button" class="ibtnDel btn btn-md btn-danger "  value="Delete"></td>';
        newRow.append(cols);
        $("table.order-list").append(newRow);
        counter++;
    });
    $("table.order-list").on("click", ".ibtnDel", function(event) {
        $(this).closest("tr").remove();
        counter -= 1
    });
});

function calculateRow(row) {
    var price = +row.find('input[name^="price"]').val();
}

function calculateGrandTotal() {
    var grandTotal = 0;
    $("table.order-list").find('input[name^="price"]').each(function() {
        grandTotal += +$(this).val();
    });
    $("#grandtotal").text(grandTotal.toFixed(2));
}

////////////////////////////////////////////////////////
//FRED
function addRow2(tabl) {
    var table = document.getElementById(tabl);
    var rowCnt = table.rows.length; // table row count.
    var tr = table.insertRow(rowCnt); // the table row.
    var td = document.createElement('td'); // table definition.
    td = tr.insertCell(0);
    var ele = document.createElement('input');
    ele.setAttribute('type', 'text');
    var series = 'series';
    series = series.concat(rowCnt - 1);
    // add a button in every new row in the first column.
    var button = document.createElement('input');
    // set input attributes.
    button.setAttribute('type', 'button');
    button.setAttribute('value', 'Delete');
    button.setAttribute('class', 'btn ');
    button.setAttribute('style', 'background-color: red; color:white');
    button.setAttribute('id', 'delete');
    var rmvbtn = 'removeRow2(this,';
    rmvbtn = rmvbtn.concat("'", tabl, "'", ')');
    button.setAttribute('onclick', rmvbtn);
    td.appendChild(button);
    //
    var td2 = document.createElement('td'); // table definition.
    td2 = tr.insertCell(0);
    ele.setAttribute('name', series);
    ele.setAttribute('id', series);
    ele.setAttribute('class', 'form-control');
    ele.setAttribute('value', "");
    td2.appendChild(ele);
}

// delete table row.
function removeRow2(oButton, tabl) {
    var table = document.getElementById(tabl);
    table.deleteRow(oButton.parentNode.parentNode.rowIndex); // button -> td -> tr.
    updateRowNumber(tabl);
}

function updateRowNumber(tabl) {
    var table = document.getElementById(tabl);
    for (var i = 1, row; row = table.rows[i]; i++) {
        var x = row.children[0];
        var y = x.children[0];
        var series = 'series';
        series = series.concat(i - 1);
        y.setAttribute('name', series);
        y.setAttribute('id', series);
    }
}
//  /FRED

// ENTITY SETTINGS
function checkChange(id) {
    var row = document.getElementById(id);
    row.setAttribute('name', id);
}
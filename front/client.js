function appendTable(tableData) {
    let table = document.createElement('table');
    table.className = "table-bordered table-info align-middle text-center"
    let tableBody = document.createElement('tbody');

    tableData.forEach(function(rowData) {
        let row = document.createElement('tr');

        rowData.forEach(function(cellData) {
            let cell = document.createElement('td');
            cell.style = "width: 20px; height: 20px;"
            cell.appendChild(document.createTextNode(cellData));
            row.appendChild(cell);
        });

        tableBody.appendChild(row);
    });

    table.appendChild(tableBody);
    let container = document.getElementById("table-content")
    container.innerHTML = table.outerHTML
}

function appendError(error) {
    let container = document.getElementById("table-content")
    let alert = document.createElement('div');
    alert.className = "alert alert-danger"
    alert.innerText = error
    container.innerHTML = alert.outerHTML
}

const getCanvas = () => {
    let canvasID = document.getElementById("canvas-id-input-current")
    if (canvasID.value.length != 0) {
        fetch("http://127.0.0.1:8080/canvas/" + canvasID.value)
            .then(function (response) {
                switch (response.status) {
                    case 200:
                        return response.text()
                    case 400:
                    case 404:
                    case 500:
                        appendError(response.statusText)
                        break;
                    default:
                        appendError("unexpected error")
                }
            })
            .then(function (body) {
                if (body != null) {
                    appendTable(JSON.parse(body).data)
                }
            })
            .catch(function (error) {
                appendError(error)
            });
    }

    setTimeout(getCanvas, 1000);
}

function changeCanvasID() {
    let canvasID = document.getElementById("canvas-id-input")
    let canvasIDCurrent = document.getElementById("canvas-id-input-current")
    canvasIDCurrent.value = canvasID.value
}

window.onload = function() {
    let observe = document.getElementById("canvas-id-observe");
    observe.addEventListener("click", changeCanvasID, false);
    getCanvas()
}

$(document).ready(function() {

    ActualizarHistorial();
    $("#txtTexto").keyup(function(event) {
        if (event.keyCode === 13) {
            $("#btnEnviar").click();
        }
    });
    $("#btnEnviar").click(function() {
        var texto = $("#txtTexto").val();
        console.log(texto);
        
        var envio = {
            palabra: texto
        };

        $.post({
            url:"/envio",
            data: JSON.stringify(envio),
            success: function(data, status, jqXHR) {
                console.log(data);
                $("#txtTexto").val('')
            },
            dataType: "json"

        }).done(function(data) {
            console.log("Petición realizada");
            ActualizarHistorial();
        
        }).fail(function(data) {
            console.log("Petición fallida");
        
        }).always(function(data){
            console.log("Petición completa");
        });
    });
});

function ActualizarHistorial() {
    var filtro = {
        fecha: moment().format('YYYY-MM-DDT00:00:00Z')
    };
    $.ajax({
        url: "/lista",
        method: "POST",
        data: JSON.stringify(filtro),
        dataType: "json",
        contentType: "application/json",
        success: function(data) {
            if(data != null)
                console.log(data.length + " objetos obtenidos");
            Historial_UI(data);
        },
        error: function(data) {
            console.log(data);
        }
    });
}

function Historial_UI(array) {
    var tbody = $("#historial tbody");
    tbody.children().remove();
    if(array != null && array.length > 0) {

        for(var x = 0; x < array.length; x++) {
            tbody.append(
                "<tr><td>" + array[x].ID + 
                "</td><td>" + array[x].Palabra + 
                "</td><td>" + moment(array[x].Fecha).format("DD-MM-YY HH:mm:ssZ") + 
                "</td></tr>");
        }
    } else {
        tbody.append('<tr><td colspan="3">No hay registros de hoy</td></tr>');
        
    }
}
function save(){
    var toSave = $("#save input").val()
    $('#tasks').append(toSave+"<br>");
    $.post("/tasks",{task:toSave},function(res,err){

    })
}

function allTask(){
    $.get("/tasks",function(res,err){
            $('#tasks').html(res);
        })
}
window.load = allTask()


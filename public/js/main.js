function save(){
    var toSave = $("#save input").val()
    $.post("/tasks",{task:toSave},function(res,err){
        $('#tasks').append(toSave);
    })
}

function allTask(){
    $.get("/tasks",function(res,err){
            $('#tasks').html(res);
        })
}
window.load = allTask()


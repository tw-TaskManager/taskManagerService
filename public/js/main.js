function save(){
    var toSave = $("#save input").val()
    $.post("/tasks",{task:toSave},function(res,err){
        alert(res);
    })
}
